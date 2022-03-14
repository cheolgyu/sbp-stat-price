package main

import (
	_ "github.com/cheolgyu/base/db"
	_ "github.com/cheolgyu/base/env"
	"github.com/cheolgyu/base/logging"
	"github.com/cheolgyu/sbp-stat-price/src/handler"
)

func main() {
	defer logging.ElapsedTime()()
	project_run()
}
func project_run() {
	handler.Handler()
}
