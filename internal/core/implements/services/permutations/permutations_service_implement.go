package permutations

import (
	services_interfaces "neversitup-test/internal/core/interfaces/services"
)

type permutationsServiceImplement struct {
}

func NewPermutationsServiceImplement() services_interfaces.PermutationsServiceInterface {
	return &permutationsServiceImplement{}
}

func (s *permutationsServiceImplement) checkExist(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

func (s *permutationsServiceImplement) Generate(input string) []string {
	if len(input) == 1 {
		return []string{input}
	}
	var permutations []string
	for i, ip := range input {
		remaining := input[:i] + input[i+1:]
		subPermutations := s.Generate(remaining)
		for _, sub := range subPermutations {
			newString := string(ip) + sub
			if newString == input {
				if s.checkExist(permutations, input) {
					continue
				}
			}
			if !s.checkExist(permutations, newString) {
				permutations = append(permutations, newString)
			}
		}
	}
	return permutations
}
