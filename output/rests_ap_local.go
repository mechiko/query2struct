package output

type RestsApLocal struct {
    Id	int64	`db:"id"` 
    ProductFullName	string	`db:"product_full_name"` 
    ProductCapacity	string	`db:"product_capacity"` 
    ProductAlcVolume	string	`db:"product_alc_volume"` 
    ProductAlcCode	string	`db:"product_alc_code"` 
    ProductCode	string	`db:"product_code"` 
    ProductUnitType	string	`db:"product_unit_type"` 
    ProductQuantity	string	`db:"product_quantity"` 
    ProductInformF1RegId	string	`db:"product_inform_f1_reg_id"` 
    ProductInformF2RegId	string	`db:"product_inform_f2_reg_id"` 
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

type RestsApLocalSlice []*RestsApLocal

func (r *RestsApLocal) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO rests_ap_local (product_full_name,product_capacity,product_alc_volume,product_alc_code,product_code,product_unit_type,product_quantity,product_inform_f1_reg_id,product_inform_f2_reg_id,producer_type,producer_client_reg_id,producer_inn,producer_kpp,producer_full_name,producer_short_name,producer_country_code,producer_region_code,producer_description) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProductQuantity,r.ProductInformF1RegId,r.ProductInformF2RegId,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription)
	return err
}

func (r *RestsApLocal) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE rests_ap_local SET product_full_name=?,product_capacity=?,product_alc_volume=?,product_alc_code=?,product_code=?,product_unit_type=?,product_quantity=?,product_inform_f1_reg_id=?,product_inform_f2_reg_id=?,producer_type=?,producer_client_reg_id=?,producer_inn=?,producer_kpp=?,producer_full_name=?,producer_short_name=?,producer_country_code=?,producer_region_code=?,producer_description=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProductQuantity,r.ProductInformF1RegId,r.ProductInformF2RegId,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription,r.Id)
	return err
}

func (r *RestsApLocal) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE rests_ap_local WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *RestsApLocal) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from rests_ap_local WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *RestsApLocal) All(db *sqlx.DB) (out RestsApLocalSlice, err error) {
	out = make(RestsApLocalSlice, 0)
	sql := `select * from rests_ap_local order by id;`
	err = db.Select(&out, sql)
	return out, err
}
