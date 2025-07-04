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

//use case - 9
// ParkCarEvenly parks a car in the lot with the fewest cars for even distribution
func (a *ParkingAttendant) ParkCarEvenly(lots []*ParkingLot, car Car) bool {
	if len(lots) == 0 {
		return false
	}
	
	// Find the lot with the fewest parked cars
	var selectedLot *ParkingLot
	minCars := -1
	
	for _, lot := range lots {
		if lot.IsFull() {
			continue // Skip full lots
		}
		
		parkedCount := lot.GetParkedCarsCount()
		if minCars == -1 || parkedCount < minCars {
			minCars = parkedCount
			selectedLot = lot
		}
	}
	
	// If no lot is available, return false
	if selectedLot == nil {
		return false
	}
	
	// Park the car in the selected lot
	return selectedLot.Park(car)
}


// ParkHandicapCar parks a handicap car in the nearest available lot, for use case-10
func (a *ParkingAttendant) ParkHandicapCar(lots []*ParkingLot, car Car) bool {
    if len(lots) == 0 {
        return false
    }
    
    // Find the first available lot (nearest)
    for _, lot := range lots {
        if !lot.IsFull() {
            return lot.Park(car)
        }
    }
    
    // No available lot found
    return false
}


//use case-11
// ParkLargeCar parks a large car in the lot with the most available space
func (a *ParkingAttendant) ParkLargeCar(lots []*ParkingLot, car Car) bool {
    if len(lots) == 0 {
        return false
    }
    
    // Find the lot with the most available space
    var selectedLot *ParkingLot
    maxAvailableSpace := -1
    
    for _, lot := range lots {
        if lot.IsFull() {
            continue // Skip full lots
        }
        
        availableSpace := lot.GetAvailableSpaces()
        if availableSpace > maxAvailableSpace {
            maxAvailableSpace = availableSpace
            selectedLot = lot
        }
    }
    
    // If no lot is available, return false
    if selectedLot == nil {
        return false
    }
    
    // Park the car in the selected lot
    return selectedLot.Park(car)
}


