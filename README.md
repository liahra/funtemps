# funtemps

Sammendrag:
Det å teste input og want viste seg å være en utfordring når det kom til flyttall (float64), på grunn av at visse verdier i titallssystemet ikke kan representeres eksakt i totallsystemet (binærtall). Dette er selvfølgelig et stort problem når mye av oppgaven går ut på å sammenligne verdier av flyttall, så vi valgte derfor å lage en funksjon som formaterte returverdiene i konverteringsfunksjonen i conv.go. Selvom det står spesifikt i oppgavespesfikasjonen at det ikke skulle gjøres formatertinger i konverteringsfunksjonene, gjorde vi dette for at funksjonene i conv skulle passere testene i conv.go. Vi prøvde først å ha funksjonen i main funksjonen, som funket, helt til vi prøvde å passere testene, som ikke gikk gjennom. I main har vi også en funksjon som formaterer output videre, for å at vi skulle få til tusentallsnotasjonen. Hvis det hadde fungert med testene, så kunne også formatReturns() vært flyttet inn til main.

Oppdatering 15/2:
Har funnet to andre måter å sammenligne flyttall på og lagt disse med som kommentarer i koden. For nærmere informasjon les koden.

Vil også bare informere om at med generell god forståelse for koding har denne oppgaven hatt et arbeidsomfang på ca. 17 timer.

Mvh,
Christine og Iselin