package output

type Form1Egais struct {
    Id	int64	`db:"id"` 
    IdRequests	int	`db:"id_requests"` 
    OriginalDocNumber	string	`db:"original_doc_number"` 
    OriginalDocDate	string	`db:"original_doc_date"` 
    BottlingDate	string	`db:"bottling_date"` 
    EgaisNumber	string	`db:"egais_number"` 
    EgaisDate	string	`db:"egais_date"` 
    GtdNumber	string	`db:"gtd_number"` 
    GtdDate	string	`db:"gtd_date"` 
    ClientType	string	`db:"client_type"` 
    ClientRegId	string	`db:"client_reg_id"` 
    ClientInn	string	`db:"client_inn"` 
    ClientKpp	string	`db:"client_kpp"` 
    ClientFullName	string	`db:"client_full_name"` 
    ClientShortName	string	`db:"client_short_name"` 
    ClientCountryCode	string	`db:"client_country_code"` 
    ClientRegionCode	string	`db:"client_region_code"` 
    ClientDescription	string	`db:"client_description"` 
    ProductFullName	string	`db:"product_full_name"` 
    ProductCapacity	string	`db:"product_capacity"` 
    ProductAlcVolume	string	`db:"product_alc_volume"` 
    ProductAlcCode	string	`db:"product_alc_code"` 
    ProductCode	string	`db:"product_code"` 
    ProductUnitType	string	`db:"product_unit_type"` 
    ProductQuantity	string	`db:"product_quantity"` 
    ProductInformF1RegId	string	`db:"product_inform_f1_reg_id"` 
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

type Form1EgaisSlice []*Form1Egais

func (r *Form1Egais) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO form1_egais (id_requests,original_doc_number,original_doc_date,bottling_date,egais_number,egais_date,gtd_number,gtd_date,client_type,client_reg_id,client_inn,client_kpp,client_full_name,client_short_name,client_country_code,client_region_code,client_description,product_full_name,product_capacity,product_alc_volume,product_alc_code,product_code,product_unit_type,product_quantity,product_inform_f1_reg_id,producer_type,producer_client_reg_id,producer_inn,producer_kpp,producer_full_name,producer_short_name,producer_country_code,producer_region_code,producer_description) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdRequests,r.OriginalDocNumber,r.OriginalDocDate,r.BottlingDate,r.EgaisNumber,r.EgaisDate,r.GtdNumber,r.GtdDate,r.ClientType,r.ClientRegId,r.ClientInn,r.ClientKpp,r.ClientFullName,r.ClientShortName,r.ClientCountryCode,r.ClientRegionCode,r.ClientDescription,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProductQuantity,r.ProductInformF1RegId,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription)
	return err
}

func (r *Form1Egais) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE form1_egais SET id_requests=?,original_doc_number=?,original_doc_date=?,bottling_date=?,egais_number=?,egais_date=?,gtd_number=?,gtd_date=?,client_type=?,client_reg_id=?,client_inn=?,client_kpp=?,client_full_name=?,client_short_name=?,client_country_code=?,client_region_code=?,client_description=?,product_full_name=?,product_capacity=?,product_alc_volume=?,product_alc_code=?,product_code=?,product_unit_type=?,product_quantity=?,product_inform_f1_reg_id=?,producer_type=?,producer_client_reg_id=?,producer_inn=?,producer_kpp=?,producer_full_name=?,producer_short_name=?,producer_country_code=?,producer_region_code=?,producer_description=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdRequests,r.OriginalDocNumber,r.OriginalDocDate,r.BottlingDate,r.EgaisNumber,r.EgaisDate,r.GtdNumber,r.GtdDate,r.ClientType,r.ClientRegId,r.ClientInn,r.ClientKpp,r.ClientFullName,r.ClientShortName,r.ClientCountryCode,r.ClientRegionCode,r.ClientDescription,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProductQuantity,r.ProductInformF1RegId,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription,r.Id)
	return err
}

func (r *Form1Egais) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE form1_egais WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *Form1Egais) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from form1_egais WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *Form1Egais) All(db *sqlx.DB) (out Form1EgaisSlice, err error) {
	out = make(Form1EgaisSlice, 0)
	sql := `select * from form1_egais order by id;`
	err = db.Select(&out, sql)
	return out, err
}
