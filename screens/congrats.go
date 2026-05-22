package screens

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

)

type Congrats struct {
	*Game
	BaseGame

	screenTextTexture *sdl.Texture
	screenTextSize int
	screenTextW int32
	screenTextH int32

	playAgainButton Button
	playAgainHoverButton Button

	scoreButton Button
	scoreHoverButton Button
}

func NewCongrats(game *Game) *Congrats {
	return &Congrats{
		Game: game,
	}
}

func (g *Congrats) LoadMedia() error {
	var err error
	g.BackgroundImage, err = img.LoadTexture(g.Renderer, "images/demo.jpg")
	if err != nil {
		return fmt.Errorf("error loading texture %v\n", err)
	}
	screenTextFont, err := g.LoadFont(80)
	if err != nil {
		return fmt.Errorf("error loading font%v\n", err)
	}
	defer screenTextFont.Close()

	congratsSurf, err := screenTextFont.RenderUTF8Blended("Congrats", sdl.Color{R:255, G:255, B:255, A:255})
	if err != nil {
		return fmt.Errorf("error loading font surface%v\n", err)
	}

	defer congratsSurf.Free()

	//prikaz
	g.screenTextTexture, err = g.Renderer.CreateTextureFromSurface(congratsSurf)
	if err != nil {
		return fmt.Errorf("error loading font texture from surface%v\n", err)
	}

	_, _, g.screenTextW, g.screenTextH, err = g.screenTextTexture.Query()
	if err != nil {
		return fmt.Errorf("error loading screen text query %v\n", err)
	}

	//obicna dugmad
	g.playAgainButton, err = g.LoadButton(
		"images/button.png", 200, 300, 200, 200,
	)
	g.SetButtonText(&g.playAgainButton, "play", 20, sdl.Color{R: 0, G: 0, B: 0, A: 255})
	g.scoreButton, err = g.LoadButton(
		"images/button.png", 400, 300, 200, 200,
	)
	g.SetButtonText(&g.scoreButton, "score", 20, sdl.Color{R: 0, G: 0, B: 0, A: 255})

	//hover dugmad
	g.playAgainHoverButton, err = g.LoadButton(
		"images/buttonHover.png", 200, 300, 200, 200,
	)
	g.SetButtonText(&g.playAgainHoverButton, "play", 20, sdl.Color{R: 255, G: 255, B: 255, A: 255})

	g.playAgainButton.hoverTexture = g.playAgainHoverButton.texture
	g.playAgainButton.hoverTextTexture = g.playAgainHoverButton.textBtnTexture

	g.scoreHoverButton, err = g.LoadButton(
		"images/buttonHover.png", 400, 300, 200, 200,
	)
	g.SetButtonText(&g.scoreHoverButton, "score", 20, sdl.Color{R: 255, G: 255, B: 255, A: 255})

	g.scoreButton.hoverTexture = g.scoreHoverButton.texture
	g.scoreButton.hoverTextTexture = g.scoreHoverButton.textBtnTexture

	return err
}

func (g *Congrats) Close() {
	if g!=nil {
		g.screenTextTexture.Destroy()
		g.screenTextTexture = nil

		g.BackgroundImage.Destroy()
		g.BackgroundImage = nil
	}
}

func (g *Congrats) Run() ScreenID {
	for true {
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
						
						if g.playAgainButton.isClicked(mouseX,mouseY) {
							fmt.Println("play again clicked")
							return StartScreen
						}
					}
			}
		}
		g.Renderer.Clear()                           //svaki frame pocinje 'praznim' ekranom i brise se sve sto je bilo sa prethodnog framea
			g.Renderer.Copy(g.BackgroundImage, nil, nil) //nil je cela tekstura
			//Copy(texture, šta uzeti iz slike, gde nacrtati)
			w := g.screenTextW
			h := g.screenTextH
			g.Renderer.Copy(g.screenTextTexture, nil, &sdl.Rect{
				X: (WindowWidth - w) / 2,
				Y: (WindowHeight- h)/2 - 70,
				W: w,
				H: h})

			mouseX, mouseY, _ := sdl.GetMouseState()
			g.playAgainButton.hovered = g.playAgainButton.isClicked(mouseX, mouseY)

			//PLAY	
			if g.playAgainButton.hovered {
				g.Renderer.Copy(g.playAgainButton.hoverTexture, nil, &g.playAgainButton.rect)
			} else {
				g.Renderer.Copy(g.playAgainButton.texture,nil,&g.playAgainButton.rect)
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
				Y: g.playAgainButton.rect.Y + (g.playAgainButton.rect.H-th)/2+30,
				W: tw,
				H: th,
			})
			/*
			if g.tryAgainButton.isClicked(mouseX, mouseY){
         	   return true  // prelaz na sledeći ekran
        	}
			*/
			g.Renderer.Present() //prikaze sve sto je nacrtano u ovom frameu
			sdl.Delay(16)        //koliko ce dugo da se prikaze igrica
	}
	return StartScreen
}