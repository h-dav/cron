package cron

import (
	"log/slog"
	"time"

	"github.com/google/uuid"
)

// TODO: Add location/timezone.
type Job struct {
	ID       uuid.UUID
	Name     string
	JobFunc  func() error
	Interval time.Duration // Time between the start of one run, and the start of the next run.
	Start    time.Time     // If you only want the job to run once, only set start.
	End      time.Time
	Logger   *slog.Logger
}

// Schedule schedules Job.JobFunc() to be ran.
func (j *Job) Schedule() {
	// nextStartTime ensures that the job will start running "almost exactly" at the right time.
	nextStartTime := j.Start

	time.Sleep(time.Until(j.Start))

	for {
		// Ensure job has not expired.
		if time.Now().After(j.End) {
			break
		}

		nextStartTime = nextStartTime.Add(j.Interval)

		if err := j.JobFunc(); err != nil {
			j.Logger.Error("job failed", "err", err, "id", j.ID, "name", j.Name)
		}

		time.Sleep(time.Until(nextStartTime))
	}
}
