package output

type TtnActs struct {
    Id	int64	`db:"id"` 
    IdTtn	int	`db:"id_ttn"` 
    ActType	string	`db:"act_type"` 
    ActNumber	string	`db:"act_number"` 
    ActDate	string	`db:"act_date"` 
    ActRegId	string	`db:"act_reg_id"` 
    ActComment	string	`db:"act_comment"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
    Xml	string	`db:"xml"` 
}

type TtnActsSlice []*TtnActs

func (r *TtnActs) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_acts (id_ttn,act_type,act_number,act_date,act_reg_id,act_comment,status,reply_id,xml) 
	VALUES(?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdTtn,r.ActType,r.ActNumber,r.ActDate,r.ActRegId,r.ActComment,r.Status,r.ReplyId,r.Xml)
	return err
}

func (r *TtnActs) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_acts SET id_ttn=?,act_type=?,act_number=?,act_date=?,act_reg_id=?,act_comment=?,status=?,reply_id=?,xml=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtn,r.ActType,r.ActNumber,r.ActDate,r.ActRegId,r.ActComment,r.Status,r.ReplyId,r.Xml,r.Id)
	return err
}

func (r *TtnActs) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_acts WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnActs) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_acts WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnActs) All(db *sqlx.DB) (out TtnActsSlice, err error) {
	out = make(TtnActsSlice, 0)
	sql := `select * from ttn_acts order by id;`
	err = db.Select(&out, sql)
	return out, err
}
