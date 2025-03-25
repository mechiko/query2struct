package output

type ChargeOnForm1 struct {
    Id	int64	`db:"id"` 
    IdChargeOnActs	int	`db:"id_charge_on_acts"` 
    DocIdentity	string	`db:"doc_identity"` 
    DocRegId	string	`db:"doc_reg_id"` 
    DocNumber	string	`db:"doc_number"` 
    Xml	string	`db:"xml"` 
}

type ChargeOnForm1Slice []*ChargeOnForm1

func (r *ChargeOnForm1) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO charge_on_form1 (id_charge_on_acts,doc_identity,doc_reg_id,doc_number,xml) 
	VALUES(?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdChargeOnActs,r.DocIdentity,r.DocRegId,r.DocNumber,r.Xml)
	return err
}

func (r *ChargeOnForm1) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE charge_on_form1 SET id_charge_on_acts=?,doc_identity=?,doc_reg_id=?,doc_number=?,xml=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdChargeOnActs,r.DocIdentity,r.DocRegId,r.DocNumber,r.Xml,r.Id)
	return err
}

func (r *ChargeOnForm1) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE charge_on_form1 WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ChargeOnForm1) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from charge_on_form1 WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ChargeOnForm1) All(db *sqlx.DB) (out ChargeOnForm1Slice, err error) {
	out = make(ChargeOnForm1Slice, 0)
	sql := `select * from charge_on_form1 order by id;`
	err = db.Select(&out, sql)
	return out, err
}
