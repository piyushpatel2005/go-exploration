module example.com/module01

go 1.21.1

require (
	github.com/pborman/uuid v1.2.1
	github.com/piyushpatel2005/cryptit v0.0.0-00010101000000-000000000000
)

require github.com/google/uuid v1.0.0 // indirect

replace github.com/piyushpatel2005/cryptit => ../cryptit
