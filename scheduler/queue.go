package scheduler

import (
	"sync"
)

type Queue struct {
	mtx  sync.Mutex
	jobs []*Job
}

func NewQueue() *Queue {
	return &Queue{
		jobs: make([]*Job, 0),
	}
}

func (q *Queue) Add(j *Job) error {
	q.mtx.Lock()
	q.jobs = append(q.jobs, j)
	q.mtx.Unlock()

	return nil
}

func (q *Queue) Pop() *Job {
	var j *Job
	q.mtx.Lock()
	if len(q.jobs) == 0 {
		q.mtx.Unlock()
		return nil
	}
	j = q.jobs[0]
	q.jobs = append(q.jobs[:0], q.jobs[1:]...) // remove the first element
	q.mtx.Unlock()
	return j
}

func (q *Queue) Count() int {
	return len(q.jobs)
}
