package output

type TtnActsTransport struct {
    Id	int64	`db:"id"` 
    IdTtnActs	int	`db:"id_ttn_acts"` 
    Ownership	string	`db:"ownership"` 
}

type TtnActsTransportSlice []*TtnActsTransport

func (r *TtnActsTransport) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_acts_transport (id_ttn_acts,ownership) 
	VALUES(?,?);`
	_, err = db.Exec(sql, r.IdTtnActs,r.Ownership)
	return err
}

func (r *TtnActsTransport) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_acts_transport SET id_ttn_acts=?,ownership=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtnActs,r.Ownership,r.Id)
	return err
}

func (r *TtnActsTransport) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_acts_transport WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnActsTransport) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_acts_transport WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnActsTransport) All(db *sqlx.DB) (out TtnActsTransportSlice, err error) {
	out = make(TtnActsTransportSlice, 0)
	sql := `select * from ttn_acts_transport order by id;`
	err = db.Select(&out, sql)
	return out, err
}
