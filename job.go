package cron

import (
	"log/slog"
	"time"

	"github.com/google/uuid"
)

// TODO: Add location/timezone.
type Job struct {
	id uuid.UUID

	// Name is used by the logger.
	Name string

	// JobFunc is the function that will be ran.
	JobFunc func() error

	// Interval is the time between the start of one run, and the start of the next run.
	// If you only want the job to run once, don't set Job.Interval (means you don't need to set Job.End either).
	Interval time.Duration

	// Start specifies the time that the job will run for the first time.
	Start time.Time

	// End specifies the time the scheduled job should terminate.
	End time.Time

	Logger *slog.Logger
}

// Schedule schedules Job.JobFunc() to be ran.
func (j *Job) Schedule() {
	uuid, _ := uuid.NewUUID()
	j.id = uuid

	j.Logger = j.Logger.With("id", j.id, "name", j.Name)

	// nextStartTime ensures that the job will start running "almost exactly" at the right time.
	nextStartTime := j.Start

	// TODO: This is a major hack and needs refactoring.
	if j.Interval == time.Second*0 {
		time.Sleep(time.Until(j.Start))

		if err := j.JobFunc(); err != nil {
			j.Logger.Error("job failed", "err", err)
		}

		return
	}

	time.Sleep(time.Until(j.Start))

	for {
		// Ensure job has not expired.
		if time.Now().After(j.End) {
			break
		}

		nextStartTime = nextStartTime.Add(j.Interval)

		if err := j.JobFunc(); err != nil {
			j.Logger.Error("job failed", "err", err)
		}

		time.Sleep(time.Until(nextStartTime))
	}
}
