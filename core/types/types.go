package types

import "time"

type Employee struct {
	Id    int
	Shift []int // список квантов времени, когда сотрудник доступен
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

type Day struct {
	Quants [24 * 60 / 5]struct{} // 24 * (60/5) - пятиминутные кванты времени
}
