package output

type ChargeOnProducts struct {
    Id	int64	`db:"id"` 
    IdChargeOnActs	int	`db:"id_charge_on_acts"` 
    ProductFullName	string	`db:"product_full_name"` 
    ProductCapacity	string	`db:"product_capacity"` 
    ProductAlcVolume	string	`db:"product_alc_volume"` 
    ProductAlcCode	string	`db:"product_alc_code"` 
    ProductCode	string	`db:"product_code"` 
    ProductUnitType	string	`db:"product_unit_type"` 
    ProductIdentity	string	`db:"product_identity"` 
    ProductQuantity	string	`db:"product_quantity"` 
    ProducerType	string	`db:"producer_type"` 
    ProducerClientRegId	string	`db:"producer_client_reg_id"` 
    ProducerInn	string	`db:"producer_inn"` 
    ProducerKpp	string	`db:"producer_kpp"` 
    ProducerFullName	string	`db:"producer_full_name"` 
    ProducerShortName	string	`db:"producer_short_name"` 
    ProducerCountryCode	string	`db:"producer_country_code"` 
    ProducerRegionCode	string	`db:"producer_region_code"` 
    ProducerDescription	string	`db:"producer_description"` 
    Form1Quantity	string	`db:"form1_quantity"` 
    Form1BottlingDate	string	`db:"form1_bottling_date"` 
    Form1TtnNumber	string	`db:"form1_ttn_number"` 
    Form1TtnDate	string	`db:"form1_ttn_date"` 
    Form1EgaisFixNumber	string	`db:"form1_egais_fix_number"` 
    Form1EgaisFixDate	string	`db:"form1_egais_fix_date"` 
}

type ChargeOnProductsSlice []*ChargeOnProducts

func (r *ChargeOnProducts) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO charge_on_products (id_charge_on_acts,product_full_name,product_capacity,product_alc_volume,product_alc_code,product_code,product_unit_type,product_identity,product_quantity,producer_type,producer_client_reg_id,producer_inn,producer_kpp,producer_full_name,producer_short_name,producer_country_code,producer_region_code,producer_description,form1_quantity,form1_bottling_date,form1_ttn_number,form1_ttn_date,form1_egais_fix_number,form1_egais_fix_date) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdChargeOnActs,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProductIdentity,r.ProductQuantity,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription,r.Form1Quantity,r.Form1BottlingDate,r.Form1TtnNumber,r.Form1TtnDate,r.Form1EgaisFixNumber,r.Form1EgaisFixDate)
	return err
}

func (r *ChargeOnProducts) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE charge_on_products SET id_charge_on_acts=?,product_full_name=?,product_capacity=?,product_alc_volume=?,product_alc_code=?,product_code=?,product_unit_type=?,product_identity=?,product_quantity=?,producer_type=?,producer_client_reg_id=?,producer_inn=?,producer_kpp=?,producer_full_name=?,producer_short_name=?,producer_country_code=?,producer_region_code=?,producer_description=?,form1_quantity=?,form1_bottling_date=?,form1_ttn_number=?,form1_ttn_date=?,form1_egais_fix_number=?,form1_egais_fix_date=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdChargeOnActs,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProductIdentity,r.ProductQuantity,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription,r.Form1Quantity,r.Form1BottlingDate,r.Form1TtnNumber,r.Form1TtnDate,r.Form1EgaisFixNumber,r.Form1EgaisFixDate,r.Id)
	return err
}

func (r *ChargeOnProducts) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE charge_on_products WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ChargeOnProducts) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from charge_on_products WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ChargeOnProducts) All(db *sqlx.DB) (out ChargeOnProductsSlice, err error) {
	out = make(ChargeOnProductsSlice, 0)
	sql := `select * from charge_on_products order by id;`
	err = db.Select(&out, sql)
	return out, err
}
