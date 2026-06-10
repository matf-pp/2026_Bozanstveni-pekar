package screens //paket u kom se nalaze ekrani igre

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// globalne konstante za prozor
const (
	WindowWidth  = 800
	WindowHeight = 600
	WindowTitle  = "Bozanstveni pekar"
)

type Screen interface {
	LoadMedia() error
	Run() ScreenID // vraca prelaz na sledeci ekran
	Events() ScreenID
	Close()
}

// Tip za
type ScreenID int

const (
	StartScreen ScreenID = iota //iota se koristi za deklarisanje niza konstanti koje se automatski uvećavaju
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
	ClickSound      *mix.Chunk
	Music           *mix.Music
}

type StartGame struct {
	*Game
	BaseGame
	titleTexture  *sdl.Texture
	title2Texture *sdl.Texture

	newGameTexture  *sdl.Texture
	newGame2Texture *sdl.Texture

	insertNameTexture *sdl.Texture
	playerName        string
	playerNameTexture *sdl.Texture

	playButton      Button
	playHoverButton Button

	bottomTextEscTexture *sdl.Texture
	bottomTextEscW       int32
	bottomTextEscH       int32

	bottomTextEnterTexture *sdl.Texture
	bottomTextEnterW       int32
	bottomTextEnterH       int32

	cursorVisible bool
	blink         uint64

	titleW int32
	titleH int32

	newGameW int32
	newGameH int32

	insertNameW int32
	insertNameH int32
	playerW     int32
	playerH     int32

	playTextW      int32
	playTextH      int32
	playHoverTextW int32
	playHoverTextH int32
}

func NewStartGame(game *Game) *StartGame {
	return &StartGame{
		Game: game,
	}
}

// Pomoćna metoda za učitavanje fonta
func (g *Game) LoadFont(size int) (*ttf.Font, error) {
	return ttf.OpenFont(
		"fonts/PixelifySans-VariableFont_wght.ttf",
		size,
	)
}

// Pomoćna metoda za učitavanje teksta u textboxu
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
		return fmt.Errorf("error loading font %v", err)
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
		return fmt.Errorf("error loading player query %v", err)
	}

	return err
}

// Pomoćna metoda za učitavanje dugmeta
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

