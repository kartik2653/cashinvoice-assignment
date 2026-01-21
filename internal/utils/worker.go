package utils

import (
	"log"
	"sync"
	"time"

	"cashinvoice-assignment/internal/database"
	"cashinvoice-assignment/internal/model"
	"cashinvoice-assignment/internal/repository"

	"github.com/google/uuid"
)

type AutoCompleteWorker struct {
	TaskChan chan uuid.UUID
	Wg       sync.WaitGroup
	Delay    time.Duration
}

func NewAutoCompleteWorker(delayMinutes int) *AutoCompleteWorker {
	return &AutoCompleteWorker{
		TaskChan: make(chan uuid.UUID, 100),
		Delay:    time.Duration(delayMinutes) * time.Minute,
	}
}

func (w *AutoCompleteWorker) Start(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		w.Wg.Add(1)
		go w.worker()
	}
}

func (w *AutoCompleteWorker) worker() {
	defer w.Wg.Done()
	for taskID := range w.TaskChan {
		go w.processTask(taskID)
	}
}

func (w *AutoCompleteWorker) processTask(taskID uuid.UUID) {
	time.Sleep(w.Delay)

	todo, err := repository.NewTodoRepository(database.DB).GetByID(taskID)
	if err != nil {
		log.Println("task not found:", taskID)
		return
	}

	if todo.Status == model.Pending || todo.Status == model.InProgress {
		todo.Status = model.Completed
		err := database.DB.Save(todo).Error // persist change
		if err != nil {
			log.Println("failed to auto-complete task:", err)
		} else {
			log.Println("task auto-completed:", taskID)
		}
	}
}

func (w *AutoCompleteWorker) Stop() {
	close(w.TaskChan)
	w.Wg.Wait()
}
