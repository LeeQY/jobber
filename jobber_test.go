package jobber

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func delayFunc(values []interface{}) bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(100)

	time.Sleep(100 * time.Microsecond)
	if num > 20 {
		return true
	} else {
		return false
	}
}

func TestJobber(t *testing.T) {
	job1 := New(delayFunc, 10)
	job2 := New(delayFunc, 20)
	thisCount := 1000
	for i := 0; i < thisCount; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		num := r.Intn(100)

		s := strconv.Itoa(i)

		time.Sleep(time.Duration(num) * time.Microsecond)

		job1.AddJob(&s)
	}

	for i := 0; i < thisCount; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		num := r.Intn(10)

		s := strconv.Itoa(i)

		time.Sleep(time.Duration(num) * time.Microsecond)

		job2.AddJob(&s)
	}

	time.Sleep(1 * time.Second)

	if thisCount != job1.count {
		t.Error("Error in handle job1 values. count: ", job1.count)
	}

	if thisCount != job2.count {
		t.Error("Error in handle job2 values. count: ", job2.count)
	}
}
