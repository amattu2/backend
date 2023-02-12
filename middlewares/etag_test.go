/*
 * Produced: Sun Feb 12 2023
 * Author: Alec M.
 * GitHub: https://amattu.com/links/github
 * Copyright: (C) 2023 Alec M.
 * License: License GNU Affero General Public License v3.0
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"placeholder-app/backend/middlewares"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

// Basically test against a bunch of different inputs and make sure the E Tag is correct
// Also assure it returns a 304 if the E Tag matches

func TestEmptyETag(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest("GET", "/", nil)

	middlewares.ETag()(context)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "W/\"d41d8cd98f00b204e9800998ecf8427e\"", recorder.Header().Get("ETag"))
}

func TestPathETagGeneration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// hash => request params
	tests := map[string][]gin.Param{
		"afc41cc0c128a10393af797ca32362f7": {
			{Key: "size", Value: "315x315"},
		},
		"dc3737b1340496be090dddbb54bf4e5f": {
			{Key: "size", Value: "400x400"},
			{Key: "bgColor", Value: "3B3B3B"},
		},
		"83874226cca6452ed03d11e9a22db707": {
			{Key: "size", Value: "200x200"},
			{Key: "bgColor", Value: "F2F2F2"},
			{Key: "txtColor", Value: "3B3B3B"},
		},
	}

	for hash, params := range tests {
		// Test Without If-None-Match
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)
		context.Request = httptest.NewRequest("GET", "/", nil)
		context.Params = params
		middlewares.ETag()(context)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "W/\""+hash+"\"", recorder.Header().Get("ETag"))

		// Test HTTP If-None-Match
		recorder = httptest.NewRecorder()
		context, _ = gin.CreateTestContext(recorder)
		context.Request = httptest.NewRequest("GET", "/", nil)
		context.Params = params
		context.Request.Header.Set("If-None-Match", hash)
		middlewares.ETag()(context)

		assert.Equal(t, http.StatusNotModified, recorder.Code)
	}
}

func TestQueryETagGeneration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// hash => query string
	tests := map[int]map[string]string{
		0: {
			"hash":  "172da84f59e1914dc7353d44c0dc22e1",
			"query": "?borderWidth=24&borderColor=3B3B3B",
		},
		1: {
			"hash":  "172da84f59e1914dc7353d44c0dc22e1",
			"query": "?borderColor=3B3B3B&borderWidth=24",
		},
		2: {
			"hash":  "b54101fbf8973b452d77a6faf07351fe",
			"query": "?text=Hello%2C%20World!&borderWidth=24&borderColor=3B3B3B",
		},
		3: {
			"hash":  "a51abaa319322ffb23537ababc115b5f",
			"query": "?font=Sans-Serif&format=bmp&text=Hello%2C%20World!&borderWidth=24&borderColor=3B3B3B",
		},
	}

	for _, tst := range tests {
		// Test Without If-None-Match
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)
		context.Request = httptest.NewRequest("GET", "/"+tst["query"], nil)
		middlewares.ETag()(context)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "W/\""+tst["hash"]+"\"", recorder.Header().Get("ETag"))

		// Test HTTP If-None-Match
		recorder = httptest.NewRecorder()
		context, _ = gin.CreateTestContext(recorder)
		context.Request = httptest.NewRequest("GET", "/"+tst["query"], nil)
		context.Request.Header.Set("If-None-Match", tst["hash"])
		middlewares.ETag()(context)

		assert.Equal(t, http.StatusNotModified, recorder.Code)
	}
}
