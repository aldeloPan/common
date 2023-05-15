// Code Generated By gen-enumer For "Enum Type: GinHttpMethod" - DO NOT EDIT;

/*
 * Copyright 2020-2023 Aldelo, LP
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ginhttpmethod

import (
	"fmt"
	"strconv"
)

// enum names constants
const (
	_GinHttpMethodName_0 = "UNKNOWN"
	_GinHttpMethodName_1 = "GET"
	_GinHttpMethodName_2 = "POST"
	_GinHttpMethodName_3 = "PUT"
	_GinHttpMethodName_4 = "DELETE"
)

// var declares of enum indexes
var (
	_GinHttpMethodIndex_0 = [...]uint8{0, 7}
	_GinHttpMethodIndex_1 = [...]uint8{0, 3}
	_GinHttpMethodIndex_2 = [...]uint8{0, 4}
	_GinHttpMethodIndex_3 = [...]uint8{0, 3}
	_GinHttpMethodIndex_4 = [...]uint8{0, 6}
)

func (i GinHttpMethod) String() string {
	switch {
	case i == UNKNOWN:
		return _GinHttpMethodName_0
	case i == GET:
		return _GinHttpMethodName_1
	case i == POST:
		return _GinHttpMethodName_2
	case i == PUT:
		return _GinHttpMethodName_3
	case i == DELETE:
		return _GinHttpMethodName_4
	default:
		return ""
	}
}

var _GinHttpMethodValues = []GinHttpMethod{
	0, // UNKNOWN
	1, // GET
	2, // POST
	3, // PUT
	4, // DELETE
}

var _GinHttpMethodNameToValueMap = map[string]GinHttpMethod{
	_GinHttpMethodName_0[0:7]: 0, // UNKNOWN
	_GinHttpMethodName_1[0:3]: 1, // GET
	_GinHttpMethodName_2[0:4]: 2, // POST
	_GinHttpMethodName_3[0:3]: 3, // PUT
	_GinHttpMethodName_4[0:6]: 4, // DELETE
}

var _GinHttpMethodValueToKeyMap = map[GinHttpMethod]string{
	0: _GinHttpMethodKey_0, // UNKNOWN
	1: _GinHttpMethodKey_1, // GET
	2: _GinHttpMethodKey_2, // POST
	3: _GinHttpMethodKey_3, // PUT
	4: _GinHttpMethodKey_4, // DELETE
}

var _GinHttpMethodValueToCaptionMap = map[GinHttpMethod]string{
	0: _GinHttpMethodCaption_0, // UNKNOWN
	1: _GinHttpMethodCaption_1, // GET
	2: _GinHttpMethodCaption_2, // POST
	3: _GinHttpMethodCaption_3, // PUT
	4: _GinHttpMethodCaption_4, // DELETE
}

var _GinHttpMethodValueToDescriptionMap = map[GinHttpMethod]string{
	0: _GinHttpMethodDescription_0, // UNKNOWN
	1: _GinHttpMethodDescription_1, // GET
	2: _GinHttpMethodDescription_2, // POST
	3: _GinHttpMethodDescription_3, // PUT
	4: _GinHttpMethodDescription_4, // DELETE
}

// Valid returns 'true' if the value is listed in the GinHttpMethod enum map definition, 'false' otherwise
func (i GinHttpMethod) Valid() bool {
	for _, v := range _GinHttpMethodValues {
		if i == v {
			return true
		}
	}

	return false
}

// ParseByName retrieves a GinHttpMethod enum value from the enum string name,
// throws an error if the param is not part of the enum
func (i GinHttpMethod) ParseByName(s string) (GinHttpMethod, error) {
	if val, ok := _GinHttpMethodNameToValueMap[s]; ok {
		// parse ok
		return val, nil
	}

	// error
	return -1, fmt.Errorf("Enum Name of %s Not Expected In GinHttpMethod Values List", s)
}

// ParseByKey retrieves a GinHttpMethod enum value from the enum string key,
// throws an error if the param is not part of the enum
func (i GinHttpMethod) ParseByKey(s string) (GinHttpMethod, error) {
	for k, v := range _GinHttpMethodValueToKeyMap {
		if v == s {
			// parse ok
			return k, nil
		}
	}

	// error
	return -1, fmt.Errorf("Enum Key of %s Not Expected In GinHttpMethod Keys List", s)
}

// Key retrieves a GinHttpMethod enum string key
func (i GinHttpMethod) Key() string {
	if val, ok := _GinHttpMethodValueToKeyMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// Caption retrieves a GinHttpMethod enum string caption
func (i GinHttpMethod) Caption() string {
	if val, ok := _GinHttpMethodValueToCaptionMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// Description retrieves a GinHttpMethod enum string description
func (i GinHttpMethod) Description() string {
	if val, ok := _GinHttpMethodValueToDescriptionMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// IntValue gets the intrinsic enum integer value
func (i GinHttpMethod) IntValue() int {
	return int(i)
}

// IntString gets the intrinsic enum integer value represented in string format
func (i GinHttpMethod) IntString() string {
	return strconv.Itoa(int(i))
}

// ValueSlice returns all values of the enum GinHttpMethod in a slice
func (i GinHttpMethod) ValueSlice() []GinHttpMethod {
	return _GinHttpMethodValues
}

// NameMap returns all names of the enum GinHttpMethod in a K:name,V:GinHttpMethod map
func (i GinHttpMethod) NameMap() map[string]GinHttpMethod {
	return _GinHttpMethodNameToValueMap
}

// KeyMap returns all keys of the enum GinHttpMethod in a K:GinHttpMethod,V:key map
func (i GinHttpMethod) KeyMap() map[GinHttpMethod]string {
	return _GinHttpMethodValueToKeyMap
}

// CaptionMap returns all captions of the enum GinHttpMethod in a K:GinHttpMethod,V:caption map
func (i GinHttpMethod) CaptionMap() map[GinHttpMethod]string {
	return _GinHttpMethodValueToCaptionMap
}

// DescriptionMap returns all descriptions of the enum GinHttpMethod in a K:GinHttpMethod,V:description map
func (i GinHttpMethod) DescriptionMap() map[GinHttpMethod]string {
	return _GinHttpMethodValueToDescriptionMap
}
