package levels

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"

	//"2026_Bozanstveni-pekar/screens"
	"github.com/matf-pp/2026_Bozanstveni-pekar/screens"
)

type Level5 struct {
	*screens.Game
	screens.BaseGame

	dante *sdl.Texture
	danteW int32
	danteH int32

	goodTunnel *sdl.Texture
	goodTunnelW int32
	goodTunnelH int32

	badTunnel *sdl.Texture
	badTunnelW int32
	badTunnelH int32
}

func NewLevel5(game *screens.Game) *Level5 {
	return &Level5{
		Game: game,
	}
}

func (g *Level5) LoadMedia() error{
	var err error
	g.BaseGame.BackgroundImage, err = img.LoadTexture(g.Game.Renderer, "images/lvl5.png")
	if err != nil {
		return fmt.Errorf("error loading texture %v\n", err)
	}
	DanteTexture, err := img.LoadTexture(g.Game.Renderer, "images/Dante.png")
	g.dante = DanteTexture

	//_, _, danteW, danteH, err = g.dante.Query()

	if err!=nil {
		return err
	}

	GoodTunnelTexture, err := img.LoadTexture(g.Game.Renderer, "images/goodTunnel.png")
	g.goodTunnel = GoodTunnelTexture

	BadTunnelTexture, err := img.LoadTexture(g.Game.Renderer, "images/badTunnel.png")
	g.badTunnel = BadTunnelTexture
	g.Music, err = mix.LoadMUS("music/spaceroad.mp3")
	if err != nil{
		return fmt.Errorf("error loading music %v\n", err)
	}

	err = g.Music.Play(0)
	if err != nil {
		return fmt.Errorf("error playing music %v\n", err)
	}
	return err
}

func (g *Level5) Close() {
	if g!=nil {
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
	}
}

func (g *Level5) Run() screens.ScreenID {
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

		g.Game.Renderer.Copy(g.dante, nil, &sdl.Rect{
			X:50,Y:10,W:60,H:60,
		})

		g.Game.Renderer.Copy(g.goodTunnel,nil,&sdl.Rect{
			X:50,Y:500,W:105,H:105,
		})

		g.Game.Renderer.Copy(g.badTunnel,nil,&sdl.Rect{
			X:225,Y:500,W:100,H:100,
		})

		g.Game.Renderer.Copy(g.badTunnel,nil,&sdl.Rect{
			X:425,Y:500,W:100,H:100,
		})

		g.Game.Renderer.Copy(g.badTunnel,nil,&sdl.Rect{
			X:650,Y:500,W:100,H:100,
		})

		g.Game.Renderer.Present()
		sdl.Delay(16)
		
	}
	return screens.Level6Screen
}