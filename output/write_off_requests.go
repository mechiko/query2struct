package output

type WriteOffRequests struct {
    Id	int64	`db:"id"` 
    IdWriteOffActs	int	`db:"id_write_off_acts"` 
    ClientRegId	string	`db:"client_reg_id"` 
    RequestNumber	string	`db:"request_number"` 
    RequestDate	string	`db:"request_date"` 
    RequestRegId	string	`db:"request_reg_id"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
}

type WriteOffRequestsSlice []*WriteOffRequests

func (r *WriteOffRequests) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO write_off_requests (id_write_off_acts,client_reg_id,request_number,request_date,request_reg_id,status,reply_id) 
	VALUES(?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdWriteOffActs,r.ClientRegId,r.RequestNumber,r.RequestDate,r.RequestRegId,r.Status,r.ReplyId)
	return err
}

func (r *WriteOffRequests) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE write_off_requests SET id_write_off_acts=?,client_reg_id=?,request_number=?,request_date=?,request_reg_id=?,status=?,reply_id=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdWriteOffActs,r.ClientRegId,r.RequestNumber,r.RequestDate,r.RequestRegId,r.Status,r.ReplyId,r.Id)
	return err
}

func (r *WriteOffRequests) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE write_off_requests WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *WriteOffRequests) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from write_off_requests WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *WriteOffRequests) All(db *sqlx.DB) (out WriteOffRequestsSlice, err error) {
	out = make(WriteOffRequestsSlice, 0)
	sql := `select * from write_off_requests order by id;`
	err = db.Select(&out, sql)
	return out, err
}
