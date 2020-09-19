package structure

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// map 是同种类型元素的无序组，键可以是任何相等性操作符支持的类型， 如整数、浮点数、复数、字符串、指针、接口（只要其动态类型支持相等性判断）、结构以及数组。 切片不能用作映射键，因为它们的相等性还未定义
// map是引用类型。 若将map传入函数中，并更改了该映射的内容，则此修改对调用者同样可见
// map中的元素并不是一个变量，而是一个值。因此，我们不能对map的元素进行取址操作

func Test_map_01(t *testing.T) {
	nameAge := make(map[string]int)
	nameAge["bob"] = 18    //增
	nameAge["tom"] = 16    //增
	delete(nameAge, "bob") //删
	nameAge["tom"] = 19    //改
	v := nameAge["tom"]    //查
	fmt.Println("v=", v)
	v, ok := nameAge["tom"] //查，推荐用法
	if ok {
		fmt.Println("v=", v, "ok=", ok)
	}
	for k, v := range nameAge { //遍历
		fmt.Println(k, v)
	}
}

// 注意事项
// 当 map 的元素为结构体类型的值，函数无法直接修改结构体中的字段值，可以通过以下方式修改
// 一是map 的 value用 struct 的指针类型
// 二是使用临时变量，每次取出来后再设置回去

//（1）将map中的元素改为struct的指针。
type person struct {
	name   string
	age    byte
	isDead bool
}

func whoIsDead(people map[string]*person) {
	for name, _ := range people {
		if people[name].age < 50 {
			people[name].isDead = true
		}
	}
}

func Test_update_map_01(t *testing.T) {
	p1 := &person{name: "zzy", age: 100}
	p2 := &person{name: "dj", age: 99}
	p3 := &person{name: "px", age: 20}
	personMap := map[string]*person{
		p1.name: p1,
		p2.name: p2,
		p3.name: p3,
	}
	whoIsDead(personMap)

	for _, v := range personMap {
		if v.isDead {
			fmt.Printf("%s is dead\n", v.name)
		}
	}
}

//（2）使用临时变量覆盖原来的元素。

type person2 struct {
	name   string
	age    byte
	isDead bool
}

func whoIsDead2(people map[string]person2) {
	for name, _ := range people {
		if people[name].age < 50 {
			tmp := people[name]
			tmp.isDead = true
			people[name] = tmp
		}
	}
}

func Test_update_map_02(t *testing.T) {
	p1 := person2{name: "zzy", age: 100}
	p2 := person2{name: "dj", age: 99}
	p3 := person2{name: "px", age: 20}
	personMap := map[string]person2{
		p1.name: p1,
		p2.name: p2,
		p3.name: p3,
	}
	whoIsDead2(personMap)

	for _, v := range personMap {
		if v.isDead {
			fmt.Printf("%s is dead\n", v.name)
		}
	}
}

func Test_Lock_map_01(t *testing.T) {
	var m = make(map[int]int)
	var rwMutex sync.RWMutex

	go func(){
		rwMutex.Lock()
		for i := 0; i < 10000; i++ {
			m[i] = i
		}
		rwMutex.Unlock()
	}()

	//一个go程读map
	go func(){
		rwMutex.RLock()
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
		}
		rwMutex.RUnlock()
	}()
	time.Sleep(time.Second*20)
}
