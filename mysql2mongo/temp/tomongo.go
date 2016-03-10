package main

import (
	"fmt"
)


func main(){
	array := []string{"a","b","c","d"}
	array = append(array,"e","f")
	//array := [5]int{1,2,3,4,5}
	for i,v := range array {
		fmt.Println(i,v)
	}
	fmt.Println(len(array))
}
