package screens

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	/*"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/ttf"*/
)

type GameOver struct {
	*Game
	BaseGame
	screenTextTexture *sdl.Texture
	//screenTextSize int
	screenTextW int32
	screenTextH int32

	smiley *sdl.Texture
	smileyW int32
	smileyH int32

	tryAgainButton Button
	tryAgainHoverButton Button

	blur *sdl.Texture
	snapped bool //za snapshot
}


func NewGameOver(game *Game) *GameOver {
	return &GameOver{
		Game: game,
	}
}
//argument metode - (scene *sdl.Texture)
func (g *GameOver) CreateBlur() error {
    var err error

    g.blur, err = g.Renderer.CreateTexture(
        sdl.PIXELFORMAT_RGBA8888,
        sdl.TEXTUREACCESS_TARGET,
        200,
        150,
    )
    if err != nil {
        return err
    }

    g.Renderer.SetRenderTarget(g.blur)
    g.Renderer.Copy(g.BackgroundImage, nil, nil) //umesto g.BackgroundImage staviti scene
    g.Renderer.SetRenderTarget(nil)

    return nil
}

func (g *GameOver) LoadMedia() error {
	var err error
	if g.BackgroundImage, err = img.LoadTexture(g.Renderer, "images/demo.jpg"); err != nil {
		return fmt.Errorf("error loading texture %v\n", err)
	}
	screenTextFont, err := g.LoadFont(80)
	if err != nil {
		return fmt.Errorf("error loading font%v\n", err)
	}
	smileyFont, err := g.LoadFont(30)
	if err != nil {
		return fmt.Errorf("error loading font%v\n", err)
	}
	defer smileyFont.Close()
	defer screenTextFont.Close()

	gameOverSurf, err := screenTextFont.RenderUTF8Blended("Game Over",sdl.Color{R:255, G:255, B:255, A:255})
	if err != nil {
		return fmt.Errorf("error loading font surface%v\n", err)
	}
	smileySurf, err := smileyFont.RenderUTF8Blended(":(",sdl.Color{R:255, G:255, B:255, A:255})
	if err != nil {
		return fmt.Errorf("error loading font surface%v\n", err)
	}

	defer gameOverSurf.Free()
	defer smileySurf.Free()

	//prikaz
	g.screenTextTexture, err = g.Renderer.CreateTextureFromSurface(gameOverSurf)
	if err != nil {
		return fmt.Errorf("error loading font texture from surface%v\n", err)
	}
	g.smiley, err = g.Renderer.CreateTextureFromSurface(smileySurf)
	if err != nil {
		return fmt.Errorf("error loading font texture from surface%v\n", err)
	}

	//izvlacenje podataka
	_, _, g.screenTextW, g.screenTextH, err = g.screenTextTexture.Query()
	if err != nil {
		return fmt.Errorf("error loading screen text query %v\n", err)
	}
	_, _, g.smileyW, g.smileyH, err = g.smiley.Query()
	if err != nil {
		return fmt.Errorf("error loading screen text query %v\n", err)
	}

	//obicno dugme
	g.tryAgainButton, err = g.LoadButton(
		"images/button.png", 300,300,200,200,
	)
	if err != nil {
		return err
	}
	g.SetButtonText(&g.tryAgainButton,"try again",20,sdl.Color{R:0, G:0, B:0, A:255})
	
	//hover dugme
	g.tryAgainHoverButton, err = g.LoadButton(
		"images/buttonHover.png", 300,300,200,200,
	)
	g.SetButtonText(&g.tryAgainHoverButton,"try again",20,sdl.Color{R:255, G:255, B:255, A:255})

	g.tryAgainButton.hoverTexture = g.tryAgainHoverButton.texture
	g.tryAgainButton.hoverTextTexture = g.tryAgainHoverButton.textBtnTexture

	return err
}

func (g *GameOver) Run() ScreenID{ 
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
						
						if g.tryAgainButton.isClicked(mouseX,mouseY) {
							fmt.Println("try again clicked")
							return StartScreen
						}
					}
			}
		}
		g.Renderer.Clear()
		
		//blur
		blrDst := sdl.Rect{X:0,Y:0,W:800,H:600}
		g.Renderer.Copy(g.blur,nil, &blrDst)

		g.Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		g.Renderer.SetDrawColor(0,0,0,150)
		g.Renderer.FillRect(nil)

		//g.Renderer.Copy(g.BackgroundImage, nil, nil) //nil je cela tekstura
		
		w := g.screenTextW
		h := g.screenTextH
		g.Renderer.Copy(g.screenTextTexture, nil, &sdl.Rect{
			X: (WindowWidth - w) / 2,
			Y: (WindowHeight- h)/2,
			W: w,
			H: h})

		g.Renderer.Copy(g.smiley, nil, &sdl.Rect{
		X: (WindowWidth - g.smileyW) / 2,
		Y: (WindowHeight- g.smileyH)/2 + 70,
		W: g.smileyW,
		H: g.smileyH})

		mouseX, mouseY, _ := sdl.GetMouseState()
		g.tryAgainButton.hovered = g.tryAgainButton.isClicked(mouseX, mouseY)

		//PLAY	
		if g.tryAgainButton.hovered {
			g.Renderer.Copy(g.tryAgainButton.hoverTexture, nil, &g.tryAgainButton.rect)
		} else {
			g.Renderer.Copy(g.tryAgainButton.texture,nil,&g.tryAgainButton.rect)
		}

		//tekst:
		var tex *sdl.Texture
		var tw, th int32

		if g.tryAgainButton.hovered {
			tex = g.tryAgainButton.hoverTextTexture
		} else {
			tex = g.tryAgainButton.textBtnTexture
		}

		_, _, tw, th, _ = tex.Query()

		g.Renderer.Copy(tex, nil, &sdl.Rect{
			X: g.tryAgainButton.rect.X + (g.tryAgainButton.rect.W-tw)/2,
			Y: g.tryAgainButton.rect.Y + (g.tryAgainButton.rect.H-th)/2+30,
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

	return ExitScreen
}

func (g *GameOver) Close() { 
	if g != nil {
		//gasimo i postavimo adresu na 0 -> free
		g.screenTextTexture.Destroy()
		g.screenTextTexture = nil

		g.blur.Destroy()
		g.blur = nil

		g.BackgroundImage.Destroy()
		g.BackgroundImage = nil

		/*
		g.Renderer.Destroy()
		g.Renderer = nil

		g.Window.Destroy()
		g.Window = nil
		*/
	}
}
