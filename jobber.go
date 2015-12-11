package jobber

import "sync"

type Jobber struct {
	m          sync.RWMutex
	jobs       []interface{}
	processing bool
	function   HandleFunc
	max        int
	//	count      int
}

// Define the handle function.
//
// batch is the passed-in jobs
// return true is the jobs are successfully handled.
type HandleFunc func(batch []interface{}) bool

// Create a New jobber
//
// f is the handle function
// maxNum is the max number of jobs handled once.
func New(f HandleFunc, maxNum int) *Jobber {
	return &Jobber{
		function: f,
		max:      maxNum,
	}
}

// Add one job.
func (j *Jobber) AddJob(v interface{}) {
	if v == nil {
		return
	}
	j.m.Lock()
	defer j.m.Unlock()

	j.jobs = append(j.jobs, v)
	if !j.processing {
		j.processing = true
		go j.start()
	}
}

// Add jobs.
func (j *Jobber) AddJobs(v []interface{}) {
	if v == nil {
		return
	}
	j.m.Lock()
	defer j.m.Unlock()

	j.jobs = append(j.jobs, v...)
	if !j.processing {
		j.processing = true
		go j.start()
	}
}

// According the max number to get jobs.
func (j *Jobber) getJobs() []interface{} {
	j.m.RLock()
	defer j.m.RUnlock()

	if len(j.jobs) > j.max {
		return j.jobs[:j.max]
	} else {
		return j.jobs
	}
}

// Delete the jobs.
//
// The beginning l jobs will be deleted.
func (j *Jobber) deleteJobs(l int) {
	j.m.Lock()
	defer j.m.Unlock()

	j.jobs = j.jobs[l:]
	//	j.count += l
}

// Start the handle.
func (j *Jobber) start() {
	jobs := j.getJobs()
	l := len(jobs)

	// Check whether there are jobs. If none, quit the goroutine.
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

	// loop the handle method.
	j.start()
}
