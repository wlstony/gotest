package main

import (
	"fmt"
	"time"
)

func calDuration()  time.Duration{
	//now := time.Now()
	//year, month, day := now.Date()
	//zero, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%d-%02d-%02d 00:00:00", year, month, day), time.Local)
	//fmt.Println("err:", err)
	//nextZero := zero.Add(5 * time.Second)
	//diff := nextZero.Sub(now)
	//fmt.Println("diff:", diff)

	now := time.Now()

	return now.Add(5 * time.Second).Sub(now)
}

func main() {
	d := calDuration()
	timer := time.NewTimer(d)
	go func() {
		for {
			<- timer.C
			d = calDuration()
			timer.Reset(d)
			fmt.Println("到时间了", time.Now())
		}
	}()
	i := 0
	for {
		i++
		fmt.Println("i", i)
		time.Sleep(10*time.Second)
	}





}