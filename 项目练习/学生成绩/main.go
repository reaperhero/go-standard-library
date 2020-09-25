package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	scoreMap = map[string]int{}
)

func TakeExam(students []string) {
	scoreMap = make(map[string]int)
	for _, name := range students {
		score := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(101) //生成0-100的随机数
		scoreMap[name] = score
	}
	return
}

func SortStudentScore(scores map[string]int) (ranking [][3]string) {
	ranking = make([][3]string, 0)
	// 把学生放入ranking
	for k, v := range scoreMap {
		scoreStr := strconv.Itoa(v)
		ranking = append(ranking, [3]string{"0", k, scoreStr})
	}

	// ranking降序排序
	for i := 0; i < len(ranking)-1; i++ {
		for j := i + 1; j < len(ranking); j++ {
			jScoreInt, _ := strconv.Atoi(ranking[j][2])
			iScoreInt, _ := strconv.Atoi(ranking[i][2])
			if jScoreInt > iScoreInt {
				ranking[i], ranking[j] = ranking[j], ranking[i]
			}
		}
	}

	// 修改序号
	for key, _ := range ranking {
		ranking[key][0] = "第" + strconv.Itoa(key) + "名" // 这里不能直接用值数组修改，但是数组是值传递
	}
	return
}

func main() {
	students := []string{"stu1", "stu2", "stu3", "stu4", "stu5"}
	TakeExam(students)
	fmt.Println(scoreMap)
	ranking := SortStudentScore(scoreMap)
	fmt.Println(ranking)
}
