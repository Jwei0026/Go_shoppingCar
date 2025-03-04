package Fileoperate

//本包可考虑优化成两个通用工具类
import (
	"ShoppingCarSystem/ShoppingCarSystem/Struct"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 用户数据保存
func Userfilesave(users Struct.Users, filename string) {
	jsonusers, userserr := json.Marshal(users)

	if userserr != nil {
		fmt.Println("users数据转json失败", userserr)
	}

	err := os.WriteFile(filename, jsonusers, 0644)
	if err != nil {
		fmt.Println("user文件写入失败", err)
	}
}

// 商品数据保存
func Goodfilesave(goods Struct.Goods, filename string) {
	jsongoods, goodserr := json.Marshal(goods)

	if goodserr != nil {
		fmt.Println("goods数据转json失败", goodserr)
	}

	err := os.WriteFile(filename, jsongoods, 0644)
	if err != nil {
		fmt.Println("goods文件写入失败", err)
	}
}

// 购物车数据保存
func Shoppingcarfilesave(shoppingcars Struct.ShoppingCars, filename string) {
	jsonshoppingcars, shoppingcarserr := json.Marshal(shoppingcars)

	if shoppingcarserr != nil {
		fmt.Println("shoppingcars数据转json失败", shoppingcarserr)
	}

	err := os.WriteFile(filename, jsonshoppingcars, 0644)
	if err != nil {
		fmt.Println("shoppingcars文件写入失败", err)
	}
}

// 订单数据保存
func Orderformsfilesave(orderforms Struct.OrderForm, filename string) {
	jsonorderforms, orderformserr := json.Marshal(orderforms)

	if orderformserr != nil {
		fmt.Println("orderforms数据转json失败", orderformserr)
	}

	err := os.WriteFile(filename, jsonorderforms, 0644)
	if err != nil {
		fmt.Println("orderforms文件写入失败", err)
	}
}

// 用户数据加载
func Usersload(users Struct.Users, filename string) {
	// fileinfo, err := os.Stat(filename)
	// fmt.Println(filename)
	// if os.IsExist(err) {
	// fmt.Println(fileinfo.Size())
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("文件读取错误", err)
	}

	if err := json.Unmarshal(data, &users); err != nil { //解码json数据到目标切片
		fmt.Println("解码失败", err)
	}
	// } else {
	// 	fmt.Println("不存在该文件")
	// }
}

// 商品数据加载
func Goodsload(goods Struct.Goods, filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("文件读取错误", err)
	}

	if err := json.Unmarshal(data, &goods); err != nil { //解码json数据到目标切片
		fmt.Println("解码失败", err)
	}
}

//购物车数据加载
//订单数据加载
