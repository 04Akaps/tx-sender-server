package auth

type App struct {
	stop chan struct{}
}

func NewAuthApp() *App {
	a := &App{
		stop: make(chan struct{}),
	}

	return a
}

func (a *App) StartAuthApp() {

}
