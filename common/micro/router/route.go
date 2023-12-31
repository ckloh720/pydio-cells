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

package router

import (
	"path"
	"regexp"
	"strings"
)

var (
	proxyRe   = regexp.MustCompile("^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*$")
	versionRe = regexp.MustCompilePOSIX("^v[0-9]+$")
)

// Translates /foo/bar/zool into api service go.micro.api.foo method Bar.Zool
// Translates /foo/bar into api service go.micro.api.foo method Foo.Bar
func apiRoute(ns, p string) (string, string) {
	p = path.Clean(p)
	p = strings.TrimPrefix(p, "/")
	parts := strings.Split(p, "/")

	// If we've got two or less parts
	// Use first part as service
	// Use all parts as method
	if len(parts) <= 2 {
		service := ns + "." + strings.Join(parts[:len(parts)-1], ".")
		method := methodName(parts)
		return service, method
	}

	// Treat /v[0-9]+ as versioning where we have 3 parts
	// /v1/foo/bar => service: v1.foo method: Foo.bar
	if len(parts) == 3 && versionRe.Match([]byte(parts[0])) {
		service := ns + "." + strings.Join(parts[:len(parts)-1], ".")
		method := methodName(parts[len(parts)-2:])
		return service, method
	}

	// Service is everything minus last two parts
	// Method is the last two parts
	service := ns + "." + strings.Join(parts[:len(parts)-2], ".")
	method := methodName(parts[len(parts)-2:])
	return service, method
}

func proxyRoute(ns string, p string) string {
	parts := strings.Split(p, "/")
	if len(parts) < 2 {
		return ""
	}

	var service string
	var alias string

	// /[service]/methods
	if len(parts) > 2 {
		// /v1/[service]
		if versionRe.MatchString(parts[1]) {
			service = parts[1] + "." + parts[2]
			alias = parts[2]
		} else {
			service = parts[1]
			alias = parts[1]
		}
		// /[service]
	} else {
		service = parts[1]
		alias = parts[1]
	}

	// check service name is valid
	if !proxyRe.MatchString(alias) {
		return ""
	}

	return ns + "." + service
}

func methodName(parts []string) string {
	for i, part := range parts {
		parts[i] = toCamel(part)
	}

	return strings.Join(parts, ".")
}

func toCamel(s string) string {
	words := strings.Split(s, "-")
	var out string
	for _, word := range words {
		out += strings.Title(word)
	}
	return out
}
