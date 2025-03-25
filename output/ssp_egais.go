package output

type SspEgais struct {
    Id	int64	`db:"id"` 
    IdRequests	int	`db:"id_requests"` 
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

type SspEgaisSlice []*SspEgais

func (r *SspEgais) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ssp_egais (id_requests,product_full_name,product_capacity,product_alc_volume,product_alc_code,product_code,product_unit_type,producer_type,producer_client_reg_id,producer_inn,producer_kpp,producer_full_name,producer_short_name,producer_country_code,producer_region_code,producer_description) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdRequests,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription)
	return err
}

func (r *SspEgais) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ssp_egais SET id_requests=?,product_full_name=?,product_capacity=?,product_alc_volume=?,product_alc_code=?,product_code=?,product_unit_type=?,producer_type=?,producer_client_reg_id=?,producer_inn=?,producer_kpp=?,producer_full_name=?,producer_short_name=?,producer_country_code=?,producer_region_code=?,producer_description=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdRequests,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription,r.Id)
	return err
}

func (r *SspEgais) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ssp_egais WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *SspEgais) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ssp_egais WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *SspEgais) All(db *sqlx.DB) (out SspEgaisSlice, err error) {
	out = make(SspEgaisSlice, 0)
	sql := `select * from ssp_egais order by id;`
	err = db.Select(&out, sql)
	return out, err
}
