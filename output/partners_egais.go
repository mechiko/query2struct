package output

type PartnersEgais struct {
    Id	int64	`db:"id"` 
    IdRequests	int	`db:"id_requests"` 
    ClientType	string	`db:"client_type"` 
    ClientRegId	string	`db:"client_reg_id"` 
    ClientInn	string	`db:"client_inn"` 
    ClientKpp	string	`db:"client_kpp"` 
    ClientFullName	string	`db:"client_full_name"` 
    ClientShortName	string	`db:"client_short_name"` 
    ClientCountryCode	string	`db:"client_country_code"` 
    ClientRegionCode	string	`db:"client_region_code"` 
    ClientDescription	string	`db:"client_description"` 
    ClientState	string	`db:"client_state"` 
    ClientWbVersion	string	`db:"client_wb_version"` 
    ClientLicense	string	`db:"client_license"` 
}

type PartnersEgaisSlice []*PartnersEgais

func (r *PartnersEgais) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO partners_egais (id_requests,client_type,client_reg_id,client_inn,client_kpp,client_full_name,client_short_name,client_country_code,client_region_code,client_description,client_state,client_wb_version,client_license) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdRequests,r.ClientType,r.ClientRegId,r.ClientInn,r.ClientKpp,r.ClientFullName,r.ClientShortName,r.ClientCountryCode,r.ClientRegionCode,r.ClientDescription,r.ClientState,r.ClientWbVersion,r.ClientLicense)
	return err
}

func (r *PartnersEgais) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE partners_egais SET id_requests=?,client_type=?,client_reg_id=?,client_inn=?,client_kpp=?,client_full_name=?,client_short_name=?,client_country_code=?,client_region_code=?,client_description=?,client_state=?,client_wb_version=?,client_license=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdRequests,r.ClientType,r.ClientRegId,r.ClientInn,r.ClientKpp,r.ClientFullName,r.ClientShortName,r.ClientCountryCode,r.ClientRegionCode,r.ClientDescription,r.ClientState,r.ClientWbVersion,r.ClientLicense,r.Id)
	return err
}

func (r *PartnersEgais) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE partners_egais WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *PartnersEgais) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from partners_egais WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *PartnersEgais) All(db *sqlx.DB) (out PartnersEgaisSlice, err error) {
	out = make(PartnersEgaisSlice, 0)
	sql := `select * from partners_egais order by id;`
	err = db.Select(&out, sql)
	return out, err
}
