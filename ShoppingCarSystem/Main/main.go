package main

import (
	"ShoppingCarSystem/ShoppingCarSystem/Admin"
	"ShoppingCarSystem/ShoppingCarSystem/Fileoperate"
	menu "ShoppingCarSystem/ShoppingCarSystem/Menu"
	"ShoppingCarSystem/ShoppingCarSystem/Orderform"
	"ShoppingCarSystem/ShoppingCarSystem/Sale"
	"ShoppingCarSystem/ShoppingCarSystem/Shopping"
	"ShoppingCarSystem/ShoppingCarSystem/Struct"
	"fmt"
)

func main() {
	var users Struct.Users //设置四个全局变量，整个过程都是对该四个变量进行操作
	var shoppingCars Struct.ShoppingCars
	var orderforms Struct.OrderForm
	var goods Struct.Goods

	model_flag := true
	for model_flag {
		var model_choice int
		menu.Modelmenu()
		fmt.Scanln(&model_choice)
		switch model_choice {
		case 1:
			admin_flag := true
			// Fileoperate.Usersload(users, "Users")
			for admin_flag {
				var admin_choice int
				menu.Adminmenu()
				fmt.Scanln(&admin_choice)
				switch admin_choice {
				case 1:
					users = Admin.Adduser(users)
				case 2:
					users = Admin.Deluser(users)
				case 3:
					users = Admin.Moduser(users)
				case 4:
					Admin.Showuser(users)
				case 0:
					//保存文件信息
					Fileoperate.Userfilesave(users, "User")
					fmt.Println("成功退出该模块...")
					admin_flag = false
				default:
					fmt.Println("请输入正确的指令数字：")
				}
			}
		case 2:
			sale_flag := true
			// Fileoperate.Goodsload(goods, "Goods")
			for sale_flag {
				var sale_choice int
				menu.Salemenu()
				fmt.Scanln(&sale_choice)
				switch sale_choice {
				case 1:
					goods = Sale.Addgood(goods)
				case 2:
					goods = Sale.Delgood(goods)
				case 3:
					goods = Sale.Modgood(goods)
				case 4:
					Sale.Showgood(goods)
				case 5:
					Orderform.ShowOrderform(orderforms, shoppingCars)
				case 6:
					orderforms = Orderform.DelOrderform(orderforms)
				case 0:
					//保存文件信息
					Fileoperate.Goodfilesave(goods, "Goods")
					fmt.Println("成功退出该模块...")
					sale_flag = false
				default:
					fmt.Println("请输入正确的指令数字：")
				}
			}
		case 3:
			shopping_flag := true
			// Fileoperate.Goodsload(goods, "Goods")
			for shopping_flag {
				var shopping_choice int
				menu.Shoppingmenu()
				fmt.Scanln(&shopping_choice)
				switch shopping_choice {
				case 1:
					Sale.Showgood(goods)
				case 2:
					shoppingCars, goods = Shopping.Addshoppingcar(goods, shoppingCars)
				case 3:
					shoppingCars = Shopping.Delshoppingcar(shoppingCars)
				case 4:
					shoppingCars = Shopping.Modshoppingcar(shoppingCars)
				case 5:
					Shopping.Showshoppingcar(shoppingCars)
				case 6:
					orderforms = Orderform.AddOrderform(shoppingCars, orderforms)
				case 7:
					Orderform.ShowOrderform(orderforms, shoppingCars)
				case 8:
					orderforms = Orderform.ModOrderform(orderforms)
				case 9:
					Orderform.ShowOrderform(orderforms, shoppingCars)
					var sure int
					fmt.Println("确认订单信息请按1，取消请按0")
					fmt.Scanln(&sure)
					if sure == 1 {
						Orderform.ShoppingPay()
						Orderform.ShoppingRespond()
					}
				case 0:
					//保存文件信息
					Fileoperate.Shoppingcarfilesave(shoppingCars, "Shoppingcars")
					Fileoperate.Orderformsfilesave(orderforms, "Orderforms")
					fmt.Println("成功退出该模块...")
					shopping_flag = false
				default:
					fmt.Println("请输入正确的指令数字：")
				}
			}
		case 0:
			fmt.Println("成功退出该系统...")
			model_flag = false
		default:
			fmt.Println("请输入正确的指令数字：")
		}
	}

}
