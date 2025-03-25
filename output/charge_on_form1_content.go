package output

type ChargeOnForm1Content struct {
    Id	int64	`db:"id"` 
    IdChargeOnForm1	int	`db:"id_charge_on_form1"` 
    ProductIdentity	string	`db:"product_identity"` 
    ProductInformF1RegId	string	`db:"product_inform_f1_reg_id"` 
    ProductInformF2RegId	string	`db:"product_inform_f2_reg_id"` 
}

type ChargeOnForm1ContentSlice []*ChargeOnForm1Content

func (r *ChargeOnForm1Content) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO charge_on_form1_content (id_charge_on_form1,product_identity,product_inform_f1_reg_id,product_inform_f2_reg_id) 
	VALUES(?,?,?,?);`
	_, err = db.Exec(sql, r.IdChargeOnForm1,r.ProductIdentity,r.ProductInformF1RegId,r.ProductInformF2RegId)
	return err
}

func (r *ChargeOnForm1Content) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE charge_on_form1_content SET id_charge_on_form1=?,product_identity=?,product_inform_f1_reg_id=?,product_inform_f2_reg_id=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdChargeOnForm1,r.ProductIdentity,r.ProductInformF1RegId,r.ProductInformF2RegId,r.Id)
	return err
}

func (r *ChargeOnForm1Content) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE charge_on_form1_content WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ChargeOnForm1Content) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from charge_on_form1_content WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ChargeOnForm1Content) All(db *sqlx.DB) (out ChargeOnForm1ContentSlice, err error) {
	out = make(ChargeOnForm1ContentSlice, 0)
	sql := `select * from charge_on_form1_content order by id;`
	err = db.Select(&out, sql)
	return out, err
}
