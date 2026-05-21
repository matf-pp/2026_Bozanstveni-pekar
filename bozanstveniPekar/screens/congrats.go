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

func newCongrats(game *Game) *Congrats {
	return &Congrats{
		Game: game,
	}
}

func (g *Congrats) LoadMedia() error {
	var err error
	g.backgroundImage, err = img.LoadTexture(g.renderer, "images/demo.jpg")
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
	g.screenTextTexture, err = g.renderer.CreateTextureFromSurface(congratsSurf)
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

		g.backgroundImage.Destroy()
		g.backgroundImage = nil
	}
}

func (g *Congrats) Run() ScreenID {
	
}