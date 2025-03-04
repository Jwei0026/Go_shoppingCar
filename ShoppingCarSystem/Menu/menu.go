package menu

import "fmt"

func Adminmenu() {
	fmt.Println("********欢迎进入管理员模块********")
	fmt.Println("1.添加用户信息")
	fmt.Println("2.删除用户信息")
	fmt.Println("3.修改用户信息")
	fmt.Println("4.查询用户信息")
	fmt.Println("请选择你需要的功能(按0退出该模块)：")
}

func Salemenu() {
	fmt.Println("********欢迎进入店主模块********")
	fmt.Println("1.添加商品")
	fmt.Println("2.删除商品")
	fmt.Println("3.修改商品")
	fmt.Println("4.查询商品")
	fmt.Println("5.查看订单信息")
	fmt.Println("6.删除订单信息")
	fmt.Println("请选择你需要的功能(按0退出该模块)：")
}

func Shoppingmenu() {
	fmt.Println("********欢迎进入购物模块********")
	fmt.Println("1.展示商品页面")
	fmt.Println("2.向购物车增加商品")
	fmt.Println("3.从购物车删除商品")
	fmt.Println("4.从购物车修改商品")
	fmt.Println("5.查看购物车")
	fmt.Println("6.生成订单信息")
	fmt.Println("7.查看订单信息")
	fmt.Println("8.修改订单信息")
	fmt.Println("9.结算订单")
	fmt.Println("请选择你需要的功能(按0退出该模块)：")
}

func Modelmenu() {
	fmt.Println("********欢迎进入购物车关系系统********")
	fmt.Println("1.进入管理员模块")
	fmt.Println("2.进入店主模块")
	fmt.Println("3.进入购物模块")
	fmt.Println("请选择你需要的功能(按0退出该系统)：")
}
