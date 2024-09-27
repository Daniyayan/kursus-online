package Coursus

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"kursus-online/Config"
	"kursus-online/Controller/Dto/Request"
	"kursus-online/Controller/Dto/Response"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func ConnectionMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println(err)
	}
	return db, mock
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

// Test CreateCoursus Success
func TestCreateCoursusSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = Request.CreateCoursus{
		Nama:        "Kursus Go Programming",
		Description: "Belajar Go Programming",
		Duration:    "2023-12-01",
		Period:      "2023",
		Price:       1500000,
	}

	query := `INSERT INTO course (nama, description, duration, periode, price) VALUES (?, ?, ?, ?, ?)`
	mock.ExpectExec(query).WithArgs(paramsMock.Nama, paramsMock.Description, paramsMock.Duration, paramsMock.Period, paramsMock.Price).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().CreateCoursus(context.Background(), paramsMock)
	assert.NoError(t, err)
}

// Test CreateCourse Failure
func TestCreateCoursusFailure(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = Request.CreateCoursus{
		Nama:        "Kursus Python Programming",
		Description: "Belajar Python",
		Duration:    "2023-12-01",
		Period:      "2023",
		Price:       0, // Harga nol untuk memicu kegagalan
	}

	query := `INSERT INTO course (nama, description, duration, periode, price) VALUES (?, ?, ?, ?, ?)`
	mock.ExpectExec(query).WithArgs(paramsMock.Nama, paramsMock.Description, paramsMock.Duration, paramsMock.Period, paramsMock.Price).WillReturnResult(
		sqlmock.NewResult(0, 0))

	err := NewRepository().CreateCoursus(context.Background(), paramsMock)
	assert.Error(t, err)
}

// Test ListCourse Success
func TestListCoursusSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var resMock = []Response.CoursusList{
		Response.CoursusList{
			IdCoursus:   1,
			Nama:        "Kursus Go Programming",
			Description: "Belajar Go Programming",
			Duration:    "2023-12-01",
			Period:      "2023",
			Price:       1500000,
		},
	}

	query := `SELECT id_coursus, nama, description, duration, periode, price FROM course WHERE created_at IS NOT NULL ORDER BY created_at DESC LIMIT 10`
	rows := sqlmock.NewRows([]string{"id_coursus", "nama", "description", "duration", "periode", "price"}).
		AddRow(resMock[0].IdCoursus, resMock[0].Nama, resMock[0].Description, resMock[0].Duration, resMock[0].Period, resMock[0].Price)

	mock.ExpectQuery(query).WillReturnRows(rows)

	res, err := NewRepository().ListCoursus(context.Background(), "id_coursus")
	assert.NoError(t, err)
	assert.Equal(t, resMock, res)
}

// Test CheckExistsCoursusNama Success
func TestCheckExistsCoursusNamaSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramMock = "Kursus Go Programming"
	var resMock bool

	query := `SELECT EXISTS (SELECT 1 FROM course WHERE LOWER(nama) = LOWER(?)) AS "exists"`

	rows := sqlmock.NewRows([]string{"exists"}).AddRow(resMock)
	mock.ExpectQuery(query).WithArgs(paramMock).WillReturnRows(rows)

	res, err := NewRepository().CheckExistsCoursusTitle(context.Background(), paramMock)
	assert.NoError(t, err)
	assert.Equal(t, resMock, res)
}

// Test UpdateCoursus Success
func TestUpdateCoursusSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = Request.UpdateCoursus{
		IdCoursus:   1,
		Nama:        "Kursus Go Programming Lanjutan",
		Description: "Go Programming Lanjutan",
		Duration:    "2024-01-01",
		Period:      "2024",
		Price:       2500000,
	}

	query := `UPDATE course SET nama = ?, description = ?, duration = ?, periode = ?, price = ?, updated_at = ? WHERE id_coursus = ?`
	mock.ExpectExec(query).WithArgs(paramsMock.Nama, paramsMock.Description, paramsMock.Duration, paramsMock.Period, paramsMock.Price, AnyTime{}, paramsMock.IdCoursus).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().UpdateCoursus(context.Background(), paramsMock)
	assert.NoError(t, err)
}

// Test DeleteCoursus Success
func TestDeleteCoursusSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = "1"

	query := `UPDATE course SET deleted_at = ? WHERE id_coursus = ?`
	mock.ExpectExec(query).WithArgs(AnyTime{}, paramsMock).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().DeleteCoursus(context.Background(), paramsMock)
	assert.NoError(t, err)
}

// Test DetailCoursus Success
func TestDetailCoursusSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var updatedAtValue = "2023-12-19 16:47:51"
	var paramMock = "1"
	var resMock = Response.CoursusDetail{
		IdCoursus:   1,
		Nama:        "Kursus Go Programming",
		Description: "Belajar Go Programming",
		Duration:    "2023-12-01",
		Period:      "2023",
		Price:       1500000,
		UpdatedAt:   &updatedAtValue,
	}

	query := `SELECT id_coursus, nama, description, duration, periode, price, updated_at FROM course WHERE id_coursus = ? AND deleted_at IS NULL`
	rows := sqlmock.NewRows([]string{"id_coursus", "nama", "description", "duration", "periode", "price", "updated_at"}).AddRow(resMock.IdCoursus,
		resMock.Nama, resMock.Description, resMock.Duration, resMock.Period, resMock.Price, resMock.UpdatedAt)
	mock.ExpectQuery(query).WithArgs(paramMock).WillReturnRows(rows)

	res, err := NewRepository().DetailCoursus(context.Background(), paramMock)
	assert.NoError(t, err)
	assert.Equal(t, resMock, res)
}
