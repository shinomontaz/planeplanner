package types

import "time"

type Employee struct {
	Id    int
	Start time.Time
	End   time.Time
}

type Brigade struct {
	Name string
	Team []*Employee
}

type Job struct {
	Id      int
	Parent  *Job
	Brigade *Brigade
	Time    int
	Count   int // число рабочих для исполнения
}

type Plane struct {
	Name    string
	Jobs    []*Job
	Arrival time.Time
}
