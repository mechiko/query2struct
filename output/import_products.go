package output

type ImportProducts struct {
    Id	int64	`db:"id"` 
    IdImportReports	int	`db:"id_import_reports"` 
    ProductFullName	string	`db:"product_full_name"` 
    ProductCapacity	string	`db:"product_capacity"` 
    ProductAlcVolume	string	`db:"product_alc_volume"` 
    ProductAlcVolumeMin	string	`db:"product_alc_volume_min"` 
    ProductAlcVolumeMax	string	`db:"product_alc_volume_max"` 
    ProductAlcCode	string	`db:"product_alc_code"` 
    ProductCode	string	`db:"product_code"` 
    ProductUnitType	string	`db:"product_unit_type"` 
    ProductIdentity	string	`db:"product_identity"` 
    ProductQuantity	string	`db:"product_quantity"` 
    ProductPlannedImportId	string	`db:"product_planned_import_id"` 
    ProductParty	string	`db:"product_party"` 
    ProductComment	string	`db:"product_comment"` 
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

type ImportProductsSlice []*ImportProducts

func (r *ImportProducts) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO import_products (id_import_reports,product_full_name,product_capacity,product_alc_volume,product_alc_volume_min,product_alc_volume_max,product_alc_code,product_code,product_unit_type,product_identity,product_quantity,product_planned_import_id,product_party,product_comment,producer_type,producer_client_reg_id,producer_inn,producer_kpp,producer_full_name,producer_short_name,producer_country_code,producer_region_code,producer_description) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdImportReports,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcVolumeMin,r.ProductAlcVolumeMax,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProductIdentity,r.ProductQuantity,r.ProductPlannedImportId,r.ProductParty,r.ProductComment,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription)
	return err
}

func (r *ImportProducts) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE import_products SET id_import_reports=?,product_full_name=?,product_capacity=?,product_alc_volume=?,product_alc_volume_min=?,product_alc_volume_max=?,product_alc_code=?,product_code=?,product_unit_type=?,product_identity=?,product_quantity=?,product_planned_import_id=?,product_party=?,product_comment=?,producer_type=?,producer_client_reg_id=?,producer_inn=?,producer_kpp=?,producer_full_name=?,producer_short_name=?,producer_country_code=?,producer_region_code=?,producer_description=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdImportReports,r.ProductFullName,r.ProductCapacity,r.ProductAlcVolume,r.ProductAlcVolumeMin,r.ProductAlcVolumeMax,r.ProductAlcCode,r.ProductCode,r.ProductUnitType,r.ProductIdentity,r.ProductQuantity,r.ProductPlannedImportId,r.ProductParty,r.ProductComment,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription,r.Id)
	return err
}

func (r *ImportProducts) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE import_products WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ImportProducts) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from import_products WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ImportProducts) All(db *sqlx.DB) (out ImportProductsSlice, err error) {
	out = make(ImportProductsSlice, 0)
	sql := `select * from import_products order by id;`
	err = db.Select(&out, sql)
	return out, err
}
