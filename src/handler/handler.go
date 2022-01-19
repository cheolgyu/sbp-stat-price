package handler

import (
	"github.com/cheolgyu/stock-write-common/logging"
	cmm_model "github.com/cheolgyu/stock-write-model/model"
	"github.com/cheolgyu/stock-write-project-52-weeks/src/dao"
	"github.com/cheolgyu/stock-write-project-52-weeks/src/model"
)

var TimeFrames []model.TimeFrame
var Configs []cmm_model.Config

var CONFIG_OP = -1
var CONFIG_CP = -1
var CONFIG_LP = -1
var CONFIG_HP = -1

func init() {
	setTimeFrames()
	configs, _ := dao.GetConfigListByUpperCode()
	Configs = configs

	for _, v := range configs {
		switch v.Code {
		case "open":
			CONFIG_OP = v.Id
		case "close":
			CONFIG_CP = v.Id
		case "low":
			CONFIG_LP = v.Id
		case "high":
			CONFIG_HP = v.Id
		}
	}
}

func Handler() {

	codes, err := dao.GetCodeAll()

	if err != nil {
		logging.Log.Panic(err)
	}
	for _, v := range codes[:10] {

		code_info, err := dao.SelectList(v.Id)
		if err != nil {
			logging.Log.Panic(err)
		}
		list := find(code_info)
		err = dao.Insert(list)
		if err != nil {
			logging.Log.Panic(err)
		}
	}

}

func find(item model.CodeInfo) []cmm_model.Tb52Weeks {

	var res []cmm_model.Tb52Weeks
	res = append(res, findPointInfo(item.Code.Id, item.OP, CONFIG_OP)...)
	res = append(res, findPointInfo(item.Code.Id, item.CP, CONFIG_CP)...)
	res = append(res, findPointInfo(item.Code.Id, item.LP, CONFIG_LP)...)
	res = append(res, findPointInfo(item.Code.Id, item.HP, CONFIG_HP)...)

	return res
}

func findPointInfo(code_id int, arr []model.PointInfo, price_type int) []cmm_model.Tb52Weeks {
	var price_info model.PriceInfo

	var res []cmm_model.Tb52Weeks

	if len(arr) > 0 {
		price_info.Cur.X = arr[0].Point.X
		price_info.Cur.Y = arr[0].Point.Y
		price_info.Min.X = arr[0].Point.X
		price_info.Min.Y = arr[0].Point.Y
		price_info.Max.X = arr[0].Point.X
		price_info.Max.Y = arr[0].Point.Y
	}
	var break_timeframes int = 0

	for _, v := range arr {

		//stop_cnt := v.Xcnt

		if v.Point.Y >= price_info.Max.Y {
			price_info.Max.Y = v.Point.Y
			price_info.Max.X = v.Point.X
		} else {
			price_info.Min.Y = v.Point.Y
			price_info.Min.X = v.Point.X
		}

		for i, t := range TimeFrames {
			if v.Xcnt > t.Day && break_timeframes <= i {
				break_timeframes = i

				max_item := cmm_model.Tb52Weeks{
					Code_id:    code_id,
					Price_type: price_type,
					Row_type:   true,
					Unit_type:  t.UnitType,
					Unit:       t.UnitVal,
					Np_dt:      price_info.Cur.X,
					Np_val:     price_info.Cur.Y,
					Op_dt:      price_info.Max.X,
					Op_val:     price_info.Max.Y,
					P_percent:  price_info.Max.Y / price_info.Cur.Y * 100,
				}
				res = append(res, max_item)

				min_item := cmm_model.Tb52Weeks{
					Code_id:    code_id,
					Price_type: price_type,
					Row_type:   false,
					Unit_type:  t.UnitType,
					Unit:       t.UnitVal,
					Np_dt:      price_info.Cur.X,
					Np_val:     price_info.Cur.Y,
					Op_dt:      price_info.Min.X,
					Op_val:     price_info.Min.Y,
					P_percent:  price_info.Min.Y / price_info.Cur.Y * 100,
				}
				res = append(res, min_item)
			}
		}

	}

	return res
}

//일자 목록 구하기
func setTimeFrames() {

	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 7, UnitType: 1, UnitVal: 1})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 14, UnitType: 1, UnitVal: 2})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 28, UnitType: 1, UnitVal: 3})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 30 * 1, UnitType: 2, UnitVal: 1})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 30 * 2, UnitType: 2, UnitVal: 2})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 30 * 3, UnitType: 2, UnitVal: 3})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 30 * 4, UnitType: 2, UnitVal: 4})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 30 * 5, UnitType: 2, UnitVal: 5})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 30 * 6, UnitType: 2, UnitVal: 6})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 30 * 7, UnitType: 2, UnitVal: 7})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 30 * 8, UnitType: 2, UnitVal: 8})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 30 * 9, UnitType: 2, UnitVal: 9})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 30 * 10, UnitType: 2, UnitVal: 10})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 30 * 11, UnitType: 2, UnitVal: 11})
	TimeFrames = append(TimeFrames, model.TimeFrame{Day: 30 * 12, UnitType: 2, UnitVal: 12})

}
