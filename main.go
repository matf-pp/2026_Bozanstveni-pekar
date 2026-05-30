package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	"github.com/matf-pp/2026_Bozanstveni-pekar/src"
	"github.com/matf-pp/2026_Bozanstveni-pekar/src/screens"
)

const (
	windowTitle = "Bozanstveni pekar"
)

func initSDL() error {
	var sdlFlags uint32 = sdl.INIT_EVERYTHING
	imgFlags := img.INIT_JPG | img.INIT_PNG
	mixFlags := mix.INIT_OGG | mix.INIT_MP3

	if err := sdl.Init(sdlFlags); err != nil {
		return fmt.Errorf("Error init sdl2 %v", err)
	}
	if err := img.Init(imgFlags); err != nil {
		return fmt.Errorf("Error init sdl img %v", err)
	}
	if err := ttf.Init(); err != nil {
		return fmt.Errorf("Error init sdl ttf %v", err)
	}
	if err := mix.Init(mixFlags); err != nil {
		return fmt.Errorf("Error init sdl mix %v", err)
	}
	if err := mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE); err != nil {
		return fmt.Errorf("Error init sdl openaudio %v", err)
	}
	return nil
}

func closeSDL() {
	mix.CloseAudio()
	mix.Quit()
	ttf.Quit()
	img.Quit()
	sdl.Quit()
}

func main() {
	defer closeSDL()

	if err := initSDL(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	engine := screens.NewGame(windowTitle)
	defer engine.Close()

	if err := engine.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	var levelVar screens.ScreenID
	var circleVar string
	var levelNum int
	var levelStr string

	var score int
	var ime string

	screen := screens.StartScreen
	for screen != screens.ExitScreen {
		switch screen {
		case screens.StartScreen:
			start := screens.NewStartGame(engine)
			start.LoadMedia()
			screen = start.Run(&ime)
			if screen == screens.TransitionScreen {
				levelVar = screens.Level1Screen
				circleVar = "prvi"
				levelNum = 1
				levelStr = "Limb"
			}

		case screens.GameOverScreen:
			gameOver := screens.NewGameOver(engine)
			gameOver.LoadMedia(score, ime)
			gameOver.CreateBlur()
			screen = gameOver.Run()

		case screens.CongratsScreen:
			congrats := screens.NewCongrats(engine)
			congrats.LoadMedia()
			congrats.CreateBlur()
			screen = congrats.Run()

		case screens.Level1Screen:
			lvl1 := src.NewLevel(engine)
			lvl1.LoadMedia("images/lvl1.png")
			screen = lvl1.Run(screens.TransitionScreen, &score)
			if screen == screens.TransitionScreen {
				levelVar = screens.Level2Screen
				circleVar = "drugi"
				levelNum = 2
				levelStr = "Pozuda"
			}

		case screens.Level2Screen:
			lvl2 := src.NewLevel(engine)
			lvl2.LoadMedia("images/lvl2.png")
			screen = lvl2.Run(screens.TransitionScreen, &score)

			if screen == screens.TransitionScreen {
				levelVar = screens.Level3Screen
				circleVar = "treci"
				levelNum = 3
				levelStr = "Prozdrljivost"
			}

		case screens.Level3Screen:
			lvl3 := src.NewLevel(engine)
			lvl3.LoadMedia("images/lvl3.png")
			screen = lvl3.Run(screens.TransitionScreen, &score)
			if screen == screens.TransitionScreen {
				levelVar = screens.Level4Screen
				circleVar = "cetvrti"
				levelNum = 4
				levelStr = "Pohlepa"
			}

		case screens.Level4Screen:
			lvl4 := src.NewLevel(engine)
			lvl4.LoadMedia("images/lvl4.png")
			screen = lvl4.Run(screens.TransitionScreen, &score)
			if screen == screens.TransitionScreen {
				levelVar = screens.Level5Screen
				circleVar = "petii"
				levelNum = 5
				levelStr = "Lenjost"
			}

		case screens.Level5Screen:
			lvl5 := src.NewLevel(engine)
			lvl5.LoadMedia("images/lvl5.png")
			screen = lvl5.Run(screens.TransitionScreen, &score)
			if screen == screens.TransitionScreen {
				levelVar = screens.Level6Screen
				circleVar = "sesti"
				levelNum = 6
				levelStr = "Jeres"
			}

		case screens.Level6Screen:
			lvl6 := src.NewLevel(engine)
			lvl6.LoadMedia("images/lvl6.png")
			screen = lvl6.Run(screens.TransitionScreen, &score)
			if screen == screens.TransitionScreen {
				levelVar = screens.Level7Screen
				circleVar = "sedmi"
				levelNum = 7
				levelStr = "Nasilje"
			}

		case screens.Level7Screen:
			lvl7 := src.NewLevel(engine)
			lvl7.LoadMedia("images/lvl7.png")
			screen = lvl7.Run(screens.TransitionScreen, &score)
			if screen == screens.TransitionScreen {
				levelVar = screens.Level8Screen
				circleVar = "osmi"
				levelNum = 8
				levelStr = "Prevara"
			}

		case screens.Level8Screen:
			lvl8 := src.NewLevel(engine)
			lvl8.LoadMedia("images/lvl8.png")
			screen = lvl8.Run(screens.TransitionScreen, &score)
			if screen == screens.TransitionScreen {
				levelVar = screens.Level9Screen
				circleVar = "deveti"
				levelNum = 9
				levelStr = "Izdaja"
			}

		case screens.Level9Screen:

			lvl9 := src.NewLevel(engine)
			lvl9.LoadMedia("images/lvl9.png")
			screen = lvl9.Run(screens.CongratsScreen, &score)

		//NewTransition(game *Game, next ScreenID, circle string, level int, name string)
		case screens.TransitionScreen:
			transition := screens.NewTransition(engine, levelVar, circleVar, levelNum, levelStr)
			transition.LoadMedia()
			screen = transition.Run()
		}
	}
}
