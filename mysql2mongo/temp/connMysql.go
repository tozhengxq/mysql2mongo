package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Order struct {
	Order_id uint64
	Dormid   []byte
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/data")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("select dorm_id,order_id from 59_order where order_id = 2014120216934128;")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		//var order1 Order
		//var s NullString
		var o Order
		//var dormid []byte
		if err := rows.Scan(&o.Dormid, &o.Order_id); err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(o.Dormid))
		fmt.Println(o.Order_id)

	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

}
