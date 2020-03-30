package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
胡扯战斗
*/

type Player struct {
	playerhp      int
	playername    string
	playerwuqi    string
	playershangyi string
	playerkuzi    string
	playerxiezi   string
}
type Monster struct {
	monsterhp      int
	monstername    string
	monsterwuqi    string
	monstershangyi string
	monsterkuzi    string
	monsterxiezi   string
}

var shuzu [2]interface{}
var xianshou [1]Player
var houshou [1]Monster
var xianshou1 [1]Monster
var houshou1 [1]Player
var rand1 = rand.New(rand.NewSource(time.Now().UnixNano()))
var playerds, monsterds int
var fa1flag bool
var fa2flag bool
var gjbw = []string{"眼睛", "太阳穴", "耳", "下颏", "喉结", "面部", "颈部", "心寓", "腹部", "肋骨",
	"裆部", "关节", "脚背", "麻筋", "脊椎", "后脑勺", "鼻", "腋窝", "脚背"}
var yhbw = []string{"眼睛", "太阳穴", "耳", "裆部", "脊椎", "后脑勺", "心寓", "颈部"}

func main() {

	Player := Player{
		playerhp:      100,
		playerwuqi:    "猎枪",
		playername:    "黄飞鸿",
		playershangyi: "猎人上衣",
		playerkuzi:    "猎人裤",
		playerxiezi:   "猎人鞋",
	}

	fmt.Println(Player)
	Monster := Monster{
		monsterhp:      100,
		monstername:    "周义翔",
		monsterwuqi:    "猎枪",
		monstershangyi: "猎人上衣",
		monsterkuzi:    "猎人裤",
		monsterxiezi:   "猎人鞋",
	}
	fmt.Println(Monster)

	fmt.Println("战斗开始")

	// 创建随机数种子  (默认将1970-1-1 00:00:00当做随机数种子)
	// 0-10之间的随机整数。 (包含0，不包含10。相当于对10求余)
	fmt.Println("双方选手开始选边")
	playerds = rand1.Intn(100)
	monsterds = rand1.Intn(100)
	i, i2 := panduan(playerds, monsterds, Player, Monster)

	fmt.Println(i, i2)
	fmt.Println(gjbw)
	fmt.Println("双方选手上场")
	fmt.Println("-----------战斗开始-------------")
	if fa1flag {
		fight1()
	} else {
		fight2()
	}
}

func panduan(i1 int, i2 int, player Player, monster Monster) (re1 int, re2 int) {

	var playerds = rand1.Intn(100)
	var monsterds = rand1.Intn(100)
	fa1flag = false
	fa2flag = false
	switch true {
	case playerds > monsterds:
		fmt.Println("方案一")
		fa1flag = true
		fmt.Println("player在左边")
		fmt.Println("monster在右边")
		xianshou[0] = player
		houshou[0] = monster
		fmt.Println("xianshou[0]", xianshou[0])
		fmt.Println("houshou[0]", houshou[0])
		break
	case monsterds > playerds:
		fmt.Println("方案二")
		fa2flag = true
		fmt.Println("monster在左边")
		fmt.Println("player在右边")
		xianshou1[0] = monster
		houshou1[0] = player
		fmt.Println("xianshou[0]", xianshou[0])
		fmt.Println("houshou[0]", houshou[0])
		break
	case monsterds == playerds:
		fmt.Println("重投点数")
		playerds = rand1.Intn(100)
		monsterds = rand1.Intn(100)
		panduan(playerds, monsterds, player, monster)
		break
	default:
		fmt.Println("...")
	}

	return playerds, monsterds
}

func fight1() {
	var monsterhp = houshou[0].monsterhp
	var playerhp = xianshou[0].playerhp
	var playergjbw string
	var monstergjbw string
	var sh int
	var mi int
	var pi int
	fmt.Println("fight1")
	for {
		playergjbw = gjbw[rand1.Intn(len(gjbw))]
		monstergjbw = gjbw[rand1.Intn(len(gjbw))]
		fmt.Println(xianshou[0].playername + "瞄准" + playergjbw + "准备击发")
		fmt.Println(houshou[0].monstername + "瞄准" + monstergjbw + "准备击发")
		if isYh(playergjbw) {
			sh = rand1.Intn(10) * 2
		} else {
			sh = rand1.Intn(10)
		}

		fmt.Println(xianshou[0].playername+"击出 , 造成", sh, "点伤害")
		mi = monsterhp - sh
		monsterhp = mi
		fmt.Println(houshou[0].monstername, "剩余", monsterhp)
		if isYh(playergjbw) {
			sh = rand1.Intn(10) * 2
		} else {
			sh = rand1.Intn(10)
		}
		if isDead(mi) {
			fmt.Println(xianshou[0].playername, "死亡 ", houshou[0].monstername, "胜")
			break
		}
		fmt.Println(houshou[0].monstername+"击出 , 造成", sh, "点伤害")
		pi = playerhp - sh
		playerhp = pi
		fmt.Println(xianshou[0].playername, "剩余", playerhp)

		if isDead(pi) {
			fmt.Println(houshou[0].monstername, "死亡 ", xianshou[0].playername, "胜")
			break
		}
	}
}

func fight2() {
	var monsterhp = houshou1[0].playerhp
	var playerhp = xianshou1[0].monsterhp
	var playergjbw string
	var monstergjbw string
	var sh int
	var mi int
	var pi int
	fmt.Println("fight2")
	for {
		playergjbw = gjbw[rand1.Intn(len(gjbw))]
		monstergjbw = gjbw[rand1.Intn(len(gjbw))]
		fmt.Println(xianshou1[0].monstername + "瞄准" + playergjbw + "准备击发")
		fmt.Println(houshou1[0].playername + "瞄准" + monstergjbw + "准备击发")
		if isYh(playergjbw) {
			sh = rand1.Intn(10) * 2
		} else {
			sh = rand1.Intn(10)
		}
		fmt.Println(xianshou1[0].monstername+"击出 , 造成", sh, "点伤害")
		pi = playerhp - sh
		playerhp = pi
		fmt.Println(houshou1[0].playername, "剩余", playerhp)

		if isYh(monstergjbw) {
			sh = rand1.Intn(10) * 2
		} else {
			sh = rand1.Intn(10)
		}
		if isDead(pi) {
			fmt.Println(houshou1[0].playername, "死亡 ", xianshou1[0].monstername, "胜")
			break
		}
		fmt.Println(houshou1[0].playername+"击出 , 造成", sh, "点伤害")
		mi = monsterhp - sh
		if isYh(playergjbw) {
			sh = rand1.Intn(10) * 2
		} else {
			sh = rand1.Intn(10)
		}
		if isDead(mi) {
			fmt.Println(xianshou1[0].monstername, "死亡 ", houshou1[0].playername, "胜")
			break
		}
		monsterhp = mi
		fmt.Println(xianshou1[0].monstername, "剩余", mi)
	}
}

func isDead(i int) (n bool) {
	if i <= 0 {
		return true
	}
	return false
}

func isYh(s string) (n bool) {
	for i := 0; i < len(yhbw); i++ {
		if s == gjbw[i] {
			return true
		}
		return false
	}
	return true
}
