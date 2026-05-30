package src

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"

	//"2026_Bozanstveni-pekar-main/screens"
	"github.com/matf-pp/2026_Bozanstveni-pekar/screens"
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

type Levels struct {
	*screens.Game
	screens.BaseGame

	//struktura koja sadrzi sve podatke o danteu
	dante Dante

	goodTunnel *sdl.Texture

	//koliko puta je dante stigao u dobar tunel (score)
	brojacDobrihTunela int32

	badTunnel *sdl.Texture

	verticalPath *sdl.Texture

	krug *sdl.Texture

	OpsegVertikalnih [4]sdl.Rect

	OpsegTunela   [4]int32
	goodTunnelPos sdl.Rect

	//da bi izbegli situaciju gde dante preskoci tunel jer se krece previse brzo, ubrzavacemo igraca tako sto povecamo framerate
	//ovo je tako zato sto je igrica dovoljno prosta da moze da radi brzo cak i na mom core 2 duo i radeon HD ???? grafickoj
	//a pametnije resenje bi zahtevalo mnogo vremena i vrv ponovno pisanje logike za puteve
	frameTime uint32
}

func NewLevel(game *screens.Game) *Levels {
	return &Levels{
		Game: game,
		dante: Dante{
			x:      82 + 16,
			y:      10,
			brzina: 1,
			w:      60,
			h:      60,
		},
		//pocetni framerate ~ 60 = (1/0.016)
		//framerate = 1/frameTime
		frameTime: 16,
	}
}

func (g *Levels) LoadMedia(putanja string) error {
	var err error
	g.BaseGame.BackgroundImage, err = img.LoadTexture(g.Game.Renderer, putanja) //putanja="images/lvl1.png"
	if err != nil {
		return fmt.Errorf("error loading texture %v\n", err)
	}
	DanteTexture, err := img.LoadTexture(g.Game.Renderer, "images/Dante.png")
	g.dante.tekstura = DanteTexture

	if err != nil {
		return err
	}

	GoodTunnelTexture, err := img.LoadTexture(g.Game.Renderer, "images/goodTunnel.png")
	g.goodTunnel = GoodTunnelTexture

	BadTunnelTexture, err := img.LoadTexture(g.Game.Renderer, "images/badTunnel.png")
	g.badTunnel = BadTunnelTexture

	verticalPathTexture, err := img.LoadTexture(g.Game.Renderer, "images/VertikalniPut.png")
	g.verticalPath = verticalPathTexture

	krugTexture, err := img.LoadTexture(g.Game.Renderer, "images/krug.png")
	g.krug = krugTexture

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

func (g *Levels) Close() {
	if g != nil {

		mix.HaltMusic()
		mix.HaltChannel(-1)
		g.Music.Free()
		g.Music = nil
		g.dante.tekstura.Destroy()
		g.dante.tekstura = nil
		g.goodTunnel.Destroy()
		g.goodTunnel = nil
		g.badTunnel.Destroy()
		g.badTunnel = nil
		g.verticalPath.Destroy()
		g.verticalPath = nil
		g.krug.Destroy()
		g.krug = nil
	}
}

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

// Pomocna funkcija
func Unutra(clickX, opX, opW int32) bool {
	return clickX >= opX && clickX <= (opX+opW)
}

// Proverava da li je klik izvrsen unutar puteva
func (g *Levels) KlikUnutra(clickX int32) bool {
	for _, opseg := range g.OpsegVertikalnih {
		if Unutra(clickX, opseg.X, opseg.W) {
			return true
		}
	}

	return false
}

// Proverava da li su tuneli susedni
func (g *Levels) DozvoljenoSpajanje(click1X, click2X int32) bool {
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

func (g *Levels) CentarVertPuta(clickX int32) int32 {
	for _, opseg := range g.OpsegVertikalnih {
		if Unutra(clickX, opseg.X, opseg.W) {
			// centar vert puta
			return opseg.X + (opseg.W / 2)
		}
	}
	return clickX //ne bi trebalo da dodje do ovde
}

func (g *Levels) RandomStartX() float64 {

	brPuta := rand.N(len(g.OpsegVertikalnih))

	// stavlja ga na x koord puta (+16 jer je to polovina sirine puta)
	return float64(g.OpsegVertikalnih[brPuta].X + 16)
}

func (g *Levels) Run(ekran screens.ScreenID) screens.ScreenID {
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
			fmt.Printf("skor: %d\n", g.brojacDobrihTunela)

			//breakuje petlju da bi otisao na drugi nivo jer ispod petlje pise retunr nesto nesto level2
			if g.brojacDobrihTunela == 5 {
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
	//return screens.Level2Screen
}
