package types

import "time"

type Employee struct {
	Id    int
	Start time.Time
	End   time.Time
	Shift []int // список квантов времени, когда сотрудник доступен
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
	Count   int // число рабочих для исполнения
}

type Plane struct {
	Name     string
	Jobs     []*Job
	Arrival  time.Time
	Arrival2 int // номер 5-минутного кванта времени, когда самолет прилетел
}

type Schedule struct {
	List  [24 * 60 / 5]map[*Employee]*Job // ключи в этом массиве есть порядковый номер 5-минутного кванта времени
	List2 map[*Employee]map[int]*Job      // ключи в этом map[int] массиве есть порядковый номер 5-минутного кванта времени
}

type Day [24 * 60 / 5]struct{}
