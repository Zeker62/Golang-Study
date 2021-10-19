package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

func dispatchCoin01() (left int) {
	var (
		coins = 50
		users = []string{
			"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
		}
		distribution = make(map[string]int, len(users))
	)

	for _, name := range users {
		getCoin := 0
		for _, n := range name {
			switch n {
			case 'e', 'E': getCoin += 1
			case 'i', 'I': getCoin += 2
			case 'o', 'O': getCoin += 3
			case 'u', 'U': getCoin += 4
			}
		}
		distribution[name] = getCoin
		coins -= getCoin
	}
	fmt.Println(distribution)
	return coins
}

func main() {
	left := dispatchCoin01()
	fmt.Println("剩下：", left)
}
//map[Aaron:3 Adriano:5 Augustus:12 Elizabeth:4 Emilie:6 Giana:2 Heidi:5 Matthew:1 Peter:2 Sarah:0]
//剩下： 10