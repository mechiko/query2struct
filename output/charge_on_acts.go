package output

type ChargeOnActs struct {
    Id	int64	`db:"id"` 
    CreateDate	string	`db:"create_date"` 
    DocIdentity	string	`db:"doc_identity"` 
    DocType	string	`db:"doc_type"` 
    DocNumber	string	`db:"doc_number"` 
    DocDate	string	`db:"doc_date"` 
    DocComment	string	`db:"doc_comment"` 
    WriteOffActRegId	string	`db:"write_off_act_reg_id"` 
    Version	string	`db:"version"` 
    State	string	`db:"state"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
    Archive	int	`db:"archive"` 
    Xml	string	`db:"xml"` 
}

type ChargeOnActsSlice []*ChargeOnActs

func (r *ChargeOnActs) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO charge_on_acts (create_date,doc_identity,doc_type,doc_number,doc_date,doc_comment,write_off_act_reg_id,version,state,status,reply_id,archive,xml) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.CreateDate,r.DocIdentity,r.DocType,r.DocNumber,r.DocDate,r.DocComment,r.WriteOffActRegId,r.Version,r.State,r.Status,r.ReplyId,r.Archive,r.Xml)
	return err
}

func (r *ChargeOnActs) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE charge_on_acts SET create_date=?,doc_identity=?,doc_type=?,doc_number=?,doc_date=?,doc_comment=?,write_off_act_reg_id=?,version=?,state=?,status=?,reply_id=?,archive=?,xml=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.CreateDate,r.DocIdentity,r.DocType,r.DocNumber,r.DocDate,r.DocComment,r.WriteOffActRegId,r.Version,r.State,r.Status,r.ReplyId,r.Archive,r.Xml,r.Id)
	return err
}

func (r *ChargeOnActs) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE charge_on_acts WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ChargeOnActs) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from charge_on_acts WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ChargeOnActs) All(db *sqlx.DB) (out ChargeOnActsSlice, err error) {
	out = make(ChargeOnActsSlice, 0)
	sql := `select * from charge_on_acts order by id;`
	err = db.Select(&out, sql)
	return out, err
}
