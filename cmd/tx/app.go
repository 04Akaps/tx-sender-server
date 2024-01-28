package tx

type App struct {
	stop chan struct{}
}

func NewTxApp() *App {
	a := &App{
		stop: make(chan struct{}),
	}

	return a
}

func (a *App) StartTxApp() {

}
