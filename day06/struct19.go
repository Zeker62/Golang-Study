package main

import (
	"encoding/json"
	"fmt"
)

type student19 struct {
	ID int `json:"id"` //在json中ID字段就会变成id
	Sex string
	name string //私有不能被json包访问
}


func main() {
	s1:=student19{
		ID:1,
		name:"aaa",
		Sex:"man",
	}
	data, err :=json.Marshal(s1)
	if err !=nil{
		fmt.Println("json序列化出现错误")
		return
	}
	fmt.Printf("json :%s\n",data)

}
//json :{"id":1,"Sex":"man"}