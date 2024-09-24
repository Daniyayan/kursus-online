package Product

import (
	"context"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ConnectionMock(t *testing.T) {
	// Buat mock connection dan mock expectation
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error saat membuat mock: %v", err)
	}
	defer db.Close()

	// Inisialisasi controller dengan mock database
	controller := &Controller{DB: db}

	// Simulasi input untuk CreateCourse
	var paramsMock = Models.Course{
		Id:          uuid.New().String(),
		Title:       "Course Title",
		Description: "Course Description",
		Instructor:  "John Doe",
		Price:       150000,
	}

	// Define query SQL yang dieksekusi dan ekspektasi dari query tersebut
	query := `INSERT INTO courses \(title, description, instructor, price\) VALUES \(\?, \?, \?, \?\)`
	mock.ExpectExec(query).WithArgs(paramsMock.Title, paramsMock.Description, paramsMock.Instructor, paramsMock.Price).WillReturnError(
		sqlmock.ErrCancelled, // Simulasi kegagalan eksekusi query
	)

	// Jalankan fungsi CreateCourse dan pastikan ada error
	err = Repository.NewRepository(controller.DB).CreateCourse(context.Background(), paramsMock)
	assert.Error(t, err)

	// Pastikan semua mock expectation terpenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectation tidak terpenuhi: %v", err)
	}
}
