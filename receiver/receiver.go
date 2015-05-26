package receiver

import (
	"github.com/my-open-falcon/transfer/receiver/rpc"
	"github.com/my-open-falcon/transfer/receiver/socket"
)

func Start() {
	go rpc.StartRpc()
	go socket.StartSocket()
}
