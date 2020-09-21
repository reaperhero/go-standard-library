package command

import (
	"fmt"
	"time"
)

// defer timeCost()() 即可打印耗时
func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("time cost = %v\n", tc)
	}
}




