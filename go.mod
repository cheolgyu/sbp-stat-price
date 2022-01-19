module github.com/cheolgyu/stock-write-project-52-weeks

go 1.16

require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/cheolgyu/stock-write-common v0.0.0
	github.com/cheolgyu/stock-write-model v0.0.0
	github.com/cheolgyu/stock-write-module-meta v0.0.0
)

replace (
	github.com/cheolgyu/stock-write-common v0.0.0 => ../stock-write-common
	github.com/cheolgyu/stock-write-model v0.0.0 => ../stock-write-model
	github.com/cheolgyu/stock-write-module-meta v0.0.0 => ../stock-write-module-meta
)
