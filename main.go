package main

import (
	_ "github.com/cheolgyu/stock-write-common/db"
	_ "github.com/cheolgyu/stock-write-common/env"
	"github.com/cheolgyu/stock-write-common/logging"
)

func main() {
	defer logging.ElapsedTime()()
	project_run()
}
func project_run() {

}
