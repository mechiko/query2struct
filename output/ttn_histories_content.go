package output

type TtnHistoriesContent struct {
    Id	int64	`db:"id"` 
    IdTtnHistories	int	`db:"id_ttn_histories"` 
    Level	string	`db:"level"` 
    ProductForm2	string	`db:"product_form2"` 
    ProductForm2Parent	string	`db:"product_form2_parent"` 
    ProductQuantity	string	`db:"product_quantity"` 
    TtnShipper	string	`db:"ttn_shipper"` 
    TtnConsignee	string	`db:"ttn_consignee"` 
    TtnRegId	string	`db:"ttn_reg_id"` 
}

type TtnHistoriesContentSlice []*TtnHistoriesContent

func (r *TtnHistoriesContent) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_histories_content (id_ttn_histories,level,product_form2,product_form2_parent,product_quantity,ttn_shipper,ttn_consignee,ttn_reg_id) 
	VALUES(?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdTtnHistories,r.Level,r.ProductForm2,r.ProductForm2Parent,r.ProductQuantity,r.TtnShipper,r.TtnConsignee,r.TtnRegId)
	return err
}

func (r *TtnHistoriesContent) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_histories_content SET id_ttn_histories=?,level=?,product_form2=?,product_form2_parent=?,product_quantity=?,ttn_shipper=?,ttn_consignee=?,ttn_reg_id=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtnHistories,r.Level,r.ProductForm2,r.ProductForm2Parent,r.ProductQuantity,r.TtnShipper,r.TtnConsignee,r.TtnRegId,r.Id)
	return err
}

func (r *TtnHistoriesContent) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_histories_content WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnHistoriesContent) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_histories_content WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnHistoriesContent) All(db *sqlx.DB) (out TtnHistoriesContentSlice, err error) {
	out = make(TtnHistoriesContentSlice, 0)
	sql := `select * from ttn_histories_content order by id;`
	err = db.Select(&out, sql)
	return out, err
}
