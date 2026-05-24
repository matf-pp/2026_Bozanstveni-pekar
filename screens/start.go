package screens

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	/*"github.com/veandco/go-sdl2/mix"*/
	"github.com/veandco/go-sdl2/ttf"
)

// globalne konstante
const (
	WindowWidth  = 800
	WindowHeight = 600
	WindowTitle  = "Bozanstveni pekar"
)

type Screen interface {
	LoadMedia() error
	Run() ScreenID // vraca true/false za prelaz na sledeci ekran
	Close()
}

type ScreenID int

const (
	StartScreen ScreenID = iota
	GameOverScreen
	ExitScreen
	CongratsScreen
	TransitionScreen
	Level1Screen
	Level2Screen
	Level3Screen
	Level4Screen
	Level5Screen
	Level6Screen
	Level7Screen
	Level8Screen
	Level9Screen
)

type Button struct {
	texture *sdl.Texture
	size    int
	rect    sdl.Rect

	hoverTexture     *sdl.Texture
	hoverTextTexture *sdl.Texture

	hovered bool

	textBtn        string
	textBtnTexture *sdl.Texture
	textBtnColor   sdl.Color
	textBtnW       int32
	textBtnH       int32
}

type BaseGame struct {
	BackgroundImage *sdl.Texture
	FontColor       sdl.Color
}

type StartGame struct {
	*Game
	BaseGame
	titleTexture      *sdl.Texture
	title2Texture *sdl.Texture

	newGameTexture *sdl.Texture

	insertNameTexture *sdl.Texture
	playerName        string
	playerNameTexture *sdl.Texture

	playButton       Button
	playHoverButton  Button
	scoreButton      Button
	scoreHoverButton Button

	escape *sdl.Texture
	escapeW int32
	escapeH int32

	cursorVisible bool
	blink         uint64

	titleW      int32
	titleH      int32

	newGameW int32
	newGameH int32

	insertNameW int32
	insertNameH int32
	playerW     int32
	playerH     int32
}

func NewStartGame(game *Game) *StartGame {
	return &StartGame{
		Game: game,
	}
}

func (g *Game) LoadFont(size int) (*ttf.Font, error) {
	return ttf.OpenFont(
		"fonts/PixelifySans-VariableFont_wght.ttf",
		size,
	)
}

func (g *StartGame) renderText(text string) error {
	if text == "" {
		if g.playerNameTexture != nil {
			g.playerNameTexture.Destroy()
			g.playerNameTexture = nil
		}
		return nil
	}

	font, err := g.LoadFont(30)
	if err != nil {
		return fmt.Errorf("error loading font %v\n", err)
	}
	defer font.Close()

	fontSurf, err := font.RenderUTF8Blended(text, sdl.Color{R: 0, G: 0, B: 0, A: 255})
	if err != nil {
		return err
	}
	defer fontSurf.Free()
	if g.playerNameTexture != nil {
		g.playerNameTexture.Destroy()
	}

	g.playerNameTexture, err = g.Renderer.CreateTextureFromSurface(fontSurf)

	_, _, g.playerW, g.playerH, err = g.playerNameTexture.Query()
	if err != nil {
		return fmt.Errorf("error loading player query %v\n", err)
	}

	return err
}

func (g *Game) LoadButton(imgPath string, x, y, w, h int32) (Button, error) {
	texture, err := img.LoadTexture(g.Renderer, imgPath)
	if err != nil {
		return Button{}, err //vraca prazno dugme
	}
	//inicijalizacija dugmeta
	button := Button{
		texture: texture,
		rect: sdl.Rect{
			X: x, Y: y, W: w, H: h,
		},
	}
	return button, nil
}

func (b *Button) isClicked(x, y int32) bool {
	return x >= b.rect.X && x <= b.rect.X+b.rect.W && y >= b.rect.Y && y <= b.rect.Y+b.rect.H
}

