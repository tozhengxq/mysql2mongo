package main

import (
	"database/sql"
	"flag" // 只是读取命令行参数，返回参数value的指针
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tozhengxq/mysql2mongo/config"
	"labix.org/v2/mgo"
)

var configFile *string = flag.String("conf", "/etc/mysql2mongo.conf", "mysql2mongo config file")

//const (
//	URL  string = "localhost:27017"
//	URL2 string = "root:123456@tcp(localhost:3306)/data"
//	SQL  string = "select Order_id,Status,type,Paytype,Paystatus,Pay_trade_no,Source,Consumption_type,Suspicious,Sid,Site_id,Dorm_id,Dormentry_id,Shop_id,Uid,Service_eva,Delivery_eva,Food_eva,Food_num,Food_amount,Ship_fee,Coupon_discount,Promotion_discount,Discount,Order_amount,Delivery_id,Add_time,Confirm_time,Send_time,Expect_date,Delivery_type,Expect_time,Expect_timeslot,Order_mark,Uname,Portrait,Phone,Phone_addr,Buy_times,Address1,Address2,Dormitory,Time_deliver,Credit,Ip,Coupon_code,Feature,Remark,Evaluation,Expect_start_time,Expect_end_time from 59_order;"
//)

type Order struct {
	Order_id           uint64
	Status             uint32
	Otype              uint32
	Paytype            string
	Paystatus          string
	Pay_trade_no       string
	Source             string
	Consumption_type   int32
	Suspicious         int32
	Sid                uint32
	Site_id            uint32
	Dorm_id            string // uint32 null
	Dormentry_id       string
	Shop_id            uint32
	Uid                uint64
	Service_eva        string
	Delivery_eva       string
	Food_eva           string
	Food_num           string
	Food_amount        float32
	Ship_fee           string
	Coupon_discount    float32
	Promotion_discount float32
	Discount           float32
	Order_amount       string
	Delivery_id        string
	Add_time           uint32
	Confirm_time       uint32
	Send_time          uint32
	Expect_date        uint32
	Delivery_type      string
	Expect_time        uint32
	Expect_timeslot    string
	Order_mark         string
	Uname              string
	Portrait           string
	Phone              string
	Phone_addr         string
	Buy_times          string
	Address1           string
	Address2           string
	Dormitory          string
	Time_deliver       string
	Credit             uint32
	Ip                 string
	Coupon_code        string
	Feature            string
	Remark             string
	Evaluation         string
	Expect_start_time  uint32
	Expect_end_time    uint32
}

var (
	orderitem Order
	//
	// 定义[]byte 类型，用以接收有可能出线null值的列
	BPaytype       []byte
	BPaystatus     []byte
	BPay_trade_no  []byte
	BSource        []byte
	BDorm_id       []byte
	BDormentry_id  []byte
	BService_eva   []byte
	BDelivery_eva  []byte
	BFood_eva      []byte
	BFood_num      []byte
	BShip_fee      []byte
	BOrder_amount  []byte
	BDelivery_id   []byte
	BDelivery_type []byte
	BOrder_mark    []byte
	BUname         []byte
	BPortrait      []byte
	BPhone_addr    []byte
	BBuy_times     []byte
	BDormitory     []byte
	BTime_deliver  []byte
	BIp            []byte
	BCoupon_code   []byte
	BFeature       []byte
	BRemark        []byte
	BEvaluation    []byte
)

func main() {
	flag.Parse() // 解析flag
	conf, err := config.ParseConfigfile(*configFile)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(conf.Mysql.NullCloums[0])

	// 连接mongo
	session, err := mgo.Dial(conf.Mongo.URL)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(conf.Mongo.DB)
	collection := db.C(conf.Mongo.Collection)

	// 连接mysql
	dbm, err := sql.Open("mysql", conf.Mysql.URL)
	if err != nil {
		fmt.Println(err)
	}
	defer dbm.Close()
	rows, err := dbm.Query(conf.Mysql.SQL)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {

		if err := rows.Scan(&orderitem.Order_id, &orderitem.Status, &orderitem.Otype, &BPaytype, &BPaystatus, &BPay_trade_no, &BSource, &orderitem.Consumption_type, &orderitem.Suspicious, &orderitem.Sid, &orderitem.Site_id, &BDorm_id, &BDormentry_id, &orderitem.Shop_id, &orderitem.Uid, &BService_eva, &BDelivery_eva, &BFood_eva, &BFood_num, &orderitem.Food_amount, &BShip_fee, &orderitem.Coupon_discount, &orderitem.Promotion_discount, &orderitem.Discount, &BOrder_amount, &BDelivery_id, &orderitem.Add_time, &orderitem.Confirm_time, &orderitem.Send_time, &orderitem.Expect_date, &BDelivery_type, &orderitem.Expect_time, &orderitem.Expect_timeslot, &BOrder_mark, &BUname, &BPortrait, &orderitem.Phone, &BPhone_addr, &BBuy_times, &orderitem.Address1, &orderitem.Address2, &BDormitory, &BTime_deliver, &orderitem.Credit, &BIp, &BCoupon_code, &BFeature, &BRemark, &BEvaluation, &orderitem.Expect_start_time, &orderitem.Expect_end_time); err != nil {
			fmt.Println(err)
		}
		//
		//插入mongo
		// 可以在这里加个if判断，如 每100次的时候，对分片数组清空，然后执行一次insert，上面的scan的数据自然是append到数组里，这样应该会提升性能
		//
		//将[]byte 类型，转换为string 插入到mongo中
		orderitem.Paytype = string(BPaytype)
		orderitem.Paystatus = string(BPaystatus)
		orderitem.Pay_trade_no = string(BPay_trade_no)
		orderitem.Source = string(BSource)
		orderitem.Dorm_id = string(BDorm_id)
		orderitem.Dormentry_id = string(BDormentry_id)
		orderitem.Service_eva = string(BService_eva)
		orderitem.Delivery_eva = string(BDelivery_eva)
		orderitem.Food_eva = string(BFood_eva)
		orderitem.Food_num = string(BFood_num)
		orderitem.Ship_fee = string(BShip_fee)
		orderitem.Order_amount = string(BOrder_amount)
		orderitem.Delivery_id = string(BDelivery_id)
		orderitem.Delivery_type = string(BDelivery_type)
		orderitem.Order_mark = string(BOrder_mark)
		orderitem.Uname = string(BUname)
		orderitem.Portrait = string(BPortrait)
		orderitem.Phone_addr = string(BPhone_addr)
		orderitem.Buy_times = string(BBuy_times)
		orderitem.Dormitory = string(BDormitory)
		orderitem.Time_deliver = string(BTime_deliver)
		orderitem.Ip = string(BIp)
		orderitem.Coupon_code = string(BCoupon_code)
		orderitem.Feature = string(BFeature)
		orderitem.Remark = string(BRemark)
		orderitem.Evaluation = string(BEvaluation)

		err = collection.Insert(orderitem)
		if err != nil {
			fmt.Println(err)
		}
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

}
