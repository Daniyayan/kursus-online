package Coursus

import (
	"context"
	"kursus-online/Config"
	"kursus-online/Controller/Dto/Request"
	"kursus-online/Controller/Dto/Response"
	"time"
)

//type course struct{}

//func NewRepository() *course {
//	return &course{}
//}

// Create Course
func (c *Coursus) CreateCourse(ctx context.Context, param Request.CreateCoursus) (err error) {
	query := `INSERT INTO course (id_coursus, nama, description, duration, periode, price) VALUES (?, ?, ?, ?, ?, ?)`
	if _, err = Config.DATABASE_MAIN.Get().ExecContext(ctx, query, param.IdCoursus, param.Nama, param.Description, param.Duration, param.Period, param.Price); err != nil {
		return err
	}
	return
}

// Check if Course Name Exists
func (c *Coursus) CheckExistsCourseName(ctx context.Context, nama string) (exists bool, err error) {
	query := `SELECT EXISTS (SELECT 1 FROM course WHERE LOWER(nama) = LOWER(?)) AS "exists"`
	err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, nama).Scan(&exists)
	return
}

// Check if Course ID Exists
func (c *Coursus) CheckExistsCourseId(ctx context.Context, id int) (exists bool, err error) {
	query := `SELECT EXISTS (SELECT 1 FROM course WHERE id_coursus = ?) AS "exists"`
	err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, id).Scan(&exists)
	return
}

// Update Course
func (c *Coursus) UpdateCourse(ctx context.Context, param Request.UpdateCoursus) (err error) {
	query := `UPDATE course SET nama = ?, description = ?, duration = ?, periode = ?, price = ?, updated_at = ? WHERE id_coursus = ?`
	if _, err = Config.DATABASE_MAIN.Get().ExecContext(ctx, query, param.Nama, param.Description, param.Duration, param.Period, param.Price, time.Now(), param.IdCoursus); err != nil {
		return err
	}
	return
}

// Delete Course
func (c *Coursus) DeleteCourse(ctx context.Context, id int) (err error) {
	query := `UPDATE course SET updated_at = ? WHERE id_coursus = ?`
	if _, err = Config.DATABASE_MAIN.Get().ExecContext(ctx, query, time.Now(), id); err != nil {
		return err
	}
	return
}

// List Courses
func (c *Coursus) ListCourse(ctx context.Context, sortBy string) (res []Response.CoursusList, err error) {
	var (
		data       Response.CoursusList
		connection = Config.DATABASE_MAIN.Get()
	)

	query := `SELECT id_coursus, nama, description, duration, periode, price FROM course ORDER BY created_at DESC LIMIT 10`
	if sortBy == "nama" {
		query = `SELECT id_coursus, nama, description, duration, periode, price FROM course ORDER BY nama ASC LIMIT 10`
	} else if sortBy == "price" {
		query = `SELECT id_coursus, nama, description, duration, periode, price FROM course ORDER BY price ASC LIMIT 10`
	}

	rows, err := connection.QueryContext(ctx, query)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&data.IdCoursus, &data.Nama, &data.Description, &data.Duration, &data.Period, &data.Price); err != nil {
			return
		}
		res = append(res, data)
	}

	return
}

// Get Course Details by ID
func (c *Coursus) DetailCourse(ctx context.Context, id int) (res Response.CoursusDetail, err error) {
	query := `SELECT id_coursus, nama, description, duration, periode, price, updated_at FROM course WHERE id_coursus = ?`
	if err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, id).Scan(&res.IdCoursus, &res.Nama, &res.Description, &res.Duration, &res.Period, &res.Price, &res.UpdatedAt); err != nil {
		return
	}
	return
}
