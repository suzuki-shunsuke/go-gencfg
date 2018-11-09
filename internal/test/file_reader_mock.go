package test

// Don't edit this file.
// This file is generated by gomic 0.5.2.
// https://github.com/suzuki-shunsuke/gomic

import (
	testing "testing"

	gomic "github.com/suzuki-shunsuke/gomic/gomic"
)

type (
	// FileReaderMock is a mock.
	FileReaderMock struct {
		t                      *testing.T
		name                   string
		callbackNotImplemented gomic.CallbackNotImplemented
		impl                   struct {
			Read func(src string) (string, error)
		}
	}
)

// NewFileReaderMock returns FileReaderMock .
func NewFileReaderMock(t *testing.T, cb gomic.CallbackNotImplemented) *FileReaderMock {
	return &FileReaderMock{
		t: t, name: "FileReaderMock", callbackNotImplemented: cb}
}

// Read is a mock method.
func (mock FileReaderMock) Read(src string) (string, error) {
	methodName := "Read" // nolint: goconst
	if mock.impl.Read != nil {
		return mock.impl.Read(src)
	}
	if mock.callbackNotImplemented != nil {
		mock.callbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroRead(src)
}

// SetFuncRead sets a method and returns the mock.
func (mock *FileReaderMock) SetFuncRead(impl func(src string) (string, error)) *FileReaderMock {
	mock.impl.Read = impl
	return mock
}

// SetReturnRead sets a fake method.
func (mock *FileReaderMock) SetReturnRead(r0 string, r1 error) *FileReaderMock {
	mock.impl.Read = func(string) (string, error) {
		return r0, r1
	}
	return mock
}

// fakeZeroRead is a fake method which returns zero values.
func (mock FileReaderMock) fakeZeroRead(src string) (string, error) {
	var (
		r0 string
		r1 error
	)
	return r0, r1
}
