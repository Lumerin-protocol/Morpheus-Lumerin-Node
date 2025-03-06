package proxyapi

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/ipfs/boxo/files"
	ipfspath "github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/client/rpc"
	iface "github.com/ipfs/kubo/core/coreiface"
)

type IpfsManager struct {
	node *rpc.HttpApi
}

// NewIpfsManager connects to the local Kubo (IPFS) node using the RPC client.
func NewIpfsManager() *IpfsManager {
	node, err := rpc.NewLocalApi()
	if err != nil {
		fmt.Println("Error creating local API:", err)
		return nil
	}
	return &IpfsManager{node: node}
}

// AddFile adds a file to IPFS by reading the file from disk.
func (i *IpfsManager) AddFile(ctx context.Context, filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	fileNode := files.NewReaderFile(f)

	addedOutput, err := i.node.Unixfs().Add(ctx, fileNode)
	if err != nil {
		return "", fmt.Errorf("failed to add file: %w", err)
	}

	return addedOutput.RootCid().String(), nil
}

// Pin pins a CID on the local IPFS node.
func (i *IpfsManager) Pin(ctx context.Context, cidStr string) error {
	c, err := cid.Decode(cidStr)
	if err != nil {
		return fmt.Errorf("invalid CID: %w", err)
	}

	p := ipfspath.FromCid(c)
	if err := i.node.Pin().Add(ctx, p); err != nil {
		return fmt.Errorf("failed to pin CID: %w", err)
	}
	return nil
}

// Unpin removes a pin for a given CID.
func (i *IpfsManager) Unpin(ctx context.Context, cidStr string) error {
	c, err := cid.Decode(cidStr)
	if err != nil {
		return fmt.Errorf("invalid CID: %w", err)
	}

	p := ipfspath.FromCid(c)
	if err := i.node.Pin().Rm(ctx, p); err != nil {
		return fmt.Errorf("failed to unpin CID: %w", err)
	}
	return nil
}

func (i *IpfsManager) GetFile(ctx context.Context, cidStr string, destinationPath string) error {
	// 1) "Probe" phase: short context to see if we can get ANY data
	probeCtx, probeCancel := context.WithTimeout(ctx, 10*time.Second)
	defer probeCancel()

	c, err := cid.Decode(cidStr)
	if err != nil {
		return fmt.Errorf("invalid CID: %w", err)
	}

	probeNode, err := i.node.Unixfs().Get(probeCtx, ipfspath.FromCid(c))
	if err != nil {
		return fmt.Errorf("failed to find file within 10s: %w", err)
	}

	probeFile, ok := probeNode.(files.File)
	if !ok {
		return fmt.Errorf("object at CID %s is not a regular file", cidStr)
	}

	// Read a small chunk to ensure data is available
	buf := make([]byte, 1)
	_, err = probeFile.Read(buf)
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to read data in probe phase: %w", err)
	}

	probeFile.Close()

	downloadCtx, downloadCancel := context.WithTimeout(ctx, 30*time.Minute)
	defer downloadCancel()

	node, err := i.node.Unixfs().Get(downloadCtx, ipfspath.FromCid(c))
	if err != nil {
		return fmt.Errorf("download phase failed: %w", err)
	}
	defer node.Close()

	fileNode, ok := node.(files.File)
	if !ok {
		return fmt.Errorf("object at CID %s is not a regular file", cidStr)
	}

	destFile, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, fileNode); err != nil {
		return fmt.Errorf("failed to copy content: %w", err)
	}

	return nil
}

func (i *IpfsManager) GetVersion(ctx context.Context) (string, error) {
	var resp struct {
		Version string `json:"Version"`
		Commit  string `json:"Commit,omitempty"`
		Repo    string `json:"Repo,omitempty"`
		System  string `json:"System,omitempty"`
		Golang  string `json:"Golang,omitempty"`
	}

	err := i.node.Request("version").Exec(ctx, &resp)
	if err != nil {
		return "", fmt.Errorf("request for version failed: %w", err)
	}

	return resp.Version, nil
}

func (i *IpfsManager) GetPinnedFiles(ctx context.Context) ([]string, error) {
	pinChan := make(chan iface.Pin)
	errChan := make(chan error, 1)

	// Call Pin().Ls in a goroutine, passing it pinChan
	// That function will close pinChan when done
	go func() {
		err := i.node.Pin().Ls(ctx, pinChan)
		errChan <- err
		// We *only* close errChan ourselves; do NOT close pinChan here
		close(errChan)
	}()

	var pinnedCIDs []string
	// Now read pins from pinChan until the library closes it
	for pin := range pinChan {
		pinnedCIDs = append(pinnedCIDs, pin.Path().RootCid().String())
	}

	// When pinChan is closed, we exit the for-loop. Check the final error:
	if err := <-errChan; err != nil {
		return nil, fmt.Errorf("failed to list pinned files: %w", err)
	}

	return pinnedCIDs, nil
}
