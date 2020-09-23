package command

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// 订单号模版
func Test_time_fun_01(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	timeValue := strings.ReplaceAll(time.Now().Format("060102150405.000000"), ".", "")
	result := fmt.Sprintf("%s%04d", timeValue, rand.Intn(1000))
	fmt.Println(result)
}
