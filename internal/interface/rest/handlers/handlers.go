package handlers

type Handlers struct {
	ProjectHandler ProjectHandlerImpl
	TaskHandler    TaskHandlerImpl
}

// New should receive ProjectHandlerImpl and TaskHandlerImpl as parameters
// but to simplify, we would not do this now
func New() *Handlers {
	return &Handlers{}
}
