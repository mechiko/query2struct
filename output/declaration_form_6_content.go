package output

type DeclarationForm6Content struct {
    Id	int64	`db:"id"` 
    IdDeclarationForm6Companies	int	`db:"id_declaration_form_6_companies"` 
    ActivityCode	string	`db:"activity_code"` 
    ActivityName	string	`db:"activity_name"` 
    ProductCode	string	`db:"product_code"` 
    ProductCodeName	string	`db:"product_code_name"` 
    PowerYear	string	`db:"power_year"` 
    PowerFirstQuarter	string	`db:"power_first_quarter"` 
    PowerSecondQuarter	string	`db:"power_second_quarter"` 
    PowerThirdQuarter	string	`db:"power_third_quarter"` 
    PowerFourthQuarter	string	`db:"power_fourth_quarter"` 
    Production	string	`db:"production"` 
    Utilization	string	`db:"utilization"` 
}

type DeclarationForm6ContentSlice []*DeclarationForm6Content

func (r *DeclarationForm6Content) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO declaration_form_6_content (id_declaration_form_6_companies,activity_code,activity_name,product_code,product_code_name,power_year,power_first_quarter,power_second_quarter,power_third_quarter,power_fourth_quarter,production,utilization) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdDeclarationForm6Companies,r.ActivityCode,r.ActivityName,r.ProductCode,r.ProductCodeName,r.PowerYear,r.PowerFirstQuarter,r.PowerSecondQuarter,r.PowerThirdQuarter,r.PowerFourthQuarter,r.Production,r.Utilization)
	return err
}

func (r *DeclarationForm6Content) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE declaration_form_6_content SET id_declaration_form_6_companies=?,activity_code=?,activity_name=?,product_code=?,product_code_name=?,power_year=?,power_first_quarter=?,power_second_quarter=?,power_third_quarter=?,power_fourth_quarter=?,production=?,utilization=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdDeclarationForm6Companies,r.ActivityCode,r.ActivityName,r.ProductCode,r.ProductCodeName,r.PowerYear,r.PowerFirstQuarter,r.PowerSecondQuarter,r.PowerThirdQuarter,r.PowerFourthQuarter,r.Production,r.Utilization,r.Id)
	return err
}

func (r *DeclarationForm6Content) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE declaration_form_6_content WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *DeclarationForm6Content) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from declaration_form_6_content WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *DeclarationForm6Content) All(db *sqlx.DB) (out DeclarationForm6ContentSlice, err error) {
	out = make(DeclarationForm6ContentSlice, 0)
	sql := `select * from declaration_form_6_content order by id;`
	err = db.Select(&out, sql)
	return out, err
}
