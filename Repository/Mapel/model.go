package Mapel

import (
	"context"
	"kursus-online/Controller/Dto/Request"
	"kursus-online/Controller/Dto/Response"
)

type RepositoryMapel interface {
	CreateMapelRepository(ctx context.Context, param Request.CreateCoursus) (err error)
	CreateMapel(ctx context.Context, param Request.CreateMapel) (err error)
	CheckExistsMapelTitle(ctx context.Context, title string) (exists bool, err error)
	CheckExistsMapelId(ctx context.Context, id string) (exists bool, err error)
	UpdateMapel(ctx context.Context, param Request.UpdateMapel) (err error)
	DeleteMapel(ctx context.Context, id string) (err error)
	ListMapel(ctx context.Context, sortBy string) (res []Response.MapelList, err error)
	DetailMapel(ctx context.Context, id string) (res Response.MapelDetail, err error)
}

type Mapel struct{}

func NewRepository() RepositoryMapel {
	return &mapel{}
}
