package output

type ProductionQueries struct {
    Id	int64	`db:"id"` 
    IdProductionReports	int	`db:"id_production_reports"` 
    QueryDate	string	`db:"query_date"` 
    QueryRegId	string	`db:"query_reg_id"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
}

type ProductionQueriesSlice []*ProductionQueries

func (r *ProductionQueries) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO production_queries (id_production_reports,query_date,query_reg_id,status,reply_id) 
	VALUES(?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdProductionReports,r.QueryDate,r.QueryRegId,r.Status,r.ReplyId)
	return err
}

func (r *ProductionQueries) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE production_queries SET id_production_reports=?,query_date=?,query_reg_id=?,status=?,reply_id=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdProductionReports,r.QueryDate,r.QueryRegId,r.Status,r.ReplyId,r.Id)
	return err
}

func (r *ProductionQueries) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE production_queries WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ProductionQueries) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from production_queries WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ProductionQueries) All(db *sqlx.DB) (out ProductionQueriesSlice, err error) {
	out = make(ProductionQueriesSlice, 0)
	sql := `select * from production_queries order by id;`
	err = db.Select(&out, sql)
	return out, err
}
