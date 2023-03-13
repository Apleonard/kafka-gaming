package kafkas

import (
	"fmt"
	"kafka-gaming/api/config"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

func Produce(config config.ServerConfig, key, value string) error {
	var logger = log.New(os.Stdout, "[Producer] ", log.LstdFlags)

	syncProducer, err := sarama.NewSyncProducer(config.KafkaBrokers, nil)
	if err != nil {
		logger.Fatalln("failed to create producer: ", err)
		return err
	}

	fmt.Println(key, value)

	partition, offset, err := syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: config.KafkaTopic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	})
	if err != nil {
		logger.Fatalln("failed to send message to", "kocak", err)
		return err
	}
	logger.Printf("wrote message at partition: %d, offset: %d", partition, offset)
	_ = syncProducer.Close()

	return nil
}

// func (k *Writer) WriteMessages(ctx context.Context, messages chan kafkago.Message, messageCommitChan chan kafkago.Message) error {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		case m := <-messages:
// 			err := k.Writer.WriteMessages(ctx, kafkago.Message{
// 				Value: m.Value,
// 			})
// 			if err != nil {
// 				return err
// 			}

// 			select {
// 			case <-ctx.Done():
// 			case messageCommitChan <- m:
// 			}
// 		}
// 	}
// }
