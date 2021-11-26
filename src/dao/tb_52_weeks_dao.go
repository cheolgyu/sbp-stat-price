package dao

import (
	"github.com/cheolgyu/stock-write-common/db"
	"github.com/cheolgyu/stock-write-common/logging"
	cmm_model "github.com/cheolgyu/stock-write-model/model"
	"github.com/cheolgyu/stock-write-project-52-weeks/src/model"
)

const query_insert = `INSERT INTO project.tb_52_weeks( ` +
	` code_id, price_type, unit_type, unit, high_price, low_price) ` +
	` VALUES ($1, $2, $3, $4, $5, $6); `

const query_select_list = `select  hp.code_id, hp.dt,  hp.op, hp.hp, hp.lp, hp.cp  ` +
	` from hist.price hp ` +
	` where hp.code_id = $1 and hp.dt >= $2 order by hp.dt asc; `

func SelecDate() (last int, new int, err error) {

	query := `select to_char( date(dt::TEXT) - interval '1 year', 'YYYYMMDD')::integer, dt::integer as dt from meta.opening order by dt desc limit 1`
	rows, err := db.Conn.Query(query)
	if err != nil {
		logging.Log.Fatalln(err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&last, &new); err != nil {
			logging.Log.Fatal(err)
			panic(err)
		}
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		logging.Log.Fatal(err)
		panic(err)
	}
	return last, new, err

}

func SelectList(code_id int, end_dt int) ([]model.List, error) {

	var res []model.List
	rows, err := db.Conn.Query(query_select_list, code_id, end_dt)
	if err != nil {
		logging.Log.Fatalln(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var item model.List
		// hp.code_id, hp.dt,  hp.op, hp.hp, hp.lp, hp.cp
		if err := rows.Scan(&item.Code.Id,
			&item.PriceMarket.Dt, &item.PriceMarket.OpenPrice, &item.PriceMarket.HighPrice, &item.PriceMarket.LowPrice, &item.PriceMarket.ClosePrice); err != nil {
			logging.Log.Fatal(err)
			panic(err)
		}
		res = append(res, item)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		logging.Log.Fatal(err)
		panic(err)
	}

	return res, err
}

func Insert(item cmm_model.Tb52Weeks) error {

	client := db.Conn
	stmt, err := client.Prepare(query_insert)
	if err != nil {
		logging.Log.Println("쿼리:Prepare 오류: ", item)
		logging.Log.Fatal(err)
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.Code_id,
		item.Price_type, item.Unit_type, item.Unit, item.High_price, item.Low_price,
	)
	if err != nil {
		logging.Log.Println("쿼리:stmt.Exec 오류: ", item)
		logging.Log.Fatal(err)
		panic(err)
	}
	return err
}
