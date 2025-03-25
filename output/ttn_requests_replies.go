package output

type TtnRequestsReplies struct {
    Id	int64	`db:"id"` 
    IdTtnRequests	int	`db:"id_ttn_requests"` 
    ReplyType	string	`db:"reply_type"` 
    ReplyNumber	string	`db:"reply_number"` 
    ReplyDate	string	`db:"reply_date"` 
    ReplyRegId	string	`db:"reply_reg_id"` 
    ReplyComment	string	`db:"reply_comment"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
}

type TtnRequestsRepliesSlice []*TtnRequestsReplies

func (r *TtnRequestsReplies) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_requests_replies (id_ttn_requests,reply_type,reply_number,reply_date,reply_reg_id,reply_comment,status,reply_id) 
	VALUES(?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdTtnRequests,r.ReplyType,r.ReplyNumber,r.ReplyDate,r.ReplyRegId,r.ReplyComment,r.Status,r.ReplyId)
	return err
}

func (r *TtnRequestsReplies) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_requests_replies SET id_ttn_requests=?,reply_type=?,reply_number=?,reply_date=?,reply_reg_id=?,reply_comment=?,status=?,reply_id=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtnRequests,r.ReplyType,r.ReplyNumber,r.ReplyDate,r.ReplyRegId,r.ReplyComment,r.Status,r.ReplyId,r.Id)
	return err
}

func (r *TtnRequestsReplies) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_requests_replies WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnRequestsReplies) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_requests_replies WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnRequestsReplies) All(db *sqlx.DB) (out TtnRequestsRepliesSlice, err error) {
	out = make(TtnRequestsRepliesSlice, 0)
	sql := `select * from ttn_requests_replies order by id;`
	err = db.Select(&out, sql)
	return out, err
}
