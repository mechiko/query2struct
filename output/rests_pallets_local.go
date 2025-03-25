package output

type RestsPalletsLocal struct {
    Id	int64	`db:"id"` 
    PalletNumber	string	`db:"pallet_number"` 
    InformF2RegId	string	`db:"inform_f2_reg_id"` 
    Archive	int	`db:"archive"` 
}

type RestsPalletsLocalSlice []*RestsPalletsLocal

func (r *RestsPalletsLocal) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO rests_pallets_local (pallet_number,inform_f2_reg_id,archive) 
	VALUES(?,?,?);`
	_, err = db.Exec(sql, r.PalletNumber,r.InformF2RegId,r.Archive)
	return err
}

func (r *RestsPalletsLocal) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE rests_pallets_local SET pallet_number=?,inform_f2_reg_id=?,archive=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.PalletNumber,r.InformF2RegId,r.Archive,r.Id)
	return err
}

func (r *RestsPalletsLocal) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE rests_pallets_local WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *RestsPalletsLocal) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from rests_pallets_local WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *RestsPalletsLocal) All(db *sqlx.DB) (out RestsPalletsLocalSlice, err error) {
	out = make(RestsPalletsLocalSlice, 0)
	sql := `select * from rests_pallets_local order by id;`
	err = db.Select(&out, sql)
	return out, err
}
