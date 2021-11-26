package handler

import (
	"strconv"
	"time"

	"github.com/cheolgyu/stock-write-common/logging"
	"github.com/cheolgyu/stock-write-project-52-weeks/src/dao"
	"github.com/cheolgyu/stock-write-project-52-weeks/src/model"
)

var TimeList []int

func Handler() {
	codes, err := dao.GetCodeAll()

	if err != nil {
		logging.Log.Panic(err)
	}

	last_dt, new_dt, err := dao.SelecDate()
	if err != nil {
		logging.Log.Panic(err)
	}
	setTimeList(new_dt)

	logging.Log.Info(last_dt, new_dt)

	for _, v := range codes[:10] {

		list, err := dao.SelectList(v.Id, last_dt)
		if err != nil {
			logging.Log.Panic(err)
		}
		find(list, new_dt)

	}

}

func find(list []model.List, new_dt int) {

	var ptype = [4]int{1, 2, 3, 4}

	for _, v := range ptype {
		loop_by_price_type(list, v)
	}

}

// ptype [1:저가,2:고가,3:시가,4:종가]
func loop_by_price_type(list []model.List, ptype int) {

	//old_price
	var omax float32
	var omax_dt int
	var omin float32
	var omin_dt int

	for _, v := range list {

		//cur_price
		var p float32
		switch ptype {
		case 1:
			p = v.PriceMarket.LowPrice
		case 2:
			p = v.PriceMarket.HighPrice
		case 3:
			p = v.PriceMarket.OpenPrice
		case 4:
			p = v.PriceMarket.ClosePrice
		}

		if omax < p {
			omax = p
			omax_dt = v.PriceMarket.Dt
		}

		if omin > p {
			omin = p
			omin_dt = v.PriceMarket.Dt
		}
	}

}

//일자 목록 구하기
func setTimeList(new_dt int) {
	dt := strconv.Itoa(new_dt)
	logging.Log.Debug(dt)
	date, err := time.Parse("20060102", dt)
	if err != nil {
		logging.Log.Error(err)
		logging.Log.Panic(err)
	}

	var list []int
	var days = [3]int{7, 7 * 2, 7 * 3}
	var months []int

	for i := 1; i < 13; i++ {
		months = append(months, i)
	}

	for _, v := range days {
		sv, _ := strconv.Atoi(date.AddDate(0, 0, v).Format("20060102"))
		list = append(list, sv)
	}

	for _, v := range months {
		sv, _ := strconv.Atoi(date.AddDate(0, v, 0).Format("20060102"))
		list = append(list, sv)
	}
	TimeList = list
}
