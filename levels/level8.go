package levels

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"

	//"2026_Bozanstveni-pekar/screens"
	"github.com/matf-pp/2026_Bozanstveni-pekar/screens"
)

type Level8 struct {
	*screens.Game
	screens.BaseGame

	dante  *sdl.Texture
	danteW int32
	danteH int32

	goodTunnel  *sdl.Texture
	goodTunnelW int32
	goodTunnelH int32

	badTunnel  *sdl.Texture
	badTunnelW int32
	badTunnelH int32

	verticalPath  *sdl.Texture
	verticalPathW int32
	verticalPathH int32
}

func NewLevel8(game *screens.Game) *Level1 {
	return &Level1{
		Game: game,
	}
}

func (g *Level8) LoadMedia() error {
	var err error
	g.BaseGame.BackgroundImage, err = img.LoadTexture(g.Game.Renderer, "images/lvl8.png")
	if err != nil {
		return fmt.Errorf("error loading texture %v\n", err)
	}
	DanteTexture, err := img.LoadTexture(g.Game.Renderer, "images/Dante.png")
	g.dante = DanteTexture

	//_, _, danteW, danteH, err = g.dante.Query()

	if err != nil {
		return err
	}

	GoodTunnelTexture, err := img.LoadTexture(g.Game.Renderer, "images/goodTunnel.png")
	g.goodTunnel = GoodTunnelTexture

	BadTunnelTexture, err := img.LoadTexture(g.Game.Renderer, "images/badTunnel.png")
	g.badTunnel = BadTunnelTexture

	g.Music, err = mix.LoadMUS("music/spaceroad.mp3")
	if err != nil {
		return fmt.Errorf("error loading music %v\n", err)
	}

	err = g.Music.Play(0)
	if err != nil {
		return fmt.Errorf("error playing music %v\n", err)
	}
	return err
}

func (g *Level8) Close() {
	if g != nil {
		mix.HaltMusic()
		mix.HaltChannel(-1)
		g.Music.Free()
		g.Music = nil
		g.dante.Destroy()
		g.dante = nil
		g.goodTunnel.Destroy()
		g.goodTunnel = nil
		g.badTunnel.Destroy()
		g.badTunnel = nil
		g.verticalPath.Destroy()
		g.verticalPath = nil
	}
}

func (g *Level8) Run() screens.ScreenID {
	for true {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent: //iks na prozoru ili crtl q
				return screens.ExitScreen
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN { //pritisnuto dugme na tastaturi
					switch e.Keysym.Scancode { //koje tacno dugme je u pitanju
					case sdl.SCANCODE_ESCAPE: //esc
						return screens.ExitScreen
					}
				}
			}
		}
		g.Game.Renderer.Clear()
		g.Game.Renderer.Copy(g.BaseGame.BackgroundImage, nil, nil)

		g.Game.Renderer.Copy(g.verticalPath, nil, &sdl.Rect{
			X: 82, Y: 1, W: 32, H: 580,
		})

		g.Game.Renderer.Copy(g.verticalPath, nil, &sdl.Rect{
			X: 257, Y: 1, W: 32, H: 580,
		})

		g.Game.Renderer.Copy(g.verticalPath, nil, &sdl.Rect{
			X: 457, Y: 1, W: 32, H: 580,
		})

		g.Game.Renderer.Copy(g.verticalPath, nil, &sdl.Rect{
			X: 672, Y: 1, W: 32, H: 580,
		})

		g.Game.Renderer.Copy(g.dante, nil, &sdl.Rect{
			X: 66, Y: 10, W: 60, H: 60,
		})

		g.Game.Renderer.Copy(g.goodTunnel, nil, &sdl.Rect{
			X: 50, Y: 500, W: 105, H: 105,
		})

		g.Game.Renderer.Copy(g.badTunnel, nil, &sdl.Rect{
			X: 225, Y: 500, W: 100, H: 100,
		})

		g.Game.Renderer.Copy(g.badTunnel, nil, &sdl.Rect{
			X: 425, Y: 500, W: 100, H: 100,
		})

		g.Game.Renderer.Copy(g.badTunnel, nil, &sdl.Rect{
			X: 640, Y: 500, W: 100, H: 100,
		})

		g.Game.Renderer.Present()
		sdl.Delay(16)

	}
	return screens.Level9Screen
}
