package Product

import (
	"context"
	"kursus-online/Controller/Dto/Request"
	"kursus-online/Controller/Dto/Response"
)

type RepositoryProduct interface {
	CreateCourse(ctx context.Context, param Request.CreateCourse) (err error)
	CheckExistsCourseTitle(ctx context.Context, title string) (exists bool, err error)
	CheckExistsCourseId(ctx context.Context, id string) (exists bool, err error)
	UpdateCourse(ctx context.Context, param Request.UpdateCourse) (err error)
	DeleteCourse(ctx context.Context, id string) (err error)
	ListCourse(ctx context.Context, sortBy string) (res []Response.CourseList, err error)
	DetailCourse(ctx context.Context, id string) (res Response.CourseDetail, err error)
}

type Course struct{}

func NewRepository() RepositoryCourse {
	return &Course{}
}
