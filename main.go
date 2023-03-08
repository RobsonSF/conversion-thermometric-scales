package main

import "fmt"

const boilingWaterInKelvin = 373.0

func main() {

	tempKelvin := boilingWaterInKelvin
	tempCelsius := tempKelvin - 273.0

	fmt.Printf("The boiling temperature of water in K = %g , boiling temperature of water in °C =%g.", tempKelvin, tempCelsius)
}
