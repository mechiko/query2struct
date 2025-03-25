package output

type Ttn struct {
    Id	int64	`db:"id"` 
    CreateDate	string	`db:"create_date"` 
    TtnType	string	`db:"ttn_type"` 
    DocIdentity	string	`db:"doc_identity"` 
    DocType	string	`db:"doc_type"` 
    DocNumber	string	`db:"doc_number"` 
    DocDate	string	`db:"doc_date"` 
    DocShippingDate	string	`db:"doc_shipping_date"` 
    DocBase	string	`db:"doc_base"` 
    DocComment	string	`db:"doc_comment"` 
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
    TranType	string	`db:"tran_type"` 
    TransportCompany	string	`db:"transport_company"` 
    TransportCustomer	string	`db:"transport_customer"` 
    TransportOwnership	string	`db:"transport_ownership"` 
    TransportType	string	`db:"transport_type"` 
    TransportDriver	string	`db:"transport_driver"` 
    TransportTrailer	string	`db:"transport_trailer"` 
    TransportRegNumber	string	`db:"transport_reg_number"` 
    TransportForwarder	string	`db:"transport_forwarder"` 
    TransportLoadPoint	string	`db:"transport_load_point"` 
    TransportUnloadPoint	string	`db:"transport_unload_point"` 
    TransportRedirect	string	`db:"transport_redirect"` 
    Version	string	`db:"version"` 
    State	string	`db:"state"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
    Archive	int	`db:"archive"` 
    Xml	string	`db:"xml"` 
}

type TtnSlice []*Ttn

func (r *Ttn) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn (create_date,ttn_type,doc_identity,doc_type,doc_number,doc_date,doc_shipping_date,doc_base,doc_comment,shipper_type,shipper_client_reg_id,shipper_inn,shipper_kpp,shipper_full_name,shipper_short_name,shipper_country_code,shipper_region_code,shipper_description,consignee_type,consignee_client_reg_id,consignee_inn,consignee_kpp,consignee_full_name,consignee_short_name,consignee_country_code,consignee_region_code,consignee_description,tran_type,transport_company,transport_customer,transport_ownership,transport_type,transport_driver,transport_trailer,transport_reg_number,transport_forwarder,transport_load_point,transport_unload_point,transport_redirect,version,state,status,reply_id,archive,xml) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.CreateDate,r.TtnType,r.DocIdentity,r.DocType,r.DocNumber,r.DocDate,r.DocShippingDate,r.DocBase,r.DocComment,r.ShipperType,r.ShipperClientRegId,r.ShipperInn,r.ShipperKpp,r.ShipperFullName,r.ShipperShortName,r.ShipperCountryCode,r.ShipperRegionCode,r.ShipperDescription,r.ConsigneeType,r.ConsigneeClientRegId,r.ConsigneeInn,r.ConsigneeKpp,r.ConsigneeFullName,r.ConsigneeShortName,r.ConsigneeCountryCode,r.ConsigneeRegionCode,r.ConsigneeDescription,r.TranType,r.TransportCompany,r.TransportCustomer,r.TransportOwnership,r.TransportType,r.TransportDriver,r.TransportTrailer,r.TransportRegNumber,r.TransportForwarder,r.TransportLoadPoint,r.TransportUnloadPoint,r.TransportRedirect,r.Version,r.State,r.Status,r.ReplyId,r.Archive,r.Xml)
	return err
}

func (r *Ttn) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn SET create_date=?,ttn_type=?,doc_identity=?,doc_type=?,doc_number=?,doc_date=?,doc_shipping_date=?,doc_base=?,doc_comment=?,shipper_type=?,shipper_client_reg_id=?,shipper_inn=?,shipper_kpp=?,shipper_full_name=?,shipper_short_name=?,shipper_country_code=?,shipper_region_code=?,shipper_description=?,consignee_type=?,consignee_client_reg_id=?,consignee_inn=?,consignee_kpp=?,consignee_full_name=?,consignee_short_name=?,consignee_country_code=?,consignee_region_code=?,consignee_description=?,tran_type=?,transport_company=?,transport_customer=?,transport_ownership=?,transport_type=?,transport_driver=?,transport_trailer=?,transport_reg_number=?,transport_forwarder=?,transport_load_point=?,transport_unload_point=?,transport_redirect=?,version=?,state=?,status=?,reply_id=?,archive=?,xml=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.CreateDate,r.TtnType,r.DocIdentity,r.DocType,r.DocNumber,r.DocDate,r.DocShippingDate,r.DocBase,r.DocComment,r.ShipperType,r.ShipperClientRegId,r.ShipperInn,r.ShipperKpp,r.ShipperFullName,r.ShipperShortName,r.ShipperCountryCode,r.ShipperRegionCode,r.ShipperDescription,r.ConsigneeType,r.ConsigneeClientRegId,r.ConsigneeInn,r.ConsigneeKpp,r.ConsigneeFullName,r.ConsigneeShortName,r.ConsigneeCountryCode,r.ConsigneeRegionCode,r.ConsigneeDescription,r.TranType,r.TransportCompany,r.TransportCustomer,r.TransportOwnership,r.TransportType,r.TransportDriver,r.TransportTrailer,r.TransportRegNumber,r.TransportForwarder,r.TransportLoadPoint,r.TransportUnloadPoint,r.TransportRedirect,r.Version,r.State,r.Status,r.ReplyId,r.Archive,r.Xml,r.Id)
	return err
}

func (r *Ttn) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *Ttn) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *Ttn) All(db *sqlx.DB) (out TtnSlice, err error) {
	out = make(TtnSlice, 0)
	sql := `select * from ttn order by id;`
	err = db.Select(&out, sql)
	return out, err
}
