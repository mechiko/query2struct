package output

type WriteOffActs struct {
    Id	int64	`db:"id"` 
    CreateDate	string	`db:"create_date"` 
    DocIdentity	string	`db:"doc_identity"` 
    DocType	string	`db:"doc_type"` 
    DocNumber	string	`db:"doc_number"` 
    DocDate	string	`db:"doc_date"` 
    DocComment	string	`db:"doc_comment"` 
    Version	string	`db:"version"` 
    State	string	`db:"state"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
    Archive	int	`db:"archive"` 
    Xml	string	`db:"xml"` 
    DocMode	string	`db:"doc_mode"` 
}

type WriteOffActsSlice []*WriteOffActs

func (r *WriteOffActs) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO write_off_acts (create_date,doc_identity,doc_type,doc_number,doc_date,doc_comment,version,state,status,reply_id,archive,xml,doc_mode) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.CreateDate,r.DocIdentity,r.DocType,r.DocNumber,r.DocDate,r.DocComment,r.Version,r.State,r.Status,r.ReplyId,r.Archive,r.Xml,r.DocMode)
	return err
}

func (r *WriteOffActs) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE write_off_acts SET create_date=?,doc_identity=?,doc_type=?,doc_number=?,doc_date=?,doc_comment=?,version=?,state=?,status=?,reply_id=?,archive=?,xml=?,doc_mode=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.CreateDate,r.DocIdentity,r.DocType,r.DocNumber,r.DocDate,r.DocComment,r.Version,r.State,r.Status,r.ReplyId,r.Archive,r.Xml,r.DocMode,r.Id)
	return err
}

func (r *WriteOffActs) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE write_off_acts WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *WriteOffActs) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from write_off_acts WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *WriteOffActs) All(db *sqlx.DB) (out WriteOffActsSlice, err error) {
	out = make(WriteOffActsSlice, 0)
	sql := `select * from write_off_acts order by id;`
	err = db.Select(&out, sql)
	return out, err
}