// Učitavanje sadržaja početnog ekrana
func (g *StartGame) LoadMedia() error {
	var err error
	if g.BaseGame.BackgroundImage, err = img.LoadTexture(g.Renderer, "images/background.jpg"); err != nil {
		return fmt.Errorf("error loading background texture %v", err)
	}
	titleFont, err := ttf.OpenFont("fonts/Early_GameBoy.ttf", 45)
	if err != nil {
		return err
	}
	newGameFont, _ := g.LoadFont(40)
	newGame2Font, _ := g.LoadFont(40)
	insertFont, _ := g.LoadFont(25)
	exitFont, _ := g.LoadFont(20)
	defer titleFont.Close()
	defer newGameFont.Close()
	defer newGame2Font.Close()
	defer insertFont.Close()
	defer exitFont.Close()

	//kreiramo tekst u sliku koju zelimo da nacrtamo
	titleSurf, err := titleFont.RenderUTF8Blended("Bozanstveni pekar", sdl.Color{R: 69, G: 40, B: 18, A: 255})
	if err != nil {
		return fmt.Errorf("error loading title font surface %v", err)
	}
	titleSurf2, err := titleFont.RenderUTF8Blended("Bozanstveni pekar", sdl.Color{R: 220, G: 162, B: 94, A: 255})
	if err != nil {
		return fmt.Errorf("error loading title font surface %v", err)
	}

	newGameSurf, err := newGameFont.RenderUTF8Blended("Start game", sdl.Color{R: 69, G: 40, B: 18, A: 255})
	if err != nil {
		return fmt.Errorf("error loading newgame font surface %v", err)
	}
	newGame2Surf, err := newGameFont.RenderUTF8Blended("Start game", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("error loading newgame font surface %v", err)
	}

	insertNameSurf, err := insertFont.RenderUTF8Blended("Insert name", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("error loading insert font surface %v", err)
	}

	exitSurf, err := exitFont.RenderUTF8Blended("Press esc to exit", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("error loading exit font surface %v", err)
	}
	enterSurf, err := exitFont.RenderUTF8Blended("Press enter to play", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("error loading enter font surface %v", err)
	}

	defer titleSurf.Free()
	defer titleSurf2.Free()
	defer newGameSurf.Free()
	defer newGame2Surf.Free()
	defer insertNameSurf.Free()
	defer exitSurf.Free()
	defer enterSurf.Free()

	//Prikaz tekstura
	g.titleTexture, err = g.Renderer.CreateTextureFromSurface(titleSurf)
	if err != nil {
		return fmt.Errorf("error creating title font texture from surface %v", err)
	}
	g.title2Texture, err = g.Renderer.CreateTextureFromSurface(titleSurf2)
	if err != nil {
		return fmt.Errorf("error creating title font texture from surface %v", err)
	}
	g.newGameTexture, err = g.Renderer.CreateTextureFromSurface(newGameSurf)
	if err != nil {
		return fmt.Errorf("error creating newgame font texture from surface %v", err)
	}
	g.newGame2Texture, err = g.Renderer.CreateTextureFromSurface(newGame2Surf)
	if err != nil {
		return fmt.Errorf("error creating newgame font texture from surface %v", err)
	}
	g.insertNameTexture, err = g.Renderer.CreateTextureFromSurface(insertNameSurf)
	if err != nil {
		return fmt.Errorf("error creating insertname font texture from surface %v", err)
	}
	g.bottomTextEscTexture, err = g.Renderer.CreateTextureFromSurface(exitSurf)
	if err != nil {
		return fmt.Errorf("error creating bottomtext esc font texture from surface %v", err)
	}
	g.bottomTextEnterTexture, err = g.Renderer.CreateTextureFromSurface(enterSurf)
	if err != nil {
		return fmt.Errorf("error creating bottomtext enter font texture from surface %v", err)
	}

	//Izvlacenje podataka za naslov i podnaslov
	_, _, g.titleW, g.titleH, err = g.titleTexture.Query()
	if err != nil {
		return fmt.Errorf("error loading title query %v", err)
	}
	_, _, g.newGameW, g.newGameH, err = g.newGameTexture.Query()
	if err != nil {
		return fmt.Errorf("error loading newgame query %v", err)
	}
	_, _, g.insertNameW, g.insertNameH, err = g.insertNameTexture.Query()
	if err != nil {
		return fmt.Errorf("error loading insert query %v", err)
	}
	_, _, g.bottomTextEscW, g.bottomTextEscH, err = g.bottomTextEscTexture.Query()
	if err != nil {
		return fmt.Errorf("error loading exit query %v", err)
	}
	_, _, g.bottomTextEnterW, g.bottomTextEnterH, err = g.bottomTextEnterTexture.Query()
	if err != nil {
		return fmt.Errorf("error loading exit query %v", err)
	}

	//obicna dugmad
	g.playButton, err = g.LoadButton(
		"images/button.png", 325, 485, 150, 55,
	)
	g.SetButtonText(&g.playButton, "play", 20, sdl.Color{R: 0, G: 0, B: 0, A: 255})

	//hover dugmad
	g.playHoverButton, err = g.LoadButton(
		"images/buttonHover.png", 325, 485, 150, 55,
	)
	g.SetButtonText(&g.playHoverButton, "play", 20, sdl.Color{R: 255, G: 255, B: 255, A: 255})

	g.playButton.hoverTexture = g.playHoverButton.texture
	g.playButton.hoverTextTexture = g.playHoverButton.textBtnTexture

	_, _, g.playTextW, g.playTextH, _ = g.playButton.textBtnTexture.Query()
	_, _, g.playHoverTextW, g.playHoverTextH, _ = g.playHoverButton.textBtnTexture.Query()

	//Zvuk za klik na dugme
	g.ClickSound, err = mix.LoadWAV("sounds/click.mp3")
	if err != nil {
		return fmt.Errorf("error loading chunk %v", err)
	}
	//Pozadinska muzika; učitavanje
	g.Music, err = mix.LoadMUS("music/thunderstorm.mp3")
	if err != nil {
		return fmt.Errorf("error loading music %v", err)
	}
	//Pokretanje pozadinske muzike
	err = g.Music.Play(-1)
	if err != nil {
		return fmt.Errorf("error playing music %v", err)
	}
	return err
}

