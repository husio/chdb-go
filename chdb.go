package chdb

import (
	"errors"
	"fmt"
	"sync"
	"unsafe"
)

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L. -lchdb
#include "chdb.h"
*/
import "C"

type Result struct {
	onclose sync.Once
	result  *C.local_result
}

func (r *Result) Bytes() []byte {
	if r.result == nil {
		return nil
	}
	return C.GoBytes(unsafe.Pointer(r.result.buf), C.int(r.result.len))
}

func (r *Result) Close() {
	r.onclose.Do(func() {
		C.free_result(r.result)
		r.result = nil
	})
}

func Query(query string, format string, path string) (*Result, error) {
	// TODO - sanitize input.
	args := []string{
		"clickhouse",
		"--multiquery",
	}
	if query != "" {
		args = append(args, "--query="+query)
	}
	if path != "" {
		args = append(args, "--path="+path)
	}
	if format == "" {
		format = "CSV"
	}
	args = append(args, "--format="+format)

	cargs := C.makeCharArray(C.int(len(args)))
	defer C.freeCharArray(cargs, C.int(len(args)))

	for i, arg := range args {
		cstr := C.CString(arg)
		defer C.free(unsafe.Pointer(cstr))
		C.setArrayString(cargs, cstr, C.int(i))
	}

	result := C.query_stable(C.int(len(args)), cargs)
	if result == nil {
		return nil, ErrOutOfMemory
	}
	return &Result{result: result}, nil
}

var (
	ErrChDB        = errors.New("chdb")
	ErrOutOfMemory = fmt.Errorf("%w: out of memory", ErrChDB)
)
