package third

import (
	"github.com/garyburd/redigo/redis"
	"testing"
	"time"
)

var pool *redis.Pool

func init() {
	host := "192.168.40.136"
	pool = &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
func Test_go_redis_01(t *testing.T) {
	pool.Get().Do("SET", "name","vlaue")
}
