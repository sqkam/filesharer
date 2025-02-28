// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/file/v1/file.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ListByAddrRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListByAddrRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListByAddrRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListByAddrRequestMultiError, or nil if none found.
func (m *ListByAddrRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListByAddrRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Addr

	// no validation rules for Path

	if len(errors) > 0 {
		return ListByAddrRequestMultiError(errors)
	}

	return nil
}

// ListByAddrRequestMultiError is an error wrapping multiple validation errors
// returned by ListByAddrRequest.ValidateAll() if the designated constraints
// aren't met.
type ListByAddrRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListByAddrRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListByAddrRequestMultiError) AllErrors() []error { return m }

// ListByAddrRequestValidationError is the validation error returned by
// ListByAddrRequest.Validate if the designated constraints aren't met.
type ListByAddrRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListByAddrRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListByAddrRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListByAddrRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListByAddrRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListByAddrRequestValidationError) ErrorName() string {
	return "ListByAddrRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListByAddrRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListByAddrRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListByAddrRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListByAddrRequestValidationError{}

// Validate checks the field values on ListByAddrReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListByAddrReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListByAddrReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListByAddrReplyMultiError, or nil if none found.
func (m *ListByAddrReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListByAddrReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListByAddrReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListByAddrReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListByAddrReplyValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListByAddrReplyMultiError(errors)
	}

	return nil
}

// ListByAddrReplyMultiError is an error wrapping multiple validation errors
// returned by ListByAddrReply.ValidateAll() if the designated constraints
// aren't met.
type ListByAddrReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListByAddrReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListByAddrReplyMultiError) AllErrors() []error { return m }

// ListByAddrReplyValidationError is the validation error returned by
// ListByAddrReply.Validate if the designated constraints aren't met.
type ListByAddrReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListByAddrReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListByAddrReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListByAddrReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListByAddrReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListByAddrReplyValidationError) ErrorName() string { return "ListByAddrReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListByAddrReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListByAddrReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListByAddrReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListByAddrReplyValidationError{}

// Validate checks the field values on GetDetailByAddrRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDetailByAddrRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDetailByAddrRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDetailByAddrRequestMultiError, or nil if none found.
func (m *GetDetailByAddrRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDetailByAddrRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Addr

	// no validation rules for Path

	if len(errors) > 0 {
		return GetDetailByAddrRequestMultiError(errors)
	}

	return nil
}

// GetDetailByAddrRequestMultiError is an error wrapping multiple validation
// errors returned by GetDetailByAddrRequest.ValidateAll() if the designated
// constraints aren't met.
type GetDetailByAddrRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDetailByAddrRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDetailByAddrRequestMultiError) AllErrors() []error { return m }

// GetDetailByAddrRequestValidationError is the validation error returned by
// GetDetailByAddrRequest.Validate if the designated constraints aren't met.
type GetDetailByAddrRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDetailByAddrRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDetailByAddrRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDetailByAddrRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDetailByAddrRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDetailByAddrRequestValidationError) ErrorName() string {
	return "GetDetailByAddrRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDetailByAddrRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDetailByAddrRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDetailByAddrRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDetailByAddrRequestValidationError{}

// Validate checks the field values on GetDetailByAddrReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDetailByAddrReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDetailByAddrReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDetailByAddrReplyMultiError, or nil if none found.
func (m *GetDetailByAddrReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDetailByAddrReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Path

	// no validation rules for Size

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return GetDetailByAddrReplyMultiError(errors)
	}

	return nil
}

// GetDetailByAddrReplyMultiError is an error wrapping multiple validation
// errors returned by GetDetailByAddrReply.ValidateAll() if the designated
// constraints aren't met.
type GetDetailByAddrReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDetailByAddrReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDetailByAddrReplyMultiError) AllErrors() []error { return m }

