package types

import "time"

type Employee struct {
	Id    int
	Start time.Time
	End   time.Time
	List  map[int]*Job // в какой квант времени какую работу сотрудник выполняет
}

type Brigade struct {
	Id   int
	Name string
	Team []*Employee
}

type Job struct {
	Id      int
	Parent  *Job
	Brigade *Brigade
	Time    int // количество 5-минутных квантов времени для выполнения работы
	Start   int // указатели на начало и окончание исполнения этой работы ( указыва.т на 5-минутный квант времени )
	End     int
	Count   int // число рабочих для исполнения - всегда считаем, что один нужен (пока что)
}

type Plane struct {
	Name     string
	Jobs     []*Job
	Arrival  time.Time
	Arrival2 int // номер 5-минутного кванта времени, когда самолет прилетел
}

const DAY_QUANTS = 24 * 60 / 5 // в одном дне вот столько 5-минутных квантов

// FindEmployee - в бригаде, что может выполнять эту работу найти первого свободного сотрудника, который это дело может делать
func (b *Brigade) FindEmployee(job *Job, acceptedStart *int) (*Employee, error) {

	// for _, e := range schedule.brigades[job.Brigade.Id].Team {
	// 	for quant := range e.List {
	// 		if quant > plane.Arrival2
	// 	}
	// }
	return nil, nil
}

func (e *Employee) AcceptJob(job *Job, acceptedStart int) error {
	for i := acceptedStart; i < job.Time; i++ {
		e.List[i] = job
	}
	return nil
}
