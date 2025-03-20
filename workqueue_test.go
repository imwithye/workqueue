package workqueue

import (
	"fmt"
	"testing"
	"time"
)

type TestTask struct {
	Name      string
	RunCount  int
	SuccessOn int
	EnqueueAt time.Time
	SuccessAt *time.Time
}

func NewTestTask(name string, successOn int) *TestTask {
	return &TestTask{
		Name:      name,
		RunCount:  0,
		SuccessOn: successOn,
		EnqueueAt: time.Now(),
	}
}

func (t *TestTask) Run() error {
	t.RunCount++
	if t.RunCount == t.SuccessOn {
		successAt := time.Now()
		t.SuccessAt = &successAt
		return nil
	}
	return fmt.Errorf("task %s failed", t.Name)
}

func (t *TestTask) DurationForSuccess() *time.Duration {
	if t.SuccessAt == nil {
		return nil
	}
	duration := t.SuccessAt.Sub(t.EnqueueAt)
	return &duration
}

func TestWorkQueue(t *testing.T) {
	queue := NewWorkQueue()

	taskSuccessOnFirstRun := NewTestTask("successOnFirstRun", 1)
	queue.Enqueue(taskSuccessOnFirstRun)

	taskSuccessOnSecondRun := NewTestTask("successOnSecondRun", 2)
	queue.Enqueue(taskSuccessOnSecondRun)

	taskSuccessOnThirdRun := NewTestTask("successOnThirdRun", 3)
	queue.Enqueue(taskSuccessOnThirdRun)

	taskNeverSucceeds := NewTestTask("neverSucceeds", 0)
	queue.Enqueue(taskNeverSucceeds)

	queue.Run()

	if taskSuccessOnFirstRun.RunCount != 1 {
		t.Errorf("Expected taskSuccessOnFirstRun.RunCount to be 1, got %d", taskSuccessOnFirstRun.RunCount)
	}
	if duration := taskSuccessOnFirstRun.DurationForSuccess(); duration == nil || duration.Seconds() > 1 {
		t.Errorf("Expected taskSuccessOnFirstRun.DurationForSuccess to be less than 1 second")
	}
	if taskSuccessOnSecondRun.RunCount != 2 {
		t.Errorf("Expected taskSuccessOnSecondRun.RunCount to be 2, got %d", taskSuccessOnSecondRun.RunCount)
	}
	if duration := taskSuccessOnSecondRun.DurationForSuccess(); duration == nil || duration.Seconds() < 2 {
		t.Errorf("Expected taskSuccessOnSecondRun.DurationForSuccess to be at least 2 seconds")
	}
	if taskSuccessOnThirdRun.RunCount != 3 {
		t.Errorf("Expected taskSuccessOnThirdRun.RunCount to be 3, got %d", taskSuccessOnThirdRun.RunCount)
	}
	if duration := taskSuccessOnThirdRun.DurationForSuccess(); duration == nil || duration.Seconds() < 4 {
		t.Errorf("Expected taskSuccessOnThirdRun.DurationForSuccess to be at least 4 seconds")
	}
	if taskNeverSucceeds.RunCount != 4 {
		t.Errorf("Expected taskNeverSucceeds.RunCount to be 4, got %d", taskNeverSucceeds.RunCount)
	}
	if duration := taskNeverSucceeds.DurationForSuccess(); duration != nil {
		t.Errorf("Expected taskNeverSucceeds.DurationForSuccess to be nil (never succeeded)")
	}
}
