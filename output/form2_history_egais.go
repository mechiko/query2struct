package output

type Form2HistoryEgais struct {
    Id	int64	`db:"id"` 
    IdRequests	int	`db:"id_requests"` 
    DocType	string	`db:"doc_type"` 
    DocId	string	`db:"doc_id"` 
    ProductQuantity	string	`db:"product_quantity"` 
    OperationName	string	`db:"operation_name"` 
    OperationDate	string	`db:"operation_date"` 
    Form2HistoryDate	string	`db:"form2_history_date"` 
}

type Form2HistoryEgaisSlice []*Form2HistoryEgais

func (r *Form2HistoryEgais) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO form2_history_egais (id_requests,doc_type,doc_id,product_quantity,operation_name,operation_date,form2_history_date) 
	VALUES(?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdRequests,r.DocType,r.DocId,r.ProductQuantity,r.OperationName,r.OperationDate,r.Form2HistoryDate)
	return err
}

func (r *Form2HistoryEgais) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE form2_history_egais SET id_requests=?,doc_type=?,doc_id=?,product_quantity=?,operation_name=?,operation_date=?,form2_history_date=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdRequests,r.DocType,r.DocId,r.ProductQuantity,r.OperationName,r.OperationDate,r.Form2HistoryDate,r.Id)
	return err
}

func (r *Form2HistoryEgais) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE form2_history_egais WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *Form2HistoryEgais) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from form2_history_egais WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *Form2HistoryEgais) All(db *sqlx.DB) (out Form2HistoryEgaisSlice, err error) {
	out = make(Form2HistoryEgaisSlice, 0)
	sql := `select * from form2_history_egais order by id;`
	err = db.Select(&out, sql)
	return out, err
}
