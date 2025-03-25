package output

type RestsMarksEgais struct {
    Id	int64	`db:"id"` 
    IdRequests	int	`db:"id_requests"` 
    Mark	string	`db:"mark"` 
    InformF2RegId	string	`db:"inform_f2_reg_id"` 
    RestsMarksDate	string	`db:"rests_marks_date"` 
}

type RestsMarksEgaisSlice []*RestsMarksEgais

func (r *RestsMarksEgais) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO rests_marks_egais (id_requests,mark,inform_f2_reg_id,rests_marks_date) 
	VALUES(?,?,?,?);`
	_, err = db.Exec(sql, r.IdRequests,r.Mark,r.InformF2RegId,r.RestsMarksDate)
	return err
}

func (r *RestsMarksEgais) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE rests_marks_egais SET id_requests=?,mark=?,inform_f2_reg_id=?,rests_marks_date=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdRequests,r.Mark,r.InformF2RegId,r.RestsMarksDate,r.Id)
	return err
}

func (r *RestsMarksEgais) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE rests_marks_egais WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *RestsMarksEgais) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from rests_marks_egais WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *RestsMarksEgais) All(db *sqlx.DB) (out RestsMarksEgaisSlice, err error) {
	out = make(RestsMarksEgaisSlice, 0)
	sql := `select * from rests_marks_egais order by id;`
	err = db.Select(&out, sql)
	return out, err
}
