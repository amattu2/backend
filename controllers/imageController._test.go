/*
 * Produced: Thu Jan 26 2023
 * Author: Alec M.
 * GitHub: https://github.com/placeholder-app
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

package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestImageController1(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	GetImage(context)

	if recorder.Code != 400 {
		t.Errorf("Expected 400, got %d", recorder.Code)
	}

	context.Params = gin.Params{{Key: "size", Value: "100x100"}}
	GetImage(context)
	if recorder.Code != 400 {
		t.Errorf("Expected 400, got %d", recorder.Code)
	}
}

func TestImageController2(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	GetImage(context)

	if recorder.Code != 400 {
		t.Errorf("Expected 400, got %d", recorder.Code)
	}

	context.Params = gin.Params{{Key: "size", Value: "fakesize"}}
	GetImage(context)
	if recorder.Code != 400 {
		t.Errorf("Expected 400, got %d", recorder.Code)
	}
}

func TestImageController3(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	GetImage(context)

	if recorder.Code != 400 {
		t.Errorf("Expected 400, got %d", recorder.Code)
	}

	context.Params = gin.Params{{Key: "size", Value: "0x0"}}
	GetImage(context)
	if recorder.Code != 400 {
		t.Errorf("Expected 400, got %d", recorder.Code)
	}
}

func TestImageController4(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	GetImage(context)

	if recorder.Code != 400 {
		t.Errorf("Expected 400, got %d", recorder.Code)
	}

	context.Params = gin.Params{{Key: "size", Value: "-450x500"}}
	GetImage(context)
	if recorder.Code != 400 {
		t.Errorf("Expected 400, got %d", recorder.Code)
	}
}

func TestImageController5(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	GetImage(context)

	if recorder.Code != 400 {
		t.Errorf("Expected 400, got %d", recorder.Code)
	}

	context.Params = gin.Params{{Key: "size", Value: "5555x240"}}
	GetImage(context)
	if recorder.Code != 400 {
		t.Errorf("Expected 400, got %d", recorder.Code)
	}
}
