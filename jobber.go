package jobber

import (
	"errors"
	"sync"
)

type Jobber struct {
	m          sync.RWMutex
	jobs       []interface{}
	processing bool
	function   HandleFunc
	max        int
	//	count      int
}

type HandleFunc func(batch []interface{}) bool

func New(f HandleFunc, maxNum int) *Jobber {
	return &Jobber{
		function: f,
		max:      maxNum,
	}
}

func (j *Jobber) AddJob(v interface{}) error {
	if v == nil {
		return errors.New("Value can't be nil.")
	}
	j.m.Lock()
	defer j.m.Unlock()

	j.jobs = append(j.jobs, v)
	if !j.processing {
		j.processing = true
		go j.start()
	}
	return nil
}

func (j *Jobber) getJobs() []interface{} {
	j.m.RLock()
	defer j.m.RUnlock()

	if len(j.jobs) > j.max {
		return j.jobs[:j.max]
	} else {
		return j.jobs
	}
}

func (j *Jobber) deleteJobs(l int) {
	j.m.Lock()
	defer j.m.Unlock()

	j.jobs = j.jobs[l:]
	//	j.count += l
}

func (j *Jobber) start() {
	jobs := j.getJobs()
	l := len(jobs)

	if l == 0 {
		j.m.Lock()
		j.processing = false
		j.m.Unlock()
		//		fmt.Println("Get out.")
		return
	}

	if j.function(jobs) {
		//		fmt.Println("Handle: ", l)
		j.deleteJobs(l)
	}

	j.start()
}
