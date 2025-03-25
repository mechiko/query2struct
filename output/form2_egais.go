package output

type Form2Egais struct {
    Id	int64	`db:"id"` 
    IdRequests	int	`db:"id_requests"` 
    TtnNumber	string	`db:"ttn_number"` 
    TtnDate	string	`db:"ttn_date"` 
    ShippingDate	string	`db:"shipping_date"` 
    ShipperType	string	`db:"shipper_type"` 
    ShipperClientRegId	string	`db:"shipper_client_reg_id"` 
    ShipperInn	string	`db:"shipper_inn"` 
    ShipperKpp	string	`db:"shipper_kpp"` 
    ShipperFullName	string	`db:"shipper_full_name"` 
    ShipperShortName	string	`db:"shipper_short_name"` 
    ShipperCountryCode	string	`db:"shipper_country_code"` 
    ShipperRegionCode	string	`db:"shipper_region_code"` 
    ShipperDescription	string	`db:"shipper_description"` 
    ConsigneeType	string	`db:"consignee_type"` 
    ConsigneeClientRegId	string	`db:"consignee_client_reg_id"` 
    ConsigneeInn	string	`db:"consignee_inn"` 
    ConsigneeKpp	string	`db:"consignee_kpp"` 
    ConsigneeFullName	string	`db:"consignee_full_name"` 
    ConsigneeShortName	string	`db:"consignee_short_name"` 
    ConsigneeCountryCode	string	`db:"consignee_country_code"` 
    ConsigneeRegionCode	string	`db:"consignee_region_code"` 
    ConsigneeDescription	string	`db:"consignee_description"` 
    ProductFullName	string	`db:"product_full_name"` 
    ProductCapacity	string	`db:"product_capacity"` 
    ProductAlcVolume	string	`db:"product_alc_volume"` 
    ProductAlcCode	string	`db:"product_alc_code"` 
    ProductCode	string	`db:"product_code"` 
    ProductUnitType	string	`db:"product_unit_type"` 
    ProductQuantity	string	`db:"product_quantity"` 
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

type Form2EgaisSlice []*Form2Egais

func (r *Form2Egais) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO form2_egais (id_requests,ttn_number,ttn_date,shipping_date,shipper_type,shipper_client_reg_id,shipper_inn,shipper_kpp,shipper_full_name,shipper_short_name,shipper_country_code,shipper_region_code,shipper_description,consignee_type,consignee_client_reg_id,consignee_inn,consignee_kpp,consignee_full_name,consignee_short_name,consignee_country_code,consignee_region_code,consignee_description,product_full_name,product_capacity,product_alc_volume,product_alc_code,product_code,product_unit_type,product_quantity,product_inform_f2_reg_id,producer_type,producer_client_reg_id,producer_inn,producer_kpp,producer_full_name,producer_short_name,producer_country_code,producer_region_code,producer_description) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdRequests,r.TtnNumber,r.TtnDate,r.ShippingDate,r.ShipperType,r.ShipperClientRegId,r.ShipperInn,r.ShipperKpp,r.ShipperFullName,r.ShipperShortName,r.ShipperCountryCode,r.ShipperRegionCode,r.ShipperDescription,r.ConsigneeType,r.ConsigneeClientRegId,r.ConsigneeInn,r.ConsigneeKpp,r.ConsigneeFullName,r.ConsigneeShortName,r.ConsigneeCountryCode,r.ConsigneeRegionCode,r.ConsigneeDescription,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProductQuantity,r.ProductInformF2RegId,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription)
	return err
}

func (r *Form2Egais) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE form2_egais SET id_requests=?,ttn_number=?,ttn_date=?,shipping_date=?,shipper_type=?,shipper_client_reg_id=?,shipper_inn=?,shipper_kpp=?,shipper_full_name=?,shipper_short_name=?,shipper_country_code=?,shipper_region_code=?,shipper_description=?,consignee_type=?,consignee_client_reg_id=?,consignee_inn=?,consignee_kpp=?,consignee_full_name=?,consignee_short_name=?,consignee_country_code=?,consignee_region_code=?,consignee_description=?,product_full_name=?,product_capacity=?,product_alc_volume=?,product_alc_code=?,product_code=?,product_unit_type=?,product_quantity=?,product_inform_f2_reg_id=?,producer_type=?,producer_client_reg_id=?,producer_inn=?,producer_kpp=?,producer_full_name=?,producer_short_name=?,producer_country_code=?,producer_region_code=?,producer_description=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdRequests,r.TtnNumber,r.TtnDate,r.ShippingDate,r.ShipperType,r.ShipperClientRegId,r.ShipperInn,r.ShipperKpp,r.ShipperFullName,r.ShipperShortName,r.ShipperCountryCode,r.ShipperRegionCode,r.ShipperDescription,r.ConsigneeType,r.ConsigneeClientRegId,r.ConsigneeInn,r.ConsigneeKpp,r.ConsigneeFullName,r.ConsigneeShortName,r.ConsigneeCountryCode,r.ConsigneeRegionCode,r.ConsigneeDescription,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProductQuantity,r.ProductInformF2RegId,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription,r.Id)
	return err
}

func (r *Form2Egais) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE form2_egais WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *Form2Egais) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from form2_egais WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *Form2Egais) All(db *sqlx.DB) (out Form2EgaisSlice, err error) {
	out = make(Form2EgaisSlice, 0)
	sql := `select * from form2_egais order by id;`
	err = db.Select(&out, sql)
	return out, err
}
