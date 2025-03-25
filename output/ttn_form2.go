package output

type TtnForm2 struct {
    Id	int64	`db:"id"` 
    IdTtn	int	`db:"id_ttn"` 
    DocIdentity	string	`db:"doc_identity"` 
    DocRegId	string	`db:"doc_reg_id"` 
    DocFixNumber	string	`db:"doc_fix_number"` 
    DocFixDate	string	`db:"doc_fix_date"` 
    DocNumber	string	`db:"doc_number"` 
    DocDate	string	`db:"doc_date"` 
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
    Xml	string	`db:"xml"` 
}

type TtnForm2Slice []*TtnForm2

func (r *TtnForm2) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_form2 (id_ttn,doc_identity,doc_reg_id,doc_fix_number,doc_fix_date,doc_number,doc_date,shipper_type,shipper_client_reg_id,shipper_inn,shipper_kpp,shipper_full_name,shipper_short_name,shipper_country_code,shipper_region_code,shipper_description,consignee_type,consignee_client_reg_id,consignee_inn,consignee_kpp,consignee_full_name,consignee_short_name,consignee_country_code,consignee_region_code,consignee_description,xml) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdTtn,r.DocIdentity,r.DocRegId,r.DocFixNumber,r.DocFixDate,r.DocNumber,r.DocDate,r.ShipperType,r.ShipperClientRegId,r.ShipperInn,r.ShipperKpp,r.ShipperFullName,r.ShipperShortName,r.ShipperCountryCode,r.ShipperRegionCode,r.ShipperDescription,r.ConsigneeType,r.ConsigneeClientRegId,r.ConsigneeInn,r.ConsigneeKpp,r.ConsigneeFullName,r.ConsigneeShortName,r.ConsigneeCountryCode,r.ConsigneeRegionCode,r.ConsigneeDescription,r.Xml)
	return err
}

func (r *TtnForm2) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_form2 SET id_ttn=?,doc_identity=?,doc_reg_id=?,doc_fix_number=?,doc_fix_date=?,doc_number=?,doc_date=?,shipper_type=?,shipper_client_reg_id=?,shipper_inn=?,shipper_kpp=?,shipper_full_name=?,shipper_short_name=?,shipper_country_code=?,shipper_region_code=?,shipper_description=?,consignee_type=?,consignee_client_reg_id=?,consignee_inn=?,consignee_kpp=?,consignee_full_name=?,consignee_short_name=?,consignee_country_code=?,consignee_region_code=?,consignee_description=?,xml=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtn,r.DocIdentity,r.DocRegId,r.DocFixNumber,r.DocFixDate,r.DocNumber,r.DocDate,r.ShipperType,r.ShipperClientRegId,r.ShipperInn,r.ShipperKpp,r.ShipperFullName,r.ShipperShortName,r.ShipperCountryCode,r.ShipperRegionCode,r.ShipperDescription,r.ConsigneeType,r.ConsigneeClientRegId,r.ConsigneeInn,r.ConsigneeKpp,r.ConsigneeFullName,r.ConsigneeShortName,r.ConsigneeCountryCode,r.ConsigneeRegionCode,r.ConsigneeDescription,r.Xml,r.Id)
	return err
}

func (r *TtnForm2) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_form2 WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnForm2) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_form2 WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnForm2) All(db *sqlx.DB) (out TtnForm2Slice, err error) {
	out = make(TtnForm2Slice, 0)
	sql := `select * from ttn_form2 order by id;`
	err = db.Select(&out, sql)
	return out, err
}
