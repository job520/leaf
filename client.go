package main
import (
	"encoding/binary"
	"net"
)
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}
	data := []byte(`{
		"Game": {
			"Event": "move:to",
			"Data": {
				"name": "a",
				"position": {
					"x": 0,
					"y": 0
				}
			}
		}
	}`)
	m := make([]byte, 2+len(data))
	binary.BigEndian.PutUint16(m, uint16(len(data)))
	copy(m[2:], data)
	conn.Write(m)
}