func (g *Game) SetButtonText(b *Button, text string, fontSize int, color sdl.Color) error {
	font, err := g.LoadFont(fontSize)
	if err != nil {
		return err
	}
	defer font.Close()

	btnSurf, err := font.RenderUTF8Blended(text, color)
	if err != nil {
		return err
	}
	defer btnSurf.Free()
	if b.textBtnTexture != nil {
		b.textBtnTexture.Destroy()
	}
	b.textBtnTexture, err = g.Renderer.CreateTextureFromSurface(btnSurf)
	if err != nil {
		return err
	}
	_, _, b.textBtnW, b.textBtnH, err = b.textBtnTexture.Query()

	return err
}

func (g *StartGame) LoadMedia() error {
	var err error
	if g.BaseGame.BackgroundImage, err = img.LoadTexture(g.Renderer, "images/background.jpg"); err != nil {
		return fmt.Errorf("error loading texture %v\n", err)
	}
	titleFont, err := ttf.OpenFont("fonts/Early_GameBoy.ttf",45)
	if err != nil {
		return err
	}
	//titleFont, _ := g.LoadFont(80)
	newGameFont, _ := g.LoadFont(50)
	insertFont, _ := g.LoadFont(30)
	exitFont, _ := g.LoadFont(20)
	defer titleFont.Close()
	defer newGameFont.Close()
	defer insertFont.Close()
	defer exitFont.Close()
	//kreiramo tekst u sliku koju zelimo da nacrtamo

	titleSurf, err := titleFont.RenderUTF8Blended("Bozanstveni pekar", sdl.Color{R: 72, G: 43, B: 6, A: 255})
	if err != nil {
		return fmt.Errorf("error loading font surface%v\n", err)
	}
	titleSurf2, err := titleFont.RenderUTF8Blended("Bozanstveni pekar", sdl.Color{R: 225, G: 130, B: 19, A: 255})
	if err != nil {
		return fmt.Errorf("error loading font surface%v\n", err)
	}

	newGameSurf, err := newGameFont.RenderUTF8Blended("Start game", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("error loading font surface%v\n", err)
	}

	insertNameSurf, err := insertFont.RenderUTF8Blended("Insert name", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("error loading font surface%v\n", err)
	}

	exitSurf, err := exitFont.RenderUTF8Blended("Press esc to exit game", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("error loading font surface%v\n", err)
	}

	defer titleSurf.Free()
	defer titleSurf2.Free()
	defer newGameSurf.Free()
	defer insertNameSurf.Free()
	defer exitSurf.Free()

	//prikaz
	g.titleTexture, err = g.Renderer.CreateTextureFromSurface(titleSurf)
	if err != nil {
		return fmt.Errorf("error loading font texture from surface%v\n", err)
	}
	g.title2Texture, err = g.Renderer.CreateTextureFromSurface(titleSurf2)
	if err != nil {
		return fmt.Errorf("error loading font texture from surface%v\n", err)
	}
	g.newGameTexture, err = g.Renderer.CreateTextureFromSurface(newGameSurf)
	if err != nil {
		return fmt.Errorf("error loading font texture from surface%v\n", err)
	}
	g.insertNameTexture, err = g.Renderer.CreateTextureFromSurface(insertNameSurf)
	if err != nil {
		return fmt.Errorf("error loading font texture from surface%v\n", err)
	}
	g.escape, err = g.Renderer.CreateTextureFromSurface(exitSurf)
	if err != nil {
		return fmt.Errorf("error loading font texture from surface%v\n", err)
	}

	//izvlacenje podataka za naslov i podnaslov
	_, _, g.titleW, g.titleH, err = g.titleTexture.Query()
	if err != nil {
		return fmt.Errorf("error loading title query %v\n", err)
	}

	_, _, g.newGameW, g.newGameH, err = g.newGameTexture.Query()
	if err != nil {
		return fmt.Errorf("error loading newgame query %v\n", err)
	}

	_, _, g.insertNameW, g.insertNameH, err = g.insertNameTexture.Query()
	if err != nil {
		return fmt.Errorf("error loading insert query %v\n", err)
	}

	_, _, g.escapeW, g.escapeH, err = g.escape.Query()
	if err != nil {
		return fmt.Errorf("error loading exit query %v\n", err)
	}

	//obicna dugmad
	g.playButton, err = g.LoadButton(
		"images/button.png", 200, 400, 200, 200,
	)
	g.SetButtonText(&g.playButton, "play", 20, sdl.Color{R: 0, G: 0, B: 0, A: 255})
	g.scoreButton, err = g.LoadButton(
		"images/button.png", 400, 400, 200, 200,
	)
	g.SetButtonText(&g.scoreButton, "score", 20, sdl.Color{R: 0, G: 0, B: 0, A: 255})

	//hover dugmad
	g.playHoverButton, err = g.LoadButton(
		"images/buttonHover.png", 200, 400, 200, 200,
	)
	g.SetButtonText(&g.playHoverButton, "play", 20, sdl.Color{R: 255, G: 255, B: 255, A: 255})

	g.playButton.hoverTexture = g.playHoverButton.texture
	g.playButton.hoverTextTexture = g.playHoverButton.textBtnTexture

	g.scoreHoverButton, err = g.LoadButton(
		"images/buttonHover.png", 400, 400, 200, 200,
	)
	g.SetButtonText(&g.scoreHoverButton, "score", 20, sdl.Color{R: 255, G: 255, B: 255, A: 255})

	g.scoreButton.hoverTexture = g.scoreHoverButton.texture
	g.scoreButton.hoverTextTexture = g.scoreHoverButton.textBtnTexture

	return err
}

