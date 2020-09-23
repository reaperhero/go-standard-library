package main


import (
	"fmt"
	"go-standard-library/项目练习/驾考系统/utils"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

/*主程序*/
func main() {
	for i := 0; i < 20; i++ {
		chNames <- utils.GetRandomName()
	}
	close(chNames)

	/*巡考*/
	go Patrol()

	/*考生并发考试*/
	for name := range chNames {
		wg.Add(1)
		go func(name string) {
			TakeExam(name)
			wg.Done()
		}(name)
	}

	wg.Wait()
	fmt.Println("考试完毕！")

	/*录入成绩*/
	wg.Add(1)
	go func() {
		utils.WriteScore2Mysql(scoreMap)
		wg.Done()
	}()
	//故意给一个时间间隔，确保WriteScore2DB先抢到数据库的读写锁
	<-time.After(1 * time.Second)

	/*考生查询成绩*/
	for _, name := range examers {
		wg.Add(1)
		go func(name string) {
			QueryScore(name)
			wg.Done()
		}(name)
	}
	<-time.After(1 * time.Second)
	for _, name := range examers {
		wg.Add(1)
		go func(name string) {
			QueryScore(name)
			wg.Done()
		}(name)
	}

	wg.Wait()
	fmt.Println("END")
}