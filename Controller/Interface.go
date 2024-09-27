package Controller

type controllerInterface interface {
	CoursusInterface
}

type Controller struct {
	controllerInterface
}
