package Controller

type controllerInterface interface {
	CourseInterface
	UserInterface
}

type Controller struct {
	controllerInterface
}
