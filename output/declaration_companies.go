package output

type DeclarationCompanies struct {
    Id	int64	`db:"id"` 
    Type	string	`db:"type"` 
    FullName	string	`db:"full_name"` 
    Inn	string	`db:"inn"` 
    Kpp	string	`db:"kpp"` 
    PostIndex	string	`db:"post_index"` 
    RegionCode	string	`db:"region_code"` 
    District	string	`db:"district"` 
    City	string	`db:"city"` 
    Locality	string	`db:"locality"` 
    Street	string	`db:"street"` 
    House	string	`db:"house"` 
    Building	string	`db:"building"` 
    Letter	string	`db:"letter"` 
    Apartment	string	`db:"apartment"` 
    Phone	string	`db:"phone"` 
    Email	string	`db:"email"` 
}

type DeclarationCompaniesSlice []*DeclarationCompanies

func (r *DeclarationCompanies) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO declaration_companies (type,full_name,inn,kpp,post_index,region_code,district,city,locality,street,house,building,letter,apartment,phone,email) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.Type,r.FullName,r.Inn,r.Kpp,r.PostIndex,r.RegionCode,r.District,r.City,r.Locality,r.Street,r.House,r.Building,r.Letter,r.Apartment,r.Phone,r.Email)
	return err
}

func (r *DeclarationCompanies) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE declaration_companies SET type=?,full_name=?,inn=?,kpp=?,post_index=?,region_code=?,district=?,city=?,locality=?,street=?,house=?,building=?,letter=?,apartment=?,phone=?,email=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.Type,r.FullName,r.Inn,r.Kpp,r.PostIndex,r.RegionCode,r.District,r.City,r.Locality,r.Street,r.House,r.Building,r.Letter,r.Apartment,r.Phone,r.Email,r.Id)
	return err
}

func (r *DeclarationCompanies) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE declaration_companies WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *DeclarationCompanies) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from declaration_companies WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *DeclarationCompanies) All(db *sqlx.DB) (out DeclarationCompaniesSlice, err error) {
	out = make(DeclarationCompaniesSlice, 0)
	sql := `select * from declaration_companies order by id;`
	err = db.Select(&out, sql)
	return out, err
}
