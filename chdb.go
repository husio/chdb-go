package chdb

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L. -lchdb
#include "chdb.h"
#include <string.h>
#include <stdio.h>
#include <stdlib.h>

char *Execute(char *query, char *format) {

    char * argv[] = {(char *)"clickhouse", (char *)"--multiquery", (char *)"--output-format=CSV", (char *)"--query="};
    char dataFormat[100];
    char *localQuery;
    // Total 4 = 3 arguments + 1 programm name
    int argc = 4;
    struct local_result *result;

    // Format
    snprintf(dataFormat, sizeof(dataFormat), "--format=%s", format);
    argv[2]=strdup(dataFormat);

    // Query - 10 characters + length of query
    localQuery = (char *) malloc(strlen(query)+10);
    if(localQuery == NULL) {
	// Out of memory.
        return NULL;
    }

    sprintf(localQuery, "--query=%s", query);
    argv[3]=strdup(localQuery);
    free(localQuery);

    // Main query and result
    result = query_stable(argc, argv);

    //Free it
    free(argv[2]);
    free(argv[3]);

    return result->buf;
}

char *Session(char *query, char *format, char* path) {

    char * argv[] = {(char *)"clickhouse", (char *)"--multiquery", (char *)"--output-format=CSV", (char *)"--query=", (char *)"--path=/tmp/"};
    char dataFormat[100];
    char dataPath[100];
    char *localQuery;
    // Total 4 = 4 arguments + 1 programm name
    int argc = 5;
    struct local_result *result;

    // Format
    snprintf(dataFormat, sizeof(dataFormat), "--format=%s", format);
    argv[2]=strdup(dataFormat);

    // Query - 10 characters + length of query
    localQuery = (char *) malloc(strlen(query)+10);
    if(localQuery == NULL) {
	// Out of memory.
        return NULL;
    }

    sprintf(localQuery, "--query=%s", query);
    argv[3]=strdup(localQuery);
    free(localQuery);

    // Path
    snprintf(dataPath, sizeof(dataPath), "--path=%s", path);
    argv[4]=strdup(dataPath);

    // Main query and result
    result = query_stable(argc, argv);

    //Free it
    free(argv[2]);
    free(argv[3]);
    free(argv[4]);

    return result->buf;
}
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

func Query(query string, format string) (string, error) {
	if format == "" {
		format = "CSV"
	}
	cquery := C.CString(query)
	defer C.free(unsafe.Pointer(cquery))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	result := C.Execute(cquery, cformat)
	if result == nil {
		return "", ErrOutOfMemory
	}
	return C.GoString(result), nil
}

func Session(query string, format string, path string) (string, error) {
	if path == "" {
		path = "/tmp/"
	}
	if format == "" {
		format = "CSV"
	}
	cquery := C.CString(query)
	defer C.free(unsafe.Pointer(cquery))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	result := C.Session(cquery, cformat, cpath)
	if result == nil {
		return "", ErrOutOfMemory
	}
	return C.GoString(result), nil
}

var (
	ErrChDB        = errors.New("chdb")
	ErrOutOfMemory = fmt.Errorf("%w: out of memory", ErrChDB)
)
