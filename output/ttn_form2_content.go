package output

type TtnForm2Content struct {
    Id	int64	`db:"id"` 
    IdTtnForm2	int	`db:"id_ttn_form2"` 
    ProductIdentity	string	`db:"product_identity"` 
    ProductInformF2RegId	string	`db:"product_inform_f2_reg_id"` 
    BottlingDate	string	`db:"bottling_date"` 
}

type TtnForm2ContentSlice []*TtnForm2Content

func (r *TtnForm2Content) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_form2_content (id_ttn_form2,product_identity,product_inform_f2_reg_id,bottling_date) 
	VALUES(?,?,?,?);`
	_, err = db.Exec(sql, r.IdTtnForm2,r.ProductIdentity,r.ProductInformF2RegId,r.BottlingDate)
	return err
}

func (r *TtnForm2Content) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_form2_content SET id_ttn_form2=?,product_identity=?,product_inform_f2_reg_id=?,bottling_date=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtnForm2,r.ProductIdentity,r.ProductInformF2RegId,r.BottlingDate,r.Id)
	return err
}

func (r *TtnForm2Content) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_form2_content WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnForm2Content) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_form2_content WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnForm2Content) All(db *sqlx.DB) (out TtnForm2ContentSlice, err error) {
	out = make(TtnForm2ContentSlice, 0)
	sql := `select * from ttn_form2_content order by id;`
	err = db.Select(&out, sql)
	return out, err
}
