package domain

import (
	"time"
)

//UC-16
type CarParkingInfo struct {
    Car        Car
    Row        string
    SlotID     int
    IsHandicap bool
}

type ParkingLot struct {
	capacity         int
	parkedCars       []Car
	ownerObserver    Owner
	securityObserver Security
	wasFull bool // to track previous full state
	parkingTimes     map[string]time.Time // Track when each car was parked for use case-8
	carParkingInfo   map[string]CarParkingInfo // Maps plate to parking info, UC-16
}

//constructor to create a new parking lot with required capacity
func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		capacity:   capacity,
		parkedCars: make([]Car, 0),
		wasFull: false,
		parkingTimes: make(map[string]time.Time), //added for use case -8
		carParkingInfo: make(map[string]CarParkingInfo),
	}
}


//to add the owner observer
func (p *ParkingLot) AddOwnerObserver(owner Owner) {
	p.ownerObserver = owner
}

// to add security observer
func (p *ParkingLot) AddSecurityObserver(security Security) {
	p.securityObserver = security
}




//function to park a car, if some car comes in a Parking lot for parking, it will first check the capacity of
// parking lot and then append the car in the parked car, and notes the car plate number along with the time at which it parked
//and return true if it parked
func (p *ParkingLot) Park(car Car) bool {
	if len(p.parkedCars) >= p.capacity {
		return false
	}

	p.parkedCars = append(p.parkedCars, car)

	// Record parking time for use case -8
    p.parkingTimes[car.Plate] = time.Now()

	// Notify owner if lot is now full
    if len(p.parkedCars) == p.capacity {
        if p.ownerObserver != nil {
            p.ownerObserver.OnLotFull("Lot is full")
        }
        if p.securityObserver != nil {
            p.securityObserver.OnLotFull("Lot is full")
        }
		p.wasFull = true
    }

	return true
}

//to unpark the car from the lot
func (p *ParkingLot) Unpark(car Car) bool {
	for i, parkedCar := range p.parkedCars {
		if parkedCar.Plate == car.Plate {
			p.parkedCars = append(p.parkedCars[:i], p.parkedCars[i+1:]...)
            
			// Remove parking time record for use case-8
            delete(p.parkingTimes, car.Plate)

			//Notify owner if lot has space available
			if p.wasFull && len(p.parkedCars) == p.capacity-1 {
				if p.ownerObserver != nil {
					p.ownerObserver.OnSpaceAvailable("Space is Available")
				}
				p.wasFull = false
			}

			return true
		}
	}
	return false
}

//to get the number of currently parked cars
func (p *ParkingLot) GetParkedCarsCount() int {
	return len(p.parkedCars)
}

//to check whether the parking lot is full or not
func (p *ParkingLot) IsFull() bool {
	return len(p.parkedCars) == p.capacity
}

// changed function name for use case-11
// to get the space available in the lot
func(p *ParkingLot) GetAvailableSpaces() int {
	return p.capacity - len(p.parkedCars)
}

// GetParkingTime returns when a car was parked, use case -8
func (p *ParkingLot) GetParkingTime(plateNumber string) time.Time {
    if parkTime, exists := p.parkingTimes[plateNumber]; exists {
        return parkTime
    }
    return time.Time{} // Zero time if not found
}

// GetParkingDuration returns how long a car has been parked, use-case 8
func (p *ParkingLot) GetParkingDuration(plateNumber string) time.Duration {
    if parkTime, exists := p.parkingTimes[plateNumber]; exists {
        return time.Since(parkTime)
    }
    return 0 // Zero duration if not found
}


//use case-12
// FindCarsByColor returns all cars of a specific color
func (p *ParkingLot) FindCarsByColor(color string) []Car {
    var matchingCars []Car
    
    for _, parkedCar := range p.parkedCars {
        if parkedCar.Color == color {
            matchingCars = append(matchingCars, parkedCar)
        }
    }
    
    return matchingCars
}

// FindCarsByMakeAndColor returns all cars of a specific make and color
func (p *ParkingLot) FindCarsByMakeAndColor(make string, color string) []Car {
    var matchingCars []Car
    
    for _, parkedCar := range p.parkedCars {
        if parkedCar.Make == make && parkedCar.Color == color {
            matchingCars = append(matchingCars, parkedCar)
        }
    }
    
    return matchingCars
}

//use case - 14
// FindCarsByMake returns all cars of a specific make
func (p *ParkingLot) FindCarsByMake(make string) []Car {
    var matchingCars []Car
    
    for _, parkedCar := range p.parkedCars {
        if parkedCar.Make == make {
            matchingCars = append(matchingCars, parkedCar)
        }
    }
    
    return matchingCars
}

//use case-15
// FindCarsParkedInLastMinutes returns all cars parked within the specified number of minutes
func (p *ParkingLot) FindCarsParkedInLastMinutes(minutes int) []Car {
    var recentCars []Car
    cutoffTime := time.Now().Add(-time.Duration(minutes) * time.Minute)
    
    for _, parkedCar := range p.parkedCars {
        if parkTime, exists := p.parkingTimes[parkedCar.Plate]; exists {
            if parkTime.After(cutoffTime) {
                recentCars = append(recentCars, parkedCar)
            }
        }
    }
    
    return recentCars
}

// SetParkingTime sets the parking time for a car (used for testing)
func (p *ParkingLot) SetParkingTime(plateNumber string, parkTime time.Time) {
    p.parkingTimes[plateNumber] = parkTime
}

//UC-16
// ParkInRow parks a car in a specific row with handicap designation
func (p *ParkingLot) ParkInRow(car Car, row string, isHandicap bool) bool {
    if len(p.parkedCars) >= p.capacity {
        return false
    }
    
    // Park the car normally
    if !p.Park(car) {
        return false
    }
    
    // Store additional parking information
    slotID := len(p.parkedCars) - 1 // Last added car's slot
    parkingInfo := CarParkingInfo{
        Car:        car,
        Row:        row,
        SlotID:     slotID,
        IsHandicap: isHandicap,
    }
    p.carParkingInfo[car.Plate] = parkingInfo
    
    return true
}

// FindSmallHandicapCarsInRows finds small handicap cars in specified rows
func (p *ParkingLot) FindSmallHandicapCarsInRows(targetRows []string) []CarParkingInfo {
    var matchingCars []CarParkingInfo
    
    // Create a map for quick row lookup
    rowMap := make(map[string]bool)
    for _, row := range targetRows {
        rowMap[row] = true
    }
    
    for _, parkingInfo := range p.carParkingInfo {
        // Check if car meets all criteria
        if parkingInfo.Car.Size == Small && 
           parkingInfo.IsHandicap && 
           rowMap[parkingInfo.Row] {
            matchingCars = append(matchingCars, parkingInfo)
        }
    }
    
    return matchingCars
}

//UC-17
// GetAllParkedCars returns all currently parked cars
func (p *ParkingLot) GetAllParkedCars() []Car {
    // Return a copy of the parked cars slice to prevent external modification
    allCars := make([]Car, len(p.parkedCars))
    copy(allCars, p.parkedCars)
    return allCars
}
