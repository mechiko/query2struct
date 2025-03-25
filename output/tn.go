package output

type Tn struct {
    Id	int64	`db:"id"` 
    IdTtn	int	`db:"id_ttn"` 
    CreateDate	string	`db:"create_date"` 
    DocNumber	string	`db:"doc_number"` 
    DocDate	string	`db:"doc_date"` 
    DocQuantity	string	`db:"doc_quantity"` 
    DocOwnership	int	`db:"doc_ownership"` 
    DocRegId	string	`db:"doc_reg_id"` 
    TranType	string	`db:"tran_type"` 
    TransportCompany	string	`db:"transport_company"` 
    TransportCar	string	`db:"transport_car"` 
    TransportCustomer	string	`db:"transport_customer"` 
    TransportDriver	string	`db:"transport_driver"` 
    TransportTrailer	string	`db:"transport_trailer"` 
    TransportForwarder	string	`db:"transport_forwarder"` 
    TransportLoadPoint	string	`db:"transport_load_point"` 
    TransportUnloadPoint	string	`db:"transport_unload_point"` 
    TransportRedirect	string	`db:"transport_redirect"` 
    Version	string	`db:"version"` 
    State	string	`db:"state"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
    Xml	string	`db:"xml"` 
}

type TnSlice []*Tn

func (r *Tn) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO tn (id_ttn,create_date,doc_number,doc_date,doc_quantity,doc_ownership,doc_reg_id,tran_type,transport_company,transport_car,transport_customer,transport_driver,transport_trailer,transport_forwarder,transport_load_point,transport_unload_point,transport_redirect,version,state,status,reply_id,xml) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdTtn,r.CreateDate,r.DocNumber,r.DocDate,r.DocQuantity,r.DocOwnership,r.DocRegId,r.TranType,r.TransportCompany,r.TransportCar,r.TransportCustomer,r.TransportDriver,r.TransportTrailer,r.TransportForwarder,r.TransportLoadPoint,r.TransportUnloadPoint,r.TransportRedirect,r.Version,r.State,r.Status,r.ReplyId,r.Xml)
	return err
}

func (r *Tn) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE tn SET id_ttn=?,create_date=?,doc_number=?,doc_date=?,doc_quantity=?,doc_ownership=?,doc_reg_id=?,tran_type=?,transport_company=?,transport_car=?,transport_customer=?,transport_driver=?,transport_trailer=?,transport_forwarder=?,transport_load_point=?,transport_unload_point=?,transport_redirect=?,version=?,state=?,status=?,reply_id=?,xml=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtn,r.CreateDate,r.DocNumber,r.DocDate,r.DocQuantity,r.DocOwnership,r.DocRegId,r.TranType,r.TransportCompany,r.TransportCar,r.TransportCustomer,r.TransportDriver,r.TransportTrailer,r.TransportForwarder,r.TransportLoadPoint,r.TransportUnloadPoint,r.TransportRedirect,r.Version,r.State,r.Status,r.ReplyId,r.Xml,r.Id)
	return err
}

func (r *Tn) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE tn WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *Tn) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from tn WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *Tn) All(db *sqlx.DB) (out TnSlice, err error) {
	out = make(TnSlice, 0)
	sql := `select * from tn order by id;`
	err = db.Select(&out, sql)
	return out, err
}
