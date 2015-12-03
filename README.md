# jobber
Used for batch processing.

##The situation
Similar jobs can be handled together. Like can be used for batch methods from AWS SQS.

User can as simple as just add job to a place. Then a worker will take out the jobs in that place to handle.
If there are two jobs, the worker will take out two. If there are three, the worker will take out three. But there may be a max number, like 10. If there are 20 jobs, the worker will take out 10 first, and start another loop.


##The features
* Simple usage. User will define a worker function, then just add jobs.
* Fast return. Because user just need to add job.
* Single worker in a seperate goroutine.
* Thread-safe.
* Cache jobs and retry the failed ones.

##Install & Update
```
go get -u github.com/LeeQY/jobber
```

##Usage

####Define the worker function:
```Go
func delayFunc(values []interface{}) bool {
	...
}
```

####Create a jobber
```Go
// pass in the function and max number.
jobber := New(delayFunc, 20)
```

####Just add job
```Go
jobber.AddJob(job)
```