package screens

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type Congrats struct {
	*Game
	BaseGame

	screenTextTexture *sdl.Texture
	screenTextSize    int
	screenTextW       int32
	screenTextH       int32

	uWonTexture *sdl.Texture
	uWonW       int32
	uWonH       int32

	playAgainButton      Button
	playAgainHoverButton Button

	blur    *sdl.Texture
	snapped bool
}

func NewCongrats(game *Game) *Congrats {
	return &Congrats{
		Game: game,
	}
}
//Pomoćna metoda za kreiranje blur overlay-a - za ekran kada pobedi igrač
func (g *Congrats) CreateBlur() error {
	var err error

	g.blur, err = g.Renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		120,
		90,
	)
	if err != nil {
		return err
	}

	g.Renderer.SetRenderTarget(g.blur)
	g.Renderer.Copy(g.BackgroundImage, nil, nil)
	g.Renderer.SetRenderTarget(nil)

	return nil
}
//Učitavanje sadržaja ekrana za pobedu
func (g *Congrats) LoadMedia() error {
	var err error
	g.BackgroundImage, err = img.LoadTexture(g.Renderer, "images/background.jpg")
	if err != nil {
		return fmt.Errorf("error loading texture %v\n", err)
	}
	screenTextFont, err := g.LoadFont(70)
	if err != nil {
		return fmt.Errorf("error loading font%v\n", err)
	}
	uWonFont, err := g.LoadFont(30)
	if err != nil {
		return fmt.Errorf("error loading font%v\n", err)
	}

	defer screenTextFont.Close()
	defer uWonFont.Close()

	congratsSurf, err := screenTextFont.RenderUTF8Blended("Congrats", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("error loading font surface%v\n", err)
	}
	uWonSurf, err := uWonFont.RenderUTF8Blended("You won the game !!!", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("error loading font surface%v\n", err)
	}

	defer congratsSurf.Free()
	defer uWonSurf.Free()

	//prikaz
	g.screenTextTexture, err = g.Renderer.CreateTextureFromSurface(congratsSurf)
	if err != nil {
		return fmt.Errorf("error loading font texture from surface%v\n", err)
	}
	g.uWonTexture, err = g.Renderer.CreateTextureFromSurface(uWonSurf)
	if err != nil {
		return fmt.Errorf("error loading font texture from surface%v\n", err)
	}

	_, _, g.screenTextW, g.screenTextH, err = g.screenTextTexture.Query()
	if err != nil {
		return fmt.Errorf("error loading screen text query %v\n", err)
	}
	_, _, g.uWonW, g.uWonH, err = g.uWonTexture.Query()
	if err != nil {
		return fmt.Errorf("error loading screen text query %v\n", err)
	}

	//obicna dugmad
	g.playAgainButton, err = g.LoadButton(
		"images/button.png", 300, 300, 200, 200,
	)
	g.SetButtonText(&g.playAgainButton, "play again", 20, sdl.Color{R: 0, G: 0, B: 0, A: 255})

	//hover dugmad
	g.playAgainHoverButton, err = g.LoadButton(
		"images/buttonHover.png", 300, 300, 200, 200,
	)
	g.SetButtonText(&g.playAgainHoverButton, "play again", 20, sdl.Color{R: 255, G: 255, B: 255, A: 255})

	g.playAgainButton.hoverTexture = g.playAgainHoverButton.texture
	g.playAgainButton.hoverTextTexture = g.playAgainHoverButton.textBtnTexture

	err = g.CreateBlur()
	if err != nil {
		return err
	}

	g.ClickSound, err = mix.LoadWAV("sounds/click.mp3")
	if err != nil {
		return fmt.Errorf("error loading chunk %v\n", err)
	}

	g.Music, err = mix.LoadMUS("sounds/win.mp3")
	if err != nil {
		return fmt.Errorf("error loading music %v\n", err)
	}

	err = g.Music.Play(0)
	if err != nil {
		return fmt.Errorf("error playing music %v\n", err)
	}

	return err
}
//Zatvaranje ekrana za gubitak
func (g *Congrats) Close() {
	if g != nil {
		mix.HaltMusic()
		mix.HaltChannel(-1)
		g.Music.Free()
		g.Music = nil
		g.ClickSound.Free()
		g.ClickSound = nil
		g.screenTextTexture.Destroy()
		g.screenTextTexture = nil

		g.uWonTexture.Destroy()
		g.uWonTexture = nil

		g.playAgainButton.texture.Destroy()
		g.playAgainButton.texture = nil
		g.playAgainHoverButton.texture.Destroy()
		g.playAgainHoverButton.texture = nil

		g.BackgroundImage.Destroy()
		g.BackgroundImage = nil

	}
}
func (g *Congrats) Events() ScreenID {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent: //iks na prozoru ili crtl q
			return ExitScreen
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN { //pritisnuto dugme na tastaturi
				switch e.Keysym.Scancode { //koje tacno dugme je u pitanju
				case sdl.SCANCODE_ESCAPE: //esc
					return ExitScreen
				}
			}
		case *sdl.MouseButtonEvent:
			if e.Type == sdl.MOUSEBUTTONDOWN {
				mouseX := e.X
				mouseY := e.Y

				if g.playAgainButton.isClicked(mouseX, mouseY) {
					fmt.Println("play again clicked")
					g.ClickSound.Play(-1, 0)
					return StartScreen
				}

			}
		}
	}
	return CongratsScreen
}

