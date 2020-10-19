package third

import (
	"fmt"
	"testing"
	"time"
	"github.com/patrickmn/go-cache"
)

func Test_gocache_test(t *testing.T) {
	// 5分钟设置一次过期状态，10分钟清理一次
	c := cache.New(5*time.Minute, 10*time.Minute)

	c.Set("foo", "bar", cache.DefaultExpiration)  // cache.NoExpiration 也可以设置不过期，重新设置key会覆盖之前的

	foo2, found := c.Get("foo")
	if found {
		fmt.Println(foo2.(string))
	}

	type User_Struct struct {
		Name string
	}
	user := User_Struct{Name: "chenqiangjun"}
	c.Set("foo", &user, cache.DefaultExpiration)
	if x, found := c.Get("foo"); found {
		foo := x.(*User_Struct)
		fmt.Println(foo) // &{chenqiangjun}
	}
}