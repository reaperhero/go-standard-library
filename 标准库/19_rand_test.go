package stand

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"text/tabwriter"
)

func Test_Rand__01(t *testing.T)  {
	// 创造并设置生成器.
	// 通常应该使用非固定种子, 如time.Now().UnixNano().
	// 使用固定的种子挥在每次运行中产生相同的输出
	r := rand.New(rand.NewSource(99))

	//这里的tabwriter帮助我们生成对齐的输出。
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	show := func(name string, v1, v2, v3 interface{}) {
		fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", name, v1, v2, v3)
	}

	// Float32 和Float64 的值在 [0, 1)之中.
	show("Float32", r.Float32(), r.Float32(), r.Float32())  //Float32     0.2635776           0.6358173           0.6718283
	show("Float64", r.Float64(), r.Float64(), r.Float64())  //Float64     0.628605430454327   0.4504798828572669  0.9562755949377957

	// ExpFloat64值的平均值为1，但是呈指数衰减。
	show("ExpFloat64", r.ExpFloat64(), r.ExpFloat64(), r.ExpFloat64())   //ExpFloat64  0.3362240648200941  1.4256072328483647  0.24354758816173044

	// NormFloat64值的平均值为0，标准差为1。
	show("NormFloat64", r.NormFloat64(), r.NormFloat64(), r.NormFloat64())  //NormFloat64 0.17233959114940064 1.577014951434847   0.04259129641113857

	// Int31，Int63和Uint32生成给定宽度的值。
	// Int方法（未显示）类似于Int31或Int63
	//取决于'int'的大小。
	show("Int31", r.Int31(), r.Int31(), r.Int31())      //Int31       1501292890          1486668269          182840835
	show("Int63", r.Int63(), r.Int63(), r.Int63())     //Int63       3546343826724305832 5724354148158589552 5239846799706671610
	show("Uint32", r.Uint32(), r.Uint32(), r.Uint32())   //Uint32      2760229429          296659907           1922395059

	// Intn，Int31n和Int63n将其输出限制为<n。
	//他们比使用r.Int（）％n更谨慎。
	show("Intn(10)", r.Intn(10), r.Intn(10), r.Intn(10))          //Intn(10)    1                   2                   5
	show("Int31n(10)", r.Int31n(10), r.Int31n(10), r.Int31n(10))     //Int31n(10)  4                   7                   8
	show("Int63n(10)", r.Int63n(10), r.Int63n(10), r.Int63n(10)) 	//Int63n(10)  7                   6                   3


	// Perm生成数字[0，n]的随机排列。
	show("Perm", r.Perm(5), r.Perm(5), r.Perm(5)) //Perm        [1 4 2 3 0]         [4 2 1 3 0]         [1 2 4 0 3]

}