package scheduler

import (
	"time"
	"sync"
	"fmt"
	"github.com/yanc0/beeping/httpcheck"
)

// Store keep all managed jobs
type Store struct {
	mtx sync.Mutex
	jobs map[string]*Job
}

// NewStore returns a new store
func NewStore() *Store{
	jobs := make(map[string]*Job, 1)
	bee := httpcheck.Check {
		URL: "https://www.skale-5.com",
	}
	jobs["j1"] = &Job{ID: "j1", Interval: 5 * time.Second, NextRun: time.Now(), Check: bee}
	jobs["j2"] = &Job{ID: "j2", Interval: 6 * time.Second, NextRun: time.Now(), Check: bee}
	jobs["j3"] = &Job{ID: "j3", Interval: 7 * time.Second, NextRun: time.Now(), Check: bee}
	jobs["j4"] = &Job{ID: "j4", Interval: 8 * time.Second, NextRun: time.Now(), Check: bee}
	jobs["j5"] = &Job{ID: "j5", Interval: 9 * time.Second, NextRun: time.Now(), Check: bee}
	jobs["j6"] = &Job{ID: "j6", Interval: 10 * time.Second, NextRun: time.Now(), Check: bee}
	jobs["j7"] = &Job{ID: "j7", Interval: 11 * time.Second, NextRun: time.Now(), Check: bee}
	jobs["j8"] = &Job{ID: "j8", Interval: 12 * time.Second, NextRun: time.Now(), Check: bee}
	jobs["j9"] = &Job{ID: "j9", Interval: 13 * time.Second, NextRun: time.Now(), Check: bee}
	jobs["j10"] = &Job{ID: "j10", Interval: 14 * time.Second, NextRun: time.Now(), Check: bee}
	jobs["j11"] = &Job{ID: "j11", Interval: 15 * time.Second, NextRun: time.Now(), Check: bee}
	
	return &Store{
		jobs: jobs,
	}	
}

// ToRun returns a slice of job to as quick as possible
func (s *Store) ToRun() []*Job {
	var toRun []*Job
	s.mtx.Lock()
	for _, j := range s.jobs {
		if j.NextRun.Before(time.Now()) {
			toRun = append(toRun,j)
		}
	}
	s.mtx.Unlock()
	return toRun
}

// Done is called after a job is done
func (s *Store) Done (jobID string) {
	s.jobs[jobID].GenNextRun()
	fmt.Println("job", jobID, "is done, next run:", s.jobs[jobID].NextRun)
}