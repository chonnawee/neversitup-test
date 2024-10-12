package odd

import services_interfaces "neversitup-test/internal/core/interfaces/services"

type oddServiceImplement struct {
}

func NewOddServiceImplement() services_interfaces.OddServiceInterface {
	return &oddServiceImplement{}
}

func (s *oddServiceImplement) groupInt(data []int) map[int]int {
	group := make(map[int]int)
	for _, d := range data {
		group[d] += 1
	}
	return group
}

func (s *oddServiceImplement) FindIntInSlice(input []int) int {
	group := s.groupInt(input)
	for _, g := range group {
		if g%2 != 0 {
			return g
		}
	}
	return 0
}
