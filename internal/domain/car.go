package domain

import(
	
)

//creating a car struct that represents a vehicle in yhe parking lot
type Car struct{
	Plate string  //Plate Number
	Make string   //Car Manufacturer
	Color string  //Car Color
	Size  CarSize   // Size category of the car, use case- 10
}

//use case-10
// CarSize enum for different car sizes
type CarSize int

const (
    Small CarSize = iota  // Small cars
    Medium                // Medium cars  
    Large                 // Large cars/SUVs
)

// String returns string representation of CarSize
func (cs CarSize) String() string {
    switch cs {
    case Small:
        return "Small"
    case Medium:
        return "Medium"
    case Large:
        return "Large"
    default:
        return "Unknown"
    }
}