//Pokretanje ekrana za gubitak
func (g *Congrats) Run() ScreenID {
	for true {
		nextScreen := g.Events()
		if nextScreen != CongratsScreen{
			return nextScreen
		}
		g.Renderer.Clear() //svaki frame pocinje 'praznim' ekranom i brise se sve sto je bilo sa prethodnog framea

		//blur
		blrDst := sdl.Rect{X: 0, Y: 0, W: 800, H: 600}
		g.Renderer.Copy(g.blur, nil, &blrDst)

		g.Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		g.Renderer.SetDrawColor(0, 0, 0, 150)
		g.Renderer.FillRect(nil)

		g.Renderer.Copy(g.screenTextTexture, nil, &sdl.Rect{
			X: (WindowWidth - g.screenTextW) / 2,
			Y: (WindowHeight - g.screenTextH) / 2,
			W: g.screenTextW,
			H: g.screenTextH})

		g.Renderer.Copy(g.uWonTexture, nil, &sdl.Rect{
			X: (WindowWidth - g.uWonW) / 2,
			Y: (WindowHeight-g.uWonH)/2 + 70,
			W: g.uWonW,
			H: g.uWonH,
		})

		mouseX, mouseY, _ := sdl.GetMouseState()
		g.playAgainButton.hovered = g.playAgainButton.isClicked(mouseX, mouseY)

		//PLAY
		if g.playAgainButton.hovered {
			g.Renderer.Copy(g.playAgainButton.hoverTexture, nil, &g.playAgainButton.rect)
		} else {
			g.Renderer.Copy(g.playAgainButton.texture, nil, &g.playAgainButton.rect)
		}

		//tekst:
		var tex *sdl.Texture
		var tw, th int32

		if g.playAgainButton.hovered {
			tex = g.playAgainButton.hoverTextTexture
		} else {
			tex = g.playAgainButton.textBtnTexture
		}

		_, _, tw, th, _ = tex.Query()

		g.Renderer.Copy(tex, nil, &sdl.Rect{
			X: g.playAgainButton.rect.X + (g.playAgainButton.rect.W-tw)/2,
			Y: g.playAgainButton.rect.Y + (g.playAgainButton.rect.H-th)/2 + 30,
			W: tw,
			H: th,
		})

		g.Renderer.Present() //prikaze sve sto je nacrtano u ovom frameu
		sdl.Delay(16)        //koliko ce dugo da se prikaze igrica
	}
	return StartScreen
}