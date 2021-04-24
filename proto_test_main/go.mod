module github.com/jeremyko/proto_test_main

go 1.16

replace github.com/jeremyko/my_proto => ../my_proto

require (
	github.com/golang/protobuf v1.5.2
	github.com/jeremyko/my_proto v0.0.0-00010101000000-000000000000
)
