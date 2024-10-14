package counter

import services_interfaces "neversitup-test/internal/core/interfaces/services"

type counterServiceImplement struct {
}

func NewCounterServiceImplement() services_interfaces.CounterServiceInterface {
	return &counterServiceImplement{}
}

func (s *counterServiceImplement) CountSmileyInSlice(input []string) int {
	var count int
	for _, ip := range input {
		if len(ip) == 2 && (ip[0] == ':' || ip[0] == ';') && (ip[1] == ')' || ip[1] == 'D') {
			count += 1
		} else if len(ip) == 3 && (ip[0] == ':' || ip[0] == ';') && (ip[1] == '-' || ip[1] == '~') && (ip[2] == ')' || ip[2] == 'D') {
			count += 1
		}
	}
	return count
}
