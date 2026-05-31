module example.com/eagle-eye/services/camera-service

go 1.23.0

require (
	example.com/eagle-eye/shared v0.0.0
	github.com/go-kit/kit v0.13.0
)

replace example.com/eagle-eye/shared => ../../shared

replace github.com/go-kit/kit => ../../third_party/go-kit
