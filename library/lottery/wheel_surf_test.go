package lottery

/*
	转盘抽奖测试
*/

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWheelSurfLottery(t *testing.T) {
	WheelSurfCfg = []*WheelSurf{{
		Id:      1,
		Rand:    1800,
		Rewards: []int{2, 1},
	}, {
		Id:      2,
		Rand:    1550,
		Rewards: []int{1, 50},
	}, {
		Id:      3,
		Rand:    1500,
		Rewards: []int{1, 200},
	}, {
		Id:      4,
		Rand:    250,
		Rewards: []int{1},
	}, {
		Id:      5,
		Rand:    1100,
		Rewards: []int{2, 3},
	}, {
		Id:      6,
		Rand:    1000,
		Rewards: []int{201, 1},
	}, {
		Id:      7,
		Rand:    1800,
		Rewards: []int{1, 150},
	}, {
		Id:      8,
		Rand:    1000,
		Rewards: []int{202, 1},
	}}
	// wg := &sync.WaitGroup{}
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		WheelSurfLottery(context.Background())
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	WheelSurfLottery(context.Background())
}

func BenchmarkWheelSurfLottery(b *testing.B) {
	WheelSurfCfg = []*WheelSurf{{
		Id:      1,
		Rand:    1800,
		Rewards: []int{2, 1},
	}, {
		Id:      2,
		Rand:    1550,
		Rewards: []int{1, 50},
	}, {
		Id:      3,
		Rand:    1500,
		Rewards: []int{1, 200},
	}, {
		Id:      4,
		Rand:    250,
		Rewards: []int{1, 1},
	}, {
		Id:      5,
		Rand:    1100,
		Rewards: []int{2, 3},
	}, {
		Id:      6,
		Rand:    1000,
		Rewards: []int{201, 1},
	}, {
		Id:      7,
		Rand:    1800,
		Rewards: []int{1, 150},
	}, {
		Id:      8,
		Rand:    1000,
		Rewards: []int{202, 1},
	}}
	ctx := context.Background()
	result := make(map[int]int)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		id := WheelSurfLottery(ctx)
		result[id] ++
	}
	// fmt.Println(result)
}

// map[1:1751 2:1601 3:1537 4:261 5:1146 6:1019 7:1681 8:1004]
// BenchmarkWheelSurfLottery-8   	  120140	      9932 ns/op	       0 B/op	       0 allocs/op
// BenchmarkWheelSurfLottery-8   	  123379	      9790 ns/op	       0 B/op	       0 allocs/op

// 随机数种子放在循环外：121048	      9984 ns/op	       0 B/op	       0 allocs/op
// 随机数种子放在循环内：24243	     49604 ns/op	       0 B/op	       0 allocs/op

func TestDailyIdea(t *testing.T) {
	err := fmt.Errorf("this is a error")
	fmt.Println(&err)
	a, err := func() (int, error) {
		return 0, nil
	}()
	fmt.Println(&err)
	fmt.Println(a, err)

	addr := "国家名称开头 接下来是省份信息 这是地级市 这个可能是小镇名称乡村名称以及街道具体名称 小区名称小区期号住宅单元楼层房间号等信息 房间号信息"
	fmt.Println("length addr: ", len(addr))

	nowDate, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	aimDate, _ := time.Parse("2006-01-02", time.Now().AddDate(0, 0, -1).Format("2006-01-02"))
	hours := nowDate.Sub(aimDate)

	fmt.Println(int(hours.Hours() / 24))
}
