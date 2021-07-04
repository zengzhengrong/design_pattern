package mediator

import "fmt"

// 火车 接口
type train interface {
	arrive()
	depart()
	permitArrival()
}

// 中介接口
type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}

// 运人火车
type passengerTrain struct {
	mediator mediator
}

func (g *passengerTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("PassengerTrain: 等待前方车辆")
		return
	}
	fmt.Println("PassengerTrain: 已到达")
}

func (g *passengerTrain) depart() {
	fmt.Println("PassengerTrain: 已离开")
	g.mediator.notifyAboutDeparture()
}

func (g *passengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: 前方已驶出, 进入车站")
	g.arrive()
}

// 运货火车
type freightTrain struct {
	mediator mediator
}

func (g *freightTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("FreightTrain: 等待前方车辆")
		return
	}
	// logic
	fmt.Println("FreightTrain: 已到达")
}

func (g *freightTrain) depart() {
	fmt.Println("FreightTrain: 已离开")
	// logic
	g.mediator.notifyAboutDeparture()
}

func (g *freightTrain) permitArrival() {
	fmt.Println("FreightTrain: 前方已驶出, 进入车站")
	g.arrive()
}

// 车站管理员
type stationManager struct {
	isPlatformFree bool
	trainQueue     []train
}

func newStationManger() *stationManager {
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
