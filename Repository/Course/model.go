package Course

import (
	"context"
	"kursus-online/Controller/Dto/Request"
)

type RepositoryCourse interface {
	CreateCourseRepository(ctx context.Context, param Request.CreateCourse) (err error)
}

type course struct{}

func NewRepository() RepositoryCourse {
	return &course{}
}
