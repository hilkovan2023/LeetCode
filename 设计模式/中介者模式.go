package main

import "fmt"

type train interface {
	arrive()
	depart()
	permitArrival()
}

type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}

type passengerTrain struct {
	mediator mediator
}

func (g *passengerTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("PassengerTrain: Arrical blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (g *passengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *passengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.arrive()
}

type freightTrain struct {
	mediator mediator
}

func (g *freightTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *freightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *freightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.arrive()
}

type stationManager struct {
	isPlatformFree bool
	trainQueue     []train
}

func newStationManager() *stationManager {
	return &stationManager{
		isPlatformFree: true,
	}
}

func (s *stationManager) canArrive(t train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *stationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}

func main() {
	stationManager := newStationManager()

	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}
	freightTrain := &freightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
