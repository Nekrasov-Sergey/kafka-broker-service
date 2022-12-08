package main

import (
	"context"
	"kafka-broker-service/internal/app/kafka"
	"kafka-broker-service/internal/app/models"
	app2 "kafka-broker-service/internal/pkg/app"
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	reader := kafka.NewKafkaReader()
	writer := kafka.NewKafkaWriter()
	ctx := context.Background()

	responseChannel := make(chan models.ResponseMessage)

	g, ctx := errgroup.WithContext(ctx)

	ctx, cancelCtx := context.WithCancel(ctx)

	app := app2.NewApplication(ctx)

	g.Go(func() error {
		return reader.FetchMessage(ctx)
	})

	g.Go(func() error {
		return writer.WriteMessages(ctx, responseChannel)
	})

	g.Go(func() error {
		err := app.Run(responseChannel)
		if err != nil {
			cancelCtx()
		}
		return err
	})

	err := g.Wait()
	if err != nil {
		log.Fatalln(err)
	}
}
