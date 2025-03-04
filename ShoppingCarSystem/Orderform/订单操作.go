package Orderform

import (
	"ShoppingCarSystem/ShoppingCarSystem/Struct"
	"fmt"
	"time"
)

// func main() {
// 	var shoppingcar Struct.ShoppingCars
// 	var orderform Struct.OrderForm

// 	shoppingcar = append(shoppingcar, struct {
// 		GoodsName  string
// 		GoodsPrice float32W
// 		GoodsNum   int
// 	}{"小米su7", 300.00, 2})
// 	shoppingcar = append(shoppingcar, struct {
// 		GoodsName  string
// 		GoodsPrice float32
// 		GoodsNum   int
// 	}{"小米su8", 400.00, 3})
// 	Sale.Showshoppingcar(shoppingcar)

// 	orderform = AddOrderform(shoppingcar, orderform)
// 	fmt.Println(orderform)
// 	ShowOrderform(orderform, shoppingcar)
// 	orderform = ModOrderform(orderform)
// 	ShowOrderform(orderform, shoppingcar)
// 	orderform = DelOrderform(orderform)
// 	fmt.Println(orderform)
// }

// 添加订单操作
func AddOrderform(shoppingcar Struct.ShoppingCars, orderform Struct.OrderForm) Struct.OrderForm {
	var name, address, tel string
	orderformid := "20240411101"
	now := time.Now()

	fmt.Println("请输入收货人姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入收货人地址：")
	fmt.Scanln(&address)
	fmt.Println("请输入收货人联系方式：")
	fmt.Scanln(&tel)
	var all float32
	for _, obj := range shoppingcar {
		all += obj.GoodsPrice * float32(obj.GoodsNum)
	}

	orderform = append(orderform, struct {
		ShoppingCar Struct.ShoppingCars "json:\"shoppingcar\""
		OrderFormId string              "json:\"id\""
		Name        string              "json:\"name\""
		Address     string              "json:\"address\""
		Tel         string              "json:\"tel\""
		All         float32             "json:\"all\""
		Time        time.Time           "json:\"time\""
	}{shoppingcar, orderformid, name, address, tel, all, now})
	return orderform
}

// 删除订单操作
func DelOrderform(orderform Struct.OrderForm) Struct.OrderForm {
	var id string
	fmt.Println("请输入你要删除订单的订单号：")
	fmt.Scanln(&id)

	for i, obj := range orderform {
		if obj.OrderFormId == id {
			var confirm int
			fmt.Printf("你确定删除订单号为%s的订单吗？\n", id)
			fmt.Println("确定请按1，取消请按0")
			fmt.Scanln(&confirm)
			if confirm == 1 {
				orderform = append(orderform[:i], orderform[i+1:]...) //使用截取法删除目标商品
				fmt.Printf("%s订单删除成功\n", obj.OrderFormId)
				return orderform
			} else {
				return orderform
			}
		}
	}
	fmt.Printf("找不到订单号为：%s的订单", id)
	return orderform
}

// 修改订单
func ModOrderform(orderform Struct.OrderForm) Struct.OrderForm {
	var id string
	fmt.Println("请输入你要修改订单的订单号：")
	fmt.Scanln(&id)

	for i, obj := range orderform {
		if obj.OrderFormId == id {
			var rename, readdress, retel string
			fmt.Println("请输入修改后的收货人：")
			fmt.Scanln(&rename)
			fmt.Println("请输入修改后的收货地址：")
			fmt.Scanln(&readdress)
			fmt.Println("请输入修改后的联系方式：")
			fmt.Scanln(&retel)
			orderform[i].Name = rename
			orderform[i].Address = readdress
			orderform[i].Tel = retel
			fmt.Printf("%s订单的基本信息修改成功", id)
			return orderform
		}
	}
	fmt.Printf("找不到订单号为：%s的订单", id)
	return orderform
}

// 查看订单操作   //先输出购买商品的信息，再输出订单的基本信息
func ShowOrderform(orderform Struct.OrderForm, shoppingcar Struct.ShoppingCars) {
	fmt.Println("当前订单的信息：")
	for _, obj := range orderform {
		for _, v := range shoppingcar {
			fmt.Printf("商品名：%s -- 价格: %v -- 购买数量：%v \n", v.GoodsName, v.GoodsPrice, v.GoodsNum)
		}

		fmt.Printf("订单号：%s -- 订单总额：%v元\n收货人：%s 收货地址：%s -- 联系电话：%s\n", obj.OrderFormId, obj.All, obj.Name, obj.Address, obj.Tel)
		fmt.Printf("订单生成时间：%s\n", obj.Time.Format("2006-01-02 15:04:05"))
	}
}

/****************这是模拟结算部分**********************/
func ShoppingPay() {
	fmt.Println("请选择你的支付方式：")
	fmt.Println("1.支付宝 2.微信 3.银行卡")
	var choice string
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		mainFunc(payAli)
	case "2":
		mainFunc(payWechat)
	case "3":
		mainFunc(payBankcard)
	default:
		fmt.Println("请正确选择支付的方式")
	}
}

func payAli() {
	//连接支付宝支付api，展示支付页面
	fmt.Println("欢迎使用支付宝支付")
	fmt.Println("支付宝支付页面\n")
}

func payWechat() {
	//连接支付宝支付api，展示支付页面
	fmt.Println("欢迎使用微信支付")
	fmt.Println("微信支付页面\n")
}

func payBankcard() {
	//连接支付宝支付api，展示支付页面
	fmt.Println("欢迎使用银行卡支付")
	fmt.Println("银行卡支付页面\n")
}
func mainFunc(f func()) {
	f() //这是一个参数
	fmt.Println("恭喜支付成功")
	fmt.Println("我们将尽快安排发货\n")
}

func ShoppingRespond() {
	var choice int //保存选择变量
	fmt.Println("确认收货请按1,其他请按0:")
	fmt.Scanln(&choice)
	if choice == 1 {
		var respond string
		fmt.Println("请输入你对该商品的评价：")
		fmt.Scanln(&respond)
		fmt.Println("本次购物已完成")
	} else {
		fmt.Println("亲爱的小主请稍后，我们会尽快送达")
	}
}

/****************这是模拟结算部分**********************/
