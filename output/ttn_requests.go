package output

type TtnRequests struct {
    Id	int64	`db:"id"` 
    IdTtn	int	`db:"id_ttn"` 
    ClientRegId	string	`db:"client_reg_id"` 
    RequestNumber	string	`db:"request_number"` 
    RequestDate	string	`db:"request_date"` 
    RequestRegId	string	`db:"request_reg_id"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
}

type TtnRequestsSlice []*TtnRequests

func (r *TtnRequests) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_requests (id_ttn,client_reg_id,request_number,request_date,request_reg_id,status,reply_id) 
	VALUES(?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdTtn,r.ClientRegId,r.RequestNumber,r.RequestDate,r.RequestRegId,r.Status,r.ReplyId)
	return err
}

func (r *TtnRequests) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_requests SET id_ttn=?,client_reg_id=?,request_number=?,request_date=?,request_reg_id=?,status=?,reply_id=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtn,r.ClientRegId,r.RequestNumber,r.RequestDate,r.RequestRegId,r.Status,r.ReplyId,r.Id)
	return err
}

func (r *TtnRequests) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_requests WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnRequests) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_requests WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnRequests) All(db *sqlx.DB) (out TtnRequestsSlice, err error) {
	out = make(TtnRequestsSlice, 0)
	sql := `select * from ttn_requests order by id;`
	err = db.Select(&out, sql)
	return out, err
}
