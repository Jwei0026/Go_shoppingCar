package Shopping

import (
	"ShoppingCarSystem/ShoppingCarSystem/Struct"
	"fmt"
)

// func main() {
// 	var good Struct.Goods
// 	var shoppingcar Struct.ShoppingCars

// 	good = append(good, struct {
// 		GoodsID    int
// 		GoodsName  string
// 		GoodsPrice float32
// 		GoodsInv   intW
// 	}{123, "小米su7", 1999.99, 100})
// 	shoppingcar = Addshoppingcar(good, shoppingcar)
// 	Showshoppingcar(shoppingcar)

// 	shoppingcar = Delshoppingcar(shoppingcar)
// 	Showshoppingcar(shoppingcar)

// 	shoppingcar = append(shoppingcar, struct {
// 		GoodsName  string
// 		GoodsPrice float32
// 		GoodsNum   int
// 	}{"小米su8", 1999.00, 20})
// 	shoppingcar = Modshoppingcar(shoppingcar)
// 	Showshoppingcar(shoppingcar)
// }

// 传入一个商品名字和购物车，商品的结构，返回一个购物车  //这里可以补充购买的数量与库存的比较
func Addshoppingcar(good Struct.Goods, shoppingcar Struct.ShoppingCars) (Struct.ShoppingCars, Struct.Goods) {
	var Goodname string
	fmt.Println("请输入你要添加的商品名字：") //读取目标信息
	fmt.Scanln(&Goodname)

	for _, obj := range good {
		if obj.GoodsName == Goodname {
			var num int
			fmt.Println("请输入你要添加的商品数量：")
			fmt.Scanln(&num)
			if obj.GoodsInv >= num {
				shoppingcar = append(shoppingcar, struct {
					GoodsName  string  "json:\"name\""
					GoodsPrice float32 "json:\"price\""
					GoodsNum   int     "json:\"price\""
				}{obj.GoodsName, obj.GoodsPrice, num})
				// obj.GoodsInv -= obj.GoodsID - num
				// fmt.Println(obj.GoodsInv) //存疑 数量减了，但是没传回去
				fmt.Printf("%s 已加入购物车！\n", obj.GoodsName)
				// fmt.Println(good)
				return shoppingcar, good
			} else {
				fmt.Println("你输入的商品数量大于库存，请重新操作")
				return shoppingcar, good
			}
		}
	}
	fmt.Println("商品不存在，请重新选择。")
	return shoppingcar, good
}

// 删除购物车某样商品
func Delshoppingcar(shoppingcar Struct.ShoppingCars) Struct.ShoppingCars {

	var goodname string
	fmt.Println("请输入你要删除的商品名字：") //读取目标信息
	fmt.Scanln(&goodname)

	for i, obj := range shoppingcar {
		if obj.GoodsName == goodname {
			shoppingcar = append(shoppingcar[:i], shoppingcar[i+1:]...) //使用截取法删除目标商品
			fmt.Printf("%s 已从购物车中移除！\n", obj.GoodsName)
			return shoppingcar
		}
	}
	fmt.Println("购物车中没有该商品。")
	return shoppingcar
}

// 修改购物车的商品数量
func Modshoppingcar(shoppingcar Struct.ShoppingCars) Struct.ShoppingCars {
	var Goodname string
	fmt.Println("请输入你要修改的商品名字：") //读取目标信息
	fmt.Scanln(&Goodname)
	var Goodnum int
	fmt.Println("请输入你要修改后的购买数量：") //读取目标信息
	fmt.Scanln(&Goodnum)

	for i, obj := range shoppingcar {
		if obj.GoodsName == Goodname {
			shoppingcar[i].GoodsNum = Goodnum
			fmt.Printf("%s 的购买数量已购物车中修改成功！\n", obj.GoodsName)
			return shoppingcar
		}
	}
	fmt.Println("购物车中没有该商品。")
	return shoppingcar
}

// 展示购物车信息
func Showshoppingcar(shoppingCar Struct.ShoppingCars) {
	fmt.Println("当前购物车信息：")
	for _, obj := range shoppingCar {
		fmt.Printf("商品名：%s -- 价格: %v -- 购买数量：%v \n", obj.GoodsName, obj.GoodsPrice, obj.GoodsNum)
	}
}
