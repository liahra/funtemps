package funfacts

type FunFacts struct {
  Sun   []string
  Luna  []string
  Terra []string
}

func GetFunFacts(about string) []string {
	facts := FunFacts{
		[]string{"Temperatur på ytre lag av Solen: ", "Temperatur i Solens kjerne: "},
		[]string{"Temperatur på Månens overflate om natten: ", "Temperatur på Månens overflate om dagen: "},
		[]string{"Høyeste temperatur målt på Jordens overflate: ", "Laveste temperatur målt på Jordens overflate: ", "Temperatur i Jordens indre kjerne: "},
	}

	if about == "sun" {
		return facts.Sun
	} else if about == "luna" {
		return facts.Luna
	} else {
		return facts.Terra
	}
}
