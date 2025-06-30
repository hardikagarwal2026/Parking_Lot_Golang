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

//to unpark the car from the lot
func (p *ParkingLot) Unpark(car Car) bool {
    for i, parkedCar := range p.parkedCars {
        if parkedCar.Plate == car.Plate {
            p.parkedCars = append(p.parkedCars[:i], p.parkedCars[i+1:]...)
            return true
        }
    }
    return false
}


//to get the number of currently parked cars
func(p *ParkingLot) GetParkedCarsCount() int {
	return len(p.parkedCars)
}