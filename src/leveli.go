package src

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"

	//"2026_Bozanstveni-pekar-main/screens"
	"github.com/matf-pp/2026_Bozanstveni-pekar/src/screens"
)

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
		//pocetni framerate ~ 60 ~ (1/0.016)
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
		if g.Music != nil {
			g.Music.Free()
			g.Music = nil
		}
		if g.dante.tekstura != nil {
			g.dante.tekstura.Destroy()
			g.dante.tekstura = nil
		}
		if g.goodTunnel != nil {
			g.goodTunnel.Destroy()
			g.goodTunnel = nil
		}
		if g.badTunnel != nil {
			g.badTunnel.Destroy()
			g.badTunnel = nil
		}
		if g.verticalPath != nil {
			g.verticalPath.Destroy()
			g.verticalPath = nil
		}
		if g.krug != nil {
			g.krug.Destroy()
			g.krug = nil
		}
		if g.BaseGame.BackgroundImage != nil {
			g.BaseGame.BackgroundImage.Destroy()
			g.BaseGame.BackgroundImage = nil
		}
	}
}
