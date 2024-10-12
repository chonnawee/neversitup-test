package permutations

import (
	services_interfaces "neversitup-test/internal/core/interfaces/services"
)

type permutationsServiceImplement struct {
}

func NewPermutationsServiceImplement() services_interfaces.PermutationsServiceInterface {
	return &permutationsServiceImplement{}
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
			permutations = append(permutations, string(ip)+sub)
		}
	}
	return permutations
}
