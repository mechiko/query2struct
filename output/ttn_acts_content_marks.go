package output

type TtnActsContentMarks struct {
    Id	int64	`db:"id"` 
    IdTtnActsContent	int	`db:"id_ttn_acts_content"` 
    Mark	string	`db:"mark"` 
}

type TtnActsContentMarksSlice []*TtnActsContentMarks

func (r *TtnActsContentMarks) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_acts_content_marks (id_ttn_acts_content,mark) 
	VALUES(?,?);`
	_, err = db.Exec(sql, r.IdTtnActsContent,r.Mark)
	return err
}

func (r *TtnActsContentMarks) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_acts_content_marks SET id_ttn_acts_content=?,mark=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtnActsContent,r.Mark,r.Id)
	return err
}

func (r *TtnActsContentMarks) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_acts_content_marks WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnActsContentMarks) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_acts_content_marks WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnActsContentMarks) All(db *sqlx.DB) (out TtnActsContentMarksSlice, err error) {
	out = make(TtnActsContentMarksSlice, 0)
	sql := `select * from ttn_acts_content_marks order by id;`
	err = db.Select(&out, sql)
	return out, err
}
