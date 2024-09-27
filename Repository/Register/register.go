package Register

import (
	"context"
	"kursus-online/Config"
	"kursus-online/Controller/Dto/Request"
)

// CreateRegister creates a new register record in the database.
func (r *register) CreateRegister(ctx context.Context, param Request.CreateRegister) (err error) {
	query := `INSERT INTO register (id_regis, email, No_telpon, status) VALUES (?, ?, ?, ?)`
	if _, err = Config.DATABASE_MAIN.Get().ExecContext(ctx, query, param.IdRegis, param.Email, param.NoTelpon, param.Status); err != nil {
		return err
	}

	return
}

//// CheckExistsRegisterEmail checks if a register record with the given email exists.
//func (r *Register) CheckExistsRegisterEmail(ctx context.Context, email string) (exists bool, err error) {
//	query := `SELECT EXISTS (SELECT 1 FROM register WHERE LOWER(email) = LOWER(?)) AS "exists"`
//	err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, email).Scan(&exists)
//	return
//}
//
//// CheckExistsRegisterId checks if a register record with the given ID exists.
//func (r *Register) CheckExistsRegisterId(ctx context.Context, id int) (exists bool, err error) {
//	query := `SELECT EXISTS (SELECT 1 FROM register WHERE id_regis = ?) AS "exists"`
//	err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, id).Scan(&exists)
//	return
//}
//
//// UpdateRegister updates the register record with the given ID.
//func (r *Register) UpdateRegister(ctx context.Context, param Request.UpdateRegister) (err error) {
//	query := `UPDATE register SET email = ?, No_telpon = ?, status = ? WHERE id_regis = ?`
//	if _, err = Config.DATABASE_MAIN.Get().ExecContext(ctx, query, param.Email, param.NoTelpon, param.Status, param.IdRegis); err != nil {
//		return err
//	}
//	return
//}
//
//// DeleteRegister deletes a register record by ID.
//func (r *Register) DeleteRegister(ctx context.Context, id int) (err error) {
//	query := `DELETE FROM register WHERE id_regis = ?`
//	if _, err = Config.DATABASE_MAIN.Get().ExecContext(ctx, query, id); err != nil {
//		return err
//	}
//
//	return
//}
//
//// ListRegister retrieves a list of register records with sorting options.
//func (r *Register) ListRegister(ctx context.Context, sortBy string) (res []Response.RegisterList, err error) {
//	var (
//		data       Response.RegisterList
//		connection = Config.DATABASE_MAIN.Get()
//	)
//
//	query := `SELECT id_regis, email, No_telpon, status FROM register ORDER BY id_regis DESC LIMIT 10`
//	if sortBy == "email" {
//		query = `SELECT id_regis, email, No_telpon, status FROM register ORDER BY email ASC LIMIT 10`
//	}
//
//	rows, err := connection.QueryContext(ctx, query)
//	if err != nil {
//		return
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		if err = rows.Scan(&data.IdRegis, &data.Email, &data.NoTelpon, &data.Status); err != nil {
//			return
//		}
//
//		res = append(res, data)
//	}
//
//	return
//}
//
//// DetailRegister retrieves the details of a specific register record by ID.
//func (r *Register) DetailRegister(ctx context.Context, id int) (res Response.RegisterDetail, err error) {
//	query := `SELECT id_regis, email, No_telpon, status FROM register WHERE id_regis = ?`
//	if err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, id).Scan(&res.IdRegis, &res.Email, &res.NoTelpon, &res.Status); err != nil {
//		return
//	}
//	return
//}
