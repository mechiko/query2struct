package output

type Tickets struct {
    Id	int64	`db:"id"` 
    TicketDate	string	`db:"ticket_date"` 
    Identity	string	`db:"identity"` 
    DocId	string	`db:"doc_id"` 
    TransportId	string	`db:"transport_id"` 
    RegId	string	`db:"reg_id"` 
    DocHash	string	`db:"doc_hash"` 
    DocType	string	`db:"doc_type"` 
    Conclusion	string	`db:"conclusion"` 
    ConclusionDate	string	`db:"conclusion_date"` 
    ConclusionComment	string	`db:"conclusion_comment"` 
    OperationName	string	`db:"operation_name"` 
    OperationResult	string	`db:"operation_result"` 
    OperationDate	string	`db:"operation_date"` 
    OperationComment	string	`db:"operation_comment"` 
    XmlReply	string	`db:"xml_reply"` 
}

type TicketsSlice []*Tickets

func (r *Tickets) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO tickets (ticket_date,identity,doc_id,transport_id,reg_id,doc_hash,doc_type,conclusion,conclusion_date,conclusion_comment,operation_name,operation_result,operation_date,operation_comment,xml_reply) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.TicketDate,r.Identity,r.DocId,r.TransportId,r.RegId,r.DocHash,r.DocType,r.Conclusion,r.ConclusionDate,r.ConclusionComment,r.OperationName,r.OperationResult,r.OperationDate,r.OperationComment,r.XmlReply)
	return err
}

func (r *Tickets) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE tickets SET ticket_date=?,identity=?,doc_id=?,transport_id=?,reg_id=?,doc_hash=?,doc_type=?,conclusion=?,conclusion_date=?,conclusion_comment=?,operation_name=?,operation_result=?,operation_date=?,operation_comment=?,xml_reply=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.TicketDate,r.Identity,r.DocId,r.TransportId,r.RegId,r.DocHash,r.DocType,r.Conclusion,r.ConclusionDate,r.ConclusionComment,r.OperationName,r.OperationResult,r.OperationDate,r.OperationComment,r.XmlReply,r.Id)
	return err
}

func (r *Tickets) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE tickets WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *Tickets) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from tickets WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *Tickets) All(db *sqlx.DB) (out TicketsSlice, err error) {
	out = make(TicketsSlice, 0)
	sql := `select * from tickets order by id;`
	err = db.Select(&out, sql)
	return out, err
}
