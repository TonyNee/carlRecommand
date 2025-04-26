package main

import (
	utils "carlRecommend/Common"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	root := utils.BuildTree(nums)
	fmt.Println("层序化打印输出：")
	utils.PrintTreeLeveled(root)
	fmt.Println("结构化打印输出：")
	utils.PrintTreeStructured(root)
}
