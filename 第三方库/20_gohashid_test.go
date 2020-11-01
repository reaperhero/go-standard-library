package third

import (
	"fmt"
	hashids "github.com/speps/go-hashids"
	"testing"
)

const hashSalt = "password_salt"

// HashIDEncode ...
func HashIDEncode(number int64) (result string, err error) {
	hd := hashids.NewData()
	hd.Salt = hashSalt                                   // 加密盐值
	hd.MinLength = 8                                     // 结果长度
	hd.Alphabet = "qwertyuiopasdfghjklzxcvbnm1234567890" // hash结果中的保含值

	id, _ := hashids.NewWithData(hd)

	result, err = id.EncodeInt64([]int64{number})
	return
}

// HashIDDecode ...
func HashIDDecode(hash string) (number int64, err error) {
	hd := hashids.NewData()
	hd.Salt = hashSalt                                   // 加密盐值
	hd.MinLength = 8                                     // 结果长度
	hd.Alphabet = "qwertyuiopasdfghjklzxcvbnm1234567890" // hash结果中的保含值

	id, _ := hashids.NewWithData(hd)

	result, err := id.DecodeInt64WithError(hash)
	if err == nil {
		if len(result) != 0 {
			number = result[0]
		}
	}
	return
}

func Test_hashids(t *testing.T) {
	result, _ := HashIDEncode(10)
	num, _ := HashIDDecode(result)
	fmt.Println(result, num) // o9dev1py  10
}