// GetDetailByAddrReplyValidationError is the validation error returned by
// GetDetailByAddrReply.Validate if the designated constraints aren't met.
type GetDetailByAddrReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDetailByAddrReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDetailByAddrReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDetailByAddrReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDetailByAddrReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDetailByAddrReplyValidationError) ErrorName() string {
	return "GetDetailByAddrReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetDetailByAddrReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDetailByAddrReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDetailByAddrReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDetailByAddrReplyValidationError{}

// Validate checks the field values on DownloadByAddrRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownloadByAddrRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadByAddrRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadByAddrRequestMultiError, or nil if none found.
func (m *DownloadByAddrRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadByAddrRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Addr

	if utf8.RuneCountInString(m.GetPath()) < 1 {
		err := DownloadByAddrRequestValidationError{
			field:  "Path",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DownloadByAddrRequestMultiError(errors)
	}

	return nil
}

// DownloadByAddrRequestMultiError is an error wrapping multiple validation
// errors returned by DownloadByAddrRequest.ValidateAll() if the designated
// constraints aren't met.
type DownloadByAddrRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadByAddrRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadByAddrRequestMultiError) AllErrors() []error { return m }

// DownloadByAddrRequestValidationError is the validation error returned by
// DownloadByAddrRequest.Validate if the designated constraints aren't met.
type DownloadByAddrRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadByAddrRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadByAddrRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadByAddrRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadByAddrRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadByAddrRequestValidationError) ErrorName() string {
	return "DownloadByAddrRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadByAddrRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadByAddrRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadByAddrRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadByAddrRequestValidationError{}

// Validate checks the field values on DownloadByAddrReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownloadByAddrReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadByAddrReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadByAddrReplyMultiError, or nil if none found.
func (m *DownloadByAddrReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadByAddrReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Data

	if len(errors) > 0 {
		return DownloadByAddrReplyMultiError(errors)
	}

	return nil
}

// DownloadByAddrReplyMultiError is an error wrapping multiple validation
// errors returned by DownloadByAddrReply.ValidateAll() if the designated
// constraints aren't met.
type DownloadByAddrReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadByAddrReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadByAddrReplyMultiError) AllErrors() []error { return m }

// DownloadByAddrReplyValidationError is the validation error returned by
// DownloadByAddrReply.Validate if the designated constraints aren't met.
type DownloadByAddrReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadByAddrReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadByAddrReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadByAddrReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadByAddrReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadByAddrReplyValidationError) ErrorName() string {
	return "DownloadByAddrReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadByAddrReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadByAddrReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadByAddrReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadByAddrReplyValidationError{}

// Validate checks the field values on DownloadDirByAddrRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownloadDirByAddrRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadDirByAddrRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadDirByAddrRequestMultiError, or nil if none found.
func (m *DownloadDirByAddrRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadDirByAddrRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Addr

	if utf8.RuneCountInString(m.GetPath()) < 1 {
		err := DownloadDirByAddrRequestValidationError{
			field:  "Path",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DownloadDirByAddrRequestMultiError(errors)
	}

	return nil
}

// DownloadDirByAddrRequestMultiError is an error wrapping multiple validation
// errors returned by DownloadDirByAddrRequest.ValidateAll() if the designated
// constraints aren't met.
type DownloadDirByAddrRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadDirByAddrRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadDirByAddrRequestMultiError) AllErrors() []error { return m }

// DownloadDirByAddrRequestValidationError is the validation error returned by
// DownloadDirByAddrRequest.Validate if the designated constraints aren't met.
type DownloadDirByAddrRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadDirByAddrRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadDirByAddrRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadDirByAddrRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadDirByAddrRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadDirByAddrRequestValidationError) ErrorName() string {
	return "DownloadDirByAddrRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadDirByAddrRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadDirByAddrRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadDirByAddrRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadDirByAddrRequestValidationError{}

// Validate checks the field values on DownloadDirByAddrReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownloadDirByAddrReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadDirByAddrReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadDirByAddrReplyMultiError, or nil if none found.
func (m *DownloadDirByAddrReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadDirByAddrReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Data

	if len(errors) > 0 {
		return DownloadDirByAddrReplyMultiError(errors)
	}

	return nil
}

// DownloadDirByAddrReplyMultiError is an error wrapping multiple validation
// errors returned by DownloadDirByAddrReply.ValidateAll() if the designated
// constraints aren't met.
type DownloadDirByAddrReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadDirByAddrReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadDirByAddrReplyMultiError) AllErrors() []error { return m }

// DownloadDirByAddrReplyValidationError is the validation error returned by
// DownloadDirByAddrReply.Validate if the designated constraints aren't met.
type DownloadDirByAddrReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadDirByAddrReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadDirByAddrReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadDirByAddrReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadDirByAddrReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadDirByAddrReplyValidationError) ErrorName() string {
	return "DownloadDirByAddrReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadDirByAddrReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadDirByAddrReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadDirByAddrReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadDirByAddrReplyValidationError{}

// Validate checks the field values on ListNodeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListNodeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListNodeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListNodeRequestMultiError, or nil if none found.
func (m *ListNodeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListNodeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListNodeRequestMultiError(errors)
	}

	return nil
}

