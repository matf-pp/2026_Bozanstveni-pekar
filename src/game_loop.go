package src

import (
	"fmt"
	"math/rand/v2"

	"github.com/matf-pp/2026_Bozanstveni-pekar/src/screens"
	"github.com/veandco/go-sdl2/sdl"
)

func (g *Levels) Run(ekran screens.ScreenID, score *int) screens.ScreenID {
	horizontalPaths := []IzgradjeniPut{}
	var klik Kliknut

	g.OpsegVertikalnih = [4]sdl.Rect{
		{X: 82, Y: 0, W: 32, H: 580},
		{X: 257, Y: 0, W: 32, H: 580},
		{X: 457, Y: 0, W: 32, H: 580},
		{X: 672, Y: 0, W: 32, H: 580},
	}

	g.OpsegTunela = [4]int32{50, 225, 425, 640}
	indeksiTunel := rand.Perm(4) //randomizovanje tunela
	g.goodTunnelPos = sdl.Rect{
		X: g.OpsegTunela[indeksiTunel[0]], Y: 500, W: 105, H: 105}

	g.dante.x = g.RandomStartX()
	g.dante.y = 10
	g.dante.ciljX = g.dante.x
	g.dante.ciljY = g.dante.y
	g.dante.naKosomPutu = false
	g.dante.trenutniPut = nil
	g.dante.poslednjiPut = nil

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
					case sdl.SCANCODE_R:
						horizontalPaths = []IzgradjeniPut{}
						g.dante.x = g.RandomStartX()
						g.dante.y = 10
						g.dante.ciljX = g.dante.x
						g.dante.ciljY = g.dante.y
						g.dante.naKosomPutu = false
						g.dante.trenutniPut = nil
						g.dante.poslednjiPut = nil
					case sdl.SCANCODE_T:
						return ekran
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

		g.Game.Renderer.Copy(g.goodTunnel, nil, &g.goodTunnelPos)

		g.Game.Renderer.Copy(g.badTunnel, nil, &sdl.Rect{
			X: g.OpsegTunela[indeksiTunel[1]], Y: 500, W: 100, H: 100,
		})

		g.Game.Renderer.Copy(g.badTunnel, nil, &sdl.Rect{
			X: g.OpsegTunela[indeksiTunel[2]], Y: 500, W: 100, H: 100,
		})

		g.Game.Renderer.Copy(g.badTunnel, nil, &sdl.Rect{
			X: g.OpsegTunela[indeksiTunel[3]], Y: 500, W: 100, H: 100,
		})

		if klik.kliknuto1 == true {
			g.Game.Renderer.Copy(g.krug, nil, &sdl.Rect{
				X: klik.klik1x - 8, Y: klik.klik1y - 8, W: 15, H: 15,
			})
		}
		if klik.kliknuto2 == true {

			centriranX1 := g.CentarVertPuta(klik.klik1x)
			centriranX2 := g.CentarVertPuta(klik.klik2x)

			flagDaLiDodatiPut := true

			for i := range horizontalPaths {
				p := horizontalPaths[i]

				//ako put koji hocemo da dodamo i trenutni put iz niza spajaju iste cevi
				//brojevi 1 i 2 u X1 i X2 samo oznacavaju koji je prvi pritisnut tkd moramo da proverimo obe kombinacije

				if centriranX1 == p.X1 && centriranX2 == p.X2 {
					// ako se putevi seku: zabrani dodavanje
					if (klik.klik1y < p.Y1 && klik.klik2y > p.Y2) || (klik.klik1y > p.Y1 && klik.klik2y < p.Y2) {
						flagDaLiDodatiPut = false
						break
					}
				} else if centriranX1 == p.X2 && centriranX2 == p.X1 { // druga kombinacija
					// ako se putevi seku: zabrani dodavanje
					if (klik.klik1y < p.Y2 && klik.klik2y > p.Y1) || (klik.klik1y > p.Y2 && klik.klik2y < p.Y1) {
						flagDaLiDodatiPut = false
						break
					}
				}
			}

			if flagDaLiDodatiPut == true {
				g.Game.Renderer.Copy(g.krug, nil, &sdl.Rect{
					X: klik.klik1x - 8, Y: klik.klik1y - 8, W: 15, H: 15,
				})

				cX, cY, sp, vp, u := g.NacrtajPut(centriranX1, klik.klik1y, centriranX2, klik.klik2y)

				horizontalPaths = append(horizontalPaths, IzgradjeniPut{
					centerX: cX, centerY: cY, sirinaPuta: sp, visinaPuta: vp, ugao: u,
					X1: centriranX1, Y1: klik.klik1y, // koord prve tacke
					X2: centriranX2, Y2: klik.klik2y, // koord druge tacke
				})
			}

			klik.brKlikova = 0
			klik.kliknuto1 = false
			klik.kliknuto2 = false
		}

		g.RenderujKosePuteve(horizontalPaths)

		g.PomeriDantea(horizontalPaths)

		tunel := g.goodTunnelPos

		// ova provera za kolizuju je vljd dovoljno precizna
		if int32(g.dante.x) >= tunel.X && int32(g.dante.x) <= (tunel.X+tunel.W) && int32(g.dante.y) >= tunel.Y {
			// stigao u dobar tunel -> uvecamo score
			g.brojacDobrihTunela++
			*score++
			fmt.Printf("skor: %d\n", g.brojacDobrihTunela)

			//breakuje petlju da bi otisao na drugi nivo jer ispod petlje pise retunr nesto nesto level2
			if g.brojacDobrihTunela >= 5 {
				break
			}

			//vrati dantea u random tunel na vrh i resetuj sve ove gluposti
			g.dante.x = g.RandomStartX()
			g.dante.y = 10
			g.dante.ciljX = g.dante.x
			g.dante.ciljY = g.dante.y
			g.dante.naKosomPutu = false
			g.dante.trenutniPut = nil
			g.dante.poslednjiPut = nil

			// neka vrednost za ubrzavanje igraca koja radi dobro
			// smanji na 2 ili 1 ako mislite da je previse brzo
			g.frameTime -= 3

			continue
		}

		// provera da li je umro, odnosno da li se nalazi ispod dobrog tunela, a nije usao u njega
		if int32(g.dante.y) >= tunel.Y+tunel.H/2 {
			fmt.Println("umro")

			g.brojacDobrihTunela = 0

			return screens.GameOverScreen
		}

		g.Game.Renderer.Copy(g.dante.tekstura, nil, &sdl.Rect{
			X: int32(g.dante.x) - g.dante.w/2,
			Y: int32(g.dante.y) - g.dante.h/2,
			W: g.dante.w,
			H: g.dante.h,
		})

		g.Game.Renderer.Present()
		sdl.Delay(g.frameTime)

	}
	return ekran
}
