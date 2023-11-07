package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
    broker := "localhost:9093"
	topic := "test-topic"
	readerConfig := kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
	}
	reader := kafka.NewReader(readerConfig)
    fmt.Println("reached")
    fmt.Println("Reached")

	go func() {
		for {
			w := kafka.NewWriter(kafka.WriterConfig{
				Brokers: []string{broker},
				Topic:   topic,
			})

            err := w.WriteMessages(context.Background(), kafka.Message{
				Value: []byte(time.Now().Format("2006-01-02 15:04:05")),
			})

            if err != nil {
                fmt.Println(err.Error())
            }

            fmt.Println("Written")
			time.Sleep(2 * time.Second)
		}
	}()

	stopChan := make(chan struct{})

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	reader = kafka.NewReader(readerConfig)

	go func() {
		for {
			select {
			case <-stopChan:
				reader.Close()
			default:
				m, err := reader.ReadMessage(context.Background())
				if err != nil {
					reader.Close()
					log.Fatal(err.Error())
				}
                fmt.Print("Read:")
				fmt.Println(string(m.Value))
			}
		}
	}()

    <-signalChan
    close(stopChan)

    fmt.Println("Exited gracefully")
}
