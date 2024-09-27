package Register

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

// Test CreateRegister Success
func TestCreateRegisterSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = Request.CreateRegister{
		Email:    "user@example.com",
		NoTelpon: 1234567890,
		Status:   "Active",
	}

	query := `INSERT INTO register (email, No_telpon, status, created_at) VALUES (?, ?, ?, ?)`
	mock.ExpectExec(query).WithArgs(paramsMock.Email, paramsMock.NoTelpon, paramsMock.Status, AnyTime{}).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().CreateRegister(context.Background(), paramsMock)
	assert.NoError(t, err)
}

// Test CreateRegister Failure
func TestCreateRegisterFailure(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = Request.CreateRegister{
		Email:    "", // Empty email to trigger failure
		NoTelpon: 1234567890,
		Status:   "Inactive",
	}

	query := `INSERT INTO register (email, No_telpon, status, created_at) VALUES (?, ?, ?, ?)`
	mock.ExpectExec(query).WithArgs(paramsMock.Email, paramsMock.NoTelpon, paramsMock.Status, AnyTime{}).WillReturnResult(
		sqlmock.NewResult(0, 0))

	err := NewRepository().CreateRegister(context.Background(), paramsMock)
	assert.Error(t, err)
}

// Test ListRegister Success
func TestListRegisterSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var resMock = []Response.RegisterList{
		Response.RegisterList{
			IdRegis:  1,
			Email:    "user@example.com",
			NoTelpon: 1234567890,
			Status:   "Active",
		},
	}

	query := `SELECT id_regis, email, No_telpon, status FROM register WHERE created_at IS NOT NULL ORDER BY created_at DESC LIMIT 10`
	rows := sqlmock.NewRows([]string{"id_regis", "email", "No_telpon", "status"}).
		AddRow(resMock[0].IdRegis, resMock[0].Email, resMock[0].NoTelpon, resMock[0].Status)

	mock.ExpectQuery(query).WillReturnRows(rows)

	res, err := NewRepository().ListRegister(context.Background(), "nil")
	assert.NoError(t, err)
	assert.Equal(t, resMock, res)
}

// Test CheckExistsRegisterEmail Success
func TestCheckExistsRegisterEmailSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramMock = "user@example.com"
	var resMock bool

	query := `SELECT EXISTS (SELECT 1 FROM register WHERE LOWER(email) = LOWER(?)) AS "exists"`

	rows := sqlmock.NewRows([]string{"exists"}).AddRow(resMock)
	mock.ExpectQuery(query).WithArgs(paramMock).WillReturnRows(rows)

	res, err := NewRepository().CheckExistsRegisterEmail(context.Background(), paramMock)
	assert.NoError(t, err)
	assert.Equal(t, resMock, res)
}

// Test UpdateRegister Success
func TestUpdateRegisterSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = Request.UpdateRegister{
		IdRegis:  1,
		Email:    "newuser@example.com",
		NoTelpon: 9876543210,
		Status:   "Inactive",
	}

	query := `UPDATE register SET email = ?, No_telpon = ?, status = ?, updated_at = ? WHERE id_regis = ?`
	mock.ExpectExec(query).WithArgs(paramsMock.Email, paramsMock.NoTelpon, paramsMock.Status, AnyTime{}, paramsMock.IdRegis).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().UpdateRegister(context.Background(), paramsMock)
	assert.NoError(t, err)
}

// Test DeleteRegister Success
func TestDeleteRegisterSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = "1"

	query := `DELETE FROM register WHERE id_regis = ?`
	mock.ExpectExec(query).WithArgs(paramsMock).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().DeleteRegister(context.Background(), paramsMock)
	assert.NoError(t, err)
}

// Test DetailRegister Success
func TestDetailRegisterSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramMock = "1"
	var resMock = Response.RegisterDetail{
		IdRegis:  1,
		Email:    "user@example.com",
		NoTelpon: 1234567890,
		Status:   "Active",
	}

	query := `SELECT id_regis, email, No_telpon, status, updated_at FROM register WHERE id_regis = ?`
	rows := sqlmock.NewRows([]string{"id_regis", "email", "No_telpon", "status", "updated_at"}).
		AddRow(resMock.IdRegis, resMock.Email, resMock.NoTelpon, resMock.Status, AnyTime{})

	mock.ExpectQuery(query).WithArgs(paramMock).WillReturnRows(rows)

	res, err := NewRepository().DetailRegister(context.Background(), paramMock)
	assert.NoError(t, err)
	assert.Equal(t, resMock, res)
}
