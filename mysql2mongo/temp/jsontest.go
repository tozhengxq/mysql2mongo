package main

// json 
// []byte json是字节数组类型的
// 可以定义一个结构体，然后用json.Marshal and json.Unmarshal 来转换为json格式和解码json
// json.Marshal 编译的json编码，是 []byte{} 字节数组格式的
import (
	"fmt"
	"encoding/json"	
)

type jdorm struct {
	item_id string `json:"_id"`
	dormentry_id string
	dorm_id string
	site_id string
}


func main(){
	dorm := jdorm{
		"a",
		"b",
		"c",
		"d",
	}
	jdorm1,err := json.Marshal(dorm)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(jdorm1)
}
