package output

type ProductionForm1Content struct {
    Id	int64	`db:"id"` 
    IdProductionForm1	int	`db:"id_production_form1"` 
    ProductIdentity	string	`db:"product_identity"` 
    ProductInformF1RegId	string	`db:"product_inform_f1_reg_id"` 
    ProductInformF2RegId	string	`db:"product_inform_f2_reg_id"` 
}

type ProductionForm1ContentSlice []*ProductionForm1Content

func (r *ProductionForm1Content) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO production_form1_content (id_production_form1,product_identity,product_inform_f1_reg_id,product_inform_f2_reg_id) 
	VALUES(?,?,?,?);`
	_, err = db.Exec(sql, r.IdProductionForm1,r.ProductIdentity,r.ProductInformF1RegId,r.ProductInformF2RegId)
	return err
}

func (r *ProductionForm1Content) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE production_form1_content SET id_production_form1=?,product_identity=?,product_inform_f1_reg_id=?,product_inform_f2_reg_id=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdProductionForm1,r.ProductIdentity,r.ProductInformF1RegId,r.ProductInformF2RegId,r.Id)
	return err
}

func (r *ProductionForm1Content) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE production_form1_content WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ProductionForm1Content) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from production_form1_content WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ProductionForm1Content) All(db *sqlx.DB) (out ProductionForm1ContentSlice, err error) {
	out = make(ProductionForm1ContentSlice, 0)
	sql := `select * from production_form1_content order by id;`
	err = db.Select(&out, sql)
	return out, err
}
