package main

import (
    "fmt"
    "os"

    "github.com/veandco/go-sdl2/img"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/ttf"

    "bozanstveniPekar/screens"
    "bozanstveniPekar/levels"
)

const (
    windowTitle = "Bozanstveni pekar"
)

func initSDL() error {
    var sdlFlags uint32 = sdl.INIT_EVERYTHING
    imgFlags := img.INIT_JPG | img.INIT_PNG

    if err := sdl.Init(sdlFlags); err != nil {
        return fmt.Errorf("Error init sdl2 %v", err)
    }
    if err := img.Init(imgFlags); err != nil {
        return fmt.Errorf("Error init sdl img %v", err)
    }
    if err := ttf.Init(); err != nil {
        return fmt.Errorf("Error init sdl ttf %v", err)
    }
    return nil
}

func closeSDL() {
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

    screen := screens.StartScreen
    for screen != screens.ExitScreen {
        switch screen {
            case screens.StartScreen:
                start := screens.NewStartGame(engine)
                start.LoadMedia()
                screen = start.Run()

            case screens.GameOverScreen:
                gameOver := screens.NewGameOver(engine)
                gameOver.LoadMedia()
                screen = gameOver.Run()
            
            case screens.CongratsScreen:
                congrats := screens.NewCongrats(engine)
                congrats.LoadMedia()
                screen = congrats.Run()

            case screens.Level1Screen:
                lvl1 := levels.NewLevel1(engine)
                lvl1.LoadMedia()
                screen = lvl1.Run()
            
            case screens.Level2Screen:
                lvl2 := levels.NewLevel2(engine)
                lvl2.LoadMedia()
                screen = lvl2.Run()
            
            case screens.Level3Screen:
                lvl3 := levels.NewLevel3(engine)
                lvl3.LoadMedia()
                screen = lvl3.Run()
            
            case screens.Level4Screen:
                lvl4 := levels.NewLevel4(engine)
                lvl4.LoadMedia()
                screen = lvl4.Run()
            
            case screens.Level5Screen:
                lvl5 := levels.NewLevel5(engine)
                lvl5.LoadMedia()
                screen = lvl5.Run()
            
            case screens.Level6Screen:
                lvl6 := levels.NewLevel6(engine)
                lvl6.LoadMedia()
                screen = lvl6.Run()

            case screens.Level7Screen:
                lvl7 := levels.NewLevel7(engine)
                lvl7.LoadMedia()
                screen = lvl7.Run()
            
            case screens.Level8Screen:
                lvl8 := levels.NewLevel8(engine)
                lvl8.LoadMedia()
                screen = lvl8.Run()
            
            case screens.Level9Screen:
                lvl9 := levels.NewLevel9(engine)
                lvl9.LoadMedia()
                screen = lvl9.Run()
        }
    }
}