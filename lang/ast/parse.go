//  Copyright 2015 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package ast

import (
	"fmt"
	"github.com/awalterschulze/picolang/lang/token"
	"github.com/awalterschulze/picolang/lang/util"
	"strconv"
	"strings"
)

func ToStringPtr(v interface{}) *string {
	s := ToString(v)
	return &s
}

func ToString(v interface{}) string {
	t := v.(*token.Token)
	s1 := string(t.Lit)
	s, err := strconv.Unquote(s1)
	if err != nil {
		return s1
	}
	return s
}

func ToFloat64(v interface{}) *float64 {
	t := v.(*token.Token)
	s1 := string(t.Lit)
	f, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		panic(err)
	}
	return &f
}

func parseBytes(s string) ([]byte, error) {
	byteElems := strings.Split(s[strings.Index(s, "{")+1:strings.LastIndex(s, "}")], ",")
	data := make([]byte, 0, len(byteElems))
	for _, b := range byteElems {
		s := strings.TrimSpace(b)
		if len(s) == 0 {
			continue
		}
		d, err := parseByte(s)
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}
	return data, nil
}

func hexToByte(c byte) byte {
	if 'a' <= c && c <= 'f' {
		return c - 'a' + 10
	}
	if 'A' <= c && c <= 'F' {
		return c - 'A' + 10
	}
	return c - '0'
}

func hexesToByte(a byte, b byte) byte {
	aa, bb := hexToByte(a), hexToByte(b)
	return byte(aa*16 + bb)
}

func parseByte(s string) (byte, error) {
	if s[0] == '\'' {
		r := util.RuneValue([]byte(s))
		if r <= 255 {
			return byte(r), nil
		}
		return 0, fmt.Errorf("rune too large %v", r)
	} else if s[0] == '0' {
		if len(s) == 1 {
			return 0, nil
		}
		if s[1] == 'x' || s[1] == 'X' {
			if len(s) == 4 {
				return hexesToByte(s[2], s[3]), nil
			} else if len(s) == 3 {
				return hexToByte(s[2]), nil
			}
			return 0, fmt.Errorf("not a hex digit %v", s)
		} else {
			switch len(s) {
			case 4:
				o := (int(s[1]-'0') * 64) + (int(s[2]-'0') * 8) + int(s[3]-'0')
				if o >= 255 {
					return 0, fmt.Errorf("octal too large %d", o)
				}
				return byte(o), nil
			case 3:
				return byte((s[1]-'0')*8 + (s[2] - '0')), nil
			case 2:
				return byte(s[1] - '0'), nil
			}
			return 0, nil
		}
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	if i >= 0 && i <= 255 {
		return byte(i), nil
	}
	return 0, fmt.Errorf("int too large %d", i)
}
