package Controller

type controllerInterface interface {
	user_testInterface
}

type Controller struct {
	controllerInterface
}