// Zatvaranje početnog ekrana
func (g *StartGame) Close() {
	if g != nil {
		mix.HaltMusic()
		mix.HaltChannel(-1) //zaustavljamo sve aktivne kanale odjednom, prekidamo sve sto rn svira

		if g.Music != nil {
			g.Music.Free()
			g.Music = nil
		}

		if g.ClickSound != nil {
			g.ClickSound.Free()
			g.ClickSound = nil
		}

		if g.insertNameTexture != nil {
			g.insertNameTexture.Destroy()
			g.insertNameTexture = nil
		}
		if g.titleTexture != nil {
			g.titleTexture.Destroy()
			g.titleTexture = nil
		}
		if g.title2Texture != nil {
			g.title2Texture.Destroy()
			g.title2Texture = nil
		}

		if g.newGameTexture != nil {
			g.newGameTexture.Destroy()
			g.newGameTexture = nil
		}

		if g.newGame2Texture != nil {
			g.newGame2Texture.Destroy()
			g.newGame2Texture = nil
		}

		if g.bottomTextEscTexture != nil {
			g.bottomTextEscTexture.Destroy()
			g.bottomTextEscTexture = nil
		}
		if g.bottomTextEnterTexture != nil {
			g.bottomTextEnterTexture.Destroy()
			g.bottomTextEnterTexture = nil
		}

		if g.playButton.texture != nil {
			g.playButton.texture.Destroy()
			g.playButton.texture = nil
		}
		if g.playHoverButton.texture != nil {
			g.playHoverButton.texture.Destroy()
			g.playHoverButton.texture = nil
		}

		if g.playerNameTexture != nil {
			g.playerNameTexture.Destroy()
			g.playerNameTexture = nil
		}

		if g.playButton.textBtnTexture != nil {
			g.playButton.textBtnTexture.Destroy()
			g.playButton.textBtnTexture = nil
		}

		if g.playHoverButton.textBtnTexture != nil {
			g.playHoverButton.textBtnTexture.Destroy()
			g.playHoverButton.textBtnTexture = nil
		}

		if g.BackgroundImage != nil {
			g.BackgroundImage.Destroy()
			g.BackgroundImage = nil
		}

		g.playButton.hoverTexture = nil
		g.playButton.hoverTextTexture = nil
		g.playHoverButton.hoverTexture = nil
	}
}

// metoda za Eventove
func (g *StartGame) Events(ime *string, brojacSlova *int32) ScreenID {
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
						*brojacSlova--
					}
					//nema potrbe da rucno pomeramo kursor nazad jer je kursor zapravo duzina teksta
				case sdl.SCANCODE_RETURN:
					if g.playerName != "" {
						*ime = g.playerName
						return TransitionScreen
					}
				}
			}
		case *sdl.TextInputEvent:
			if *brojacSlova < 6 {
				g.playerName += e.GetText()
				g.renderText(g.playerName)
				*brojacSlova++
			}
		case *sdl.MouseButtonEvent:
			if e.Type == sdl.MOUSEBUTTONDOWN {
				mouseX := e.X
				mouseY := e.Y

				if g.playButton.isClicked(mouseX, mouseY) && g.playerName != "" {
					g.ClickSound.Play(-1, 0)
					fmt.Println("play clicked")
					*ime = g.playerName
					return TransitionScreen
				}
			}
		}
	}
	return StartScreen
}

