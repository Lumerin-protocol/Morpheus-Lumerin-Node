package tcphandlers

import (
	"bufio"
	"context"
	"encoding/json"
	"net"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/proxyapi"
	morrpc "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/proxyapi/morrpcmessage"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/transport"
)

func NewTCPHandler(
	log, connLog lib.ILogger,
	schedulerLogFactory func(contractID string) (lib.ILogger, error),
	morRpcHandler *proxyapi.MORRPCController,
) transport.Handler {
	return func(ctx context.Context, conn net.Conn) {
		addr := conn.RemoteAddr().String()
		sourceLog := connLog.Named("SRC").With("SrcAddr", addr)

		defer func() {
			sourceLog.Info("Closing connection")
			conn.Close()
		}()

		msg, err := getMessageV2(conn)
		if err != nil {
			sourceLog.Error("Error reading message", err)
			return
		}

		err = morRpcHandler.Handle(ctx, *msg, sourceLog, func(resp *morrpc.RpcResponse) error {
			_, err := sendMsg(conn, resp)
			if err != nil {
				sourceLog.Error("Error sending message", err)
				return err
			}
			sourceLog.Debug("sent message")
			return err
		})
		if err != nil {
			sourceLog.Error("Error handling message", err)
			return
		}
	}
}

func sendMsg(conn net.Conn, msg *morrpc.RpcResponse) (int, error) {
	msgJson, err := json.Marshal(msg)
	if err != nil {
		return 0, err
	}
	return conn.Write(msgJson)
}

func getMessageV2(conn net.Conn) (*morrpc.RPCMessageV2, error) {
	reader := bufio.NewReader(conn)
	d := json.NewDecoder(reader)

	var msg *morrpc.RPCMessageV2
	err := d.Decode(&msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
