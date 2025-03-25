package output

type SpLocal struct {
    Id	int64	`db:"id"` 
    ProductFullName	string	`db:"product_full_name"` 
    ProductCapacity	string	`db:"product_capacity"` 
    ProductAlcVolume	string	`db:"product_alc_volume"` 
    ProductAlcCode	string	`db:"product_alc_code"` 
    ProductCode	string	`db:"product_code"` 
    ProductUnitType	string	`db:"product_unit_type"` 
}

type SpLocalSlice []*SpLocal

func (r *SpLocal) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO sp_local (product_full_name,product_capacity,product_alc_volume,product_alc_code,product_code,product_unit_type) 
	VALUES(?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType)
	return err
}

func (r *SpLocal) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE sp_local SET product_full_name=?,product_capacity=?,product_alc_volume=?,product_alc_code=?,product_code=?,product_unit_type=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.Id)
	return err
}

func (r *SpLocal) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE sp_local WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *SpLocal) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from sp_local WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *SpLocal) All(db *sqlx.DB) (out SpLocalSlice, err error) {
	out = make(SpLocalSlice, 0)
	sql := `select * from sp_local order by id;`
	err = db.Select(&out, sql)
	return out, err
}
