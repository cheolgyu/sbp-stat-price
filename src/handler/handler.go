package handler

import (
	"github.com/cheolgyu/stock-write-common/logging"
	"github.com/cheolgyu/stock-write-project-52-weeks/src/dao"
	"github.com/cheolgyu/stock-write-project-52-weeks/src/model"
)

var TimeFrames []model.TimeFrame

func init() {
	setTimeFrames()
}

func Handler() {

	codes, err := dao.GetCodeAll()

	if err != nil {
		logging.Log.Panic(err)
	}

	for _, v := range codes[:10] {

		res, err := dao.SelectList(v.Id)
		if err != nil {
			logging.Log.Panic(err)
		}
		pc := split_by_price_type(res)

	}

}

func split_by_price_type(list []model.Res) model.PriceCode {
	var pcode model.PriceCode

	for _, v := range list {
		pobjs := v.Convert_PriceObject()
		pcode.PriceArr[0].PriceResObjects = append(pcode.PriceArr[0].PriceResObjects, pobjs[0])
		pcode.PriceArr[1].PriceResObjects = append(pcode.PriceArr[1].PriceResObjects, pobjs[1])
		pcode.PriceArr[2].PriceResObjects = append(pcode.PriceArr[2].PriceResObjects, pobjs[2])
		pcode.PriceArr[3].PriceResObjects = append(pcode.PriceArr[3].PriceResObjects, pobjs[3])
	}
	return pcode
}

func search(pcode model.PriceCode) {

	for i, v := range pcode.PriceArr {
		plh := loop_by_priceArr(v)
		plh.PriceType = i
	}

}

func loop_by_priceArr(parr model.PriceArr) model.PriceLH {
	var plh model.PriceLH

	var (
		minDt  int
		minVal float32
		maxDt  int
		maxVal float32
	)

	if len(parr.PriceResObjects) > 0 {
		minDt = parr.PriceResObjects[0].Date
		minVal = parr.PriceResObjects[0].Price
		maxDt = parr.PriceResObjects[0].Date
		maxVal = parr.PriceResObjects[0].Price
	}

	for _, v := range parr.PriceResObjects {

		breakVal := v.DayCnt

		if v.Price > maxVal {
			maxDt = v.Date
			maxVal = v.Price
		} else {
			minDt = v.Date
			minVal = v.Price
		}
	}
	plh.Cur.Date = parr.PriceResObjects[0].Date
	plh.Cur.Price = parr.PriceResObjects[0].Price
	plh.Max.Date = maxDt
	plh.Max.Price = maxVal
	plh.Min.Date = minDt
	plh.Min.Price = minVal

	return plh
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
