package model

import (
	cmm_model "github.com/cheolgyu/stock-write-model/model"
)

/*
DayCnt : 해당 일자가 365일에서 얼마나 경과된 일자인지
*/
type Res struct {
	cmm_model.Code
	cmm_model.PriceMarket
	DayCnt int
}

// 0:op 1:cp 2:lp 3:hp
func (o *Res) StuctByPrice() [4]PriceResObject {
	var list [4]PriceResObject

	dt := o.PriceMarket.Dt

	list[0] = PriceResObject{
		Date:   dt,
		DayCnt: o.DayCnt,
		Price:  o.PriceMarket.OpenPrice,
	}
	list[1] = PriceResObject{
		Date:   dt,
		DayCnt: o.DayCnt,
		Price:  o.PriceMarket.ClosePrice,
	}
	list[2] = PriceResObject{
		Date:   dt,
		DayCnt: o.DayCnt,
		Price:  o.PriceMarket.LowPrice,
	}
	list[3] = PriceResObject{
		Date:   dt,
		DayCnt: o.DayCnt,
		Price:  o.PriceMarket.HighPrice,
	}

	return list
}

type PriceObject struct {
	Date  int
	Price float32
}

type PriceResObject struct {
	Date   int
	Price  float32
	DayCnt int
}

type PriceArr struct {
	PriceResObjects []PriceResObject
}

// 0:op 1:cp 2:lp 3:hp
type PriceCode struct {
	PriceArr []PriceArr
}

type PriceLH struct {
	PriceType int
	Cur       PriceObject
	Min       PriceObject
	Max       PriceObject
}

type TimeFrame struct {
	Day      int
	UnitType int
	UnitVal  int
}
