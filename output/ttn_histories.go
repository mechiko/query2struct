package output

type TtnHistories struct {
    Id	int64	`db:"id"` 
    IdTtn	int	`db:"id_ttn"` 
    Identity	string	`db:"identity"` 
}

type TtnHistoriesSlice []*TtnHistories

func (r *TtnHistories) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_histories (id_ttn,identity) 
	VALUES(?,?);`
	_, err = db.Exec(sql, r.IdTtn,r.Identity)
	return err
}

func (r *TtnHistories) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_histories SET id_ttn=?,identity=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtn,r.Identity,r.Id)
	return err
}

func (r *TtnHistories) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_histories WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnHistories) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_histories WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnHistories) All(db *sqlx.DB) (out TtnHistoriesSlice, err error) {
	out = make(TtnHistoriesSlice, 0)
	sql := `select * from ttn_histories order by id;`
	err = db.Select(&out, sql)
	return out, err
}
