package conv

import "math"

// Konverterer Farhenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	return formatReturns((value - 32) * (5.0 / 9.0))
}

// Konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(value float64) float64 {
  return formatReturns(value*(9.0/5.0) + 32)

}

// Konverterer Kelvin til Fahrenheit
func KelvinToFahrenheit(value float64) float64 {
  return formatReturns((value-273.15)*(9.0/5.0) + 32)

}

// Konverterer Kelvin til Celsius
func KelvinToCelsius(value float64) float64 {
  return formatReturns(value - 273.15)

}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) float64 {
  return formatReturns(value + 273.15)

}

// Konverterer Fahrenheit til Kelvin
func FahrenheitToKelvin(value float64) float64 {
  return formatReturns((value-32)*(5.0/9.0) + 273.15)
}

// Formaterer verdiene til riktig antall desimaltall (0 eller 2)
func formatReturns(value float64) float64 {

	// Deler value opp i heltall og desimaltall
	int, float := math.Modf(value)

	// Er desimaltalldelen lik 0?
	if float == 0 {
		var prec uint = 0
		ratio := math.Pow(10, float64(prec))
		return math.Round(int*ratio) / ratio

	} else {
		var prec uint = 2
		ratio := math.Pow(10, float64(prec))
		return math.Round(value*ratio) / ratio
	}

}