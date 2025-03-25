package output

type ChargeOnProductsMarks struct {
    Id	int64	`db:"id"` 
    IdChargeOnProducts	int	`db:"id_charge_on_products"` 
    Mark	string	`db:"mark"` 
}

type ChargeOnProductsMarksSlice []*ChargeOnProductsMarks

func (r *ChargeOnProductsMarks) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO charge_on_products_marks (id_charge_on_products,mark) 
	VALUES(?,?);`
	_, err = db.Exec(sql, r.IdChargeOnProducts,r.Mark)
	return err
}

func (r *ChargeOnProductsMarks) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE charge_on_products_marks SET id_charge_on_products=?,mark=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdChargeOnProducts,r.Mark,r.Id)
	return err
}

func (r *ChargeOnProductsMarks) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE charge_on_products_marks WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ChargeOnProductsMarks) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from charge_on_products_marks WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ChargeOnProductsMarks) All(db *sqlx.DB) (out ChargeOnProductsMarksSlice, err error) {
	out = make(ChargeOnProductsMarksSlice, 0)
	sql := `select * from charge_on_products_marks order by id;`
	err = db.Select(&out, sql)
	return out, err
}
