package Admin

import (
	"ShoppingCarSystem/ShoppingCarSystem/Struct"
	"fmt"
)

// 增加新用户的函数
func Adduser(user Struct.Users) Struct.Users {
	var ac, pd string
	fmt.Println("请输入你注册的账户：")
	fmt.Scanln(&ac)
	fmt.Println("请输入你注册的密码：")
	fmt.Scanln(&pd)

	user = append(user, struct {
		Account  string "json:\"account\""
		Password string "json:\"pasword\""
	}{ac, pd})
	return user
}

// 删除用户的函数
func Deluser(user Struct.Users) Struct.Users {
	var ac string
	fmt.Println("请输入需要删除用户的账户")
	fmt.Scanln(&ac)
	for i, obj := range user {
		if obj.Account == ac {
			user = append(user[:i], user[i+1:]...) //使用截取法删除目标商品
			fmt.Printf("%s 用户已移除\n", ac)
			return user
		}
	}
	fmt.Println("该用户不存在")
	return user
}

// 修改用户数据
func Moduser(user Struct.Users) Struct.Users {
	var ac string
	fmt.Println("请输入需要修改用户密码的账户")
	fmt.Scanln(&ac)
	for i, obj := range user {
		if obj.Account == ac {
			var pd string
			fmt.Println("请输入需要修改后的用户的密码")
			fmt.Scanln(&pd)
			user[i].Password = pd //使用截取法删除目标商品
			fmt.Printf("%s 用户密码已修改\n", ac)
			return user
		}
	}
	fmt.Println("该用户不存在")
	return user
}

// 显示所有用户数据
func Showuser(user Struct.Users) Struct.Users {
	for i, obj := range user {
		fmt.Printf("用户%v -- 账户: %s -- 密码：%s \n", i+1, obj.Account, obj.Password)
	}
	return user
}
