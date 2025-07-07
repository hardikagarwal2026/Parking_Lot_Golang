package domain

import "time"

// PoliceDepartment represents law enforcement investigation capabilities
type PoliceDepartment struct {
    departmentName string
}

// NewPoliceDepartment creates a new police department instance
func NewPoliceDepartment(name string) *PoliceDepartment {
    return &PoliceDepartment{
        departmentName: name,
    }
}

// InvestigateWhiteCars finds all white cars across multiple lots for bomb threat investigation
func (pd *PoliceDepartment) InvestigateWhiteCars(lots []*ParkingLot) []CarLocation {
    var allWhiteCars []CarLocation
    
    for i, lot := range lots {
        whiteCars := lot.FindCarsByColor("White")
        for _, car := range whiteCars {
            location := CarLocation{
                Car:    car,
                LotID:  i,
                SlotID: lot.FindCar(car.Plate),
            }
            allWhiteCars = append(allWhiteCars, location)
        }
    }
    
    return allWhiteCars
}

// CarLocation represents a car's location information for police investigations
type CarLocation struct {
    Car    Car
    LotID  int
    SlotID int
}

//for use case-13
// InvestigateBlueToyotas finds all blue Toyota cars with complete investigation details
func (pd *PoliceDepartment) InvestigateBlueToyotas(lots []*ParkingLot, attendant *ParkingAttendant) []RobberyInvestigation {
    var allBlueToyotas []RobberyInvestigation
    
    for i, lot := range lots {
        blueToyotas := lot.FindCarsByMakeAndColor("Toyota", "Blue")
        for _, car := range blueToyotas {
            investigation := RobberyInvestigation{
                Car:           car,
                LotID:         i,
                SlotID:        lot.FindCar(car.Plate),
                AttendantName: attendant.GetName(),
            }
            allBlueToyotas = append(allBlueToyotas, investigation)
        }
    }
    
    return allBlueToyotas
}

// RobberyInvestigation represents complete information for robbery case investigation
type RobberyInvestigation struct {
    Car           Car
    LotID         int
    SlotID        int
    AttendantName string
}

//use case- 14
// InvestigateBMWCars finds all BMW cars for security enhancement purposes
func (pd *PoliceDepartment) InvestigateBMWCars(lots []*ParkingLot) []SecurityInvestigation {
    var allBMWCars []SecurityInvestigation
    
    for i, lot := range lots {
        bmwCars := lot.FindCarsByMake("BMW")
        for _, car := range bmwCars {
            investigation := SecurityInvestigation{
                Car:    car,
                LotID:  i,
                SlotID: lot.FindCar(car.Plate),
            }
            allBMWCars = append(allBMWCars, investigation)
        }
    }
    
    return allBMWCars
}

// SecurityInvestigation represents information for security enhancement purposes
type SecurityInvestigation struct {
    Car    Car
    LotID  int
    SlotID int
}

//use case-15
// InvestigateRecentlyParkedCars finds all cars parked within specified minutes for bomb threat investigation
func (pd *PoliceDepartment) InvestigateRecentlyParkedCars(lots []*ParkingLot, minutes int) []BombThreatInvestigation {
    var allRecentCars []BombThreatInvestigation
    
    for i, lot := range lots {
        recentCars := lot.FindCarsParkedInLastMinutes(minutes)
        for _, car := range recentCars {
            investigation := BombThreatInvestigation{
                Car:         car,
                LotID:       i,
                SlotID:      lot.FindCar(car.Plate),
                ParkingTime: lot.GetParkingTime(car.Plate),
            }
            allRecentCars = append(allRecentCars, investigation)
        }
    }
    
    return allRecentCars
}

// BombThreatInvestigation represents information for bomb threat investigation
type BombThreatInvestigation struct {
    Car         Car
    LotID       int
    SlotID      int
    ParkingTime time.Time
}

//UC-16
// InvestigateHandicapPermitFraud finds small handicap cars in specific rows for fraud investigation
func (pd *PoliceDepartment) InvestigateHandicapPermitFraud(lots []*ParkingLot, targetRows []string) []HandicapFraudInvestigation {
    var allFraudCars []HandicapFraudInvestigation
    
    for i, lot := range lots {
        suspiciousCars := lot.FindSmallHandicapCarsInRows(targetRows)
        for _, carInfo := range suspiciousCars {
            investigation := HandicapFraudInvestigation{
                CarInfo: carInfo,
                LotID:   i,
            }
            allFraudCars = append(allFraudCars, investigation)
        }
    }
    
    return allFraudCars
}

// HandicapFraudInvestigation represents information for handicap permit fraud investigation
type HandicapFraudInvestigation struct {
    CarInfo CarParkingInfo
    LotID   int
}

//UC-17
// InvestigateFraudulentPlates gets all cars in a specific lot for plate fraud investigation
func (pd *PoliceDepartment) InvestigateFraudulentPlates(lot *ParkingLot) []PlateInvestigation {
    var allPlateInvestigations []PlateInvestigation
    
    allCars := lot.GetAllParkedCars()
    for _, car := range allCars {
        investigation := PlateInvestigation{
            Car:         car,
            SlotID:      lot.FindCar(car.Plate),
            ParkingTime: lot.GetParkingTime(car.Plate),
        }
        allPlateInvestigations = append(allPlateInvestigations, investigation)
    }
    
    return allPlateInvestigations
}

// PlateInvestigation represents information for fraudulent plate number investigation
type PlateInvestigation struct {
    Car         Car
    SlotID      int
    ParkingTime time.Time
}

