package helper

import (
	"reflect"
	"time"
)

type Entity struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NullifyProperties(data any, fields []string) any {
	fieldSet := make(map[string]struct{}, len(fields))
	for _, f := range fields {
		fieldSet[f] = struct{}{}
	}

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// work on an addressable copy so we never mutate shared state unexpectedly
	copyVal := reflect.New(v.Type()).Elem()
	copyVal.Set(v)

	nullifyRecursive(copyVal, fieldSet)
	return copyVal.Interface()
}

func nullifyRecursive(v reflect.Value, fields map[string]struct{}) {
	switch v.Kind() {
	case reflect.Ptr:
		if !v.IsNil() {
			nullifyRecursive(v.Elem(), fields)
		}
	case reflect.Struct:
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			if !fv.CanSet() {
				continue // unexported field
			}
			if _, match := fields[t.Field(i).Name]; match {
				fv.Set(reflect.Zero(fv.Type()))
				continue
			}
			switch fv.Kind() {
			case reflect.Ptr, reflect.Struct:
				nullifyRecursive(fv, fields)
			case reflect.Slice, reflect.Array:
				for j := 0; j < fv.Len(); j++ {
					nullifyRecursive(fv.Index(j), fields)
				}
			}
		}
	}
}
