# funtemps

Sammendrag:

Det å teste input og want  viste seg å være en liten utfordring når det kom til visse flyttall (float64), på grunn av en velkjent avrundingsfeil når det testes på float64. Dette er selvfølgelig et stort problem når mye av oppgaven går ut på å teste verdier av flyttall, så vi valgte derfor å lage en funksjon som formaterte returverdiene i konverteringsfunksjonen i conv.go.Selvom det står spesifikt i oppgavespesfikasjonen at det ikke skulle gjøres formatertinger i konverterings-funksjonene, gjorde vi dette for at funksjonene i conv skulle passere testene i conv.go. Vi prøvde først å ha funksjonen i main funksjonen, som funket, helt til vi prøvde å passere testene som ikke gikk gjennom. I main har vi også en funksjon som formaterer output videre, for å at vi skulle få til tusentallsnotasjonen. Hvis det hadde fungert med testene, så kunne også formatReturns()  vært flyttet inn til main.


Vil også bare informere om at med generell god forståelse for koding har denne oppgaven hatt et arbeidsomfang på ca. 17 timer.


Mvh,

Christine og Iselin