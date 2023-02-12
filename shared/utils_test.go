/*
 * Produced: Thu Jan 26 2023
 * Author: Alec M., James A.
 * GitHub: https://github.com/placeholder-app
 * Copyright: (C) 2023 Alec M., James A.
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

package shared_test

import (
	"placeholder-app/backend/shared"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRoundTo1(t *testing.T) {
	r := shared.RoundTo(100, 5)
	assert.Equal(t, r, 100)
}

func TestRoundTo2(t *testing.T) {
	r := shared.RoundTo(106, 10)
	assert.Equal(t, r, 100)
}

func TestRoundTo3(t *testing.T) {
	r := shared.RoundTo(94, 10)
	assert.Equal(t, r, 90)
}

func TestRoundTo4(t *testing.T) {
	r := shared.RoundTo(0, 10)
	assert.Equal(t, r, 0)
}

func TestSplitSize1(t *testing.T) {
	w, h := shared.SplitSize("100x100")
	if (w != 100) || (h != 100) {
		t.Error("Did not split 100x100 into 100x100")
	}
}

func TestSplitSize2(t *testing.T) {
	w, h := shared.SplitSize("100x100x5")
	if (w != 100) || (h != 100) {
		t.Error("Did not split 100x100x5 into 100x100")
	}
}

func TestSplitSize3(t *testing.T) {
	w, h := shared.SplitSize("302x949")
	if (w != 300) || (h != 950) {
		t.Error("Did not split 302x949 into 300x950")
	}
}

func TestSplitSize4(t *testing.T) {
	w, h := shared.SplitSize("500x")
	if (w != 500) || (h != 0) {
		t.Error("Did not split 500x nothing into 500x0")
	}
}

func TestCoerceInt(t *testing.T) {
	i := shared.CoerceInt("100")
	assert.Equal(t, i, 100)
}

func TestCoerceInt2(t *testing.T) {
	i := shared.CoerceInt("abc")
	assert.Equal(t, i, 0)
}

func TestGetEnvExisting(t *testing.T) {
	t.Setenv("TEST", "100")

	assert.Equal(t, "100", shared.GetEnv("TEST", "0"))
}

func TestGetEnvDefault(t *testing.T) {
	assert.Equal(t, "DEFAULT", shared.GetEnv("TEST", "DEFAULT"))
}

func TestGenerateHash(t *testing.T) {
	// word => hash
	tests := map[string]string{
		"password":                          "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8",
		"":                                  "da39a3ee5e6b4b0d3255bfef95601890afd80709",
		"example-123-9838-2910110":          "e1f634558c277e7abbfbce688013779cbdea7fdb",
		"What is an MD5 hash?":              "7e5faa8267e0fdcecb267a53efeb6437240b4ac0",
		`\11\11`:                            "f49fc131293990e9c17ea29e9d32ec7381b33663",
		`https://www.md5hashgenerator.com/`: "169f6bd54cd4174b025371c6636f0cb78728a0e7",
		"👍":                                 "78654ffdf2db3ef8dd605074250103f770177eb6",
		"example":                           "c3499c2729730a7f807efb8676a92dcb6f8a3f8f",
		"   ":                               "088fb1a4ab057f4fcf7d487006499060c7fe5773",
	}

	for word, hash := range tests {
		assert.Equal(t, hash, shared.GenerateHash(word))
	}
}
