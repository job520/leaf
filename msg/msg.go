package msg
import (
	"github.com/name5566/leaf/network/json"
)
var Processor = json.NewProcessor()
func init() {
	Processor.Register(&Login{})
	Processor.Register(&Game{})
	Processor.Register(&Return{})
}
// 接收值
type Login struct {
	Event string
	Data interface{}
}
type Game struct {
	Event string
	Data interface{}
}
// 返回值
type Return struct {
	Event string
	Msg string
	Data interface{}
}