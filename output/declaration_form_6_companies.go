package output

type DeclarationForm6Companies struct {
    Id	int64	`db:"id"` 
    IdDeclarationForm6	int	`db:"id_declaration_form_6"` 
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

type DeclarationForm6CompaniesSlice []*DeclarationForm6Companies

func (r *DeclarationForm6Companies) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO declaration_form_6_companies (id_declaration_form_6,type,full_name,inn,kpp,post_index,region_code,district,city,locality,street,house,building,letter,apartment,phone,email) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.IdDeclarationForm6,r.Type,r.FullName,r.Inn,r.Kpp,r.PostIndex,r.RegionCode,r.District,r.City,r.Locality,r.Street,r.House,r.Building,r.Letter,r.Apartment,r.Phone,r.Email)
	return err
}

func (r *DeclarationForm6Companies) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE declaration_form_6_companies SET id_declaration_form_6=?,type=?,full_name=?,inn=?,kpp=?,post_index=?,region_code=?,district=?,city=?,locality=?,street=?,house=?,building=?,letter=?,apartment=?,phone=?,email=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdDeclarationForm6,r.Type,r.FullName,r.Inn,r.Kpp,r.PostIndex,r.RegionCode,r.District,r.City,r.Locality,r.Street,r.House,r.Building,r.Letter,r.Apartment,r.Phone,r.Email,r.Id)
	return err
}

func (r *DeclarationForm6Companies) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE declaration_form_6_companies WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *DeclarationForm6Companies) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from declaration_form_6_companies WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *DeclarationForm6Companies) All(db *sqlx.DB) (out DeclarationForm6CompaniesSlice, err error) {
	out = make(DeclarationForm6CompaniesSlice, 0)
	sql := `select * from declaration_form_6_companies order by id;`
	err = db.Select(&out, sql)
	return out, err
}
