package test

// Don't edit this file.
// This file is generated by gomic 0.5.1.
// https://github.com/suzuki-shunsuke/gomic

import (
	testing "testing"

	test "github.com/suzuki-shunsuke/go-gencfg/internal/domain"
	gomic "github.com/suzuki-shunsuke/gomic/gomic"
)

type (
	// CodeGeneratorMock is a mock.
	CodeGeneratorMock struct {
		t                      *testing.T
		name                   string
		callbackNotImplemented gomic.CallbackNotImplemented
		impl                   struct {
			Exec func(dest, tmpl string, config test.TemplateData) error
		}
	}
)

// NewCodeGeneratorMock returns CodeGeneratorMock .
func NewCodeGeneratorMock(t *testing.T, cb gomic.CallbackNotImplemented) *CodeGeneratorMock {
	return &CodeGeneratorMock{t: t, callbackNotImplemented: cb}
}

// Exec is a mock method.
func (mock CodeGeneratorMock) Exec(dest, tmpl string, config test.TemplateData) error {
	methodName := "Exec" // nolint: goconst
	if mock.impl.Exec != nil {
		return mock.impl.Exec(dest, tmpl, config)
	}
	if mock.callbackNotImplemented != nil {
		mock.callbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroExec(dest, tmpl, config)
}

// SetFuncExec sets a method and returns the mock.
func (mock *CodeGeneratorMock) SetFuncExec(impl func(dest, tmpl string, config test.TemplateData) error) *CodeGeneratorMock {
	mock.impl.Exec = impl
	return mock
}

// SetReturnExec sets a fake method.
func (mock *CodeGeneratorMock) SetReturnExec(r0 error) *CodeGeneratorMock {
	mock.impl.Exec = func(string, string, test.TemplateData) error {
		return r0
	}
	return mock
}

// fakeZeroExec is a fake method which returns zero values.
func (mock CodeGeneratorMock) fakeZeroExec(dest, tmpl string, config test.TemplateData) error {
	var (
		r0 error
	)
	return r0
}
