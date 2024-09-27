package Mapel

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

// Test CreateMapel Success
func TestCreateMapelSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = Request.CreateMapel{
		NamaMapel: "Matematika",
		Materi:    "Aljabar",
		NamaDosen: "Dosen A",
	}

	query := `INSERT INTO mapel (nama_mapel, materi, nama_dosen, created_at) VALUES (?, ?, ?, ?)`
	mock.ExpectExec(query).WithArgs(paramsMock.NamaMapel, paramsMock.Materi, paramsMock.NamaDosen, AnyTime{}).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().CreateMapel(context.Background(), paramsMock)
	assert.NoError(t, err)
}

// Test CreateMapel Failure
func TestCreateMapelFailure(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = Request.CreateMapel{
		NamaMapel: "Fisika",
		Materi:    "", // Empty materi to trigger failure
		NamaDosen: "Dosen B",
	}

	query := `INSERT INTO mapel (nama_mapel, materi, nama_dosen, created_at) VALUES (?, ?, ?, ?)`
	mock.ExpectExec(query).WithArgs(paramsMock.NamaMapel, paramsMock.Materi, paramsMock.NamaDosen, AnyTime{}).WillReturnResult(
		sqlmock.NewResult(0, 0))

	err := NewRepository().CreateMapel(context.Background(), paramsMock)
	assert.Error(t, err)
}

// Test ListMapel Success
func TestListMapelSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var resMock = []Response.MapelList{
		Response.MapelList{
			IdMapel:   1,
			NamaMapel: "Matematika",
			Materi:    "Aljabar",
			NamaDosen: "Dosen A",
		},
	}

	query := `SELECT id_mapel, nama_mapel, materi, nama_dosen FROM mapel WHERE created_at IS NOT NULL ORDER BY created_at DESC LIMIT 10`
	rows := sqlmock.NewRows([]string{"id_mapel", "nama_mapel", "materi", "nama_dosen"}).
		AddRow(resMock[0].IdMapel, resMock[0].NamaMapel, resMock[0].Materi, resMock[0].NamaDosen)

	mock.ExpectQuery(query).WillReturnRows(rows)

	res, err := NewRepository().ListMapel(context.Background(), "mapel")
	assert.NoError(t, err)
	assert.Equal(t, resMock, res)
}

// Test CheckExistsMapelNama Success
func TestCheckExistsMapelNamaSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramMock = "Matematika"
	var resMock bool

	query := `SELECT EXISTS (SELECT 1 FROM mapel WHERE LOWER(nama_mapel) = LOWER(?)) AS "exists"`

	rows := sqlmock.NewRows([]string{"exists"}).AddRow(resMock)
	mock.ExpectQuery(query).WithArgs(paramMock).WillReturnRows(rows)

	res, err := NewRepository().CheckExistsMapelId(context.Background(), paramMock)
	assert.NoError(t, err)
	assert.Equal(t, resMock, res)
}

// Test UpdateMapel Success
func TestUpdateMapelSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = Request.UpdateMapel{
		IdMapel:   1,
		NamaMapel: "Matematika Lanjutan",
		Materi:    "Persamaan Linear",
		NamaDosen: "Dosen A",
	}

	query := `UPDATE mapel SET nama_mapel = ?, materi = ?, nama_dosen = ?, updated_at = ? WHERE id_mapel = ?`
	mock.ExpectExec(query).WithArgs(paramsMock.NamaMapel, paramsMock.Materi, paramsMock.NamaDosen, AnyTime{}, paramsMock.IdMapel).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().UpdateMapel(context.Background(), paramsMock)
	assert.NoError(t, err)
}

// Test DeleteMapel Success
func TestDeleteMapelSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = "1"

	query := `DELETE FROM mapel WHERE id_mapel = ?`
	mock.ExpectExec(query).WithArgs(paramsMock).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().DeleteMapel(context.Background(), paramsMock)
	assert.NoError(t, err)
}

// Test DetailMapel Success
func TestDetailMapelSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramMock = "1"
	var resMock = Response.MapelDetail{
		IdMapel:   1,
		NamaMapel: "Matematika",
		Materi:    "Aljabar",
		NamaDosen: "Dosen A",
	}

	query := `SELECT id_mapel, nama_mapel, materi, nama_dosen, updated_at FROM mapel WHERE id_mapel = ?`
	rows := sqlmock.NewRows([]string{"id_mapel", "nama_mapel", "materi", "nama_dosen", "updated_at"}).AddRow(resMock.IdMapel,
		resMock.NamaMapel, resMock.Materi, resMock.NamaDosen, AnyTime{})
	mock.ExpectQuery(query).WithArgs(paramMock).WillReturnRows(rows)

	res, err := NewRepository().DetailMapel(context.Background(), paramMock)
	assert.NoError(t, err)
	assert.Equal(t, resMock, res)
}
