package internal
import (
	"github.com/name5566/leaf/gate"
	"reflect"
	"server/msg"
	"server/base"
)
func init() {
	handler(&msg.Login{}, handleLogin)
}
func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
func handleLogin(args []interface{}) {
	data := args[0].(*msg.Login)
	conn := args[1].(gate.Agent)
	switch data.Event {
		case "do:login":
			input_data := data.Data.(map[string]interface{})
			name := input_data["name"].(string)
			exists := base.UserExists(name)
			// 初始位置
			position := map[string]float64{
				"x": 0,
				"y": 0,
			}
			if exists == true{
				conn.WriteMsg(&msg.Return{
					Event: "login:fail",
					Msg: "登录失败",
					Data: []interface{}{},
				})
				conn.Close()
			}else {
				base.AddUser(name,position)
				users := base.GetAllUser()
				conn.WriteMsg(&msg.Return{
					Event: "login:success",
					Msg: "登录成功",
					Data: users,
				})
				base.Broadcast("system",&msg.Return{
					Event: "new:user",
					Msg: "新用户加入",
					Data: map[string]interface{}{
						"name": name,
						"position": position,
					},
				})
			}
	}
	///////////////////////////////////////////
	///////////////////////////////////////////
	///////////////////////////////////////////
	//// todo: Leaf ChanRPC（模块间调用）
	//game.ChanRPC.Go("Do", "param1","param2")
	//// todo: Leaf Go（leaf 自带 go程）
	//var count int
	//skeleton.Go(func() {
	//	for i := 0;i < 1000;i ++{
	//		count += i;
	//	}
	//}, func() {
	//	fmt.Println("==========================")
	//	fmt.Println("========= leaf go =========")
	//	fmt.Println(count)
	//	fmt.Println("==========================")
	//})
	//// todo: Leaf timer（计时器）
	//skeleton.AfterFunc(3 * time.Second, func() {
	//	fmt.Println("==========================")
	//	fmt.Println("======= leaf timer =======")
	//	fmt.Println("延迟 3秒输出的消息")
	//	fmt.Println("==========================")
	//})
	//// todo: Leaf recordfile（CSV 文本读取）
	//fmt.Println("==========================")
	//fmt.Println("===== leaf recordfile =====")
	//type Test struct {
	//	Id  int "index"
	//	Str string
	//}
	//var RfTest = gamedata.ReadRf(Test{})
	//r := RfTest.Index(1)
	//if r != nil {
	//	row := r.(*Test)
	//	fmt.Println(row)
	//}
	//fmt.Println("==========================")
}

