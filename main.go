package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/matf-pp/2026_Bozanstveni-pekar/src"
	"github.com/matf-pp/2026_Bozanstveni-pekar/src/screens"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Zakljucava izvrsavanje na glavnoj niti operativnog sistema, go automatski pokrece
func init() {
	runtime.LockOSThread()
}

const (
	windowTitle = "Bozanstveni pekar"
)

type LevelData struct {
	ImgPath    string
	NextScreen screens.ScreenID
	CircleVar  string
	LevelNum   int
	LevelStr   string
}

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

	var score int
	var ime string
	var levels = map[screens.ScreenID]LevelData{
		screens.Level1Screen: {"images/lvl1.png", screens.Level2Screen, "drugi", 2, "Pozuda"},
		screens.Level2Screen: {"images/lvl2.png", screens.Level3Screen, "treci", 3, "Prozdrljivost"},
		screens.Level3Screen: {"images/lvl3.png", screens.Level4Screen, "cetvrti", 4, "Pohlepa"},
		screens.Level4Screen: {"images/lvl4.png", screens.Level5Screen, "peti", 5, "Lenjost"},
		screens.Level5Screen: {"images/lvl5.png", screens.Level6Screen, "sesti", 6, "Jeres"},
		screens.Level6Screen: {"images/lvl6.png", screens.Level7Screen, "sedmi", 7, "Nasilje"},
		screens.Level7Screen: {"images/lvl7.png", screens.Level8Screen, "osmi", 8, "Prevara"},
		screens.Level8Screen: {"images/lvl8.png", screens.Level9Screen, "deveti", 9, "Izdaja"},
	}

	var levelVar screens.ScreenID
	var circleVar string
	var levelNum int
	var levelStr string

	screen := screens.StartScreen
	for screen != screens.ExitScreen {
		switch screen {
		case screens.StartScreen:
			score = 0
			start := screens.NewStartGame(engine)
			start.LoadMedia()
			screen = start.Run(&ime)
			start.Close()

			if screen == screens.TransitionScreen {
				levelVar, circleVar, levelNum, levelStr = screens.Level1Screen, "prvi", 1, "Limb"
			}

		case screens.GameOverScreen:
			gameOver := screens.NewGameOver(engine)
			gameOver.LoadMedia(score, ime)
			gameOver.CreateBlur()
			screen = gameOver.Run()
			gameOver.Close()

		case screens.CongratsScreen:
			congrats := screens.NewCongrats(engine)
			congrats.LoadMedia()
			congrats.CreateBlur()
			screen = congrats.Run()
			congrats.Close()

		case screens.Level9Screen:
			lvl9 := src.NewLevel(engine)
			lvl9.LoadMedia("images/lvl9.png")
			screen = lvl9.Run(screens.CongratsScreen, &score)
			lvl9.Close()

		case screens.TransitionScreen:
			transition := screens.NewTransition(engine, levelVar, circleVar, levelNum, levelStr)
			transition.LoadMedia()
			screen = transition.Run()
			transition.Close()

		//ostali nivoi
		default:
			if data, exists := levels[screen]; exists {
				lvl := src.NewLevel(engine)
				lvl.LoadMedia(data.ImgPath)
				screen = lvl.Run(screens.TransitionScreen, &score)
				lvl.Close()

				if screen == screens.TransitionScreen {
					levelVar = data.NextScreen
					circleVar = data.CircleVar
					levelNum = data.LevelNum
					levelStr = data.LevelStr
				}
			}
		}
	}
}
