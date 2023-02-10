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

func TestGenerateMD5Hash(t *testing.T) {
	// word => hash
	tests := map[string]string{
		"password":                          "5f4dcc3b5aa765d61d8327deb882cf99",
		"":                                  "d41d8cd98f00b204e9800998ecf8427e",
		"example-123-9838-2910110":          "6ac00b045dc3b7ecfc14099cb06326d6",
		"What is an MD5 hash?":              "7b96e636e4bd247fc6dfe3371a194766",
		`\11\11`:                            "d013f250fc146dfc90b6401ff86b9727",
		`https://www.md5hashgenerator.com/`: "df30cb178eb8e37728f39b3e6551c8de",
		"👍":                                 "0215ac4dab1ecaf71d83f98af5726984",
		"example":                           "1a79a4d60de6718e8e5b326e338ae533",
	}

	for word, hash := range tests {
		assert.Equal(t, hash, shared.GenerateMD5Hash(word))
	}
}
