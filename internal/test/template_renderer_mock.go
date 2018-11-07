package test

// Don't edit this file.
// This file is generated by gomic 0.4.0.
// https://github.com/suzuki-shunsuke/gomic

import (
	testing "testing"

	gomic "github.com/suzuki-shunsuke/gomic/gomic"
)

type (
	// TemplateRendererMock is a mock.
	TemplateRendererMock struct {
		t                      *testing.T
		name                   string
		CallbackNotImplemented gomic.CallbackNotImplemented
		Impl                   TemplateRendererMockImpl
	}

	// TemplateRendererMockImpl holds functions which implement interface's methods.
	TemplateRendererMockImpl struct {
		Render func(txt string, data interface{}) (string, error)
	}
)

// NewTemplateRendererMock returns TemplateRendererMock .
func NewTemplateRendererMock(t *testing.T, cb gomic.CallbackNotImplemented) *TemplateRendererMock {
	return &TemplateRendererMock{t: t, CallbackNotImplemented: cb}
}

// Render is a mock method.
func (mock TemplateRendererMock) Render(txt string, data interface{}) (string, error) {
	methodName := "Render" // nolint: goconst
	if mock.Impl.Render != nil {
		return mock.Impl.Render(txt, data)
	}
	if mock.CallbackNotImplemented != nil {
		mock.CallbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroRender(txt, data)
}

// SetFakeRender sets a fake method.
func (mock *TemplateRendererMock) SetFakeRender(r0 string, r1 error) {
	mock.Impl.Render = func(txt string, data interface{}) (string, error) {
		return r0, r1
	}
}

// fakeZeroRender is a fake method which returns zero values.
func (mock TemplateRendererMock) fakeZeroRender(txt string, data interface{}) (string, error) {
	var (
		r0 string
		r1 error
	)
	return r0, r1
}