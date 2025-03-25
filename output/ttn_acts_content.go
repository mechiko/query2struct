package output

type TtnActsContent struct {
    Id	int64	`db:"id"` 
    IdTtnActs	int	`db:"id_ttn_acts"` 
    ProductIdentity	string	`db:"product_identity"` 
    ProductInformF2RegId	string	`db:"product_inform_f2_reg_id"` 
    ProductQuantity	string	`db:"product_quantity"` 
}

type TtnActsContentSlice []*TtnActsContent

func (r *TtnActsContent) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_acts_content (id_ttn_acts,product_identity,product_inform_f2_reg_id,product_quantity) 
	VALUES(?,?,?,?);`
	_, err = db.Exec(sql, r.IdTtnActs,r.ProductIdentity,r.ProductInformF2RegId,r.ProductQuantity)
	return err
}

func (r *TtnActsContent) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_acts_content SET id_ttn_acts=?,product_identity=?,product_inform_f2_reg_id=?,product_quantity=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtnActs,r.ProductIdentity,r.ProductInformF2RegId,r.ProductQuantity,r.Id)
	return err
}

func (r *TtnActsContent) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_acts_content WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnActsContent) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_acts_content WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnActsContent) All(db *sqlx.DB) (out TtnActsContentSlice, err error) {
	out = make(TtnActsContentSlice, 0)
	sql := `select * from ttn_acts_content order by id;`
	err = db.Select(&out, sql)
	return out, err
}
