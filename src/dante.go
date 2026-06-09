package src

import (
	"crypto/rand"
	"math"
	"math/big"

	"github.com/veandco/go-sdl2/sdl"
)

// sadrzi podatke o danteu
type Dante struct {
	tekstura *sdl.Texture
	w        int32
	h        int32
	x        float64
	y        float64
	brzina   float64

	naKosomPutu bool
	trenutniPut *IzgradjeniPut

	ciljX float64
	ciljY float64

	poslednjiPut *IzgradjeniPut
}

func (g *Levels) PomeriDantea(putevi []IzgradjeniPut) {
	// na kosom putu -> pomeraj ka ciljnoj tacki
	if g.dante.naKosomPutu {

		//izracuna pravac, normalizuje vektor, pomnozi ga sa brzinom i to doda na poziciju
		dx := g.dante.ciljX - g.dante.x
		dy := g.dante.ciljY - g.dante.y
		duzina := math.Sqrt(dx*dx + dy*dy)

		if duzina > 0 {
			g.dante.x += (dx / duzina) * g.dante.brzina
			g.dante.y += (dy / duzina) * g.dante.brzina
		}

		// ako je u blizini cilja
		if math.Abs(g.dante.y-g.dante.ciljY) < 1 && math.Abs(g.dante.x-g.dante.ciljX) < 1 {

			// zapamti put da bi ga ignorisali sledeci put da se ne bi vracao nazad istim putem
			g.dante.poslednjiPut = g.dante.trenutniPut

			g.dante.naKosomPutu = false
			g.dante.trenutniPut = nil

			// pomeri ga na dole za svaki slucaj da ne krene opet na isti put
			g.dante.y += g.dante.brzina
		}
		return
	}

	// prodji kroz sve staze i proveri na kojoj se nalazimo
	for i := range putevi {
		p := &putevi[i]

		// ako smo naisli na stari put, ignorisi ga
		if g.dante.poslednjiPut == p {
			continue
		}

		//dve provere: jedna ako prilazi sa jedne, a druga ako prilazi sa druge strane puta
		//PROVERA 1
		if math.Abs(float64(p.X1)-g.dante.x) < 1 && math.Abs(float64(p.Y1)-g.dante.y) < 2 && g.dante.y < float64(p.Y1) {
			g.dante.poslednjiPut = nil
			g.dante.naKosomPutu = true
			g.dante.trenutniPut = p
			g.dante.ciljX = float64(p.X2)
			g.dante.ciljY = float64(p.Y2)
			return
		}
		//PROVERA 2
		if math.Abs(float64(p.X2)-g.dante.x) < 1 && math.Abs(float64(p.Y2)-g.dante.y) < 2 && g.dante.y < float64(p.Y2) {
			g.dante.poslednjiPut = nil
			g.dante.naKosomPutu = true
			g.dante.trenutniPut = p
			g.dante.ciljX = float64(p.X1)
			g.dante.ciljY = float64(p.Y1)
			return
		}

	}

	//ako zaobidje sve ifove, treba da se krece samo na dole
	g.dante.y += g.dante.brzina
}

func (g *Levels) RandomStartX() float64 {

	brPuta, _ := rand.Int(rand.Reader, big.NewInt(int64(len(g.OpsegVertikalnih))))

	// stavlja ga na x koord puta (+16 jer je to polovina sirine puta)
	return float64(g.OpsegVertikalnih[brPuta.Int64()].X + 16)
}
