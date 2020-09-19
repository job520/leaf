package base

import (
	"github.com/name5566/leaf/gate"
	"sync"
)


var Conns map[string][]gate.Agent = map[string][]gate.Agent{}

/**
 * 加入房间
 * @param room 房间名
 * @param conn 当前连接
 */
func Join(room string,conn gate.Agent) {
	Conns[room] = append(Conns[room],conn)
}

/**
 * 离开房间
 * @param room 房间名
 * @param conn 当前连接
 */
func Leave(room string,conn gate.Agent) {
	Conns[room] = SliceDelete(Conns[room],conn)
}

/**
 * 广播消息
 * @param room 房间名
 * @param msg 内容
 */
func Broadcast(room string,msg interface{}) {
	wg := sync.WaitGroup{}
	_len := len(Conns[room])
	wg.Add(_len)
	for _,conn := range Conns[room]{
		go  func(conn gate.Agent,wg *sync.WaitGroup) {
			conn.WriteMsg(msg)
			wg.Done()
		}(conn,&wg)
	}
	wg.Wait()
}