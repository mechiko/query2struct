package output

type ChargeOnRequests struct {
    Id	int64	`db:"id"` 
    IdChargeOnActs	int	`db:"id_charge_on_acts"` 
    ClientRegId	string	`db:"client_reg_id"` 
    RequestNumber	string	`db:"request_number"` 
    RequestDate	string	`db:"request_date"` 
    RequestRegId	string	`db:"request_reg_id"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
}

type ChargeOnRequestsSlice []*ChargeOnRequests

func (r *ChargeOnRequests) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO charge_on_requests (id_charge_on_acts,client_reg_id,request_number,request_date,request_reg_id,status,reply_id) 
	VALUES(?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdChargeOnActs,r.ClientRegId,r.RequestNumber,r.RequestDate,r.RequestRegId,r.Status,r.ReplyId)
	return err
}

func (r *ChargeOnRequests) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE charge_on_requests SET id_charge_on_acts=?,client_reg_id=?,request_number=?,request_date=?,request_reg_id=?,status=?,reply_id=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdChargeOnActs,r.ClientRegId,r.RequestNumber,r.RequestDate,r.RequestRegId,r.Status,r.ReplyId,r.Id)
	return err
}

func (r *ChargeOnRequests) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE charge_on_requests WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ChargeOnRequests) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from charge_on_requests WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ChargeOnRequests) All(db *sqlx.DB) (out ChargeOnRequestsSlice, err error) {
	out = make(ChargeOnRequestsSlice, 0)
	sql := `select * from charge_on_requests order by id;`
	err = db.Select(&out, sql)
	return out, err
}
