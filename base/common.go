package base

import "github.com/name5566/leaf/gate"

/**
* 删除切片中指定的元素
* @param arr 原切片
* @param val 要删除的值
* @return arr 删除后的切片
 */
func SliceDelete(arr []gate.Agent,val gate.Agent) []gate.Agent {
	for k,v := range arr{
		if v == val{
			arr = append(arr[:k],arr[k + 1:]...)
		}
	}
	return arr
}