package output

type RestsMarksLocal struct {
    Id	int64	`db:"id"` 
    IdRestsBoxesLocal	int	`db:"id_rests_boxes_local"` 
    Mark	string	`db:"mark"` 
    InformF2RegId	string	`db:"inform_f2_reg_id"` 
    Archive	int	`db:"archive"` 
}

type RestsMarksLocalSlice []*RestsMarksLocal

func (r *RestsMarksLocal) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO rests_marks_local (id_rests_boxes_local,mark,inform_f2_reg_id,archive) 
	VALUES(?,?,?,?);`
	_, err = db.Exec(sql, r.IdRestsBoxesLocal,r.Mark,r.InformF2RegId,r.Archive)
	return err
}

func (r *RestsMarksLocal) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE rests_marks_local SET id_rests_boxes_local=?,mark=?,inform_f2_reg_id=?,archive=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdRestsBoxesLocal,r.Mark,r.InformF2RegId,r.Archive,r.Id)
	return err
}

func (r *RestsMarksLocal) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE rests_marks_local WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *RestsMarksLocal) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from rests_marks_local WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *RestsMarksLocal) All(db *sqlx.DB) (out RestsMarksLocalSlice, err error) {
	out = make(RestsMarksLocalSlice, 0)
	sql := `select * from rests_marks_local order by id;`
	err = db.Select(&out, sql)
	return out, err
}
