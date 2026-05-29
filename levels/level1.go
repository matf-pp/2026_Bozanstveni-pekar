package levels

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"

	//"2026_Bozanstveni-pekar/screens"
	"github.com/matf-pp/2026_Bozanstveni-pekar/screens"
)

type Level1 struct {
	*screens.Game
	screens.BaseGame

	dante  *sdl.Texture
	danteW int32
	danteH int32

	goodTunnel  *sdl.Texture
	goodTunnelW int32
	goodTunnelH int32

	badTunnel  *sdl.Texture
	badTunnelW int32
	badTunnelH int32

	verticalPath *sdl.Texture

	OpsegVertikalnih [4]sdl.Rect
}

func NewLevel1(game *screens.Game) *Level1 {
	return &Level1{
		Game: game,
	}
}

func (g *Level1) LoadMedia() error {
	var err error
	g.BaseGame.BackgroundImage, err = img.LoadTexture(g.Game.Renderer, "images/lvl1.png")
	if err != nil {
		return fmt.Errorf("error loading texture %v\n", err)
	}
	DanteTexture, err := img.LoadTexture(g.Game.Renderer, "images/Dante.png")
	g.dante = DanteTexture

	//_, _, danteW, danteH, err = g.dante.Query()

	if err != nil {
		return err
	}

	GoodTunnelTexture, err := img.LoadTexture(g.Game.Renderer, "images/goodTunnel.png")
	g.goodTunnel = GoodTunnelTexture

	BadTunnelTexture, err := img.LoadTexture(g.Game.Renderer, "images/badTunnel.png")
	g.badTunnel = BadTunnelTexture

	verticalPathTexture, err := img.LoadTexture(g.Game.Renderer, "images/VertikalniPut.png")
	g.verticalPath = verticalPathTexture

	g.Music, err = mix.LoadMUS("music/spaceroad.mp3")
	if err != nil {
		return fmt.Errorf("error loading music %v\n", err)
	}

	err = g.Music.Play(0)
	if err != nil {
		return fmt.Errorf("error playing music %v\n", err)
	}
	return err
}

func (g *Level1) Close() {
	if g != nil {

		mix.HaltMusic()
		mix.HaltChannel(-1)
		g.Music.Free()
		g.Music = nil
		g.dante.Destroy()
		g.dante = nil
		g.goodTunnel.Destroy()
		g.goodTunnel = nil
		g.badTunnel.Destroy()
		g.badTunnel = nil
		g.verticalPath.Destroy()
		g.verticalPath = nil
	}
}

