package command

import (
	"errors"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

// log err
func Test_print_err(t *testing.T) {
	//dataPath := os.Getenv("DATA_PATH")
	//file, _ := os.OpenFile(dataPath+"main.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	//logrus.SetOutput(file)
	err := errors.New("err")
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano})
	logrus.WithError(err).Warnln("err msg") //{"error":"err","level":"warning","msg":"err msg","time":"2020-09-19T15:52:37.605883+08:00"}
	logrus.WithError(err).Infoln("err msg") //{"error":"err","level":"info","msg":"err msg","time":"2020-09-19T15:52:37.606083+08:00"}
	logrus.WithFields(logrus.Fields{
		"key1":     "value1",
		"key2":     "value2",
	}).Infoln("err msg") // {"key1":"value1","key2":"value2","level":"info","msg":"err msg","time":"2020-09-19T15:55:36.682854+08:00"}

	logrus.WithError(err).WithFields(logrus.Fields{
		"key1":     "value1",
		"key2":     "value2",
	}).Errorln("err msg") // {"error":"err","key1":"value1","key2":"value2","level":"error","msg":"err msg","time":"2020-09-19T15:56:41.166771+08:00"}
}