func (g *StartGame) Close() {
	if g != nil {
		g.insertNameTexture.Destroy()
		g.insertNameTexture = nil
		g.titleTexture.Destroy()
		g.titleTexture = nil
		g.title2Texture.Destroy()
		g.title2Texture = nil
		g.escape.Destroy()
		g.escape = nil
		g.BackgroundImage.Destroy()
		g.BackgroundImage = nil
	}
}

// pokretanje igrice - game loop
func (g *StartGame) Run() ScreenID {
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

					case sdl.SCANCODE_BACKSPACE:
						if len(g.playerName) > 0 {
							g.playerName = g.playerName[:len(g.playerName)-1]
							g.renderText(g.playerName)
						}
						//nema potrbe da rucno pomeramo kursor nazad jer je kursor zapravo duzina teksta
					case sdl.SCANCODE_RETURN:
						return TransitionScreen
					}
				}
			case *sdl.TextInputEvent:
				g.playerName += e.GetText()
				g.renderText(g.playerName)
			case *sdl.MouseButtonEvent:
				if e.Type == sdl.MOUSEBUTTONDOWN {
					mouseX := e.X
					mouseY := e.Y

					if g.playButton.isClicked(mouseX, mouseY) {
						fmt.Println("play clicked")
						return TransitionScreen
					}

					if g.scoreButton.isClicked(mouseX, mouseY) {
						fmt.Println("score clicked")
						return CongratsScreen
					}
				}
			}
		}
		g.Renderer.Clear()                           //svaki frame pocinje 'praznim' ekranom i brise se sve sto je bilo sa prethodnog framea
		g.Renderer.Copy(g.BackgroundImage, nil, nil) //nil je cela tekstura
		//Copy(texture, šta uzeti iz slike, gde nacrtati)
		w := g.titleW
		h := g.titleH
		g.Renderer.Copy(g.titleTexture, nil, &sdl.Rect{
			X: (WindowWidth - w) / 2,
			Y: (WindowHeight-h)/2 - 45,
			W: w,
			H: h})

		g.Renderer.Copy(g.title2Texture, nil, &sdl.Rect{
			X: (WindowWidth - w) / 2 - 5,
			Y: (WindowHeight-h)/2 - 45,
			W: w,
			H: h})

		g.Renderer.Copy(g.newGameTexture, nil, &sdl.Rect{
			X: (WindowWidth - g.newGameW) / 2,
			Y: (WindowHeight- g.newGameH)/ 2 + 35 ,
			W: g.newGameW,
			H: g.newGameH})

		w = g.insertNameW
		h = g.insertNameH

		g.Renderer.Copy(g.insertNameTexture, nil, &sdl.Rect{
			X: (WindowWidth - w) / 2,
			Y: (WindowHeight - h) / 2 + 90,
			W: w,
			H: h,
		})

		//input box
		g.Renderer.SetDrawColor(255, 255, 255, 255)
		boxW := int32(400)
		boxH := int32(50)

		boxX := int32((WindowWidth - boxW) / 2)
		boxY := int32(410)

		g.Renderer.FillRect(&sdl.Rect{
			X: boxX,
			Y: boxY,
			W: boxW,
			H: boxH,
		})
		w, h = 0, 0
		if g.playerNameTexture != nil {
			//ako postoji tekst onda pozicioniramo trenutne promenljve na velicinu fonta za igraca
			w = g.playerW
			h = g.playerH

			g.Renderer.Copy(g.playerNameTexture, nil, &sdl.Rect{
				X: boxX + 10,
				Y: boxY + (boxH-h)/2,
				W: w,
				H: h,
			})
		}
		now := sdl.GetTicks64()
		if now-uint64(g.blink) > 500 {
			g.cursorVisible = !g.cursorVisible
			g.blink = now
		}

		if g.cursorVisible {
			cursorH := h
			if cursorH == 0 {
				cursorH = boxH - 10 //duzina kursora
			}
			cursorX := boxX + w + 10
			cursorY := boxY + (boxH-cursorH)/2
			g.Renderer.SetDrawColor(0, 0, 0, 255)

			g.Renderer.DrawLine(
				cursorX,
				cursorY,
				cursorX, //width nam je jedan piksel odnosno pozicija horizontalno
				cursorY+cursorH,
			)
		}

		mouseX, mouseY, _ := sdl.GetMouseState()
		g.playButton.hovered = g.playButton.isClicked(mouseX, mouseY)
		g.scoreButton.hovered = g.scoreButton.isClicked(mouseX, mouseY)

		//PLAY
		if g.playButton.hovered {
			g.Renderer.Copy(g.playButton.hoverTexture, nil, &g.playButton.rect)
		} else {
			g.Renderer.Copy(g.playButton.texture, nil, &g.playButton.rect)
		}

		//tekst:
		var tex *sdl.Texture
		var tw, th int32

		if g.playButton.hovered {
			tex = g.playButton.hoverTextTexture
		} else {
			tex = g.playButton.textBtnTexture
		}

		_, _, tw, th, _ = tex.Query()

		g.Renderer.Copy(tex, nil, &sdl.Rect{
			X: g.playButton.rect.X + (g.playButton.rect.W-tw)/2,
			Y: g.playButton.rect.Y + (g.playButton.rect.H-th)/2 + 30,
			W: tw,
			H: th,
		})

		//SCORE
		if g.scoreButton.hovered {
			g.Renderer.Copy(g.scoreButton.hoverTexture, nil, &g.scoreButton.rect)
		} else {
			g.Renderer.Copy(g.scoreButton.texture, nil, &g.scoreButton.rect)
		}
		//score tekst
		if g.scoreButton.hovered {
			tex = g.scoreButton.hoverTextTexture
		} else {
			tex = g.scoreButton.textBtnTexture
		}

		_, _, tw, th, _ = tex.Query()

		g.Renderer.Copy(tex, nil, &sdl.Rect{
			X: g.scoreButton.rect.X + (g.scoreButton.rect.W-tw)/2,
			Y: g.scoreButton.rect.Y + (g.scoreButton.rect.H-th)/2 + 30,
			W: tw,
			H: th,
		})

		g.Renderer.Copy(g.escape, nil, &sdl.Rect{
			X: (WindowWidth - g.newGameW) / 2 + 40,
			Y: (WindowHeight- g.newGameH)/ 2 + 290 ,
			W: g.escapeW,
			H: g.escapeH})

		g.Renderer.Present() //prikaze sve sto je nacrtano u ovom frameu
		sdl.Delay(16)        //koliko ce dugo da se prikaze igrica
	}
	return TransitionScreen
}
