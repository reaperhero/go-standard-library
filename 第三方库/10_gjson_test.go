package third

import (
	"github.com/tidwall/gjson"
	"testing"
)

const jsonstr = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func Test_gjson_01(t *testing.T) {
	value := gjson.Get(jsonstr, "name.last")
	println(value.String())
}

var test = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`



func Test_gjson_02(t *testing.T) {
	//"age"                >> 37
	//"children"           >> ["Sara","Alex","Jack"]
	//"children.#"         >> 3
	//"children.1"         >> "Alex"
	//"child*.2"           >> "Jack"
	//"c?ildren.0"         >> "Sara"
	//"fav\.movie"         >> "Deer Hunter"
	//"friends.#.first"    >> ["Dale","Roger","Jane"]
	//"friends.1.last"     >> "Craig"
	//friends.#(last=="Murphy").first    >> "Dale"
	//friends.#(last=="Murphy")#.first   >> ["Dale","Jane"]
	//friends.#(age>45)#.last            >> ["Craig","Murphy"]
	//friends.#(first%"D*").last         >> "Murphy"
	//friends.#(first!%"D*").last        >> "Craig"
	//friends.#(nets.#(=="fb"))#.first   >> ["Dale","Roger"]
}

//方法
//result.Exists() bool
//result.Value() interface{}
//result.Int() int64
//result.Uint() uint64
//result.Float() float64
//result.String() string
//result.Bool() bool
//result.Time() time.Time
//result.Array() []gjson.Result
//result.Map() map[string]gjson.Result
//result.Get(path string) Result
//result.ForEach(iterator func(key, value Result) bool)
//result.Less(token Result, caseSensitive bool) bool