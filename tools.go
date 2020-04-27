// +build tools

package users

// используется для добавления и компиляции внешних приложений
import (
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/mwitkow/go-proto-validators/protoc-gen-govalidators"
)
