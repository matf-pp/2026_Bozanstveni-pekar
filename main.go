package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/veandco/go-sdl2/mix"

	"github.com/matf-pp/2026_Bozanstveni-pekar/levels"
	"github.com/matf-pp/2026_Bozanstveni-pekar/screens"
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
	if err := mix.OpenAudio(mix.DEFAULT_FREQUENCY,mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE); err != nil {
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

	screen := screens.StartScreen
	for screen != screens.ExitScreen {
		switch screen {
		case screens.StartScreen:
			start := screens.NewStartGame(engine)
			start.LoadMedia()
			screen = start.Run()
			if screen == screens.TransitionScreen {
				levelVar = screens.Level1Screen
				circleVar = "prvi"
				levelNum = 1
				levelStr = "Limb"
			}

		case screens.GameOverScreen:
			gameOver := screens.NewGameOver(engine)
			gameOver.LoadMedia()
			gameOver.CreateBlur()
			screen = gameOver.Run()

		case screens.CongratsScreen:
			congrats := screens.NewCongrats(engine)
			congrats.LoadMedia()
			congrats.CreateBlur()
			screen = congrats.Run()

		case screens.Level1Screen:
			levelVar = screens.Level1Screen
			circleVar = "prvi"
			levelNum = 1
			levelStr = "Limb"
			lvl1 := levels.NewLevel1(engine)
			lvl1.LoadMedia()
			screen = lvl1.Run()

			if screen == screens.TransitionScreen {
				levelVar = screens.Level2Screen
				circleVar = "drugi"
				levelNum = 2
				levelStr = "Pozuda"
			}

		case screens.Level2Screen:
			levelVar = screens.Level2Screen
			circleVar = "drugi"
			levelNum = 2
			levelStr = "Pozuda"
			lvl2 := levels.NewLevel2(engine)
			lvl2.LoadMedia()
			screen = lvl2.Run()

			if screen == screens.TransitionScreen {
				levelVar = screens.Level3Screen
				circleVar = "treci"
				levelNum = 3
				levelStr = "Prozdrljivost"
			}

		case screens.Level3Screen:
			levelVar = screens.Level3Screen
			circleVar = "treci"
			levelNum = 3
			levelStr = "Prozdrljivost"
			lvl3 := levels.NewLevel3(engine)
			lvl3.LoadMedia()
			screen = lvl3.Run()
			if screen == screens.TransitionScreen {
				levelVar = screens.Level4Screen
				circleVar = "cetvrti"
				levelNum = 4
				levelStr = "Pohlepa"
			}

		case screens.Level4Screen:
			levelVar = screens.Level4Screen
			circleVar = "cetvrti"
			levelNum = 4
			levelStr = "Pohlepa"
			lvl4 := levels.NewLevel4(engine)
			lvl4.LoadMedia()
			screen = lvl4.Run()
			if screen == screens.TransitionScreen {
				levelVar = screens.Level5Screen
				circleVar = "petii"
				levelNum = 5
				levelStr = "Lenjost"
			}

		case screens.Level5Screen:
			levelVar = screens.Level5Screen
			circleVar = "peti"
			levelNum = 5
			levelStr = "Lenjost"
			lvl5 := levels.NewLevel5(engine)
			lvl5.LoadMedia()
			screen = lvl5.Run()
			if screen == screens.TransitionScreen {
				levelVar = screens.Level6Screen
				circleVar = "sesti"
				levelNum = 6
				levelStr = "Jeres"
			}

		case screens.Level6Screen:
			levelVar = screens.Level6Screen
			circleVar = "sesti"
			levelNum = 6
			levelStr = "Jeres"
			lvl6 := levels.NewLevel6(engine)
			lvl6.LoadMedia()
			screen = lvl6.Run()
			if screen == screens.TransitionScreen {
				levelVar = screens.Level7Screen
				circleVar = "sedmi"
				levelNum = 7
				levelStr = "Nasilje"
			}

		case screens.Level7Screen:
			levelVar = screens.Level7Screen
			circleVar = "sedmi"
			levelNum = 7
			levelStr = "Nasilje"
			lvl7 := levels.NewLevel7(engine)
			lvl7.LoadMedia()
			screen = lvl7.Run()
			if screen == screens.TransitionScreen {
				levelVar = screens.Level8Screen
				circleVar = "osmi"
				levelNum = 8
				levelStr = "Prevara"
			}

		case screens.Level8Screen:
			levelVar = screens.Level8Screen
			circleVar = "osmi"
			levelNum = 8
			levelStr = "Prevara"
			lvl8 := levels.NewLevel8(engine)
			lvl8.LoadMedia()
			screen = lvl8.Run()
			if screen == screens.TransitionScreen {
				levelVar = screens.Level9Screen
				circleVar = "deveti"
				levelNum = 9
				levelStr = "Izdaja"
			}

		case screens.Level9Screen:
			levelVar = screens.Level9Screen
			circleVar = "deveti"
			levelNum = 9
			levelStr = "Izdaja"
			lvl9 := levels.NewLevel9(engine)
			lvl9.LoadMedia()
			screen = lvl9.Run()

		//NewTransition(game *Game, next ScreenID, circle string, level int, name string)
		case screens.TransitionScreen:
			transition := screens.NewTransition(engine, levelVar, circleVar, levelNum, levelStr)
			transition.LoadMedia()
			screen = transition.Run()
		}
	}
}