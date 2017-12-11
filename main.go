package main

import (
	//"fmt"
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
	User string            //用户名称
	Price float64          //可以接受的价格
	Count float64          //货物数量
	CTime int64            //上架时间，先使用秒为单位
}
var BuyerList =make([]Order,0,200)
var SellerList =make([]Order,0,200)

//将买家需要的货物信息插入买入链表中
func AddBuyer(goods *Order) {
	//time.Now().UnixNano()
	goods.CTime= time.Now().Unix()   // //以后可能会使用纳妙
	BuyerList = append(BuyerList,*goods)
}
//将卖家需要卖出的货物信息插入出售链表中
func AddSeller(goods *Order) {
	//time.Now().UnixNano()
	goods.CTime= time.Now().Unix()   // //以后可能会使用纳妙
	SellerList = append(SellerList,*goods)
}
func main() {
	tuDou0 := Order{"土豆",2.3,90,0} 	//1斤土豆等于2.3斤大米
	AddSeller(&tuDou0)
	time.Sleep(1)
	tuDou1 := Order{"土豆",2.32,73,0}
	AddSeller(&tuDou1)
	 
	time.Sleep(1)
	tuDou2 := Order{"土豆",2.41,102,0}
	AddSeller(&tuDou2)
 
	yuMi0 := Order{"玉米",2.7,45,0} 
	AddSeller(&yuMi0)
	time.Sleep(1)
	yuMi1 := Order{"玉米",2.6,32,0} 
	AddSeller(&yuMi1)
	time.Sleep(1)
	yuMi2 := Order{"玉米",2.68,65,0} 
	AddSeller(&yuMi2)

	huaSheng0 := Order{"花生",4.6,21,0}
	AddSeller(&huaSheng0)
	time.Sleep(1)
	huaSheng1 := Order{"花生",4.65,64,0}
	AddSeller(&huaSheng1)
	time.Sleep(1)
	huaSheng2 := Order{"花生",4.727,6,0}
	AddSeller(&huaSheng2)


	diGua0 := Order{"地瓜",1.4,54,0}
	AddSeller(&diGua0)
	time.Sleep(1)
	diGua1 := Order{"地瓜",1.28,87,0}
	AddSeller(&diGua1)
	time.Sleep(1)
	diGua2 := Order{"地瓜",1.32,29,0}
	AddSeller(&diGua2)

}
