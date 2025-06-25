package task

import (
	"fmt"
	"math/rand"
	"time"
)

func NewTask() *Task {
	id := store.generateID()
	task := &Task{
		ID:        id,
		CreatedAt: time.Now(),
		Status:    StatusPending,
	}

	store.save(task)

	go run(task.ID)

	return task
}

func GetTask(id string) (*Task, error) {
	return store.get(id)
}

func DeleteTask(id string) error {
	return store.delete(id)
}

func run(id string) {
	start := time.Now()
	_ = store.update(id, func(t *Task) {
		t.StartedAt = &start
		t.Status = StatusRunning
	})

	delay := time.Duration(rand.Intn(30)) * time.Second // иммитируем задержку 
	time.Sleep(delay)

	end := time.Now()
	_ = store.update(id, func(t *Task) {
		t.FinishedAt = &end
		t.Status = StatusCompleted
		t.Duration = end.Sub(*t.StartedAt).String()
		t.Result = fmt.Sprintf("Task %s finished in %s", t.ID, t.Duration)
	})
}
