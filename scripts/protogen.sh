#!/usr/bin/env sh

cd $(dirname $0)/..
proto_path="api/protobuf-spec"
substitutes="Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api"

go mod vendor
protoc \
	-I$proto_path -Ivendor \
	--gogo_out=plugins=grpc,$substitutes:. \
	--govalidators_out=gogoimport=true,$substitutes:. \
	$proto_path/*.proto