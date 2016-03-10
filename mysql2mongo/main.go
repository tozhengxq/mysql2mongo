package main

import (
	"flag" // 只是读取命令行参数，返回参数value的指针
	"fmt"
	"io/ioutil"
	"mysql2mongo/config"
)

var configFile *string = flag.String("conf", "/etc/test.conf", "mysql2mongo config file")

func main() {
	conf := config.ParseConfigfile(*configFile)
	fmt.Println(conf.Collection, conf.DB, conf.URL)

}
