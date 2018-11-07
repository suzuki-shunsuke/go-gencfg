package test

// Don't edit this file.
// This file is generated by gomic 0.4.0.
// https://github.com/suzuki-shunsuke/gomic

import (
	testing "testing"

	gomic "github.com/suzuki-shunsuke/gomic/gomic"
)

type (
	// FileWriterMock is a mock.
	FileWriterMock struct {
		t                      *testing.T
		name                   string
		CallbackNotImplemented gomic.CallbackNotImplemented
		Impl                   FileWriterMockImpl
	}

	// FileWriterMockImpl holds functions which implement interface's methods.
	FileWriterMockImpl struct {
		Write func(dest string, data []byte) error
	}
)

// NewFileWriterMock returns FileWriterMock .
func NewFileWriterMock(t *testing.T, cb gomic.CallbackNotImplemented) *FileWriterMock {
	return &FileWriterMock{t: t, CallbackNotImplemented: cb}
}

// Write is a mock method.
func (mock FileWriterMock) Write(dest string, data []byte) error {
	methodName := "Write" // nolint: goconst
	if mock.Impl.Write != nil {
		return mock.Impl.Write(dest, data)
	}
	if mock.CallbackNotImplemented != nil {
		mock.CallbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroWrite(dest, data)
}

// SetFakeWrite sets a fake method.
func (mock *FileWriterMock) SetFakeWrite(r0 error) {
	mock.Impl.Write = func(dest string, data []byte) error {
		return r0
	}
}

// fakeZeroWrite is a fake method which returns zero values.
func (mock FileWriterMock) fakeZeroWrite(dest string, data []byte) error {
	var (
		r0 error
	)
	return r0
}