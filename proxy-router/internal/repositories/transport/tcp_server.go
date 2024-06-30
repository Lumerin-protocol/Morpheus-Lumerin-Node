package transport

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/netip"
	"sync"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
)

type TCPServer struct {
	serverAddr string
	handler    Handler
	log        lib.ILogger
}

func NewTCPServer(serverAddr string, log lib.ILogger) *TCPServer {
	return &TCPServer{
		serverAddr: serverAddr,
		log:        log,
	}
}

func (p *TCPServer) SetConnectionHandler(handler Handler) {
	p.handler = handler
}

func (p *TCPServer) Run(ctx context.Context) error {
	add, err := netip.ParseAddrPort(p.serverAddr)
	if err != nil {
		return fmt.Errorf("invalid server address %s %w", p.serverAddr, err)
	}

	listener, err := net.Listen("tcp", add.String())

	if err != nil {
		return fmt.Errorf("listener error %s %w", p.serverAddr, err)
	}

	p.log.Infof("tcp server is listening: %s", p.serverAddr)

	serverErr := make(chan error, 1)

	go func() {
		serverErr <- p.startAccepting(ctx, listener)
	}()

	select {
	case <-ctx.Done():
		err := listener.Close()
		if err != nil {
			return err
		}
		err = ctx.Err()
		p.log.Infof("tcp server closed: %s", p.serverAddr)
		<-serverErr
		return err
	case err = <-serverErr:
	}

	return err
}

func (p *TCPServer) startAccepting(ctx context.Context, listener net.Listener) error {
	wg := sync.WaitGroup{} // waits for all handlers to finish to ensure proper cleanup
	defer wg.Wait()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		conn, err := listener.Accept()

		if errors.Is(err, net.ErrClosed) {
			return fmt.Errorf("incoming connection listener was closed")
		}

		if err != nil {
			p.log.Errorf("incoming connection accept error: %s", err)
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			p.log.Debugf("incoming connection accepted: %s", conn.RemoteAddr().String())
			p.handler(ctx, conn)

			err = conn.Close()
			if errors.Is(err, net.ErrClosed) {
				return
			}
			if err != nil {
				p.log.Warnf("error during closing connection: %s", err)
				return
			}
			p.log.Debugf("incoming connection closed")
		}()

	}
}
