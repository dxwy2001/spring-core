/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package testdata

const BitCount = `
{
	"session": "df3b64266ebe4e63a464e135000a07cd",
	"inbound": {},
	"actions": [{
		"protocol": "redis",
		"request": "SET mykey foobar",
		"response": "OK"
	}, {
		"protocol": "redis",
		"request": "BITCOUNT mykey",
		"response": 26
	}, {
		"protocol": "redis",
		"request": "BITCOUNT mykey 0 0",
		"response": 4
	}, {
		"protocol": "redis",
		"request": "BITCOUNT mykey 1 1",
		"response": 6
	}]
}`

const BitOpAnd = `
{
	"session": "df3b64266ebe4e63a464e135000a07cd",
	"inbound": {},
	"actions": [{
		"protocol": "redis",
		"request": "SET key1 foobar",
		"response": "OK"
	}, {
		"protocol": "redis",
		"request": "SET key2 abcdef",
		"response": "OK"
	}, {
		"protocol": "redis",
		"request": "BITOP AND dest key1 key2",
		"response": 6
	}, {
		"protocol": "redis",
		"request": "GET dest",
		"response": "` + "`bc`ab\"" + `
	}]
}`

const BitPos = `
{
	"session": "df3b64266ebe4e63a464e135000a07cd",
	"inbound": {},
	"actions": [{
		"protocol": "redis",
		"request": "SET mykey @\"\\xff\\xf0\\x00\"@",
		"response": "OK"
	}, {
		"protocol": "redis",
		"request": "BITPOS mykey 0",
		"response": 12
	}, {
		"protocol": "redis",
		"request": "SET mykey @\"\\x00\\xff\\xf0\"@",
		"response": "OK"
	}, {
		"protocol": "redis",
		"request": "BITPOS mykey 1 0",
		"response": 8
	}, {
		"protocol": "redis",
		"request": "BITPOS mykey 1 2",
		"response": 16
	}, {
		"protocol": "redis",
		"request": "SET mykey \u0000\u0000\u0000",
		"response": "OK"
	}, {
		"protocol": "redis",
		"request": "BITPOS mykey 1",
		"response": -1
	}]
}`

const GetBit = `
{
	"session": "df3b64266ebe4e63a464e135000a07cd",
	"inbound": {},
	"actions": [{
		"protocol": "redis",
		"request": "SETBIT mykey 7 1",
		"response": 0
	}, {
		"protocol": "redis",
		"request": "GETBIT mykey 0",
		"response": 0
	}, {
		"protocol": "redis",
		"request": "GETBIT mykey 7",
		"response": 1
	}, {
		"protocol": "redis",
		"request": "GETBIT mykey 100",
		"response": 0
	}]
}`

const SetBit = `
{
	"session": "df3b64266ebe4e63a464e135000a07cd",
	"inbound": {},
	"actions": [{
		"protocol": "redis",
		"request": "SETBIT mykey 7 1",
		"response": 0
	}, {
		"protocol": "redis",
		"request": "SETBIT mykey 7 0",
		"response": 1
	}, {
		"protocol": "redis",
		"request": "GET mykey",
		"response": "\u0000"
	}]
}`
