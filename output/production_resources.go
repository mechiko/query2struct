package output

type ProductionResources struct {
    Id	int64	`db:"id"` 
    IdProductionReports	int	`db:"id_production_reports"` 
    IdProductionProducts	int	`db:"id_production_products"` 
    ResourceFullName	string	`db:"resource_full_name"` 
    ResourceCapacity	string	`db:"resource_capacity"` 
    ResourceAlcVolume	string	`db:"resource_alc_volume"` 
    ResourceAlcCode	string	`db:"resource_alc_code"` 
    ResourceCode	string	`db:"resource_code"` 
    ResourceUnitType	string	`db:"resource_unit_type"` 
    ResourceIdentity	string	`db:"resource_identity"` 
    ResourceQuantity	string	`db:"resource_quantity"` 
    ResourceInformF1RegId	string	`db:"resource_inform_f1_reg_id"` 
    ResourceInformF2RegId	string	`db:"resource_inform_f2_reg_id"` 
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

type ProductionResourcesSlice []*ProductionResources

func (r *ProductionResources) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO production_resources (id_production_reports,id_production_products,resource_full_name,resource_capacity,resource_alc_volume,resource_alc_code,resource_code,resource_unit_type,resource_identity,resource_quantity,resource_inform_f1_reg_id,resource_inform_f2_reg_id,producer_type,producer_client_reg_id,producer_inn,producer_kpp,producer_full_name,producer_short_name,producer_country_code,producer_region_code,producer_description) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdProductionReports,r.IdProductionProducts,r.ResourceFullName,r.ResourceCapacity,r.ResourceAlcVolume,r.ResourceAlcCode,r.ResourceCode,r.ResourceUnitType,r.ResourceIdentity,r.ResourceQuantity,r.ResourceInformF1RegId,r.ResourceInformF2RegId,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription)
	return err
}

func (r *ProductionResources) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE production_resources SET id_production_reports=?,id_production_products=?,resource_full_name=?,resource_capacity=?,resource_alc_volume=?,resource_alc_code=?,resource_code=?,resource_unit_type=?,resource_identity=?,resource_quantity=?,resource_inform_f1_reg_id=?,resource_inform_f2_reg_id=?,producer_type=?,producer_client_reg_id=?,producer_inn=?,producer_kpp=?,producer_full_name=?,producer_short_name=?,producer_country_code=?,producer_region_code=?,producer_description=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdProductionReports,r.IdProductionProducts,r.ResourceFullName,r.ResourceCapacity,r.ResourceAlcVolume,r.ResourceAlcCode,r.ResourceCode,r.ResourceUnitType,r.ResourceIdentity,r.ResourceQuantity,r.ResourceInformF1RegId,r.ResourceInformF2RegId,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription,r.Id)
	return err
}

func (r *ProductionResources) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE production_resources WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ProductionResources) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from production_resources WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ProductionResources) All(db *sqlx.DB) (out ProductionResourcesSlice, err error) {
	out = make(ProductionResourcesSlice, 0)
	sql := `select * from production_resources order by id;`
	err = db.Select(&out, sql)
	return out, err
}
