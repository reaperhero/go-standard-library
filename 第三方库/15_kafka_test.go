package third

import (
	"context"
	"github.com/segmentio/kafka-go"
	"encoding/json"
)

type kafkaRepositoryImpl struct {
	kafkaWriter *kafka.Writer
	kafkaReader *kafka.Reader
}

type kafkaRepository interface {
}

func NewkafkaRepository(kafkaAddr []string, topic string) kafkaRepository {

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  kafkaAddr, // []string{"192.168.1.1:9200","192.168.1.2:9200"}
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  kafkaAddr,
		GroupID:  "consumer_group_name",
		Topic:    topic,
		MinBytes: 1e3, // 1KB
		MaxBytes: 1e6, // 1MB
	})

	repo := kafkaRepositoryImpl{kafkaWriter: writer, kafkaReader: reader}

	return &repo
}

type user struct {
	Name string
}

func (d *kafkaRepositoryImpl) send() error {
	user := user{
		Name: "chen",
	}
	taskBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return d.kafkaWriter.WriteMessages(context.Background(), kafka.Message{Value: taskBytes})
}

func (d *kafkaRepositoryImpl) receive() () {
	m, err := d.kafkaReader.ReadMessage(context.Background())
	if err != nil {
		return
	}

	user := &user{}

	err = json.Unmarshal(m.Value, &user)

	return
}