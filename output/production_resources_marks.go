package output

type ProductionResourcesMarks struct {
    Id	int64	`db:"id"` 
    IdProductionResources	int	`db:"id_production_resources"` 
    IdProductionResourcesBoxes	int	`db:"id_production_resources_boxes"` 
    Mark	string	`db:"mark"` 
}

type ProductionResourcesMarksSlice []*ProductionResourcesMarks

func (r *ProductionResourcesMarks) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO production_resources_marks (id_production_resources,id_production_resources_boxes,mark) 
	VALUES(?,?,?);`
	_, err = db.Exec(sql, r.IdProductionResources,r.IdProductionResourcesBoxes,r.Mark)
	return err
}

func (r *ProductionResourcesMarks) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE production_resources_marks SET id_production_resources=?,id_production_resources_boxes=?,mark=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdProductionResources,r.IdProductionResourcesBoxes,r.Mark,r.Id)
	return err
}

func (r *ProductionResourcesMarks) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE production_resources_marks WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ProductionResourcesMarks) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from production_resources_marks WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ProductionResourcesMarks) All(db *sqlx.DB) (out ProductionResourcesMarksSlice, err error) {
	out = make(ProductionResourcesMarksSlice, 0)
	sql := `select * from production_resources_marks order by id;`
	err = db.Select(&out, sql)
	return out, err
}
