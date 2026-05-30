package screens

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
)

type Game struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
}

func NewGame(title string) *Game {
	return &Game{}
}

func (g *Game) Init() error {
	var err error
	//naslov, hor i vert pozicija, sirina i visina u pikselima, prikazi odmah prozor(ako zelimo da ga povecamo -> |sdl.WINDOW_RESIZABLE)
	if g.Window, err = sdl.CreateWindow(WindowTitle, sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED, WindowWidth, WindowHeight, sdl.WINDOW_SHOWN); err != nil {
		return fmt.Errorf("Error creating window %v", err)
	}
	//prozor, automatski renderer (0-2 su konkretni drajveri poput opengla itd), primena gpua
	if g.Renderer, err = sdl.CreateRenderer(g.Window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		return fmt.Errorf("Error creating render %v", err)
	}
	iconSurf, err := img.Load("images/leb.png")
	if err != nil {
		return fmt.Errorf("Error creating icon srf %v", err)
	}
	defer iconSurf.Free()
	g.Window.SetIcon(iconSurf)
	return err
}

func (g *Game) Close() {
	if g!=nil {
        g.Renderer.Destroy()
        g.Renderer = nil
        g.Window.Destroy()
        g.Window = nil
	}
}