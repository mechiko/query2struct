package output

type DeclarationPowers struct {
    ProductCode	string	`db:"product_code"` 
    ProductName	string	`db:"product_name"` 
    PowerYear	string	`db:"power_year"` 
    PowerFirstQuarter	string	`db:"power_first_quarter"` 
    PowerSecondQuarter	string	`db:"power_second_quarter"` 
    PowerThirdQuarter	string	`db:"power_third_quarter"` 
    PowerFourthQuarter	string	`db:"power_fourth_quarter"` 
}

type DeclarationPowersSlice []*DeclarationPowers

func (r *DeclarationPowers) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO declaration_powers (product_code,product_name,power_year,power_first_quarter,power_second_quarter,power_third_quarter,power_fourth_quarter) 
	VALUES(?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.ProductCode,r.ProductName,r.PowerYear,r.PowerFirstQuarter,r.PowerSecondQuarter,r.PowerThirdQuarter,r.PowerFourthQuarter)
	return err
}

func (r *DeclarationPowers) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE declaration_powers SET product_code=?,product_name=?,power_year=?,power_first_quarter=?,power_second_quarter=?,power_third_quarter=?,power_fourth_quarter=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.ProductCode,r.ProductName,r.PowerYear,r.PowerFirstQuarter,r.PowerSecondQuarter,r.PowerThirdQuarter,r.PowerFourthQuarter,)
	return err
}

func (r *DeclarationPowers) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE declaration_powers WHERE =?;`
	_, err = db.Exec(sql, )
	return err
}

func (r *DeclarationPowers) Find(db *sqlx.DB) (err error) {
	if  == 0 {
		return fmt.Errorf(" = 0")
	}
	sql := `select * from declaration_powers WHERE =?;`
	err = db.Get(r, sql, )
	return err
}

func (r *DeclarationPowers) All(db *sqlx.DB) (out DeclarationPowersSlice, err error) {
	out = make(DeclarationPowersSlice, 0)
	sql := `select * from declaration_powers order by ;`
	err = db.Select(&out, sql)
	return out, err
}
