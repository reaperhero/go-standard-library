package stand

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func Test_Os_01(t *testing.T) {
	out, err := exec.Command("date").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The date is %s\n", out)
}

func Test_Os_02(t *testing.T) {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func Test_Os_03(t *testing.T) {
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}