type IzgradjeniPut struct {
	centerX, centerY       int32
	sirinaPuta, visinaPuta int32
	ugao                   float64
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

// funkcija koja crta krug - Midpoint Circle Algorithm
func (g *Level1) nacrtajKrug(pikselX int32, pikselY int32) {
	x := int32(8)
	y := int32(0)
	t := int32(0)
	g.Game.Renderer.DrawPoint(pikselX, pikselY)
	for x >= y {
		g.Game.Renderer.DrawPoint(pikselX+x, pikselY+y)
		g.Game.Renderer.DrawPoint(pikselX+y, pikselY+x)
		g.Game.Renderer.DrawPoint(pikselX-y, pikselY+x)
		g.Game.Renderer.DrawPoint(pikselX-x, pikselY+y)
		g.Game.Renderer.DrawPoint(pikselX-x, pikselY-y)
		g.Game.Renderer.DrawPoint(pikselX-y, pikselY-x)
		g.Game.Renderer.DrawPoint(pikselX+y, pikselY-x)
		g.Game.Renderer.DrawPoint(pikselX+x, pikselY-y)

		if t <= 0 {
			y++
			t += 2*y + 1
		}
		if t > 0 {
			x--
			t -= 2*x + 1
		}
	}
}

func (g *Level1) NacrtajPut(x1, y1, x2, y2 int32) (int32, int32, int32, int32, float64) {
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

func (g *Level1) RenderujKosePuteve(putevi []IzgradjeniPut) {

	for _, kos := range putevi {
		g.Game.Renderer.CopyEx(g.verticalPath, nil, &sdl.Rect{
			X: kos.centerX, Y: kos.centerY, W: kos.sirinaPuta, H: kos.visinaPuta,
		}, kos.ugao, nil, sdl.FLIP_NONE)
	}
}

// Pomocna funkcija
func Unutra(clickX, opX, opW int32) bool {
	return clickX >= opX && clickX <= (opX+opW)
}

// Proverava da li je klik izvrsen unutar puteva
func (g *Level1) KlikUnutra(clickX int32) bool {
	for _, opseg := range g.OpsegVertikalnih {
		if Unutra(clickX, opseg.X, opseg.W) {
			return true
		}
	}

	return false
}

// Proverava da li su tuneli susedni
func (g *Level1) DozvoljenoSpajanje(click1X, click2X int32) bool {
	if Unutra(click1X, g.OpsegVertikalnih[0].X, g.OpsegVertikalnih[0].W) && Unutra(click2X, g.OpsegVertikalnih[1].X, g.OpsegVertikalnih[1].W) {
		return true
	} else if Unutra(click1X, g.OpsegVertikalnih[1].X, g.OpsegVertikalnih[1].W) && (Unutra(click2X, g.OpsegVertikalnih[0].X, g.OpsegVertikalnih[0].W) || Unutra(click2X, g.OpsegVertikalnih[2].X, g.OpsegVertikalnih[2].W)) {
		return true
	} else if Unutra(click1X, g.OpsegVertikalnih[2].X, g.OpsegVertikalnih[2].W) && (Unutra(click2X, g.OpsegVertikalnih[1].X, g.OpsegVertikalnih[1].W) || Unutra(click2X, g.OpsegVertikalnih[3].X, g.OpsegVertikalnih[3].W)) {
		return true
	} else if Unutra(click1X, g.OpsegVertikalnih[3].X, g.OpsegVertikalnih[3].W) && Unutra(click2X, g.OpsegVertikalnih[2].X, g.OpsegVertikalnih[2].W) {
		return true
	}

	return false
}

func (g *Level1) Run() screens.ScreenID {
	horizontalPaths := []IzgradjeniPut{}
	var klik Kliknut

	g.OpsegVertikalnih = [4]sdl.Rect{
		{X: 82, Y: 0, W: 32, H: 580},
		{X: 257, Y: 0, W: 32, H: 580},
		{X: 457, Y: 0, W: 32, H: 580},
		{X: 672, Y: 0, W: 32, H: 580},
	}

	for true {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent: //iks na prozoru ili crtl q
				return screens.ExitScreen
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN { //pritisnuto dugme na tastaturi
					switch e.Keysym.Scancode { //koje tacno dugme je u pitanju
					case sdl.SCANCODE_ESCAPE: //esc
						return screens.ExitScreen
					}
				}

			case *sdl.MouseButtonEvent:
				if e.Type == sdl.MOUSEBUTTONDOWN {
					if e.Clicks == 1 && e.Button == sdl.BUTTON_LEFT && (g.KlikUnutra(e.X) == true) {
						if klik.brKlikova == 0 {
							klik.klik1x = e.X
							klik.klik1y = e.Y
							klik.brKlikova++
							klik.kliknuto1 = true

						} else if klik.brKlikova == 1 && (g.DozvoljenoSpajanje(klik.klik1x, e.X) == true) {
							klik.klik2x = e.X
							klik.klik2y = e.Y
							klik.brKlikova++
							klik.kliknuto2 = true

						}
					} else if e.Clicks == 1 && e.Button == sdl.BUTTON_RIGHT {
						klik.brKlikova = 0
						klik.kliknuto1 = false
					}
				}
			}
		}
		g.Game.Renderer.Clear()
		g.Game.Renderer.Copy(g.BaseGame.BackgroundImage, nil, nil)

		g.Game.Renderer.Copy(g.verticalPath, nil, &sdl.Rect{
			X: 82, Y: 0, W: 32, H: 580,
		})

		g.Game.Renderer.Copy(g.verticalPath, nil, &sdl.Rect{
			X: 257, Y: 0, W: 32, H: 580,
		})

		g.Game.Renderer.Copy(g.verticalPath, nil, &sdl.Rect{
			X: 457, Y: 0, W: 32, H: 580,
		})

		g.Game.Renderer.Copy(g.verticalPath, nil, &sdl.Rect{
			X: 672, Y: 0, W: 32, H: 580,
		})

		g.Game.Renderer.Copy(g.goodTunnel, nil, &sdl.Rect{
			X: 50, Y: 500, W: 105, H: 105,
		})

		g.Game.Renderer.Copy(g.badTunnel, nil, &sdl.Rect{
			X: 225, Y: 500, W: 100, H: 100,
		})

		g.Game.Renderer.Copy(g.badTunnel, nil, &sdl.Rect{
			X: 425, Y: 500, W: 100, H: 100,
		})

		g.Game.Renderer.Copy(g.badTunnel, nil, &sdl.Rect{
			X: 640, Y: 500, W: 100, H: 100,
		})

		if klik.kliknuto1 == true {
			g.Game.Renderer.SetDrawColor(255, 0, 0, 255) // crvena
			g.nacrtajKrug(klik.klik1x, klik.klik1y)      //g.Game.Renderer.DrawPoint(klik.klik1x, klik.klik1y)

		}
		if klik.kliknuto2 == true {
			g.Game.Renderer.SetDrawColor(255, 0, 0, 255) // crvena
			g.nacrtajKrug(klik.klik2x, klik.klik2y)      //g.Game.Renderer.DrawPoint(klik.klik2x, klik.klik2y)

			cX, cY, sp, vp, u := g.NacrtajPut(klik.klik1x, klik.klik1y, klik.klik2x, klik.klik2y)

			horizontalPaths = append(horizontalPaths, IzgradjeniPut{
				centerX: cX, centerY: cY, sirinaPuta: sp, visinaPuta: vp, ugao: u,
			})

			klik.brKlikova = 0
			klik.kliknuto1 = false
			klik.kliknuto2 = false
		}

		g.RenderujKosePuteve(horizontalPaths)

		g.Game.Renderer.Copy(g.dante, nil, &sdl.Rect{
			X: 66, Y: 10, W: 60, H: 60,
		})

		g.Game.Renderer.Present()
		sdl.Delay(16)

	}
	return screens.Level2Screen
}
