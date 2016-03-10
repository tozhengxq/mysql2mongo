package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// 里面的东西必须大写，如 NAME ，要不然插入为空
// 一开始查询的时候，都没有返回结果，后面就好了，注意大小写就好
//
//
type Man struct {
	NAME string
	SEX  string
	AGE  uint32
}

const URL = "localhost:27017"

func main() {

	session, err := mgo.Dial(URL)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("test")
	collection := db.C("t1")

	Person := &Man{
		NAME: "zheng123",
		SEX:  "manw",
		AGE:  10,
	}

	err = collection.Insert(Person)
	if err != nil {
		fmt.Println(err)
	}
	///
	result := Man{}
	err = collection.Find(bson.M{"age": 29}).One(&result)
	fmt.Println("result:", result.NAME, result.SEX, result.AGE)

}
