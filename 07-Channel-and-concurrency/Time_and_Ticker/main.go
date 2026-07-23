package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	defer ticker.Stop()

	var alertTimerChan <-chan time.Time
	alertTimerActive := false

	go func() {
		for {
			select {
			case <-ticker.C:
				cpuUsage := rand.Intn(100)
				fmt.Printf("CPU Usage : %d%%\n", cpuUsage)

				if cpuUsage > 80 {
					if !alertTimerActive {
						fmt.Println("High CPU usage detected sarting alert timmer")
						alertTimer := time.NewTimer(10 * time.Second)

						alertTimerChan = alertTimer.C
						alertTimerActive = true
					}
				} else {
					if alertTimerActive {
						fmt.Println("CPU usage normalized stopping alert timer")
						alertTimerActive = false
						alertTimerChan = nil
					}
				}
			case <-alertTimerChan:
				if alertTimerActive {
					fmt.Println("ALERT: High CPU usage sustained for 10s!")
					alertTimerActive = false
					alertTimerChan = nil
				}
			}

		}
	}()

	select {}
}
