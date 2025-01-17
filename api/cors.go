// Copyright 2023 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/rs/cors"
	"github.com/tsuru/config"
)

func corsMiddleware() *cors.Cors {
	allowedOrigins, _ := config.GetList("cors:allowed-origins")

	if len(allowedOrigins) == 0 {
		return nil
	}
	return cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
	})
}
