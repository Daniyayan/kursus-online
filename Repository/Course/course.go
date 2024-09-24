package Course

import (
	"context"
	"kursus-online/Config"
	"kursus-online/Controller/Dto/Request"
)

func (r course) CreateCourseRepository(ctx context.Context, param Request.CreateCourse) (err error) {
	query := `INSERT INTO courses (name, description, duration, price) VALUES (?, ?, ?, ?)`
	_, err = Config.DATABASE_MAIN.Get().ExecContext(ctx, query, param.Name, param.Description, param.Duration, param.Price)
	if err != nil {
		return
	}

	return
}
