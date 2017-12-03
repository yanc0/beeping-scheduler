package scheduler

import (
	"time"
	"fmt"
)

// Scheduler is responsible of scheduling job when their
// NextRun time has come
type Scheduler struct {
	queue *Queue
	store *Store
	beepingUrls []string
}

// NewScheduler return a Scheduler
func NewScheduler() *Scheduler{
	queue := NewQueue()
	store := NewStore()
	beepingUrls := []string{"http://localhost:8080/check"}
	return &Scheduler{
		queue: queue,
		store: store,
		beepingUrls: beepingUrls,
	}
}

// Run the main scheluler loop
func (s *Scheduler) Run() {
	// Launch queue filler
    go func(){
		for {			
			s.FillQueue()
			time.Sleep(time.Second)
		}
	}()

	// Launch Job workers
	for {
		job := s.queue.Pop()
		if job != nil {			
			go job.Do(s.beepingUrls[0])
			s.store.Done(job.ID)
		}
		time.Sleep(time.Millisecond * 200)		
	}
}


// FillQueue gets all job expired and put them
// in queue
func (s *Scheduler) FillQueue() {
	jobs := s.store.ToRun()
	fmt.Println(len(jobs), "to fill")
	for _, j := range jobs {		
		s.queue.Add(j)
	}
}
