package output

type SpEgais struct {
    Id	int64	`db:"id"` 
    IdRequests	int	`db:"id_requests"` 
    ProductFullName	string	`db:"product_full_name"` 
    ProductCapacity	string	`db:"product_capacity"` 
    ProductAlcVolume	string	`db:"product_alc_volume"` 
    ProductAlcCode	string	`db:"product_alc_code"` 
    ProductCode	string	`db:"product_code"` 
    ProductUnitType	string	`db:"product_unit_type"` 
}

type SpEgaisSlice []*SpEgais

func (r *SpEgais) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO sp_egais (id_requests,product_full_name,product_capacity,product_alc_volume,product_alc_code,product_code,product_unit_type) 
	VALUES(?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdRequests,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType)
	return err
}

func (r *SpEgais) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE sp_egais SET id_requests=?,product_full_name=?,product_capacity=?,product_alc_volume=?,product_alc_code=?,product_code=?,product_unit_type=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdRequests,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.Id)
	return err
}

func (r *SpEgais) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE sp_egais WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *SpEgais) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from sp_egais WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *SpEgais) All(db *sqlx.DB) (out SpEgaisSlice, err error) {
	out = make(SpEgaisSlice, 0)
	sql := `select * from sp_egais order by id;`
	err = db.Select(&out, sql)
	return out, err
}
