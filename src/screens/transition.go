package screens

import (
	"fmt"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

// Novi tip za fazu tranzicije
type TransitionPhase int

const (
	FadeIn TransitionPhase = iota
	Hold                   //koliko ce tekst biti prikazan dugo na ekranu
	FadeOut
)

type Transition struct {
	*Game
	BaseGame

	transitionTextTexture    *sdl.Texture
	transitionSubtextTexture *sdl.Texture

	tW  int32
	tH  int32
	tsW int32
	tsH int32

	alpha uint8    //providnost teksta
	next  ScreenID //koji ekran ide posle tranzicije
	done  bool     //da li je tranzicija završena

	//promenljive koje se menjaju u zavisnosti od nivoa
	circleName  string
	levelNumber int
	levelName   string

	phase     TransitionPhase
	holdStart uint64 //početak zadržavanja teksta
}

func NewTransition(game *Game, next ScreenID, circle string, level int, name string) *Transition {
	return &Transition{
		Game:  game,
		next:  next,
		alpha: 0,
		done:  false,

		circleName:  circle,
		levelNumber: level,
		levelName:   name,
	}
}

// Učitavanje sadržaja tranzicije
func (g *Transition) LoadMedia() error {
	var err error
	transitionFont, _ := g.LoadFont(50)
	transitionSubFont, _ := g.LoadFont(25)
	defer transitionFont.Close()
	defer transitionSubFont.Close()

	//Izvlačimo tekstove koji će biti ispisani na ekranu tokom tranzicije
	text := fmt.Sprintf(
		"Silazak u %s krug pakla...", g.circleName)
	textsub := fmt.Sprintf("Level %d - %s", g.levelNumber, g.levelName)

	//surface
	transitionSurf, err := transitionFont.RenderUTF8Blended(text, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("error loading transition text font from surface %v ", err)
	}
	defer transitionSurf.Free()

	transitionSubSurf, err := transitionSubFont.RenderUTF8Blended(textsub, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("error loading transition subtext font from surface %v ", err)
	}
	defer transitionSubSurf.Free()

	//texture
	g.transitionTextTexture, err = g.Renderer.CreateTextureFromSurface(transitionSurf)
	if err != nil {
		return fmt.Errorf("error loading transition text font texture from surface %v ", err)
	}

	g.transitionSubtextTexture, err = g.Renderer.CreateTextureFromSurface(transitionSubSurf)
	if err != nil {
		return fmt.Errorf("error loading transition subtext font texture from surface %v ", err)
	}

	//Izvlačenje dimenzija teksta
	_, _, g.tW, g.tH, err = g.transitionTextTexture.Query()
	_, _, g.tsW, g.tsH, err = g.transitionSubtextTexture.Query()

	//Učitavanje pozadinske muzike
	g.Music, err = mix.LoadMUS("sounds/transition.mp3")
	if err != nil {
		return fmt.Errorf("error loading music %v\n", err)
	}
	err = g.Music.Play(0)
	if err != nil {
		return fmt.Errorf("error playing music %v\n", err)
	}
	//koliko pozadinska muzika radi
	mix.FadeOutMusic(5000)
	return err
}

// Zatvaranje ekrana tranzicije
func (g *Transition) Close() {
	if g != nil {
		mix.HaltMusic()
		mix.HaltChannel(-1)

		if g.Music != nil {
			g.Music.Free()
			g.Music = nil
		}

		if g.transitionTextTexture != nil {
			g.transitionTextTexture.Destroy()
			g.transitionTextTexture = nil
		}

		if g.transitionSubtextTexture != nil {
			g.transitionSubtextTexture.Destroy()
			g.transitionSubtextTexture = nil
		}
	}
}

// Pokretanje tranzicije
func (g *Transition) Run() ScreenID {
	for !g.done {
		g.Renderer.Clear()

		switch g.phase {
		case FadeIn:
			if g.alpha < 255 {
				g.alpha += 3
			} else {
				g.phase = Hold
				g.holdStart = sdl.GetTicks64()
			}
		case Hold:
			//nakon sto istekne vreme prelazimo na fade out
			if sdl.GetTicks64()-g.holdStart > 2000 {
				g.phase = FadeOut
			}

		case FadeOut:
			if g.alpha > 0 {
				g.alpha -= 3
			} else {
				g.done = true
			}
		}

		//pozadina
		g.Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		g.Renderer.SetDrawColor(0, 0, 0, g.alpha)
		g.Renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: 800, H: 600})

		//tekst
		//ovim dozvoljavamo da mesamo teksturu teksta sa pozadinom i da reaguje na vrednost providnosti
		g.transitionTextTexture.SetBlendMode(sdl.BLENDMODE_BLEND)
		g.transitionSubtextTexture.SetBlendMode(sdl.BLENDMODE_BLEND)

		//dodela vrednosti providnosti
		g.transitionTextTexture.SetAlphaMod(g.alpha)
		g.transitionSubtextTexture.SetAlphaMod(g.alpha)

		g.Renderer.Copy(g.transitionTextTexture, nil, &sdl.Rect{
			X: (WindowWidth - g.tW) / 2,
			Y: (WindowHeight - g.tH) / 2,
			W: g.tW,
			H: g.tH,
		})

		g.Renderer.Copy(g.transitionSubtextTexture, nil, &sdl.Rect{
			X: (WindowWidth - g.tsW) / 2,
			Y: (WindowHeight-g.tsH)/2 + 60,
			W: g.tsW,
			H: g.tsH,
		})

		g.Renderer.Present()
		sdl.Delay(16)
	}
	return g.next
}
