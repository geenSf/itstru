/*
 * Copyright 2022 Eugene Shchemeleff
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
	"encoding/json"
	"errors"
	"sync"
)

/*
type Devtype string

const (
	Sysunit   Devtype = "SYSUN" //system unit of computer
	Display   Devtype = "DISPL" //display, monitor
	Keyboard  Devtype = "KEYBR" //keyboard
	Mouse     Devtype = "MOUSE" //mouse
	MFD       Devtype = "PRMFD" //multifunctional device
	Laptop    Devtype = "LPTOP" //notebook, laptop
	Scanner   Devtype = "SCAND" //flatbed scanner device, other scanner device
	UPS       Devtype = "UPSUN" //uninterruptible power supply
	Moniputer Devtype = "MONPU" //all-in-one computer
	PBX       Devtype = "PBXUN" //private branch exchange unit
	Router    Devtype = "NARTR" //network active equipment, router
	Server    Devtype = "SERVR" //net server
	Switch    Devtype = "NASWT" //network active equipment, switch
	Token     Devtype = "CRTOK" //hardware crypto token
	Modem     Devtype = "NAMDM" //network active equipment, modem
)
*/

//["devname", "invnumb", "devtype", "mnfdate", "mnfcountry", "mnfname", "expldate", "devname_acc", "b_acc", "sernumb","included","location"]

//device description
type Device struct {
	ID           string `json:"id"`
	Name         string `json:"devname"`
	Number       string `json:"invnumb"`
	Type         string `json:"devtype"`
	Manufactured string `json:"mnfdate"`
	Country      string `json:"mnfcountry"`
	Fabricator   string `json:"mnfname"`
	ExploitFrom  string `json:"expldate"`
	AccName      string `json:"devname_acc"`
	Account      string `json:"account"`
	SerialNumber string `json:"sernumb"`
	Included     string `json:"included"`
	Location     string `json:"location"`
	User         string `json:"user"`
}

var dev *Device

var store = struct {
	sync.RWMutex
	m map[string]*Device
}{m: make(map[string]*Device)}

var ErrorNoSuchKey = errors.New("no such key")

func Delete(key string) error {
	store.Lock()
	delete(store.m, key)
	store.Unlock()

	return nil
}

func Get(key string) (string, error) {

	store.RLock()
	value, ok := store.m[key]
	store.RUnlock()

	if !ok {
		return "", ErrorNoSuchKey
	}

	json_data, err := json.Marshal(&value)

	if err != nil {
		return "", err
	}

	return string(json_data), nil

}

func GetCollection() map[string]*Device {

	s := make(map[string]*Device)

	store.RLock()
	s = store.m
	store.RUnlock()

	return s

}

func Put(key string, value string) error {

	err := json.Unmarshal([]byte(value), &dev)
	if err != nil {
		return err
	}

	store.Lock()
	store.m[key] = dev
	store.Unlock()

	return nil
}
