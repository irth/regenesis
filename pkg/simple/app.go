package simple

import (
	"fmt"
)

type App struct {
	currentScene *Scene
	nextScene    *Scene
	renderLock   OneAtATime
}

func NewApp(scene *Scene) *App {
	return &App{
		nextScene: scene,
	}
}

func (a *App) Render() error {
	if !a.renderLock.Lock() {
		return fmt.Errorf("you cannot render while another rendering is in progress")
	}
	defer a.renderLock.Unlock()

	a.currentScene = a.nextScene
	eventHandlers, err := a.currentScene.Render()
	if err != nil {
		return fmt.Errorf("while rendering the scene: %w", err)
	}

	for _, eventHandler := range eventHandlers {
		err = eventHandler(a)
		if err != nil {
			return fmt.Errorf("while running an event handler: %w", err)
		}
	}

	return nil
}

func (a *App) NextScene(scene *Scene) {
	a.nextScene = scene
}

func (a *App) RunForever() error {
	for {
		err := a.Render()
		if err != nil {
			return err
		}
	}
}
