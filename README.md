# 2026_Božanstveni-pekar

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/65b6dffcaea94ec9b793c0d48ca77379)](https://app.codacy.com/gh/matf-pp/2026_Bozanstveni-pekar?utm_source=github.com&utm_medium=referral&utm_content=matf-pp/2026_Bozanstveni-pekar&utm_campaign=Badge_Grade)

## Opis
“Božanstveni pekar” je igra inspirisana “Božanstvenom komedijom” i ljubavlju autora igre prema hlebu. Igrač navodi Dantea kroz devet krugova pakla kako bi iz njega spasao sve hlebove koje na svom put nađe. Igra je u 2D-u, a grafika dočarana piksel artom i *fotorealizmom*.
## Uputsvo
Na početnom ekranu nepohodno je uneti ime u naznačenom polju. Kada je ime uneto, pritiskom “Enter” tastera na tastaturi ili “play” dugmeta prikazanog na ekranu započinje se igra. Igrač može u bilo kom trenutku izaći iz igre pritiskom “Esc” tastera.

Tokom igre, Dante se spušta niz jedan od četiri vertikalnih puteva kroz devet nivoa. Na kraju tri puta nalaze se pećnice, dok je na kraju jednog puta tost. Tost i pećnice se raspoređuju nasumično na svakom nivou. Cilj igrača je da usmeri Dantea ka tostu gradeći kose puteve između već postojećih veritkalnih. Ako Dante upadne u pećnicu, igrač gubi igru.

Svaki nivo ima 5 iteracija. Na početku svake iteracije, Dante počinje kretanje sa nasumičnog puta. Na kraju svake iteracije u okviru nivoa, Dante ubrzava svoje kretanje.

Igrač gradi kose puteve tako što levim klikom na mišu pritisne na početnu i krajnju tačku puta koji želi da izgradi. Pritisnute tačke se moraju nalaziti na dva susedna vertikalna puta. Desnim klikom igrač otklanja izbor početne tačke puta koji želi da izgradi. Izgrađeni putevi se ne uklanjaju do prelaska na naredni nivo.

## Jezici i korišćene tehnologije

    - Go
    - SDL2
    - Pixilart
    - Piskel
    - Visual Studio Code


## Prevođenje i pokretanje projekta
Da biste samostalno preveli i pokrenuli projekat, potrebno je da na vasem računaru imate instalirane [Go v1.13+](https://go.dev/dl/) i [SDL2](https://github.com/libsdl-org/SDL/releases).

Pokretanje možete izvršiti narednim komandama u vašem CLI:
```
$ git clone github.com/matf-pp/2026_Bozanstveni-pekar.git
$ cd  2026_Bozanstveni-pekar
$ go build main.go
$ ./main
```

## Pokretanje izvršnog programa
Preuzmite Bozanstveni_pekar.zip iz Releases taba. Kada otpakujete arhivu, pozicionirajte se unutar otpakovanog direktorijuma. Program možete pokrenuti duplim klikom na izvršni fajl “Bozanstveni pekar” ili komadnom:
```
$ ./'Bozanstveni pekar'
```

## Operativni sistemi

    - Linux

## Autori
Lazar Beljić - https://github.com/Beldzik

Nađa Kostić - https://github.com/djadjaa

https://github.com/microslop-pp
