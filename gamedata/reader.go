package gamedata

import (
	"fmt"
	"github.com/name5566/leaf/recordfile"
	"reflect"
)

func ReadRf(st interface{}) *recordfile.RecordFile {
	rf, err := recordfile.New(st)
	if err != nil {
		fmt.Println("读取文件出错[1]：",err)
	}
	fn := reflect.TypeOf(st).Name() + ".csv"
	err = rf.Read("gamedata/" + fn)
	if err != nil {
		fmt.Println("读取文件出错[2]：",err)
	}

	return rf
}

