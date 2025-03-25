package output

type ProductionReports struct {
    Id	int64	`db:"id"` 
    CreateDate	string	`db:"create_date"` 
    DocIdentity	string	`db:"doc_identity"` 
    DocType	string	`db:"doc_type"` 
    DocNumber	string	`db:"doc_number"` 
    DocDate	string	`db:"doc_date"` 
    DocProducedDate	string	`db:"doc_produced_date"` 
    DocComment	string	`db:"doc_comment"` 
    ProducerType	string	`db:"producer_type"` 
    ProducerClientRegId	string	`db:"producer_client_reg_id"` 
    ProducerInn	string	`db:"producer_inn"` 
    ProducerKpp	string	`db:"producer_kpp"` 
    ProducerFullName	string	`db:"producer_full_name"` 
    ProducerShortName	string	`db:"producer_short_name"` 
    ProducerCountryCode	string	`db:"producer_country_code"` 
    ProducerRegionCode	string	`db:"producer_region_code"` 
    ProducerDescription	string	`db:"producer_description"` 
    Version	string	`db:"version"` 
    State	string	`db:"state"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
    Archive	int	`db:"archive"` 
    Xml	string	`db:"xml"` 
}

type ProductionReportsSlice []*ProductionReports

func (r *ProductionReports) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO production_reports (create_date,doc_identity,doc_type,doc_number,doc_date,doc_produced_date,doc_comment,producer_type,producer_client_reg_id,producer_inn,producer_kpp,producer_full_name,producer_short_name,producer_country_code,producer_region_code,producer_description,version,state,status,reply_id,archive,xml) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.CreateDate,r.DocIdentity,r.DocType,r.DocNumber,r.DocDate,r.DocProducedDate,r.DocComment,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription,r.Version,r.State,r.Status,r.ReplyId,r.Archive,r.Xml)
	return err
}

func (r *ProductionReports) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE production_reports SET create_date=?,doc_identity=?,doc_type=?,doc_number=?,doc_date=?,doc_produced_date=?,doc_comment=?,producer_type=?,producer_client_reg_id=?,producer_inn=?,producer_kpp=?,producer_full_name=?,producer_short_name=?,producer_country_code=?,producer_region_code=?,producer_description=?,version=?,state=?,status=?,reply_id=?,archive=?,xml=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.CreateDate,r.DocIdentity,r.DocType,r.DocNumber,r.DocDate,r.DocProducedDate,r.DocComment,r.ProducerType,r.ProducerClientRegId,r.ProducerInn,r.ProducerKpp,r.ProducerFullName,r.ProducerShortName,r.ProducerCountryCode,r.ProducerRegionCode,r.ProducerDescription,r.Version,r.State,r.Status,r.ReplyId,r.Archive,r.Xml,r.Id)
	return err
}

func (r *ProductionReports) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE production_reports WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ProductionReports) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from production_reports WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ProductionReports) All(db *sqlx.DB) (out ProductionReportsSlice, err error) {
	out = make(ProductionReportsSlice, 0)
	sql := `select * from production_reports order by id;`
	err = db.Select(&out, sql)
	return out, err
}
