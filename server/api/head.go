// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"github.com/labstack/echo"
)

// HeadOpts defines options for the Head middleward
type HeadOpts struct {
	CacheControl string
}

// Head defines middleware for setting miscellaneous headers
func Head(opts *HeadOpts) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			// Set default values for opts.PoweredBy
			cacheControl := opts.CacheControl
			if cacheControl == "" {
				cacheControl = "no-cache"
			}

			c.Response().Header().Set("Cache-control", cacheControl)

			return h(c)

		}
	}
}