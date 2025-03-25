package output

type DocsGrapesSort struct {
    Sort	string	`db:"sort"` 
    Code	string	`db:"code"` 
}

type DocsGrapesSortSlice []*DocsGrapesSort

func (r *DocsGrapesSort) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO docs_grapes_sort (sort,code) 
	VALUES(?,?);`
	_, err = db.Exec(sql, r.Sort,r.Code)
	return err
}

func (r *DocsGrapesSort) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE docs_grapes_sort SET sort=?,code=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.Sort,r.Code,)
	return err
}

func (r *DocsGrapesSort) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE docs_grapes_sort WHERE =?;`
	_, err = db.Exec(sql, )
	return err
}

func (r *DocsGrapesSort) Find(db *sqlx.DB) (err error) {
	if  == 0 {
		return fmt.Errorf(" = 0")
	}
	sql := `select * from docs_grapes_sort WHERE =?;`
	err = db.Get(r, sql, )
	return err
}

func (r *DocsGrapesSort) All(db *sqlx.DB) (out DocsGrapesSortSlice, err error) {
	out = make(DocsGrapesSortSlice, 0)
	sql := `select * from docs_grapes_sort order by ;`
	err = db.Select(&out, sql)
	return out, err
}
