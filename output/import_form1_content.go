package output

type ImportForm1Content struct {
    Id	int64	`db:"id"` 
    IdImportForm1	int	`db:"id_import_form1"` 
    ProductIdentity	string	`db:"product_identity"` 
    ProductInformF1RegId	string	`db:"product_inform_f1_reg_id"` 
    ProductInformF2RegId	string	`db:"product_inform_f2_reg_id"` 
}

type ImportForm1ContentSlice []*ImportForm1Content

func (r *ImportForm1Content) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO import_form1_content (id_import_form1,product_identity,product_inform_f1_reg_id,product_inform_f2_reg_id) 
	VALUES(?,?,?,?);`
	_, err = db.Exec(sql, r.IdImportForm1,r.ProductIdentity,r.ProductInformF1RegId,r.ProductInformF2RegId)
	return err
}

func (r *ImportForm1Content) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE import_form1_content SET id_import_form1=?,product_identity=?,product_inform_f1_reg_id=?,product_inform_f2_reg_id=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdImportForm1,r.ProductIdentity,r.ProductInformF1RegId,r.ProductInformF2RegId,r.Id)
	return err
}

func (r *ImportForm1Content) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE import_form1_content WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ImportForm1Content) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from import_form1_content WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ImportForm1Content) All(db *sqlx.DB) (out ImportForm1ContentSlice, err error) {
	out = make(ImportForm1ContentSlice, 0)
	sql := `select * from import_form1_content order by id;`
	err = db.Select(&out, sql)
	return out, err
}
