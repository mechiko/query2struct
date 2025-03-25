package output

type TnRequests struct {
    Id	int64	`db:"id"` 
    IdTn	int	`db:"id_tn"` 
    RequestDate	string	`db:"request_date"` 
    RequestRegId	string	`db:"request_reg_id"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
}

type TnRequestsSlice []*TnRequests

func (r *TnRequests) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO tn_requests (id_tn,request_date,request_reg_id,status,reply_id) 
	VALUES(?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdTn,r.RequestDate,r.RequestRegId,r.Status,r.ReplyId)
	return err
}

func (r *TnRequests) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE tn_requests SET id_tn=?,request_date=?,request_reg_id=?,status=?,reply_id=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTn,r.RequestDate,r.RequestRegId,r.Status,r.ReplyId,r.Id)
	return err
}

func (r *TnRequests) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE tn_requests WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TnRequests) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from tn_requests WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TnRequests) All(db *sqlx.DB) (out TnRequestsSlice, err error) {
	out = make(TnRequestsSlice, 0)
	sql := `select * from tn_requests order by id;`
	err = db.Select(&out, sql)
	return out, err
}
