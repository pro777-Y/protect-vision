package main

// 需要使用windows弹框提醒
import (
	"time"
	"github.com/lxn/walk"
)

var ticker *time.Ticker

func main() {
	ticker = time.NewTicker(time.Second * 3)
	for {
		select {
		case <-ticker.C:
			walk.MsgBox(nil, "提示", "                     预防及控制近视发展,请遵循20 20 20 法则\n使用电子产品20分钟，向20英尺（约6米）以外远眺，并眨眼20次，以不产生视觉疲劳为宜.", walk.MsgBoxIconInformation)
			ticker = time.NewTicker(time.Minute * 20)
		}
	}
}
