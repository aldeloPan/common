// Code Generated By gen-enumer For "Enum Type: RedisSetCondition" - DO NOT EDIT;

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

package redissetcondition

import (
	"fmt"
	"strconv"
)

const (
	_RedisSetConditionKey_0 = "UNKNOWN"
	_RedisSetConditionKey_1 = "Normal"
	_RedisSetConditionKey_2 = "SetIfExists"
	_RedisSetConditionKey_3 = "SetIfNotExists"
)

const (
	_RedisSetConditionCaption_0 = "UNKNOWN"
	_RedisSetConditionCaption_1 = "Normal"
	_RedisSetConditionCaption_2 = "SetIfExists"
	_RedisSetConditionCaption_3 = "SetIfNotExists"
)

const (
	_RedisSetConditionDescription_0 = "UNKNOWN"
	_RedisSetConditionDescription_1 = "Normal"
	_RedisSetConditionDescription_2 = "SetIfExists"
	_RedisSetConditionDescription_3 = "SetIfNotExists"
)

// enum names constants
const (
	_RedisSetConditionName_0 = "UNKNOWN"
	_RedisSetConditionName_1 = "Normal"
	_RedisSetConditionName_2 = "SetIfExists"
	_RedisSetConditionName_3 = "SetIfNotExists"
)

// var declares of enum indexes
var (
	_RedisSetConditionIndex_0 = [...]uint8{0, 7}
	_RedisSetConditionIndex_1 = [...]uint8{0, 6}
	_RedisSetConditionIndex_2 = [...]uint8{0, 11}
	_RedisSetConditionIndex_3 = [...]uint8{0, 14}
)

func (i RedisSetCondition) String() string {
	switch {
	case i == UNKNOWN:
		return _RedisSetConditionName_0
	case i == Normal:
		return _RedisSetConditionName_1
	case i == SetIfExists:
		return _RedisSetConditionName_2
	case i == SetIfNotExists:
		return _RedisSetConditionName_3
	default:
		return ""
	}
}

var _RedisSetConditionValues = []RedisSetCondition{
	0, // UNKNOWN
	1, // Normal
	2, // SetIfExists
	3, // SetIfNotExists
}

var _RedisSetConditionNameToValueMap = map[string]RedisSetCondition{
	_RedisSetConditionName_0[0:7]:  0, // UNKNOWN
	_RedisSetConditionName_1[0:6]:  1, // Normal
	_RedisSetConditionName_2[0:11]: 2, // SetIfExists
	_RedisSetConditionName_3[0:14]: 3, // SetIfNotExists
}

var _RedisSetConditionValueToKeyMap = map[RedisSetCondition]string{
	0: _RedisSetConditionKey_0, // UNKNOWN
	1: _RedisSetConditionKey_1, // Normal
	2: _RedisSetConditionKey_2, // SetIfExists
	3: _RedisSetConditionKey_3, // SetIfNotExists
}

var _RedisSetConditionValueToCaptionMap = map[RedisSetCondition]string{
	0: _RedisSetConditionCaption_0, // UNKNOWN
	1: _RedisSetConditionCaption_1, // Normal
	2: _RedisSetConditionCaption_2, // SetIfExists
	3: _RedisSetConditionCaption_3, // SetIfNotExists
}

var _RedisSetConditionValueToDescriptionMap = map[RedisSetCondition]string{
	0: _RedisSetConditionDescription_0, // UNKNOWN
	1: _RedisSetConditionDescription_1, // Normal
	2: _RedisSetConditionDescription_2, // SetIfExists
	3: _RedisSetConditionDescription_3, // SetIfNotExists
}

// Valid returns 'true' if the value is listed in the RedisSetCondition enum map definition, 'false' otherwise
func (i RedisSetCondition) Valid() bool {
	for _, v := range _RedisSetConditionValues {
		if i == v {
			return true
		}
	}

	return false
}

// ParseByName retrieves a RedisSetCondition enum value from the enum string name,
// throws an error if the param is not part of the enum
func (i RedisSetCondition) ParseByName(s string) (RedisSetCondition, error) {
	if val, ok := _RedisSetConditionNameToValueMap[s]; ok {
		// parse ok
		return val, nil
	}

	// error
	return -1, fmt.Errorf("Enum Name of %s Not Expected In RedisSetCondition Values List", s)
}

// ParseByKey retrieves a RedisSetCondition enum value from the enum string key,
// throws an error if the param is not part of the enum
func (i RedisSetCondition) ParseByKey(s string) (RedisSetCondition, error) {
	for k, v := range _RedisSetConditionValueToKeyMap {
		if v == s {
			// parse ok
			return k, nil
		}
	}

	// error
	return -1, fmt.Errorf("Enum Key of %s Not Expected In RedisSetCondition Keys List", s)
}

// Key retrieves a RedisSetCondition enum string key
func (i RedisSetCondition) Key() string {
	if val, ok := _RedisSetConditionValueToKeyMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// Caption retrieves a RedisSetCondition enum string caption
func (i RedisSetCondition) Caption() string {
	if val, ok := _RedisSetConditionValueToCaptionMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// Description retrieves a RedisSetCondition enum string description
func (i RedisSetCondition) Description() string {
	if val, ok := _RedisSetConditionValueToDescriptionMap[i]; ok {
		// found
		return val
	} else {
		// not found
		return ""
	}
}

// IntValue gets the intrinsic enum integer value
func (i RedisSetCondition) IntValue() int {
	return int(i)
}

// IntString gets the intrinsic enum integer value represented in string format
func (i RedisSetCondition) IntString() string {
	return strconv.Itoa(int(i))
}

// ValueSlice returns all values of the enum RedisSetCondition in a slice
func (i RedisSetCondition) ValueSlice() []RedisSetCondition {
	return _RedisSetConditionValues
}

// NameMap returns all names of the enum RedisSetCondition in a K:name,V:RedisSetCondition map
func (i RedisSetCondition) NameMap() map[string]RedisSetCondition {
	return _RedisSetConditionNameToValueMap
}

// KeyMap returns all keys of the enum RedisSetCondition in a K:RedisSetCondition,V:key map
func (i RedisSetCondition) KeyMap() map[RedisSetCondition]string {
	return _RedisSetConditionValueToKeyMap
}

// CaptionMap returns all captions of the enum RedisSetCondition in a K:RedisSetCondition,V:caption map
func (i RedisSetCondition) CaptionMap() map[RedisSetCondition]string {
	return _RedisSetConditionValueToCaptionMap
}

// DescriptionMap returns all descriptions of the enum RedisSetCondition in a K:RedisSetCondition,V:description map
func (i RedisSetCondition) DescriptionMap() map[RedisSetCondition]string {
	return _RedisSetConditionValueToDescriptionMap
}
