package output

type SspLocal struct {
    Id	int64	`db:"id"` 
    ProductFullName	string	`db:"product_full_name"` 
    ProductCapacity	string	`db:"product_capacity"` 
    ProductAlcVolume	string	`db:"product_alc_volume"` 
    ProductAlcCode	string	`db:"product_alc_code"` 
    ProductCode	string	`db:"product_code"` 
    ProductUnitType	string	`db:"product_unit_type"` 
    ProducerType	string	`db:"producer_type"` 
    ProducerClientRegId	string	`db:"producer_client_reg_id"` 
    ProducerInn	string	`db:"producer_inn"` 
    ProducerKpp	string	`db:"producer_kpp"` 
    ProducerFullName	string	`db:"producer_full_name"` 
    ProducerShortName	string	`db:"producer_short_name"` 
    ProducerCountryCode	string	`db:"producer_country_code"` 
    ProducerRegionCode	string	`db:"producer_region_code"` 
    ProducerDescription	string	`db:"producer_description"` 
}

type SspLocalSlice []*SspLocal

func (r *SspLocal) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ssp_local (product_full_name,product_capacity,product_alc_volume,product_alc_code,product_code,product_unit_type,producer_type,producer_client_reg_id,producer_inn,producer_kpp,producer_full_name,producer_short_name,producer_country_code,producer_region_code,producer_description) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription)
	return err
}

func (r *SspLocal) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ssp_local SET product_full_name=?,product_capacity=?,product_alc_volume=?,product_alc_code=?,product_code=?,product_unit_type=?,producer_type=?,producer_client_reg_id=?,producer_inn=?,producer_kpp=?,producer_full_name=?,producer_short_name=?,producer_country_code=?,producer_region_code=?,producer_description=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription,r.Id)
	return err
}

func (r *SspLocal) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ssp_local WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *SspLocal) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ssp_local WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *SspLocal) All(db *sqlx.DB) (out SspLocalSlice, err error) {
	out = make(SspLocalSlice, 0)
	sql := `select * from ssp_local order by id;`
	err = db.Select(&out, sql)
	return out, err
}
