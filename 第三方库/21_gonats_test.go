package third

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

var Nats *NatsService

type NatsService struct {
	conn *nats.Conn
}

func InitNats(natsAddr string) {
	Nats = new(NatsService)
	Nats.Init(natsAddr)
}

func (n *NatsService) Init(natsAddr string) {
	var err error
	n.conn, err = nats.Connect(natsAddr)
	if err != nil {
		logrus.Fatalln("Could not connect to nats, err  ", err)
	}
}

func (n *NatsService) Publish(topic string, data string) error {
	return n.conn.Publish(topic, []byte(data))
}


func GetNatsConnection(url string, options ...nats.Option) (conn *nats.Conn, err error) {
	opts := nats.GetDefaultOptions()

	// 断开自动重连
	opts.AllowReconnect = true
	// 重连间隔
	opts.ReconnectWait = time.Second
	// 尝试重连的次数，负数代表一直尝试重连
	opts.MaxReconnect = -1
	// 连接关闭，无法重连
	opts.ClosedCB = func(nc *nats.Conn) {
		logrus.WithError(nc.LastError()).Errorln("nats connection closed")
	}
	// 重新连接成功
	opts.ReconnectedCB = func(nc *nats.Conn) {
		logrus.WithField("statistics", nc.Stats).Infoln("nats reconnect success")
	}
	// 连接断开
	opts.DisconnectedCB = func(nc *nats.Conn) {
		logrus.WithError(nc.LastError()).Errorln("nats connection lost")
	}

	opts.Servers = processUrlString(url)
	for _, opt := range options {
		if opt != nil {
			if err := opt(&opts); err != nil {
				return nil, err
			}
		}
	}

	conn, err = opts.Connect()
	if err != nil {
		logrus.WithError(err).Errorln("nats connect failed")
	}

	return
}

func processUrlString(url string) []string {
	urls := strings.Split(url, ",")
	for i, s := range urls {
		urls[i] = strings.TrimSpace(s)
	}
	return urls
}
