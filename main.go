package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/liahra/funtemps/conv"
	"github.com/liahra/funtemps/funfacts"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var cels float64
var kelv float64
var out string
var funfact string
var temp string

func init() {
	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celcius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")

	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfact, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&temp, "t", "C", "temperaturskala i C - celsius, F - farhenheit, K- Kelvin")
}

func main() {
	flag.Parse()

	// Fahrenheit til celsius
	if out == "C" && isFlagPassed("F") {
		fmt.Print(formatNumber(fahr), "°F is ", formatNumber(conv.FahrenheitToCelsius(fahr)), "°C\n")
	}

	// Kelvin til celsius
	if out == "C" && isFlagPassed("K") {
		fmt.Print(formatNumber(kelv), "K is ", formatNumber(conv.KelvinToCelsius(kelv)), "°C\n")
	}

	// Celsius til fahrenheit
	if out == "F" && isFlagPassed("C") {
		fmt.Print(formatNumber(cels), "°C is ", formatNumber(conv.CelsiusToFahrenheit(cels)), "°F\n")
	}

	// Kelvin til fahrenheit
	if out == "F" && isFlagPassed("K") {
		fmt.Print(formatNumber(kelv), "K is ", formatNumber(conv.KelvinToFahrenheit(kelv)), "°F\n")
	}

	// Celsius til kelvin
	if out == "K" && isFlagPassed("C") {
		fmt.Print(formatNumber(cels), "°C is ", formatNumber(conv.CelsiusToKelvin(cels)), "K\n")
	}

	// Fahrenheit til kelvin
	if out == "K" && isFlagPassed("F") {
		fmt.Print(formatNumber(fahr), "°F is ", formatNumber(conv.FahrenheitToKelvin(fahr)), "K\n")
	}

	// Sjekk funfacts
	if temp == "C" && isFlagPassed("funfacts") {
		fact := funfacts.GetFunFacts(funfact)

		if funfact == "sun" {
			fmt.Print(fact[0], formatNumber(conv.KelvinToCelsius(5778)), "°C\n", fact[1], "15 000 000°C\n")
		}

		if funfact == "luna" {
			fmt.Print(fact[0], "-183°C\n", fact[1], "106°C\n")
		}

		if funfact == "terra" {
			fmt.Print(fact[0], "56.7°C\n", fact[1], "-89.4°C\n")
		}
	}

	if temp == "F" && isFlagPassed("funfacts") {
		fact := funfacts.GetFunFacts(funfact)

		if funfact == "sun" {
			fmt.Print(fact[0], formatNumber(conv.KelvinToFahrenheit(5778)), "°F\n", fact[1], formatNumber(conv.CelsiusToFahrenheit(15000000)), "°F\n")
		}

		if funfact == "luna" {
			fmt.Print(fact[0], formatNumber(conv.CelsiusToFahrenheit(-183)), "°F\n", fact[1], formatNumber(conv.CelsiusToFahrenheit(106)), "°F\n")
		}

		if funfact == "terra" {
			fmt.Print(fact[0], "134°F\n", fact[1], formatNumber(conv.CelsiusToFahrenheit(-89.4)), "°F\n")
		}
	}

	if temp == "K" && isFlagPassed("funfacts") {
		fact := funfacts.GetFunFacts(funfact)

		if funfact == "sun" {
			fmt.Print(fact[0], "5 778", "K\n", fact[1], formatNumber(conv.CelsiusToKelvin(15000000)), "K\n")
		}

		if funfact == "luna" {
			fmt.Print(fact[0], formatNumber(conv.CelsiusToKelvin(-183)), "K\n", fact[1], formatNumber(conv.CelsiusToKelvin(106)), "K\n")
		}

		if funfact == "terra" {
			fmt.Print(fact[0], "329.82K\n", fact[1], formatNumber(conv.CelsiusToKelvin(-89.4)), "K\n")
		}
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// Formaterer tall til ønsket format
func formatNumber(convValue float64) string {
	convValueString := fmt.Sprintf("%.2f", convValue)
	sign := ""

	if convValueString[0:1] == "-" {
		sign = "-"
		convValueString = convValueString[1:]
	}

	numSlice := strings.Split(convValueString, ".")
	mainNum := numSlice[0]
	outputString := ""

	for i := len(mainNum); i > 0; i = i - 3 {
		if i-3 < 0 {
			outputString = mainNum[0:i] + " " + outputString
		} else {
			outputString = mainNum[i-3:i] + " " + outputString
		}
	}

	outputString = outputString[0 : len(outputString)-1] // Fordi outputstring av en eller annen grunn alltid har et mellomrom på slutten
	if numSlice[1] != "00" {
		outputString = outputString + "." + numSlice[1]
	}

	outputString = sign + outputString
	return outputString
}
