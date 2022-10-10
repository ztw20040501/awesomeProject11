package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	randNun := rand.Intn(900) + 100
	random := make([]int, 3)
	random[0] = randNun / 100
	random[1] = randNun / 10 % 10
	random[2] = randNun % 10

	userNum := make([]int, 3)
	var num int
	var flag = 0
	for {
		for { //输入三位数字，如果有误，则继续
			fmt.Println("请输入一个三位数：")
			fmt.Scan(&num)
			if num >= 100 && num <= 999 {
				break
			}
			fmt.Println("输入有误，请重新输入：")
		}
		userNum[0] = num / 100
		userNum[1] = num / 10 % 10
		userNum[2] = num % 10
		for i := 0; i < 3; i++ { //循环比较输入数与随机数大小
			if userNum[i] > random[i] {
				fmt.Printf("您输入的第%d位数太大了\n", i+1)
			} else if userNum[i] < random[i] {
				fmt.Printf("您输入的第%d位数太小了\n", i+1)
			} else {
				fmt.Printf("恭喜你，第%d位数字相同\n", i+1)
				flag++
			}
		}
		if flag == 3 {
			fmt.Println("成功")
			break
		} else {
			flag = 0
		}
	}
}
