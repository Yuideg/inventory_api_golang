package main

import (
	"flag"
	initializer "github.com/Yideg/inventory_app/initiator"
	_ "github.com/lib/pq"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	_ "google.golang.org/protobuf/runtime/protoimpl"
	"os"
	_ "reflect"
	_"sync"
)

var testInit bool

func init() {
	flag.BoolVar(&testInit, "test", false, "initialize test mode without serving")
	flag.Parse()

	os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	initializer.User(testInit)

}
