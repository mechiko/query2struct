package output

type TnRoutes struct {
    Id	int64	`db:"id"` 
    IdTn	int	`db:"id_tn"` 
    Route	string	`db:"route"` 
}

type TnRoutesSlice []*TnRoutes

func (r *TnRoutes) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO tn_routes (id_tn,route) 
	VALUES(?,?);`
	_, err = db.Exec(sql, r.IdTn,r.Route)
	return err
}

func (r *TnRoutes) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE tn_routes SET id_tn=?,route=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTn,r.Route,r.Id)
	return err
}

func (r *TnRoutes) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE tn_routes WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TnRoutes) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from tn_routes WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TnRoutes) All(db *sqlx.DB) (out TnRoutesSlice, err error) {
	out = make(TnRoutesSlice, 0)
	sql := `select * from tn_routes order by id;`
	err = db.Select(&out, sql)
	return out, err
}
