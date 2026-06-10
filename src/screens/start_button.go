package screens

//Fajl u kome se nalaze podaci vezani za rad sa dugmetom na pocetnom ekranu
import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

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
