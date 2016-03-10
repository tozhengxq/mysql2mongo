package main

import (
	"flag" // 只是读取命令行参数，返回参数value的指针
	"fmt"
	"github.com/tozhengxq/mysql2mongo/config"
)

var configFile *string = flag.String("conf", "/etc/test.conf", "mysql2mongo config file")

func main() {
	conf, err := config.ParseConfigfile(*configFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf.Collection)

}
