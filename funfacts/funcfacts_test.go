package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string // her må du skrive riktig type for input
		want string // her må du skrive riktig type for returverdien
	}

	// Her må du legge inn korrekte testverdier
	tests := []test{
		{input: "sun", want: "Temperatur på ytre lag av Solen 5506.85°C.\nTemperatur i Solens kjerne er 15 000 000°C."},
	 	{input: "luna", want: "Temperatur på Månens overflate om natten: -183\nTemperatur på Månens overflate om dagen: 106."},
		{input: "terra", want: "Høyeste temperatur målt på Jordens overflate: 56.7\nLaveste temperatur målt på Jordens overflate: -89.4\nTemperatur i Jordens indre kjerne: 9392K."},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
