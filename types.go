package main

type column struct {
	Cid       int    `db:"cid"`
	Name      string `db:"name"`
	NameCamel string
	Type      string `db:"type"`
	NotNull   int    `db:"notnull"`
	Pk        int    `db:"pk"`
}

type columnSlice []*column
