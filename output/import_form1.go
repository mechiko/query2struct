package output

type ImportForm1 struct {
    Id	int64	`db:"id"` 
    IdImportReports	int	`db:"id_import_reports"` 
    DocIdentity	string	`db:"doc_identity"` 
    DocRegId	string	`db:"doc_reg_id"` 
    ClientType	string	`db:"client_type"` 
    ClientRegId	string	`db:"client_reg_id"` 
    ClientInn	string	`db:"client_inn"` 
    ClientKpp	string	`db:"client_kpp"` 
    ClientFullName	string	`db:"client_full_name"` 
    ClientShortName	string	`db:"client_short_name"` 
    ClientCountryCode	string	`db:"client_country_code"` 
    ClientRegionCode	string	`db:"client_region_code"` 
    ClientDescription	string	`db:"client_description"` 
    Xml	string	`db:"xml"` 
}

type ImportForm1Slice []*ImportForm1

func (r *ImportForm1) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO import_form1 (id_import_reports,doc_identity,doc_reg_id,client_type,client_reg_id,client_inn,client_kpp,client_full_name,client_short_name,client_country_code,client_region_code,client_description,xml) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdImportReports,r.DocIdentity,r.DocRegId,r.ClientType,r.ClientRegId,r.ClientInn,r.ClientKpp,r.ClientFullName,r.ClientShortName,r.ClientCountryCode,r.ClientRegionCode,r.ClientDescription,r.Xml)
	return err
}

func (r *ImportForm1) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE import_form1 SET id_import_reports=?,doc_identity=?,doc_reg_id=?,client_type=?,client_reg_id=?,client_inn=?,client_kpp=?,client_full_name=?,client_short_name=?,client_country_code=?,client_region_code=?,client_description=?,xml=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdImportReports,r.DocIdentity,r.DocRegId,r.ClientType,r.ClientRegId,r.ClientInn,r.ClientKpp,r.ClientFullName,r.ClientShortName,r.ClientCountryCode,r.ClientRegionCode,r.ClientDescription,r.Xml,r.Id)
	return err
}

func (r *ImportForm1) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE import_form1 WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ImportForm1) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from import_form1 WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ImportForm1) All(db *sqlx.DB) (out ImportForm1Slice, err error) {
	out = make(ImportForm1Slice, 0)
	sql := `select * from import_form1 order by id;`
	err = db.Select(&out, sql)
	return out, err
}
