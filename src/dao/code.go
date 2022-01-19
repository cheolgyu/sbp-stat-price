package dao

import (
	"log"

	"github.com/cheolgyu/stock-write-common/db"
	"github.com/cheolgyu/stock-write-model/model"
	mod_code "github.com/cheolgyu/stock-write-module-meta/code"
	mod_config "github.com/cheolgyu/stock-write-module-meta/config"
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
	res, err := mod_code.GetCodeList(db.Conn)
	return res, err
}

func GetConfigListByUpperCode() ([]model.Config, error) {
	res, err := mod_config.GetConfigListByUpperCode(db.Conn, mod_config.CONFIG_UPPER_CODE_PRICE_TYPE)
	return res, err
}
