package Mapel

import (
	"context"
	"kursus-online/Config"
	"kursus-online/Controller/Dto/Request"
	"kursus-online/Controller/Dto/Response"
	"time"
)

func (m *Mapel) CreateMapel(ctx context.Context, param Request.CreateMapel) (err error) {
	query := `INSERT INTO mapel (id_mapel, nama_mapel, materi, nama_dosen, created_at) VALUES (?, ?, ?, ?, ?)`
	if _, err = Config.DATABASE_MAIN.Get().ExecContext(ctx, query, param.IdMapel, param.NamaMapel, param.Materi, param.NamaDosen, time.Now()); err != nil {
		return err
	}

	return
}

func (m *Mapel) CheckExistsMapelNama(ctx context.Context, nama string) (exists bool, err error) {
	query := `SELECT EXISTS (SELECT 1 FROM mapel WHERE LOWER(nama_mapel) = LOWER(?)) AS "exists"`
	err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, nama).Scan(&exists)
	return
}

func (m *Mapel) CheckExistsMapelId(ctx context.Context, id int) (exists bool, err error) {
	query := `SELECT EXISTS (SELECT 1 FROM mapel WHERE id_mapel = ?) AS "exists"`
	err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, id).Scan(&exists)
	return
}

func (m *Mapel) UpdateMapel(ctx context.Context, param Request.UpdateMapel) (err error) {
	query := `UPDATE mapel SET nama_mapel = ?, materi = ?, nama_dosen = ?, updated_at = ? WHERE id_mapel = ?`
	if _, err = Config.DATABASE_MAIN.Get().ExecContext(ctx, query, param.NamaMapel, param.Materi, param.NamaDosen, time.Now(), param.IdMapel); err != nil {
		return err
	}
	return
}

func (m *Mapel) DeleteMapel(ctx context.Context, id int) (err error) {
	query := `DELETE FROM mapel WHERE id_mapel = ?`
	if _, err = Config.DATABASE_MAIN.Get().ExecContext(ctx, query, id); err != nil {
		return err
	}

	return
}

func (m *Mapel) ListMapel(ctx context.Context, sortBy string) (res []Response.MapelList, err error) {
	var (
		data       Response.MapelList
		connection = Config.DATABASE_MAIN.Get()
	)

	query := `SELECT id_mapel, nama_mapel, materi, nama_dosen FROM mapel ORDER BY created_at DESC LIMIT 10`
	if sortBy == "nama_mapel" {
		query = `SELECT id_mapel, nama_mapel, materi, nama_dosen FROM mapel ORDER BY nama_mapel ASC LIMIT 10`
	}

	rows, err := connection.QueryContext(ctx, query)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&data.IdMapel, &data.NamaMapel, &data.Materi, &data.NamaDosen); err != nil {
			return
		}

		res = append(res, data)
	}

	return
}

func (m *Mapel) DetailMapel(ctx context.Context, id int) (res Response.MapelDetail, err error) {
	query := `SELECT id_mapel, nama_mapel, materi, nama_dosen, updated_at FROM mapel WHERE id_mapel = ?`
	if err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, id).Scan(&res.IdMapel, &res.NamaMapel, &res.Materi,
		&res.NamaDosen, &res.UpdatedAt); err != nil {
		return
	}
	return
}
