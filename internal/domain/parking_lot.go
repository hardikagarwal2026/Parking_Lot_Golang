package domain

import ()

type ParkingLot struct {
	capacity int
	parkedCars []Car
}

//constructor to create a new parking lot with required capacity
func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		capacity: capacity,
		parkedCars: make([]Car, 0),
	}
}

//function to park a car, if some car comes in a Parking lot for parking, it will first check the capacity of 
// parking lot and then append the car in the parked car, and notes the car plate number along with the time at which it parked
//and return true if it parked
func (p *ParkingLot) Park(car Car) bool {
	if len(p.parkedCars) >= p.capacity{
		return false
	}

	p.parkedCars = append(p.parkedCars, car)
	return true
}
