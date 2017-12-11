package main

import (
	"fmt"
	"time"
)
//用农作物替代币
//价格以1斤大米的价格为单位
const (
	DAMI = iota
	YUMI
	TUDOU
	DIGUA
	HUASHENG
	GAOLIANG
)

type Order struct {
	User string            //名称
	Price float64          //价格
	Count float64          //数量
	CTime int64            //上架时间，先使用秒为单位
}


func AddBuyer(goods *Order) {

	

}

func AddSeller(goods *Order) {

}
func main() {
	tuDou0 := Order{"土豆",2.3,0} 	//1斤土豆等于2.3斤大米
	//time.Now().UnixNano()
	tuDou0.CTime = time.Now().Unix()  //以后可能会使用纳妙
	time.Sleep(1)
	tuDou1 := Order{"土豆",2.32,0}
	tuDou1.CTime = time.Now().Unix()  
	time.Sleep(1)
	tuDou2 := Order{"土豆",2.41,0}
	tuDou2.CTime = time.Now().Unix()  

	yuMi0 := Order{"玉米",2.7,0} 
	yuMi0.CTime = time.Now().Unix()
	time.Sleep(1)
	yuMi1 := Order{"玉米",2.6,0} 
	yuMi1.CTime = time.Now().Unix()  
	time.Sleep(1)
	yuMi2 := Order{"玉米",2.68,0} 
	yuMi2.CTime = time.Now().Unix() 

	huaSheng0 := Order{"花生",4.6,0}
	huaSheng0.CTime = time.Now().Unix()
	time.Sleep(1)
	huaSheng1 := Order{"花生",4.65,0}
	huaSheng1.CTime = time.Now().Unix()
	time.Sleep(1)
	huaSheng2 := Order{"花生",4.72,0}
	huaSheng2.CTime = time.Now().Unix()

	diGua0 := Order{"地瓜",1.4,0}
	diGua0.CTime = time.Now().Unix()
	time.Sleep(1)
	diGua1 := Order{"地瓜",1.28,0}
	diGua1.CTime = time.Now().Unix()
	time.Sleep(1)
	diGua2 := Order{"地瓜",1.32,0}
	diGua2.CTime = time.Now().Unix()



}
