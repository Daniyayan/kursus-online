package Coursus

import (
	"context"
	"kursus-online/Controller/Dto/Request"
)

type RepositoryCoursus interface {
	CreateCourse(ctx context.Context, param Request.CreateCoursus) (err error)
	//CreateCoursus(ctx context.Context, param Request.CreateCoursus) (err error)
	//CheckExistsCoursusTitle(ctx context.Context, title string) (exists bool, err error)
	//CheckExistsCoursusId(ctx context.Context, Id string) (exists bool, err error)
	//UpdateCoursus(ctx context.Context, param Request.UpdateCoursus) (err error)
	//DeleteCoursus(ctx context.Context, Id string) (err error)
	//ListCoursus(ctx context.Context, sortBy string) (res []Response.CoursusList, err error)
	//DetailCoursus(ctx context.Context, Id string) (res Response.CoursusDetail, err error)
}

// Define the Coursus struct
type coursus struct{}

// NewRepository returns an instance of RepositoryCoursus
func NewRepository() RepositoryCoursus {
	return &coursus{}
}
