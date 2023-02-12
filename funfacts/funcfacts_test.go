package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	
	type test struct {
		input string 
		want []string 
	}

	facts := [7]string{
		"Temperatur på ytre lag av Solen: ", 
		"Temperatur i Solens kjerne: ", 
		"Temperatur på Månens overflate om natten: ", 
		"Temperatur på Månens overflate om dagen: ", 
		"Høyeste temperatur målt på Jordens overflate: ", 
		"Laveste temperatur målt på Jordens overflate: ", 
		"Temperatur i Jordens indre kjerne: "}
	
	tests := []test{
		{
			input: "sun",
			want:  facts[0:2],
		},
		{
			input: "luna",
			want:  facts[2:4],
		},
		{
			input: "terra",
			want:  facts[4:],
		},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
