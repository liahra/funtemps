package funfacts

import "strings"

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra

  Sett inn alle Funfucts i en struktur
  type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
  }
*/

type FunFacts struct {
  Sun []string
  Luna []string
  Terra []string
}

func GetFunFacts(about string) string {
  facts := FunFacts {
    []string {"Temperatur på ytre lag av Solen 5506.85°C.", "Temperatur i Solens kjerne er 15 000 000°C."}, 
    []string {"Temperatur på Månens overflate om natten: -183", "Temperatur på Månens overflate om dagen: 106."}, 
    []string {"Høyeste temperatur målt på Jordens overflate: 56.7", "Laveste temperatur målt på Jordens overflate: -89.4", "Temperatur i Jordens indre kjerne: 9392K."},
  }
	
  
  if about == "sun" {
		return strings.Join(facts.Sun, "\n")
	} else if about == "luna"{
		return strings.Join(facts.Luna, "\n")
	} else {
    return strings.Join(facts.Terra, "\n")
  }
}
