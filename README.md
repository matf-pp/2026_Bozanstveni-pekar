# 2026_Božanstveni-pekar

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/41bc129bd46c4a589452bcda783eab64)](https://app.codacy.com/gh/matf-pp/2026_Bozanstveni-pekar/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)

## Opis

“Božanstveni pekar” je igra inspirisana “Božanstvenom komedijom” i ljubavlju autora igre prema hlebu. Igrač navodi Dantea kroz devet krugova pakla kako bi iz njega spasao sve hlebove koje na svom put nađe. Igra je u 2D-u, a grafika dočarana piksel artom i *fotorealizmom*.

## Uputsvo

Na početnom ekranu nepohodno je uneti ime u naznačenom polju. Kada je ime uneto, pritiskom “Enter” tastera na tastaturi ili “play” dugmeta prikazanog na ekranu započinje se igra. Igrač može u bilo kom trenutku izaći iz igre pritiskom “Esc” tastera.

<https://github.com/user-attachments/assets/cd2982c1-9c46-4c40-a693-03afcf7c8d1d>

Tokom igre, Dante se spušta niz jedan od četiri vertikalnih puteva kroz devet nivoa. Na kraju tri puta nalaze se pećnice, dok je na kraju jednog puta tost. Tost i pećnice se raspoređuju nasumično na svakom nivou. Cilj igrača je da usmeri Dantea ka tostu gradeći kose puteve između već postojećih veritkalnih. Ako Dante upadne u pećnicu, igrač gubi igru.

<https://github.com/user-attachments/assets/7f8437e8-8c00-45e3-bad6-838870485a21>

Svaki nivo ima 5 iteracija. Na početku svake iteracije, Dante počinje kretanje sa nasumičnog puta. Na kraju svake iteracije u okviru nivoa, Dante ubrzava svoje kretanje.

Igrač gradi kose puteve tako što levim klikom na mišu pritisne na početnu i krajnju tačku puta koji želi da izgradi. Pritisnute tačke se moraju nalaziti na dva susedna vertikalna puta. Desnim klikom igrač otklanja izbor početne tačke puta koji želi da izgradi. Izgrađeni putevi se ne uklanjaju do prelaska na naredni nivo.

<https://github.com/user-attachments/assets/1f0ae506-2dd7-4cc9-8455-dff039e550d8>

## Jezici i korišćene tehnologije

- Go
- SDL2
- Pixilart
- Piskel
- Visual Studio Code

## Prevođenje i pokretanje projekta

Da biste samostalno preveli i pokrenuli projekat, potrebno je da na vasem računaru imate instalirane [Go v1.13+](https://go.dev/dl/) i [SDL2](https://github.com/libsdl-org/SDL/releases).

Pokretanje možete izvršiti narednim komandama u vašem CLI:

```bash
git clone github.com/matf-pp/2026_Bozanstveni-pekar.git
cd  2026_Bozanstveni-pekar
go build main.go
./main
```

## Pokretanje izvršnog programa

Preuzmite Bozanstveni_pekar.zip iz Releases taba. Kada otpakujete arhivu, pozicionirajte se unutar otpakovanog direktorijuma. Program možete pokrenuti duplim klikom na izvršni fajl “Bozanstveni pekar” ili komandom:

```bash
./'Bozanstveni pekar'
```

## Operativni sistemi

- Linux

## Autori

Lazar Beljić - <https://github.com/Beldzik>

Nađa Kostić - <https://github.com/djadjaa>

<https://github.com/microslop-pp>
