package funfacts

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

func GetFunFacts(about string) string {
	var sunFact string = "Temperatur på ytre lag av Solen 5506.85°C.\nTemperatur i Solens kjerne er 15 000 000°C."
	var lunaFact string = "Temperatur på Månens overflate om natten: -183\nTemperatur på Månens overflate om dagen: 106."
  var terraFact string = "Høyeste temperatur målt på Jordens overflate: 56.7\nLaveste temperatur målt på Jordens overflate: -89.4\nTemperatur i Jordens indre kjerne: 9392K."

	if about == "sun" {
		return sunFact
	} else if about == "luna"{
		return lunaFact
	} else {
    return terraFact
  }

}
