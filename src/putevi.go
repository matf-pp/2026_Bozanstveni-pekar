package src

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type IzgradjeniPut struct {
	centerX, centerY       int32
	sirinaPuta, visinaPuta int32
	ugao                   float64

	X1, Y1 int32 // pocetak kosog puta (gde klikne)
	X2, Y2 int32 // kraj
}

// Pomocna struktura za pravljenje kosih puteva, broji kliktaje na vertikalnim putevima
type Kliknut struct {
	klik1x    int32
	klik1y    int32
	kliknuto1 bool

	klik2x    int32
	klik2y    int32
	kliknuto2 bool

	brKlikova int32
}

// Pomocna funkcija
func Unutra(clickX, clickY, opX, opW int32) bool {
	return clickX >= opX && clickX <= (opX+opW) && clickY < 550
}

// Proverava da li je klik izvrsen unutar puteva
func (g *Levels) KlikUnutra(clickX, clickY int32) bool {
	for _, opseg := range g.OpsegVertikalnih {
		if Unutra(clickX, clickY, opseg.X, opseg.W) {
			return true
		}
	}

	return false
}

// Proverava da li su tuneli susedni
func (g *Levels) DozvoljenoSpajanje(click1X, click1Y, click2X, click2Y int32) bool {
	if Unutra(click1X, click1Y, g.OpsegVertikalnih[0].X, g.OpsegVertikalnih[0].W) && Unutra(click2X, click2Y, g.OpsegVertikalnih[1].X, g.OpsegVertikalnih[1].W) {
		return true
	} else if Unutra(click1X, click1Y, g.OpsegVertikalnih[1].X, g.OpsegVertikalnih[1].W) && (Unutra(click2X, click2Y, g.OpsegVertikalnih[0].X, g.OpsegVertikalnih[0].W) || Unutra(click2X, click2Y, g.OpsegVertikalnih[2].X, g.OpsegVertikalnih[2].W)) {
		return true
	} else if Unutra(click1X, click1Y, g.OpsegVertikalnih[2].X, g.OpsegVertikalnih[2].W) && (Unutra(click2X, click2Y, g.OpsegVertikalnih[1].X, g.OpsegVertikalnih[1].W) || Unutra(click2X, click2Y, g.OpsegVertikalnih[3].X, g.OpsegVertikalnih[3].W)) {
		return true
	} else if Unutra(click1X, click1Y, g.OpsegVertikalnih[3].X, g.OpsegVertikalnih[3].W) && Unutra(click2X, click2Y, g.OpsegVertikalnih[2].X, g.OpsegVertikalnih[2].W) {
		return true
	}

	return false
}

func (g *Levels) CentarVertPuta(clickX, clickY int32) int32 {
	for _, opseg := range g.OpsegVertikalnih {
		if Unutra(clickX, clickY, opseg.X, opseg.W) {
			// centar vert puta
			return opseg.X + (opseg.W / 2)
		}
	}
	return clickX //ne bi trebalo da dodje do ovde
}

func (g *Levels) NacrtajPut(x1, y1, x2, y2 int32) (int32, int32, int32, int32, float64) {
	sirinaPuta := int32(32)

	//duzine stranica
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)

	// Pitagorina teorema - duzina puta
	visinaPuta := int32(math.Sqrt(dx*dx + dy*dy))

	// Centar puta
	centerX := (x1 + x2) / 2
	centerY := (y1 + y2) / 2

	// Ugao rotacije slike puta
	radiani := math.Atan2(dy, dx)
	ugao := (radiani * (180.0 / math.Pi)) - 90.0

	return centerX - sirinaPuta/2, centerY - (visinaPuta / 2), sirinaPuta, visinaPuta, ugao

}

func (g *Levels) RenderujKosePuteve(putevi []IzgradjeniPut) {

	for _, kos := range putevi {
		g.Game.Renderer.CopyEx(g.verticalPath, nil, &sdl.Rect{
			X: kos.centerX, Y: kos.centerY, W: kos.sirinaPuta, H: kos.visinaPuta,
		}, kos.ugao, nil, sdl.FLIP_NONE)
	}
}
