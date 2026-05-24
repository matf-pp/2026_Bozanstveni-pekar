package screens

import (
	"fmt"

	//"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	/*"github.com/veandco/go-sdl2/mix"*/
	//"github.com/veandco/go-sdl2/ttf"
	//"2026_Bozanstveni-pekar/levels"
)

type TransitionPhase int

const (
	FadeIn TransitionPhase = iota
	Hold //koliko ce tekst biti prikazan dugo na ekranu
	FadeOut
)

type Transition struct {
	*Game
	BaseGame

	transitionTextTexture *sdl.Texture
	transitionSubtextTexture *sdl.Texture

	tW int32
	tH int32
	tsW int32
	tsH int32
	
	alpha uint8
    next  ScreenID
	done bool

	circleName string
    levelNumber int
    levelName string

	phase TransitionPhase
	holdStart uint64
}

func NewTransition(game *Game, next ScreenID, circle string, level int, name string) *Transition {
	return &Transition{
		Game: game,
		next: next,
		alpha: 0,
		done: false,

		circleName: circle,
		levelNumber: level,
		levelName: name,
	}
}

func (g *Transition) LoadMedia() error {
	var err error
	transitionFont, _ := g.LoadFont(60)
	transitionSubFont, _ := g.LoadFont(30)
	defer transitionFont.Close()
	defer transitionSubFont.Close()

	text := fmt.Sprintf(
    "Silazak u %s krug pakla...", g.circleName)

	textsub := fmt.Sprintf("Level %d - %s",g.levelNumber, g.levelName)

	//surface
	transitionSurf, err := transitionFont.RenderUTF8Blended(text, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("error loading font from surface%v\n", err)
	}
	transitionSubSurf, err := transitionSubFont.RenderUTF8Blended(textsub, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("error loading font from surface%v\n", err)
	}

	//texture
	g.transitionTextTexture, err = g.Renderer.CreateTextureFromSurface(transitionSurf)
	if err!=nil {
		return fmt.Errorf("error loading font texture from surface%v\n", err)
	}

	g.transitionSubtextTexture, err = g.Renderer.CreateTextureFromSurface(transitionSubSurf)
	if err != nil {
		return fmt.Errorf("error loading font texture from surface%v\n", err)
	}

	_, _, g.tW, g.tH, err = g.transitionTextTexture.Query()
	_, _, g.tsW, g.tsH, err = g.transitionSubtextTexture.Query()


	return err
}


func (g *Transition) Close() {
	if g != nil {
		g.transitionTextTexture.Destroy()
		g.transitionTextTexture = nil
		g.transitionSubtextTexture.Destroy()
		g.transitionSubtextTexture = nil
	}
}

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
				if g.alpha >0 {
					g.alpha -= 3
				} else {
					g.done = true	
				}
		}

		//pozadina
		g.Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		g.Renderer.SetDrawColor(0,0,0,g.alpha)
		g.Renderer.FillRect(&sdl.Rect{X:0,Y:0,W:800,H:600})

		//tekst
		//ovim dozvoljavamo da mesamo teksturu teksta sa pozadinom i da reaguje na vrednost providnosti
		g.transitionTextTexture.SetBlendMode(sdl.BLENDMODE_BLEND)
		g.transitionSubtextTexture.SetBlendMode(sdl.BLENDMODE_BLEND)

		//dodela vrednosti providnosti
		g.transitionTextTexture.SetAlphaMod(g.alpha)
        g.transitionSubtextTexture.SetAlphaMod(g.alpha)

		g.Renderer.Copy(g.transitionTextTexture, nil, &sdl.Rect{
			X: (WindowWidth - g.tW) / 2,
			Y: (WindowHeight- g.tH)/2 - 70,
			W: g.tW,
			H: g.tH,
		})

		g.Renderer.Copy(g.transitionSubtextTexture, nil, &sdl.Rect{
			X: (WindowWidth - g.tsW) / 2,
			Y: (WindowHeight - g.tsH) / 2,
			W: g.tsW,
			H: g.tsH,
		})

		g.Renderer.Present()
		sdl.Delay(16)
	}	
	return g.next
}