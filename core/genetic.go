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
	jobs     []*types.Job
	planes   []*types.Plane
	brigades map[int]*types.Brigade
}

func algorithm(planes []*types.Plane, jobs []*types.Job, brigades []*types.Brigade) *Schedule {
	res := &Schedule{}

	return res
}

type ScheduleFactory struct {
	jobs     []*types.Job
	planes   []*types.Plane
	brigades map[int]*types.Brigade
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
	}

	for _, plane := range schedule.planes {
		for _, job := range plane.Jobs {
			acceptedStart := plane.Arrival2 // время, с которого задача может быть взята в работу, но может быть взята и позже

			empl, err := schedule.brigades[job.Brigade.Id].FindEmployee(job, &acceptedStart)
			if err != nil {
				panic("!")
			}

			if empl != nil {
				empl.AcceptJob(job, acceptedStart)
			}

		}
	}

	return schedule
}

func (s *Schedule) Clone() ga.Individual {
	return s
}

func (s *Schedule) Crossover(p ga.Individual) ga.Individual {
	// оставить на месте те работы, которые выполняют в одно и тоже время
	// (определить только кто именно их выполняет, но на самом деле это не важно ибо все сотрудники одинаковы)
	return s
}

func (s *Schedule) Educate() {

}

func (s *Schedule) Fitness() float64 {
	return 0.0
}

func (s *Schedule) Mutate() ga.Individual {
	clone := s.Clone().(*Schedule)
	for _, plane := range clone.planes {
		for _, job := range plane.Jobs {
			acceptedStart := plane.Arrival2 // время, с которого задача может быть взята в работу, но может быть взята и позже

			empl, err := clone.brigades[job.Brigade.Id].FindEmployee(job, &acceptedStart)

			if empl != nil {
				empl.AcceptJob(job, acceptedStart)
			}

		}
	}

	return s
}
