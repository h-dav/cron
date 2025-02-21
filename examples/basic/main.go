package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/h-dav/cron"
)

func main() {
	counter := 0

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	uuid, _ := uuid.NewUUID()

	job := cron.Job{
		ID:   uuid,
		Name: "Simple counter",
		JobFunc: func() error {
			counter = counter + 1

			fmt.Printf("Counter: %v\n", counter)
			fmt.Printf("Time: %s\n", time.Now())
			return nil
		},
		Interval: time.Second * 3,
		Start:    time.Now(),
		End:      time.Now().Add(time.Second * 15),
		Logger:   logger,
	}

	job.Schedule()
}
