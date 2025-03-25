package output

type TtnActsTickets struct {
    Id	int64	`db:"id"` 
    IdTtnActs	int	`db:"id_ttn_acts"` 
    TicketType	string	`db:"ticket_type"` 
    TicketNumber	string	`db:"ticket_number"` 
    TicketDate	string	`db:"ticket_date"` 
    TicketRegId	string	`db:"ticket_reg_id"` 
    TicketComment	string	`db:"ticket_comment"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
    Xml	string	`db:"xml"` 
}

type TtnActsTicketsSlice []*TtnActsTickets

func (r *TtnActsTickets) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_acts_tickets (id_ttn_acts,ticket_type,ticket_number,ticket_date,ticket_reg_id,ticket_comment,status,reply_id,xml) 
	VALUES(?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdTtnActs,r.TicketType,r.TicketNumber,r.TicketDate,r.TicketRegId,r.TicketComment,r.Status,r.ReplyId,r.Xml)
	return err
}

func (r *TtnActsTickets) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_acts_tickets SET id_ttn_acts=?,ticket_type=?,ticket_number=?,ticket_date=?,ticket_reg_id=?,ticket_comment=?,status=?,reply_id=?,xml=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtnActs,r.TicketType,r.TicketNumber,r.TicketDate,r.TicketRegId,r.TicketComment,r.Status,r.ReplyId,r.Xml,r.Id)
	return err
}

func (r *TtnActsTickets) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_acts_tickets WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnActsTickets) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_acts_tickets WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnActsTickets) All(db *sqlx.DB) (out TtnActsTicketsSlice, err error) {
	out = make(TtnActsTicketsSlice, 0)
	sql := `select * from ttn_acts_tickets order by id;`
	err = db.Select(&out, sql)
	return out, err
}
