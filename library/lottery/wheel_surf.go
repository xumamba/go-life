package lottery

import (
	"context"
	"math/rand"
	"time"
)

/*
	转盘式抽奖：
	抽奖前，用户已知全部奖品信息
	后端设置各个奖品的中奖概率和数量限制
	更新奖品库存的时候存在并发安全性问题
*/

var WheelSurfCfg []*WheelSurf

// WheelSurfCfg 抽奖奖励配置
type WheelSurf struct {
	Id      int   `json:"id"`      // 奖励配置唯一标识
	Rand    int   `json:"rand"`    // 中奖概率
	Rewards []int `json:"rewards"` // 奖励列表
}

func WheelSurfLottery(_ context.Context) int {
	allProbability := 10000
	// for _, ws := range WheelSurfCfg {
	// 	allProbability += ws.Rand
	// }
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		result := rand.Intn(allProbability)
		temp := 0
		// fmt.Println("=====")
		for _, ws := range WheelSurfCfg {
			temp += ws.Rand
			if result < temp {
				// fmt.Println("恭喜你，获得了：", ws.Rewards)
				break
				// return ws.Id
			}
		}
	}
	// rand.Seed(time.Now().UnixNano())
	// result := rand.Intn(allProbability)
	// temp := 0
	// for _, ws := range WheelSurfCfg {
	// 	temp += ws.Rand
	// 	if result < temp {
	// 		fmt.Println("恭喜你，获得了：", ws.Rewards)
	// 		return ws.Id
	// 	}
	// }
	return 0
}
