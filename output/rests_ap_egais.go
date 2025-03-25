package output

type RestsApEgais struct {
    Id	int64	`db:"id"` 
    IdRequests	int	`db:"id_requests"` 
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
    RestsApDate	string	`db:"rests_ap_date"` 
}

type RestsApEgaisSlice []*RestsApEgais

func (r *RestsApEgais) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO rests_ap_egais (id_requests,product_full_name,product_capacity,product_alc_volume,product_alc_code,product_code,product_unit_type,product_quantity,product_inform_f1_reg_id,product_inform_f2_reg_id,producer_type,producer_client_reg_id,producer_inn,producer_kpp,producer_full_name,producer_short_name,producer_country_code,producer_region_code,producer_description,rests_ap_date) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdRequests,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProductQuantity,r.ProductInformF1RegId,r.ProductInformF2RegId,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription,r.RestsApDate)
	return err
}

func (r *RestsApEgais) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE rests_ap_egais SET id_requests=?,product_full_name=?,product_capacity=?,product_alc_volume=?,product_alc_code=?,product_code=?,product_unit_type=?,product_quantity=?,product_inform_f1_reg_id=?,product_inform_f2_reg_id=?,producer_type=?,producer_client_reg_id=?,producer_inn=?,producer_kpp=?,producer_full_name=?,producer_short_name=?,producer_country_code=?,producer_region_code=?,producer_description=?,rests_ap_date=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdRequests,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProductQuantity,r.ProductInformF1RegId,r.ProductInformF2RegId,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription,r.RestsApDate,r.Id)
	return err
}

func (r *RestsApEgais) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE rests_ap_egais WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *RestsApEgais) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from rests_ap_egais WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *RestsApEgais) All(db *sqlx.DB) (out RestsApEgaisSlice, err error) {
	out = make(RestsApEgaisSlice, 0)
	sql := `select * from rests_ap_egais order by id;`
	err = db.Select(&out, sql)
	return out, err
}
