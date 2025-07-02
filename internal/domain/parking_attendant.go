package domain

// ParkingAttendant represents an employee who parks cars
type ParkingAttendant struct {
	name string // Name of the attendant
}

// NewParkingAttendant creates a new parking attendant
func NewParkingAttendant(name string) *ParkingAttendant {
	return &ParkingAttendant{
		name: name,
	}
}

// GetName returns the attendant's name
func (a *ParkingAttendant) GetName() string {
	return a.name
}

// ParkCar parks a car in the given parking lot
func (a *ParkingAttendant) ParkCar(lot *ParkingLot, car Car) bool {
	return lot.Park(car)
}

// UnparkCar removes a car from the given parking lot
func (a *ParkingAttendant) UnparkCar(lot *ParkingLot, car Car) bool {
	return lot.Unpark(car)
}


//use case-6
// FindCar returns the slot number where the car is parked, or -1 if not found
func (p *ParkingLot) FindCar(plateNumber string) int {
	for i, parkedCar := range p.parkedCars {
		if parkedCar.Plate == plateNumber {
			return i
		}
	}
	return -1
}

