package services_interfaces

type CounterServiceInterface interface {
	CountSmileyInSlice(input []string) int
}
