package Struct

import "time"

//用切片构造商品结构体
type Goods []struct {
	GoodsID    int     `json:"id"`
	GoodsName  string  `json:"name"`
	GoodsPrice float32 `json:"price"`
	GoodsInv   int     `json:"inv"`
	// time       time.Time //操作时间
}

type ShoppingCars []struct {
	GoodsName  string  `json:"name"`
	GoodsPrice float32 `json:"price"`
	GoodsNum   int     `json:"price"` //购买的数量
}

type Users []struct {
	Account  string `json:"account"` //用户账户
	Password string `json:"pasword"`
}

type OrderForm []struct {
	ShoppingCar ShoppingCars `json:"shoppingcar"` //购买商品的信息
	OrderFormId string       `json:"id"`          //订单号
	Name        string       `json:"name"`
	Address     string       `json:"address"`
	Tel         string       `json:"tel"`
	All         float32      `json:"all"`
	Time        time.Time    `json:"time"` //订单生成时间
}
