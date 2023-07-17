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

package model

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/gobwas/glob"
)

var (
	defaultIgnores = []glob.Glob{
		glob.MustCompile("**/$buckets.json", GlobSeparator),
		glob.MustCompile("**/$multiparts-session.json", GlobSeparator),
		glob.MustCompile("**/.DS_Store", GlobSeparator),
		glob.MustCompile("**/.minio.sys", GlobSeparator),
		glob.MustCompile("**/.minio.sys/**", GlobSeparator),
	}
)

func IgnoreMatcher(ignores ...glob.Glob) func(string) bool {
	mm := append(defaultIgnores, ignores...)
	return func(s string) bool {
		s = InternalPathSeparator + strings.TrimLeft(s, InternalPathSeparator)
		for _, i := range mm {
			if i.Match(s) {
				return true
			}
		}
		return false
	}
}

func IsIgnoredFile(path string, ignores ...glob.Glob) (ignored bool) {
	// For comparing, we make sure it has a left slash
	path = InternalPathSeparator + strings.TrimLeft(path, InternalPathSeparator)
	for _, i := range append(defaultIgnores, ignores...) {
		if i.Match(path) {
			return true
		}
	}
	return false
}

func NodeRequiresChecksum(node Node) bool {
	return node.IsLeaf() && (node.GetEtag() == "" || node.GetEtag() == DefaultEtag || strings.Contains(node.GetEtag(), "-"))
}

func StringContentToETag(uuid string) string {
	h := md5.New()
	io.Copy(h, strings.NewReader(uuid))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func ZapEndpoint(key string, e Endpoint) zapcore.Field {
	return zap.String(key, e.GetEndpointInfo().URI)
}
