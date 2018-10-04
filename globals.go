package oci8

/*
#include "oci8.go.h"
#cgo !noPkgConfig pkg-config: oci8
*/
import "C"

// noPkgConfig is a Go tag for disabling using pkg-config and using environmental settings like CGO_CFLAGS and CGO_LDFLAGS instead

import (
	"database/sql"
	"database/sql/driver"
	"regexp"
	"time"
	"unsafe"
)

const blobBufSize = 4000
const useOCISessionBegin = true
const sizeOfNilPointer = unsafe.Sizeof(unsafe.Pointer(nil))

// DSN is Oracle Data Source Name
type DSN struct {
	Connect                string
	Username               string
	Password               string
	prefetch_rows          uint32
	prefetch_memory        uint32
	Location               *time.Location
	transactionMode        C.ub4
	enableQMPlaceholders   bool
	operationMode          C.ub4
	externalauthentication bool
}

// OCI8Driver is Oracle driver
type OCI8Driver struct {
}

// OCI8Conn is Oracle connection
type OCI8Conn struct {
	svc                  unsafe.Pointer
	srv                  unsafe.Pointer
	env                  unsafe.Pointer
	err                  *C.OCIError
	usr_session          unsafe.Pointer
	prefetch_rows        uint32
	prefetch_memory      uint32
	location             *time.Location
	transactionMode      C.ub4
	operationMode        C.ub4
	inTransaction        bool
	enableQMPlaceholders bool
	closed               bool
}

// OCI8Tx is Oracle transaction
type OCI8Tx struct {
	conn *OCI8Conn
}

type namedValue struct {
	Name    string
	Ordinal int
	Value   driver.Value
}

type outValue struct {
	Dest interface{}
	In   bool
}

// OCI8Stmt is Oracle statement
type OCI8Stmt struct {
	conn   *OCI8Conn
	s      unsafe.Pointer
	closed bool
	bp     **C.OCIBind
	defp   **C.OCIDefine
	pbind  []oci8bind //bind params
}

// OCI8Result is Oracle result
type OCI8Result struct {
	n     int64
	errn  error
	id    int64
	errid error
	stmt  *OCI8Stmt
}

type oci8col struct {
	name string
	kind C.ub2
	size int
	ind  *C.sb2
	rlen *C.ub2
	pbuf unsafe.Pointer
}

type oci8bind struct {
	kind C.ub2
	pbuf unsafe.Pointer
	clen C.sb4
	out  interface{} // original binded data type
}

// OCI8Rows is Oracle rows
type OCI8Rows struct {
	stmt       *OCI8Stmt
	cols       []oci8col
	e          bool
	indrlenptr unsafe.Pointer
	closed     bool
	done       chan struct{}
	cls        bool
}

var phre = regexp.MustCompile(`\?`)

func init() {
	sql.Register("oci8", &OCI8Driver{})
}
