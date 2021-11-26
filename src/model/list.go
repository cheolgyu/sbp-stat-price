package model

import (
	cmm_model "github.com/cheolgyu/stock-write-model/model"
)

type List struct {
	cmm_model.Code
	cmm_model.PriceMarket
}