// ListNodeRequestMultiError is an error wrapping multiple validation errors
// returned by ListNodeRequest.ValidateAll() if the designated constraints
// aren't met.
type ListNodeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListNodeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListNodeRequestMultiError) AllErrors() []error { return m }

// ListNodeRequestValidationError is the validation error returned by
// ListNodeRequest.Validate if the designated constraints aren't met.
type ListNodeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListNodeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListNodeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListNodeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListNodeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListNodeRequestValidationError) ErrorName() string { return "ListNodeRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListNodeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListNodeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListNodeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListNodeRequestValidationError{}

// Validate checks the field values on ListNodeReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListNodeReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListNodeReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListNodeReplyMultiError, or
// nil if none found.
func (m *ListNodeReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListNodeReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListNodeReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListNodeReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListNodeReplyValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListNodeReplyMultiError(errors)
	}

	return nil
}

// ListNodeReplyMultiError is an error wrapping multiple validation errors
// returned by ListNodeReply.ValidateAll() if the designated constraints
// aren't met.
type ListNodeReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListNodeReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListNodeReplyMultiError) AllErrors() []error { return m }

// ListNodeReplyValidationError is the validation error returned by
// ListNodeReply.Validate if the designated constraints aren't met.
type ListNodeReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListNodeReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListNodeReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListNodeReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListNodeReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListNodeReplyValidationError) ErrorName() string { return "ListNodeReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListNodeReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListNodeReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListNodeReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListNodeReplyValidationError{}

// Validate checks the field values on ListByAddrReplyItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListByAddrReplyItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListByAddrReplyItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListByAddrReplyItemMultiError, or nil if none found.
func (m *ListByAddrReplyItem) ValidateAll() error {
	return m.validate(true)
}

func (m *ListByAddrReplyItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Path

	// no validation rules for IsDir

	// no validation rules for Size

	if len(errors) > 0 {
		return ListByAddrReplyItemMultiError(errors)
	}

	return nil
}

// ListByAddrReplyItemMultiError is an error wrapping multiple validation
// errors returned by ListByAddrReplyItem.ValidateAll() if the designated
// constraints aren't met.
type ListByAddrReplyItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListByAddrReplyItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListByAddrReplyItemMultiError) AllErrors() []error { return m }

// ListByAddrReplyItemValidationError is the validation error returned by
// ListByAddrReplyItem.Validate if the designated constraints aren't met.
type ListByAddrReplyItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListByAddrReplyItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListByAddrReplyItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListByAddrReplyItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListByAddrReplyItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListByAddrReplyItemValidationError) ErrorName() string {
	return "ListByAddrReplyItemValidationError"
}

// Error satisfies the builtin error interface
func (e ListByAddrReplyItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListByAddrReplyItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListByAddrReplyItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListByAddrReplyItemValidationError{}

// Validate checks the field values on ListNodeReplyItem with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListNodeReplyItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListNodeReplyItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListNodeReplyItemMultiError, or nil if none found.
func (m *ListNodeReplyItem) ValidateAll() error {
	return m.validate(true)
}

func (m *ListNodeReplyItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ServiceID

	// no validation rules for ServiceAddress

	// no validation rules for ServicePort

	if len(errors) > 0 {
		return ListNodeReplyItemMultiError(errors)
	}

	return nil
}

// ListNodeReplyItemMultiError is an error wrapping multiple validation errors
// returned by ListNodeReplyItem.ValidateAll() if the designated constraints
// aren't met.
type ListNodeReplyItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListNodeReplyItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListNodeReplyItemMultiError) AllErrors() []error { return m }

// ListNodeReplyItemValidationError is the validation error returned by
// ListNodeReplyItem.Validate if the designated constraints aren't met.
type ListNodeReplyItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListNodeReplyItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListNodeReplyItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListNodeReplyItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListNodeReplyItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListNodeReplyItemValidationError) ErrorName() string {
	return "ListNodeReplyItemValidationError"
}

// Error satisfies the builtin error interface
func (e ListNodeReplyItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListNodeReplyItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListNodeReplyItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListNodeReplyItemValidationError{}
