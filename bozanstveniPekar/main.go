package main

import (
    "fmt"
    "os"

    "github.com/veandco/go-sdl2/img"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/ttf"

    "bozanstveniPekar/screens"
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
        }
    }
}