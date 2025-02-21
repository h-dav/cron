/*
Package github.com/h-dav/cron is a job scheduling package.

Define a job and schedule it like so:

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

To make the job non-blocking, just put `job.Schedule()` in a Goroutine.
*/
package cron
