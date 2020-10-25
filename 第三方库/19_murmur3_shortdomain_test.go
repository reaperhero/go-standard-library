package third

/**
  CREATE TABLE `short_url_map` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `lurl` varchar(160) DEFAULT NULL COMMENT '长地址',
    `surl` varchar(10) DEFAULT NULL COMMENT '短地址',
    `gmt_create` int(11) DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

import (
	"bytes"
	"fmt"
	"github.com/spaolacci/murmur3"
	"math"
	"testing"
)

// characters used for conversion
const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// converts number to base62
func Encode(number int) string {
	if number == 0 {
		return string(alphabet[0])
	}

	chars := make([]byte, 0)

	length := len(alphabet)

	for number > 0 {
		result := number / length
		remainder := number % length
		chars = append(chars, alphabet[remainder])
		number = result
	}

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

// converts base62 token to int
func Decode(token string) int {
	number := 0
	idx := 0.0
	chars := []byte(alphabet)

	charsLength := float64(len(chars))
	tokenLength := float64(len(token))

	for _, c := range []byte(token) {
		power := tokenLength - (idx + 1)
		index := bytes.IndexByte(chars, c)
		number += index * int(math.Pow(charsLength, power))
		idx++
	}

	return number
}

func Test_short_01(t *testing.T) {
	incr := murmur3.Sum32([]byte("https://u.geekbang.org/subject/python/100038901?utm_source=wechat&utm_medium=pyq02282300&utm_term=wechatpyq02282300"))
	fmt.Println(incr)
	fmt.Println(Encode(int(incr)))
}
