package output

type RestsBoxesLocal struct {
    Id	int64	`db:"id"` 
    IdRestsPalletsLocal	int	`db:"id_rests_pallets_local"` 
    BoxNumber	string	`db:"box_number"` 
    InformF2RegId	string	`db:"inform_f2_reg_id"` 
    Archive	int	`db:"archive"` 
}

type RestsBoxesLocalSlice []*RestsBoxesLocal

func (r *RestsBoxesLocal) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO rests_boxes_local (id_rests_pallets_local,box_number,inform_f2_reg_id,archive) 
	VALUES(?,?,?,?);`
	_, err = db.Exec(sql, r.IdRestsPalletsLocal,r.BoxNumber,r.InformF2RegId,r.Archive)
	return err
}

func (r *RestsBoxesLocal) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE rests_boxes_local SET id_rests_pallets_local=?,box_number=?,inform_f2_reg_id=?,archive=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdRestsPalletsLocal,r.BoxNumber,r.InformF2RegId,r.Archive,r.Id)
	return err
}

func (r *RestsBoxesLocal) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE rests_boxes_local WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *RestsBoxesLocal) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from rests_boxes_local WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *RestsBoxesLocal) All(db *sqlx.DB) (out RestsBoxesLocalSlice, err error) {
	out = make(RestsBoxesLocalSlice, 0)
	sql := `select * from rests_boxes_local order by id;`
	err = db.Select(&out, sql)
	return out, err
}
