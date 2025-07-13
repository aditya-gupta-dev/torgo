package torgo

import (
	"fmt"
	"time"
)

func ChangeIPRepeatedly(torPath string, interval, count int64) {
	if count < 1 {
		for {
			status, err := IsTorRunning(torPath)
			if err != nil {
				panic(err)
			}
			if !status {
				fmt.Println("tor stopped running")
				break
			}
			fmt.Println("Changed IP <''> ", count, interval)
			time.Sleep(time.Duration(interval * int64(time.Second)))
		}
	} else {
		for range count {
			status, err := IsTorRunning(torPath)
			if err != nil {
				panic(err)
			}
			if !status {
				fmt.Println("tor stopped running")
				break
			}
			fmt.Println("Changed IP <..> ", count, interval)
			time.Sleep(time.Duration(interval * int64(time.Second)))
		}
	}
}
