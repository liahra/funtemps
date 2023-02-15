package conv

import (
	"reflect"
	"testing"
)

func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
		{input: 275, want: 135},
	}

	for _, tc := range tests {
		// Alternativ 1
		// Konverter got til string med 2 desimaler, og så tilbake til float. Gir riktig svar med DeepEqual
		// got, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", FahrenheitToCelsius(tc.input)), 64)

		// Alternativ 2
		// Tester på differanse mellom got og want lik mindre enn en gitt toleranse.
		// got := FahrenheitToCelsius(tc.input)
		// tol := 0.005
		// diff := math.Abs(got - tc.want)
		// if diff > tol {
		// 	t.Errorf("Tolerance set to %v, got difference %v", tol, diff)
		// 	t.Errorf("Wanted: %v, got: %v", tc.want, got)
		// }

		// Alternativ 3 (original)
		// Tester på likhet med DeepEqual
		got := FahrenheitToCelsius(tc.input)
		 if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		} 
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 90, want: 194},
		{input: 0, want: 32},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 700, want: 800.33},
		{input: 273.15, want: 32},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: -273.15},
		{input: 313.15, want: 40},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: 273.15},
		{input: 40, want: 313.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 800.33, want: 700},
		{input: 32, want: 273.15},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}