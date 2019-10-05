package main

import (
	"planeplanner/types"
	"sort"

	"github.com/shinomontaz/ga"
)

// type Schedule struct {
// 	jobs     []*types.Job
// 	planes   []*types.Plane
// 	brigades []*types.Brigade
// 	List     map[*types.Plane]map[*types.Job]*types.Employee //
// }

type Schedule struct {
	List  [24 * 60 / 5]map[*types.Employee]*types.Job // ключи в этом массиве есть порядковый номер 5-минутного кванта времени
	List2 map[*types.Employee]map[int]*types.Job      // ключи в этом map[int] массиве есть порядковый номер 5-минутного кванта времени
	jobs2 map[*types.Job]int                          // значение тут есть порядковый номер 5-минутного кванта

	jobs     []*types.Job
	planes   []*types.Plane
	brigades []*types.Brigade
}

func algorithm(planes []*types.Plane, jobs []*types.Job, brigades []*types.Brigade) *Schedule {
	res := &Schedule{}

	return res
}

type ScheduleFactory struct {
	jobs     []*types.Job
	planes   []*types.Plane
	brigades []*types.Brigade
}

func (sf *ScheduleFactory) Create() ga.Individual {
	// собираем случайную особь так, чтобы уложиться в лимиты в конфиге
	// отсортировать работы по порядку времени начала - время возможного начала определяется прилетом самолета
	// по сути надо отсортировать самолеты
	// затем берем последовательно работы данного самолета и назначаем на свободных сотрудников
	sort.Slice(sf.planes, func(i, j int) bool {
		//		return sf.planes[i].Arrival.Before(sf.planes[j].Arrival)
		return sf.planes[i].Arrival.Before(sf.planes[j].Arrival)
	})

	schedule := &Schedule{
		jobs:     sf.jobs,
		planes:   sf.planes,
		brigades: sf.brigades,
		// 	List  [24 * 60 / 5]map[*types.Employee]*types.Job // ключи в этом массиве есть порядковый номер 5-минутного кванта времени
		// List2 map[*types.Employee]map[int]*types.Job      // ключи в этом map[int] массиве есть порядковый номер 5-минутного кванта времени
		// jobs2 map[*types.Job]int  // значение тут есть порядковый номер 5-минутного кванта
	}

	for _, plane := range sf.planes {
		for _, job := range plane.Jobs {
			schedule.jobs2[job] = plane.Arrival2
		}
	}

	for j, quant := range schedule.jobs2 {

	}

	return schedule
}

func (s *Schedule) Clone() ga.Individual {
	return s
}

func (s *Schedule) Crossover(p ga.Individual) ga.Individual {
	return s
}

func (s *Schedule) Educate() {

}

func (s *Schedule) Fitness() float64 {
	return 0.0
}

func (s *Schedule) Mutate() ga.Individual {
	return s
}
