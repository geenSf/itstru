/*
 * Copyright 2020 Matthew A. Titmus
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/go-errors/errors"
	"testing"
)

func TestPut(t *testing.T) {
	const key = "create-key"

	var value Device = Device{
		ID:           "SYSUN0001",
		Name:         "NONAME (MB GIGABYTE GA-H61M-S1)",
		Number:       "101240000001",
		Type:         "SYSUN",
		Manufactured: "",
		Country:      "RUSSIA",
		Fabricator:   "NONAME",
		ExploitFrom:  "2011-09-23",
		AccName:      "",
		Account:      "10020000000000244210124",
		SerialNumber: "2/A/ECA424 00496",
		Included:     "WS11",
		Location:     "К-4",
		User:         "Пастаджян Ксения Сергеевна",
	}

	var val *Device
	var contains bool

	defer delete(store.m, key)

	// Sanity check
	_, contains = store.m[key]
	if contains {
		t.Error("key/value already exists")
	}

	// err should be nil
	err := Put(key, &value)
	if err != nil {
		t.Error(err)
	}

	val, contains = store.m[key]
	if !contains {
		t.Error("create failed")
	}

	if val != &value {
		t.Error("val/value mismatch")
	}
}

//["devname", "invnumb", "devtype", "mnfdate", "mnfcountry", "mnfname", "expldate", "devname_acc", "b_acc", "sernumb","included","location"],
//"1": ["SYSUN0001","NONAME (MB GIGABYTE GA-H61M-S1)", "101240000001", "SYSUN", "", "RUSSIA", "NONAME", "2011-09-23", "", "10020000000000244210124", "2/A/ECA424 00496"],
//"2": ["eMachines E200HV b", "101240000001", "DISPL", "2011-05-01", "CHINA", "ACER", "2011-09-23", "", "10020000000000244210124", "ETQ3Z0W004121047C34302"],
//"3": ["Oklick 100M Black PS/2", "101240000001", "KEYBR", "", "CHINA", "OKLICK", "2011-09-23", "", "10020000000000244210124", "KK100MBKX2ACX00670"],
//"4": ["OKLICK OPTICAL MOUSE 125M", "101240000001", "MOUSE", "", "CHINA", "OKLICK", "2011-09-23", "", "10020000000000244210124", "51659200186"]

func TestGet(t *testing.T) {
	const key = "read-key"

	var value Device = Device{
		ID:           "SYSUN0001",
		Name:         "NONAME (MB GIGABYTE GA-H61M-S1)",
		Number:       "101240000001",
		Type:         "SYSUN",
		Manufactured: "",
		Country:      "RUSSIA",
		Fabricator:   "NONAME",
		ExploitFrom:  "2011-09-23",
		AccName:      "",
		Account:      "10020000000000244210124",
		SerialNumber: "2/A/ECA424 00496",
		Included:     "WS11",
		Location:     "К-4",
		User:         "Пастаджян Ксения Сергеевна",
	}

	var val *Device
	var err error

	defer delete(store.m, key)

	// Read a non-thing
	val, err = Get(key)
	if err == nil {
		t.Error("expected an error")
	}
	if !errors.Is(err, ErrorNoSuchKey) {
		t.Error("unexpected error:", err)
	}

	store.m[key] = &value

	val, err = Get(key)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	if val != &value {
		t.Error("val/value mismatch")
	}
}

func TestDelete(t *testing.T) {
	const key = "delete-key"
	var value Device = Device{
		ID:           "SYSUN0001",
		Name:         "NONAME (MB GIGABYTE GA-H61M-S1)",
		Number:       "101240000001",
		Type:         "SYSUN",
		Manufactured: "",
		Country:      "RUSSIA",
		Fabricator:   "NONAME",
		ExploitFrom:  "2011-09-23",
		AccName:      "",
		Account:      "10020000000000244210124",
		SerialNumber: "2/A/ECA424 00496",
		Included:     "WS11",
		Location:     "К-4",
		User:         "Пастаджян Ксения Сергеевна",
	}
	var contains bool

	defer delete(store.m, key)

	store.m[key] = &value

	_, contains = store.m[key]
	if !contains {
		t.Error("key/value doesn't exist")
	}

	Delete(key)

	_, contains = store.m[key]
	if contains {
		t.Error("Delete failed")
	}
}

func TestGetCollection(t *testing.T) {
	const key1 = "one"
	const key2 = "two"
	const key3 = "free"

	var value1 Device = Device{
		ID:           "0001",
		Name:         "NONAME (MB GIGABYTE GA-H61M-S1)",
		Number:       "101240000001",
		Type:         "SYSUN",
		Manufactured: "",
		Country:      "RUSSIA",
		Fabricator:   "NONAME",
		ExploitFrom:  "2011-09-23",
		AccName:      "",
		Account:      "10020000000000244210124",
		SerialNumber: "2/A/ECA424 00496",
		Included:     "WS11",
		Location:     "К-4",
		User:         "Пастаджян Ксения Сергеевна",
	}

	var value2 Device = Device{
		ID:           "0002",
		Name:         "NONAME (MB GIGABYTE GA-H61M-S1)",
		Number:       "101240000001",
		Type:         "SYSUN",
		Manufactured: "",
		Country:      "RUSSIA",
		Fabricator:   "NONAME",
		ExploitFrom:  "2011-09-23",
		AccName:      "",
		Account:      "10020000000000244210124",
		SerialNumber: "2/A/ECA424 00496",
		Included:     "WS11",
		Location:     "К-4",
		User:         "Пастаджян Ксения Сергеевна",
	}

	var value3 Device = Device{
		ID:           "0003",
		Name:         "NONAME (MB GIGABYTE GA-H61M-S1)",
		Number:       "101240000001",
		Type:         "SYSUN",
		Manufactured: "",
		Country:      "RUSSIA",
		Fabricator:   "NONAME",
		ExploitFrom:  "2011-09-23",
		AccName:      "",
		Account:      "10020000000000244210124",
		SerialNumber: "2/A/ECA424 00496",
		Included:     "WS11",
		Location:     "К-4",
		User:         "Пастаджян Ксения Сергеевна",
	}

	//m := make(map[string]*Device)
	//m[key1] = &value1
	//m[key2] = &value2
	//m[key3] = &value3

	//var val *Device
	var err error

	defer delete(store.m, key1)
	defer delete(store.m, key2)
	defer delete(store.m, key3)

	// Sanity check
	_, contains := store.m[key1]
	if contains {
		t.Error("key1/value already exists")
	}

	_, contains = store.m[key2]
	if contains {
		t.Error("key2/value already exists")
	}

	_, contains = store.m[key3]
	if contains {
		t.Error("key3/value already exists")
	}

	// err should be nil
	err = Put(key1, &value1)
	if err != nil {
		t.Error(err)
	}

	err = Put(key2, &value2)
	if err != nil {
		t.Error(err)
	}

	err = Put(key3, &value3)
	if err != nil {
		t.Error(err)
	}

	// Read a non-thing
	s := GetCollection()

	if s[key1] != &value1 {
		t.Error("on store[key1] expected an error")
	}
	if s[key2] != &value2 {
		t.Error("on store[key2] expected an error")
	}
	if s[key3] != &value3 {
		t.Error("on store[key3] expected an error")
	}

	/*
		if err == nil {
			t.Error("expected an error")
		}
		if !errors.Is(err, ErrorNoSuchKey) {
			t.Error("unexpected error:", err)
		}

		store.m[key] = &value

		val, err = Get(key)
		if err != nil {
			t.Error("unexpected error:", err)
		}

		if val != &value {
			t.Error("val/value mismatch")
		}
	*/

}
