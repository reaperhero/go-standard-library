package third

import (
	"fmt"
	"github.com/gansidui/skiplist"
	"testing"
)

type User struct {
	score float64
	id    string
}

func (u *User) Less(other interface{}) bool {
	if u.score > other.(*User).score {
		return true
	}
	if u.score == other.(*User).score && len(u.id) > len(other.(*User).id) {
		return true
	}
	return false
}

func Test_User_01(t *testing.T) {
	us := make([]*User, 7)
	us[0] = &User{6.6, "hi"}
	us[1] = &User{4.4, "hello"}
	us[2] = &User{2.2, "world"}
	us[3] = &User{3.3, "go"}
	us[4] = &User{1.1, "skip"}
	us[5] = &User{2.2, "list"}
	us[6] = &User{3.3, "lang"}

	sl := skiplist.New()
	for i := 0; i < len(us); i++ {
		sl.Insert(us[i])
	}

	// traverse
	for e := sl.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*User).id, "-->", e.Value.(*User).score)
	}
	// rank
	rank1 := sl.GetRank(&User{2.2, "list"})
	rank2 := sl.GetRank(&User{6.6, "hi"})
	fmt.Println(rank1) // 6
	fmt.Println(rank2) // 1

	// get value by rank
	e := sl.GetElementByRank(2)
	fmt.Println(e.Value) // {4.4 hello}
}
