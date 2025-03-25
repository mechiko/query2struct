package output

type ImportQueries struct {
    Id	int64	`db:"id"` 
    IdImportReports	int	`db:"id_import_reports"` 
    QueryDate	string	`db:"query_date"` 
    QueryRegId	string	`db:"query_reg_id"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
}

type ImportQueriesSlice []*ImportQueries

func (r *ImportQueries) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO import_queries (id_import_reports,query_date,query_reg_id,status,reply_id) 
	VALUES(?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdImportReports,r.QueryDate,r.QueryRegId,r.Status,r.ReplyId)
	return err
}

func (r *ImportQueries) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE import_queries SET id_import_reports=?,query_date=?,query_reg_id=?,status=?,reply_id=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdImportReports,r.QueryDate,r.QueryRegId,r.Status,r.ReplyId,r.Id)
	return err
}

func (r *ImportQueries) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE import_queries WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ImportQueries) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from import_queries WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ImportQueries) All(db *sqlx.DB) (out ImportQueriesSlice, err error) {
	out = make(ImportQueriesSlice, 0)
	sql := `select * from import_queries order by id;`
	err = db.Select(&out, sql)
	return out, err
}
