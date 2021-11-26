package dao

import (
	"log"

	"github.com/cheolgyu/stock-write-common/db"
	"github.com/cheolgyu/stock-write-model/model"
	"github.com/cheolgyu/stock-write-project-52-weeks/src/c"
)

func Update_info() {
	query := `INSERT INTO public.info( name, updated) VALUES ('`
	query += c.INFO_NAME_UPDATED
	query += `', now()) ON CONFLICT ("name") DO UPDATE SET  updated= now()  `

	_, err := db.Conn.Exec(query)
	if err != nil {
		log.Fatalln(err, query)
		panic(err)
	}
}

func GetCodeAll() ([]model.Code, error) {
	var res []model.Code
	rows, err := db.Conn.Query("select id, code, code_type from meta.code   order by id  ")
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		c := model.Code{}
		if err := rows.Scan(&c.Id, &c.Code, &c.Code_type); err != nil {
			log.Fatal(err)
			panic(err)
		}
		res = append(res, c)
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		panic(err)
	}
	return res, err
}
