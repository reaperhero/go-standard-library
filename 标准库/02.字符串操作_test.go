package stand

import (
	"fmt"
	"strings"
	"testing"
)

func Test_String01(t *testing.T) {
	s1 := "helloworld"
	//1.是否包含指定的内容-->bool
	fmt.Println(strings.Contains(s1, "abc"))
	//2.是否包含chars中任意的一个字符即可
	fmt.Println(strings.ContainsAny(s1, "abcd"))
	//3.统计substr在s中出现的次数
	fmt.Println(strings.Count(s1, "lloo"))
	//4.以xxx前缀开头，以xxx后缀结尾
	s2 := "20190525课堂笔记.txt"
	if strings.HasPrefix(s2, "201905") {
		fmt.Println("19年5月的文件。。")
	}
	if strings.HasSuffix(s2, ".txt") {
		fmt.Println("文本文档。。")
	}

	fmt.Println(strings.Index(s1, "lloo"))     //查找substr在s中的位置，如果不存在就返回-1
	fmt.Println(strings.IndexAny(s1, "abcdh")) //查找chars中任意的一个字符，出现在s中的位置
	fmt.Println(strings.LastIndex(s1, "l"))    //查找substr在s中最后一次出现的位置

	//字符串的拼接
	ss1 := []string{"abc", "world", "hello", "ruby"}
	s3 := strings.Join(ss1, "-")
	fmt.Println(s3)

	//切割
	s4 := "123,4563,aaa,49595,45"
	ss2 := strings.Split(s4, ",")
	//fmt.Println(ss2)
	for i := 0; i < len(ss2); i++ {
		fmt.Println(ss2[i])
	}

	//重复，自己拼接自己count次
	s5 := strings.Repeat("hello", 5)
	fmt.Println(s5)

	//替换
	//helloworld
	s6 := strings.Replace(s1, "l", "*", -1)
	fmt.Println(s6)
	//fmt.Println(strings.Repeat("hello",5))

	s7 := "heLLo WOrlD**123.."
	fmt.Println(strings.ToLower(s7))
	fmt.Println(strings.ToUpper(s7))

	fmt.Println(s1)
	s8 := s1[:5]
	fmt.Println(s8)
	fmt.Println(s1[5:])
}

// str := "Yinzhengjie"
// fmt.Println(strings.EqualFold(str,"YINZHENGJIE")) //忽略大小写，但是如果除了大小写的差异之外，还有其他的差异就会判定为false.
//fmt.Println(strings.HasPrefix(str,"Yinz"))         //判断字符串是否以Yinz开头
//fmt.Println(strings.HasSuffix(str,"jie"))       //判断字符串是否以“到此一游”结尾

// name := "yinzhengjie"
// str := "尹正杰到此一游"
// fmt.Println(strings.Index(str,"杰")) //注意，一个汉字战友三个字节，在“杰”前面有2个汉字，占据了0-5的索引，因此“杰”所对应的下班索引应该是“6”
// fmt.Println(strings.Index(name,"i"))  //找到第一个匹配到的“i”的索引下标。
//fmt.Println(strings.Index(name,"haha")) //如果没有找到的话就会返回“-1”
// list := strings.Split(name,"i")  //表示以字符串中的字母“i”为分隔符，将这个字符串进行分离,i被删除
// fmt.Println(strings.SplitAfter(name,"i"))  //SplitAfter这个方法表示在字符串中的字母“i”之后进行切割， 但是并不会覆盖到字母“i”,这一点跟Split方法是有所不同的哟！

// str := "#尹#正#杰#is#a#good#boy#"
// fmt.Println(strings.Trim(str,"#"))  //该方法可以去掉字符串左右两边的符号，但是字符串之间的是去不掉“#”的哟
// strings.TrimSpace(str))  //该方法可以脱去两边的空格和换行符。

func Test_string_03(t *testing.T) {
	str := `{"level":"info","msg":"\ufffd强军告警测\ufffd {\"sessionId\":\"haBhCtroXrRF7kQtJaOVYYeu\",\"alarmStatus\":\"1\",\"alarmType\":\"metric\",\"alarmObjInfo\":{\"region\":\"sh\",\"namespace\":\"qce/cvm\",\"dimensions\":{\"deviceName\":\"prod-myun-tx-mixer-tmp005\",\"objId\":\"5d62a04b-60f7-4587-aa6c-7c5efc630087\",\"objName\":\"10.200.1.2#3983277\",\"unInstanceId\":\"ins-0swvuraf\"}},\"alarmPolicyInfo\":{\"policyId\":\"policy-8dyufjzp\",\"policyType\":\"cvm_device\",\"policyName\":\"陈强军告警测试\",\"policyTypeCName\":\"云服务器-基础监控\",\"policyTypeEname\":\"\",\"conditions\":{\"metricName\":\"disk_usage\",\"metricShowName\":\"磁盘利用率 \",\"calcType\":\"\u003e\",\"calcValue\":\"10\",\"currentValue\":\"10.009\",\"unit\":\"%\",\"period\":\"300\",\"periodNum\":\"300\",\"alarmNotifyType\":\"singleAlarm\",\"alarmNotifyPeriod\":1}},\"firstOccurTime\":\"2020-10-21 13:10:00\",\"durationTime\":0,\"recoverTime\":\"0\"}\n","time":"2020-10-21T13:11:33+08:00"}`
	result := strings.SplitAfterN(str, "{", 10)
	str1 := strings.Join(result[2:], "")
	result = strings.SplitAfterN(str1, "}", 10)
	str2 := []string{"{"}
	for i := 0; i < len(result)-2; i++ {
		str2 = append(str2, result[i])
	}
	msg := strings.Replace(strings.Join(str2, ""), "\\", "", -1)
	fmt.Println(msg)
}

func Test_string_04(t *testing.T) {
	// 原生拼接字符串的方式会导致大量的string创建、销毁和内存分配，Builder为了解决性能问题，从1.10引入
	ss := []string{
		"A",
		"B",
		"C",
	}
	var b strings.Builder
	for _, s := range ss {
		fmt.Fprint(&b, s)
	}

	print(b.String())
}
