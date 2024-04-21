package utils

import (
	"fmt"
	"time"
)

func Loader(delay time.Duration, done <-chan struct{}) {
	dots := []string{
		"⠷",
		"⠯",
		"⠟",
		"⠻",
		"⠽",
		"⠾",
	}
	for {
		select {
		case <-done:
			fmt.Print("\r")
			return
		default:
			for _, r := range dots {
				fmt.Printf("\rfetching %s ", r)
				time.Sleep(delay)
			}
		}
	}
}
