/*
 * Copyright (c) 2019-2021. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

package grpc

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/micro/go-micro/registry"
)

func extractValue(v reflect.Type, d int) *registry.Value {
	if d == 3 {
		return nil
	}
	if v == nil {
		return nil
	}

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	arg := &registry.Value{
		Name: v.Name(),
		Type: v.Name(),
	}

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			val := extractValue(f.Type, d+1)
			if val == nil {
				continue
			}

			// if we can find a json tag use it
			if tags := f.Tag.Get("json"); len(tags) > 0 {
				parts := strings.Split(tags, ",")
				val.Name = parts[0]
			}

			// if there's no name default it
			if len(val.Name) == 0 {
				val.Name = v.Field(i).Name
			}

			arg.Values = append(arg.Values, val)
		}
	case reflect.Slice:
		p := v.Elem()
		if p.Kind() == reflect.Ptr {
			p = p.Elem()
		}
		arg.Type = "[]" + p.Name()
		val := extractValue(v.Elem(), d+1)
		if val != nil {
			arg.Values = append(arg.Values, val)
		}
	}

	return arg
}

func extractEndpoint(method reflect.Method) *registry.Endpoint {
	if method.PkgPath != "" {
		return nil
	}

	var rspType, reqType reflect.Type
	var stream bool
	mt := method.Type

	switch mt.NumIn() {
	case 3:
		reqType = mt.In(1)
		rspType = mt.In(2)
	case 4:
		reqType = mt.In(2)
		rspType = mt.In(3)
	default:
		return nil
	}

	// are we dealing with a stream?
	switch rspType.Kind() {
	case reflect.Func, reflect.Interface:
		stream = true
	}

	request := extractValue(reqType, 0)
	response := extractValue(rspType, 0)

	return &registry.Endpoint{
		Name:     method.Name,
		Request:  request,
		Response: response,
		Metadata: map[string]string{
			"stream": fmt.Sprintf("%v", stream),
		},
	}
}

func extractSubValue(typ reflect.Type) *registry.Value {
	var reqType reflect.Type
	switch typ.NumIn() {
	case 1:
		reqType = typ.In(0)
	case 2:
		reqType = typ.In(1)
	case 3:
		reqType = typ.In(2)
	default:
		return nil
	}
	return extractValue(reqType, 0)
}
