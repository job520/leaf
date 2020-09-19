package internal

import (
	"reflect"
	"server/base"
	"server/msg"
)
func init() {
	handler(&msg.Game{}, handleGame)
}
func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
func handleGame(args []interface{}) {
	data := args[0].(*msg.Game)
	switch data.Event {
		case "move:to":
			input_data := data.Data.(map[string]interface{})
			name := input_data["name"].(string)
			position := input_data["position"].(map[string]interface{})
			base.SetUser(name, map[string]float64{
				"x": position["x"].(float64),
				"y": position["y"].(float64),
			})
			base.Broadcast("system",&msg.Return{
				Event: "move:to",
				Msg: "移动",
				Data: map[string]interface{}{
					"name": name,
					"position": position,
				},
			})
	}
}

