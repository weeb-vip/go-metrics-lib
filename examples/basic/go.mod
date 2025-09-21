module example.com/basic

go 1.23.0

toolchain go1.23.11

replace github.com/weeb-vip/go-metrics-lib => ../../

require github.com/weeb-vip/go-metrics-lib v0.0.0-00010101000000-000000000000

require (
	github.com/DataDog/datadog-go/v5 v5.3.0 // indirect
	github.com/Microsoft/go-winio v0.5.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
)
