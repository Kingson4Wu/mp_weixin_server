package timingwheel_test

import (
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/timingwheel"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {

	//初始化一个tick是1s，wheelSize是32的时间轮：
	tw := timingwheel.NewTimingWheel(time.Second, 32)
	tw.Start()
	// 添加任务
	//通过AfterFunc方法添加一个15s的定时任务，如果到期了，那么执行传入的函数。
	tw.AfterFunc(time.Second*15, func() {
		fmt.Println("The timer fires")
	})
}
