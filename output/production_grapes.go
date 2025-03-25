package output

type ProductionGrapes struct {
    Id	int64	`db:"id"` 
    IdProductionReports	int	`db:"id_production_reports"` 
    IdProductionProducts	int	`db:"id_production_products"` 
    GrapeSort	string	`db:"grape_sort"` 
    GrapeCode	string	`db:"grape_code"` 
    GrapeWeight	string	`db:"grape_weight"` 
    VineyardNumber	string	`db:"vineyard_number"` 
    ReceiptDate	string	`db:"receipt_date"` 
    GrapeIdentity	string	`db:"grape_identity"` 
    SupplierType	string	`db:"supplier_type"` 
    SupplierClientRegId	string	`db:"supplier_client_reg_id"` 
    SupplierInn	string	`db:"supplier_inn"` 
    SupplierKpp	string	`db:"supplier_kpp"` 
    SupplierFullName	string	`db:"supplier_full_name"` 
    SupplierShortName	string	`db:"supplier_short_name"` 
    SupplierCountryCode	string	`db:"supplier_country_code"` 
    SupplierRegionCode	string	`db:"supplier_region_code"` 
    SupplierDescription	string	`db:"supplier_description"` 
}

type ProductionGrapesSlice []*ProductionGrapes

func (r *ProductionGrapes) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO production_grapes (id_production_reports,id_production_products,grape_sort,grape_code,grape_weight,vineyard_number,receipt_date,grape_identity,supplier_type,supplier_client_reg_id,supplier_inn,supplier_kpp,supplier_full_name,supplier_short_name,supplier_country_code,supplier_region_code,supplier_description) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdProductionReports,r.IdProductionProducts,r.GrapeSort,r.GrapeCode,r.GrapeWeight,r.VineyardNumber,r.ReceiptDate,r.GrapeIdentity,r.SupplierType,r.SupplierClientRegId,r.SupplierInn,r.SupplierKpp,r.SupplierFullName,r.SupplierShortName,r.SupplierCountryCode,r.SupplierRegionCode,r.SupplierDescription)
	return err
}

func (r *ProductionGrapes) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE production_grapes SET id_production_reports=?,id_production_products=?,grape_sort=?,grape_code=?,grape_weight=?,vineyard_number=?,receipt_date=?,grape_identity=?,supplier_type=?,supplier_client_reg_id=?,supplier_inn=?,supplier_kpp=?,supplier_full_name=?,supplier_short_name=?,supplier_country_code=?,supplier_region_code=?,supplier_description=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdProductionReports,r.IdProductionProducts,r.GrapeSort,r.GrapeCode,r.GrapeWeight,r.VineyardNumber,r.ReceiptDate,r.GrapeIdentity,r.SupplierType,r.SupplierClientRegId,r.SupplierInn,r.SupplierKpp,r.SupplierFullName,r.SupplierShortName,r.SupplierCountryCode,r.SupplierRegionCode,r.SupplierDescription,r.Id)
	return err
}

func (r *ProductionGrapes) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE production_grapes WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ProductionGrapes) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from production_grapes WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ProductionGrapes) All(db *sqlx.DB) (out ProductionGrapesSlice, err error) {
	out = make(ProductionGrapesSlice, 0)
	sql := `select * from production_grapes order by id;`
	err = db.Select(&out, sql)
	return out, err
}