// Pokretanje početnog ekrana
func (g *StartGame) Run(ime *string) ScreenID {
	var brojacSlova int32
	for true {
		nextScreen := g.Events(ime, &brojacSlova)
		if nextScreen != StartScreen {
			return nextScreen
		}
		g.Renderer.Clear()                           //svaki frame pocinje 'praznim' ekranom i brise se sve sto je bilo sa prethodnog framea
		g.Renderer.Copy(g.BackgroundImage, nil, nil) //nil je cela tekstura
		//Copy(texture, šta uzeti iz slike, gde nacrtati)

		g.Renderer.Copy(g.titleTexture, nil, &sdl.Rect{
			X: (WindowWidth - g.titleW) / 2,
			Y: (WindowHeight - g.titleH) / 2,
			W: g.titleW,
			H: g.titleH,
		})

		g.Renderer.Copy(g.title2Texture, nil, &sdl.Rect{
			X: (WindowWidth-g.titleW)/2 - 7,
			Y: (WindowHeight - g.titleH) / 2,
			W: g.titleW,
			H: g.titleH,
		})

		g.Renderer.Copy(g.newGameTexture, nil, &sdl.Rect{
			X: (WindowWidth - g.newGameW) / 2,
			Y: (WindowHeight-g.newGameH)/2 + 50,
			W: g.newGameW,
			H: g.newGameH,
		})

		g.Renderer.Copy(g.newGame2Texture, nil, &sdl.Rect{
			X: (WindowWidth-g.newGameW)/2 - 5,
			Y: (WindowHeight-g.newGameH)/2 + 50,
			W: g.newGameW,
			H: g.newGameH,
		})

		g.Renderer.Copy(g.insertNameTexture, nil, &sdl.Rect{
			X: (WindowWidth - g.insertNameW) / 2,
			Y: (WindowHeight-g.insertNameH)/2 + 90,
			W: g.insertNameW,
			H: g.insertNameH,
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
		w := int32(0)
		h := int32(0)
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

		var bottomText *sdl.Texture

		//PLAY
		if g.playButton.hovered {
			g.Renderer.Copy(g.playButton.hoverTexture, nil, &g.playButton.rect)
			bottomText = g.bottomTextEnterTexture
		} else {
			g.Renderer.Copy(g.playButton.texture, nil, &g.playButton.rect)
			bottomText = g.bottomTextEscTexture
		}

		//tekst:
		var tex *sdl.Texture
		var tw, th int32

		if g.playButton.hovered {
			tex = g.playButton.hoverTextTexture
			tw = g.playHoverTextW
			th = g.playHoverTextH
		} else {
			tex = g.playButton.textBtnTexture
			tw = g.playTextW
			th = g.playTextH
		}

		g.Renderer.Copy(tex, nil, &sdl.Rect{
			X: g.playButton.rect.X + (g.playButton.rect.W-tw)/2,
			Y: g.playButton.rect.Y + (g.playButton.rect.H-th)/2,
			W: tw,
			H: th,
		})

		btw := g.bottomTextEscW
		bth := g.bottomTextEscH

		if g.playButton.hovered {
			btw = g.bottomTextEnterW
			bth = g.bottomTextEnterH
		}

		g.Renderer.Copy(bottomText, nil, &sdl.Rect{
			X: (WindowWidth - btw) / 2,
			Y: (WindowHeight-bth)/2 + 280,
			W: btw,
			H: bth})

		g.Renderer.Present() //prikaze sve sto je nacrtano u ovom frameu
		sdl.Delay(16)        //koliko ce dugo da se prikaze igrica
	}
	return TransitionScreen
}
