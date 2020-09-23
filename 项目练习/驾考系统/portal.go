package main

import (
	"fmt"
	"go-standard-library/项目练习/驾考系统/utils"
	"time"
)

var(
	chNames = make(chan string, 100)
	examers = make([]string, 0)

	//信号量，只有5条车道
	chLanes = make(chan int, 5)
	//违纪者
	chFouls = make(chan string, 100)

	//考试成绩
	scoreMap = make(map[string]int)
)

/*巡考逻辑*/
func Patrol() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		//fmt.Println("战狼正在巡考...")
		select {
		case name := <-chFouls:
			fmt.Println(name, "考试违纪!!!!! ")
		default:
			fmt.Println("考场秩序良好")
		}
		<-ticker.C
	}
}

/*考试逻辑*/
func TakeExam(name string) {
	chLanes <- 123
	fmt.Println(name, "正在考试...")

	//记录参与考试的考生姓名
	examers = append(examers, name)

	//生成考试成绩
	score := utils.GetRandomInt(0, 100)
	//fmt.Println(score)
	scoreMap[name] = score
	if score < 10 {
		score = 0
		chFouls <- name
		//fmt.Println(name, "考试违纪！！！", score)
	}

	//考试持续5秒
	<-time.After(400 * time.Millisecond)

	<-chLanes
	//wg.Done()
}

/*二级缓存查询成绩*/
func QueryScore(name string) {
	score, err := utils.QueryScoreFromRedis(name)
	if err != nil {
		fmt.Println(err)
		//score, _ = utils.QueryScoreFromMysql(name)

		scores := make([]utils.ExamScore, 0)
		argsMap := make(map[string]interface{})
		argsMap["name"] = name
		//argsMap["score"] = 50
		err = utils.QueryFromMysql("score", argsMap, &scores)
		utils.HandlerError(err,`utils.QueryFromMysql("score", argsMap, &scores)`)
		fmt.Println("Mysql成绩：", name, ":", scores[0].Score)

		/*将数据写入Redis*/
		utils.WriteScore2Redis(name, scores[0].Score)

	} else {
		fmt.Println("Redis成绩：", name, ":", score)
	}
	//wg.Done()
}
