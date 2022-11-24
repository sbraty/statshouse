// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/statshouse/internal/vkgo/basictl"
)

var _ = basictl.NatWrite

type Int int32

func (Int) TLName() string { return "int" }
func (Int) TLTag() uint32  { return 0xa8509bda }

func (item *Int) Reset() {
	ptr := (*int32)(item)
	*ptr = 0
}

func (item *Int) Read(w []byte) (_ []byte, err error) {
	ptr := (*int32)(item)
	return basictl.IntRead(w, ptr)
}

func (item *Int) Write(w []byte) (_ []byte, err error) {
	ptr := (*int32)(item)
	return basictl.IntWrite(w, *ptr), nil
}

func (item *Int) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xa8509bda); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *Int) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xa8509bda)
	return item.Write(w)
}

func (item Int) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func Int__ReadJSON(item *Int, j interface{}) error { return item.readJSON(j) }
func (item *Int) readJSON(j interface{}) error {
	ptr := (*int32)(item)
	if err := JsonReadInt32(j, ptr); err != nil {
		return err
	}
	return nil
}

func (item *Int) WriteJSON(w []byte) (_ []byte, err error) {
	ptr := (*int32)(item)
	w = basictl.JSONWriteInt32(w, *ptr)
	return w, nil
}
func (item *Int) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *Int) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("int", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("int", err.Error())
	}
	return nil
}

func TupleInt4Reset(vec *[4]int32) {
	for i := range *vec {
		(*vec)[i] = 0
	}
}

func TupleInt4Read(w []byte, vec *[4]int32) (_ []byte, err error) {
	for i := range *vec {
		if w, err = basictl.IntRead(w, &(*vec)[i]); err != nil {
			return w, err
		}
	}
	return w, nil
}

func TupleInt4Write(w []byte, vec *[4]int32) (_ []byte, err error) {
	for _, elem := range *vec {
		w = basictl.IntWrite(w, elem)
	}
	return w, nil
}

func TupleInt4ReadJSON(j interface{}, vec *[4]int32) error {
	_, _arr, err := JsonReadArrayFixedSize("[4]int32", j, 4)
	if err != nil {
		return err
	}
	for i := range *vec {
		if err := JsonReadInt32(_arr[i], &(*vec)[i]); err != nil {
			return err
		}
	}
	return nil
}

func TupleInt4WriteJSON(w []byte, vec *[4]int32) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range *vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteInt32(w, elem)
	}
	return append(w, ']'), nil
}

func VectorInt0Read(w []byte, vec *[]int32) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]int32, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = basictl.IntRead(w, &(*vec)[i]); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorInt0Write(w []byte, vec []int32) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = basictl.IntWrite(w, elem)
	}
	return w, nil
}

func VectorInt0ReadJSON(j interface{}, vec *[]int32) error {
	l, _arr, err := JsonReadArray("[]int32", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]int32, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := JsonReadInt32(_arr[i], &(*vec)[i]); err != nil {
			return err
		}
	}
	return nil
}

func VectorInt0WriteJSON(w []byte, vec []int32) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteInt32(w, elem)
	}
	return append(w, ']'), nil
}
