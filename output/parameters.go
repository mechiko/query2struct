package output

type Parameters struct {
    DbVersion	string	`db:"db_version"` 
}

type ParametersSlice []*Parameters

func (r *Parameters) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO parameters (db_version) 
	VALUES(?);`
	_, err = db.Exec(sql, r.DbVersion)
	return err
}

func (r *Parameters) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE parameters SET db_version=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.DbVersion,)
	return err
}

func (r *Parameters) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE parameters WHERE =?;`
	_, err = db.Exec(sql, )
	return err
}

func (r *Parameters) Find(db *sqlx.DB) (err error) {
	if  == 0 {
		return fmt.Errorf(" = 0")
	}
	sql := `select * from parameters WHERE =?;`
	err = db.Get(r, sql, )
	return err
}

func (r *Parameters) All(db *sqlx.DB) (out ParametersSlice, err error) {
	out = make(ParametersSlice, 0)
	sql := `select * from parameters order by ;`
	err = db.Select(&out, sql)
	return out, err
}
