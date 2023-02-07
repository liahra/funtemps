package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
  // celsius := (value - 32)*(5.0/9)
	return (value - 32)*(5.0/9.0)
}

// De andre konverteringsfunksjonene implementere her
// ...

func CelsiusToFahrenheit(value float64) float64 {
  return (value * (9.0/5) + 32)

  // Farhrenheit = Celsius*(9/5) + 32
}

func KelvinToFarhenheit(value float64) float64 {
  return (value - 273.15) * (9.0/5) + 32

  // Farhrenheit = (Kelvin - 273.15)*(9/5) + 32
}

func KelvinToCelsius(value float64) float64 {
  return (value - 273.15)

  // Celsius = Kelvin - 273.15
}

func CelsiusToKelvin(value float64) float64 {
  return (value + 273.15)

  // Celsius = Kelvin + 273.15
}

func FarhenheitToKelvin(value float64) float64 {
  return (value - 32) * (5.0/9) + 273.15

  // Kelvin = (Farhenheit - 32) * (5/9) + 273.15
}
