package util

import (
	"reflect"
	"strings"
)

//	func GetGormFields(stc any) []string {
//		typ := reflect.TypeOf(stc)
//		if typ.Kind() == reflect.Ptr {
//			typ = typ.Elem()
//		}
//		if typ.Kind() == reflect.Struct {
//			columns := make([]string, 0, typ.NumField())
//			for i := 0; i < typ.NumField(); i++ {
//				fieldType := typ.Field(i)
//				if fieldType.IsExported() {
//					if fieldType.Tag.Get("gorm") == "-" {
//						continue
//					}
//					name := Came12Snake(fieldType.Name)
//					if len(fieldType.Tag.Get("gorm")) > 0 {
//						content := fieldType.Tag.Get("gorm")
//						if strings.HasPrefix(content, "column:") {
//							content = content[7:]
//							pos := strings.Index(content, ";")
//							if pos > 0 {
//								name = content[0:pos]
//							} else if pos < 0 {
//								name = content
//							}
//						}
//					}
//					columns = append(columns, name)
//				}
//			}
//			return columns
//		}
//	}
//
// GetGormFields returns the GORM fields of a struct
func GetGormFields(stc any) []string {
	typ := reflect.TypeOf(stc)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Struct {
		columns := make([]string, 0, typ.NumField())
		for i := 0; i < typ.NumField(); i++ {
			fieldType := typ.Field(i)
			if fieldType.IsExported() {
				if fieldType.Tag.Get("gorm") == "-" {
					continue
				}
				name := Camel2Snake(fieldType.Name)
				if len(fieldType.Tag.Get("gorm")) > 0 {
					content := fieldType.Tag.Get("gorm")
					if strings.HasPrefix(content, "column:") {
						content = content[7:]
						pos := strings.Index(content, ";")
						if pos > 0 {
							name = content[0:pos]
						} else if pos < 0 {
							name = content
						}
					}
				}
				columns = append(columns, name)
			}
		}
		return columns
	}
	return nil
}

// Camel2Snake converts CamelCase to snake_case
func Camel2Snake(str string) string {
	var result []byte
	for i, v := range str {
		if i > 0 && v >= 'A' && v <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, byte(v))
	}
	return strings.ToLower(string(result))
}
