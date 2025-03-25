package output

type ImportReports struct {
    Id	int64	`db:"id"` 
    CreateDate	string	`db:"create_date"` 
    DocIdentity	string	`db:"doc_identity"` 
    DocNumber	string	`db:"doc_number"` 
    DocDate	string	`db:"doc_date"` 
    DocImportedDate	string	`db:"doc_imported_date"` 
    ContractNumber	string	`db:"contract_number"` 
    ContractDate	string	`db:"contract_date"` 
    GtdNumber	string	`db:"gtd_number"` 
    GtdDate	string	`db:"gtd_date"` 
    DocComment	string	`db:"doc_comment"` 
    ImporterType	string	`db:"importer_type"` 
    ImporterClientRegId	string	`db:"importer_client_reg_id"` 
    ImporterInn	string	`db:"importer_inn"` 
    ImporterKpp	string	`db:"importer_kpp"` 
    ImporterFullName	string	`db:"importer_full_name"` 
    ImporterShortName	string	`db:"importer_short_name"` 
    ImporterCountryCode	string	`db:"importer_country_code"` 
    ImporterRegionCode	string	`db:"importer_region_code"` 
    ImporterDescription	string	`db:"importer_description"` 
    SupplierType	string	`db:"supplier_type"` 
    SupplierClientRegId	string	`db:"supplier_client_reg_id"` 
    SupplierInn	string	`db:"supplier_inn"` 
    SupplierKpp	string	`db:"supplier_kpp"` 
    SupplierFullName	string	`db:"supplier_full_name"` 
    SupplierShortName	string	`db:"supplier_short_name"` 
    SupplierCountryCode	string	`db:"supplier_country_code"` 
    SupplierRegionCode	string	`db:"supplier_region_code"` 
    SupplierDescription	string	`db:"supplier_description"` 
    Version	string	`db:"version"` 
    State	string	`db:"state"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
    Archive	int	`db:"archive"` 
    Xml	string	`db:"xml"` 
}

type ImportReportsSlice []*ImportReports

func (r *ImportReports) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO import_reports (create_date,doc_identity,doc_number,doc_date,doc_imported_date,contract_number,contract_date,gtd_number,gtd_date,doc_comment,importer_type,importer_client_reg_id,importer_inn,importer_kpp,importer_full_name,importer_short_name,importer_country_code,importer_region_code,importer_description,supplier_type,supplier_client_reg_id,supplier_inn,supplier_kpp,supplier_full_name,supplier_short_name,supplier_country_code,supplier_region_code,supplier_description,version,state,status,reply_id,archive,xml) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.CreateDate,r.DocIdentity,r.DocNumber,r.DocDate,r.DocImportedDate,r.ContractNumber,r.ContractDate,r.GtdNumber,r.GtdDate,r.DocComment,r.ImporterType,r.ImporterClientRegId,r.ImporterInn,r.ImporterKpp,r.ImporterFullName,r.ImporterShortName,r.ImporterCountryCode,r.ImporterRegionCode,r.ImporterDescription,r.SupplierType,r.SupplierClientRegId,r.SupplierInn,r.SupplierKpp,r.SupplierFullName,r.SupplierShortName,r.SupplierCountryCode,r.SupplierRegionCode,r.SupplierDescription,r.Version,r.State,r.Status,r.ReplyId,r.Archive,r.Xml)
	return err
}

func (r *ImportReports) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE import_reports SET create_date=?,doc_identity=?,doc_number=?,doc_date=?,doc_imported_date=?,contract_number=?,contract_date=?,gtd_number=?,gtd_date=?,doc_comment=?,importer_type=?,importer_client_reg_id=?,importer_inn=?,importer_kpp=?,importer_full_name=?,importer_short_name=?,importer_country_code=?,importer_region_code=?,importer_description=?,supplier_type=?,supplier_client_reg_id=?,supplier_inn=?,supplier_kpp=?,supplier_full_name=?,supplier_short_name=?,supplier_country_code=?,supplier_region_code=?,supplier_description=?,version=?,state=?,status=?,reply_id=?,archive=?,xml=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.CreateDate,r.DocIdentity,r.DocNumber,r.DocDate,r.DocImportedDate,r.ContractNumber,r.ContractDate,r.GtdNumber,r.GtdDate,r.DocComment,r.ImporterType,r.ImporterClientRegId,r.ImporterInn,r.ImporterKpp,r.ImporterFullName,r.ImporterShortName,r.ImporterCountryCode,r.ImporterRegionCode,r.ImporterDescription,r.SupplierType,r.SupplierClientRegId,r.SupplierInn,r.SupplierKpp,r.SupplierFullName,r.SupplierShortName,r.SupplierCountryCode,r.SupplierRegionCode,r.SupplierDescription,r.Version,r.State,r.Status,r.ReplyId,r.Archive,r.Xml,r.Id)
	return err
}

func (r *ImportReports) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE import_reports WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ImportReports) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from import_reports WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ImportReports) All(db *sqlx.DB) (out ImportReportsSlice, err error) {
	out = make(ImportReportsSlice, 0)
	sql := `select * from import_reports order by id;`
	err = db.Select(&out, sql)
	return out, err
}
