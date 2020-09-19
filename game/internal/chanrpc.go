package internal
import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"server/base"
)
func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("Do", do)
}
func rpcNewAgent(args []interface{}) {
	fmt.Println("[game - chanrpc]创建时自动调用")
	conn := args[0].(gate.Agent)
	base.Join("system",conn)
}
func rpcCloseAgent(args []interface{}) {
	fmt.Println("[game - chanrpc]销毁时自动调用")
	conn := args[0].(gate.Agent)
	base.Leave("system",conn)
}
func do(args []interface{}) {
	fmt.Println("==========================")
	fmt.Println("======= leaf chanrpc =======")
	fmt.Println(args)
	fmt.Println("==========================")
}