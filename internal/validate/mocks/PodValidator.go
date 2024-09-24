// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/api/core/v1"

	validate "github.com/kyma-project/warden/internal/validate"
)

// PodValidator is an autogenerated mock type for the PodValidator type
type PodValidator struct {
	mock.Mock
}

// ValidatePod provides a mock function with given fields: ctx, pod, ns
func (_m *PodValidator) ValidatePod(ctx context.Context, pod *v1.Pod, ns *v1.Namespace) (validate.ValidationResult, error) {
	ret := _m.Called(ctx, pod, ns)

	if len(ret) == 0 {
		panic("no return value specified for ValidatePod")
	}

	var r0 validate.ValidationResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Pod, *v1.Namespace) (validate.ValidationResult, error)); ok {
		return rf(ctx, pod, ns)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Pod, *v1.Namespace) validate.ValidationResult); ok {
		r0 = rf(ctx, pod, ns)
	} else {
		r0 = ret.Get(0).(validate.ValidationResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.Pod, *v1.Namespace) error); ok {
		r1 = rf(ctx, pod, ns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPodValidator creates a new instance of PodValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPodValidator(t interface {
	mock.TestingT
	Cleanup(func())
}) *PodValidator {
	mock := &PodValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
