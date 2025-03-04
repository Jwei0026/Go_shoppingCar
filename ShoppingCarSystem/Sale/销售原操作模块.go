package Sale

import (
	"ShoppingCarSystem/ShoppingCarSystem/Struct"
	"fmt"
)

// func main() {
// 	var good Struct.Goods

// 	good = Addgood(good)
// 	good = Addgood(good)
// 	Showgood(good)
// 	good = Modgood(good)
// 	Showgood(good)
// 	good = Delgood(good)
// 	Showgood(good)
// }

// 增加新商品的函数
func Addgood(good Struct.Goods) Struct.Goods {
	var id, inv int
	var name string
	var price float32

	fmt.Println("请输入你要增加商品的ID:")
	fmt.Scanln(&id)
	fmt.Println("请输入你要增加商品的名称：")
	fmt.Scanln(&name)
	fmt.Println("请输入你要增加商品的价格：")
	fmt.Scanln(&price)
	fmt.Println("请输入你要增加商品的库存：")
	fmt.Scanln(&inv)

	good = append(good, struct {
		GoodsID    int     "json:\"id\""
		GoodsName  string  "json:\"name\""
		GoodsPrice float32 "json:\"price\""
		GoodsInv   int     "json:\"inv\""
	}{id, name, price, inv})
	return good
}

// 删除商品的函数
func Delgood(good Struct.Goods) Struct.Goods {
	var name string
	fmt.Println("请输入需要删除商品的名称：")
	fmt.Scanln(&name)
	for i, obj := range good {
		if obj.GoodsName == name {
			good = append(good[:i], good[i+1]) //这个删除方法不能删除最后一个元素，有待优化
			fmt.Println(name + "商品删除成功")
			return good
		}
	}
	fmt.Println(name + "该商品不存在")
	return good
}

// 修改商品价格和库存数据
func Modgood(good Struct.Goods) Struct.Goods {
	var name string
	fmt.Println("请输入需要修改商品的名称：")
	fmt.Scanln(&name)
	for i, obj := range good {
		if obj.GoodsName == name {
			var reprice float32
			var reinv int
			fmt.Println("请输入需要修改后的商品价格")
			fmt.Scanln(&reprice)
			fmt.Println("请输入需要修改后的商品库存")
			fmt.Scanln(&reinv)
			good[i].GoodsPrice = reprice //找到目标位置进行替换
			good[i].GoodsInv = reinv
			fmt.Printf("%s 商品信息已修改\n", name)
			return good
		}
	}
	fmt.Println("该商品不存在")
	return good
}

// 显示所有商品数据
func Showgood(good Struct.Goods) Struct.Goods {
	for _, obj := range good {
		fmt.Printf("商品编码：%v  商品名：%s -- 价格: %v -- 库存：%v \n", obj.GoodsID, obj.GoodsName, obj.GoodsPrice, obj.GoodsInv)
	}
	return good
}
