// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	firefly "github.com/randaalex/finance_bot/pkg/firefly"

	mock "github.com/stretchr/testify/mock"

	os "os"
)

// AttachmentsApi is an autogenerated mock type for the AttachmentsApi type
type AttachmentsApi struct {
	mock.Mock
}

// DeleteAttachment provides a mock function with given fields: ctx, id
func (_m *AttachmentsApi) DeleteAttachment(ctx context.Context, id int32) firefly.ApiDeleteAttachmentRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiDeleteAttachmentRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiDeleteAttachmentRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiDeleteAttachmentRequest)
	}

	return r0
}

// DeleteAttachmentExecute provides a mock function with given fields: r
func (_m *AttachmentsApi) DeleteAttachmentExecute(r firefly.ApiDeleteAttachmentRequest) (*http.Response, error) {
	ret := _m.Called(r)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(firefly.ApiDeleteAttachmentRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(firefly.ApiDeleteAttachmentRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DownloadAttachment provides a mock function with given fields: ctx, id
func (_m *AttachmentsApi) DownloadAttachment(ctx context.Context, id int32) firefly.ApiDownloadAttachmentRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiDownloadAttachmentRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiDownloadAttachmentRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiDownloadAttachmentRequest)
	}

	return r0
}

// DownloadAttachmentExecute provides a mock function with given fields: r
func (_m *AttachmentsApi) DownloadAttachmentExecute(r firefly.ApiDownloadAttachmentRequest) (*os.File, *http.Response, error) {
	ret := _m.Called(r)

	var r0 *os.File
	if rf, ok := ret.Get(0).(func(firefly.ApiDownloadAttachmentRequest) *os.File); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*os.File)
		}
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiDownloadAttachmentRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiDownloadAttachmentRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAttachment provides a mock function with given fields: ctx, id
func (_m *AttachmentsApi) GetAttachment(ctx context.Context, id int32) firefly.ApiGetAttachmentRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiGetAttachmentRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiGetAttachmentRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiGetAttachmentRequest)
	}

	return r0
}

// GetAttachmentExecute provides a mock function with given fields: r
func (_m *AttachmentsApi) GetAttachmentExecute(r firefly.ApiGetAttachmentRequest) (firefly.AttachmentSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.AttachmentSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiGetAttachmentRequest) firefly.AttachmentSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.AttachmentSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiGetAttachmentRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiGetAttachmentRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListAttachment provides a mock function with given fields: ctx
func (_m *AttachmentsApi) ListAttachment(ctx context.Context) firefly.ApiListAttachmentRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiListAttachmentRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiListAttachmentRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiListAttachmentRequest)
	}

	return r0
}

// ListAttachmentExecute provides a mock function with given fields: r
func (_m *AttachmentsApi) ListAttachmentExecute(r firefly.ApiListAttachmentRequest) (firefly.AttachmentArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.AttachmentArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListAttachmentRequest) firefly.AttachmentArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.AttachmentArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListAttachmentRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListAttachmentRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// StoreAttachment provides a mock function with given fields: ctx
func (_m *AttachmentsApi) StoreAttachment(ctx context.Context) firefly.ApiStoreAttachmentRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiStoreAttachmentRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiStoreAttachmentRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiStoreAttachmentRequest)
	}

	return r0
}

// StoreAttachmentExecute provides a mock function with given fields: r
func (_m *AttachmentsApi) StoreAttachmentExecute(r firefly.ApiStoreAttachmentRequest) (firefly.AttachmentSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.AttachmentSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiStoreAttachmentRequest) firefly.AttachmentSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.AttachmentSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiStoreAttachmentRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiStoreAttachmentRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateAttachment provides a mock function with given fields: ctx, id
func (_m *AttachmentsApi) UpdateAttachment(ctx context.Context, id int32) firefly.ApiUpdateAttachmentRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiUpdateAttachmentRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiUpdateAttachmentRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiUpdateAttachmentRequest)
	}

	return r0
}

// UpdateAttachmentExecute provides a mock function with given fields: r
func (_m *AttachmentsApi) UpdateAttachmentExecute(r firefly.ApiUpdateAttachmentRequest) (firefly.AttachmentSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.AttachmentSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiUpdateAttachmentRequest) firefly.AttachmentSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.AttachmentSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiUpdateAttachmentRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiUpdateAttachmentRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UploadAttachment provides a mock function with given fields: ctx, id
func (_m *AttachmentsApi) UploadAttachment(ctx context.Context, id int32) firefly.ApiUploadAttachmentRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiUploadAttachmentRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiUploadAttachmentRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiUploadAttachmentRequest)
	}

	return r0
}

// UploadAttachmentExecute provides a mock function with given fields: r
func (_m *AttachmentsApi) UploadAttachmentExecute(r firefly.ApiUploadAttachmentRequest) (*http.Response, error) {
	ret := _m.Called(r)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(firefly.ApiUploadAttachmentRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(firefly.ApiUploadAttachmentRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
