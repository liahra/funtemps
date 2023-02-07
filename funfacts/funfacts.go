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
  var fact string = "Temperatur på ytre lag av Solen 5506.85°C.\nTemperatur i Solens kjerne er 15 000 000°C."
  return fact
}
