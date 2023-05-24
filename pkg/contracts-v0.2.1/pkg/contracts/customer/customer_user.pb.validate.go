// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: customer_user.proto

package customer

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

// define the regex for a UUID once up-front
var _customer_user_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserRequestMultiError, or nil if none found.
func (m *CreateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 6 {
		err := CreateUserRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 6 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateUuid(m.GetOfficeUuid()); err != nil {
		err = CreateUserRequestValidationError{
			field:  "OfficeUuid",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateUserRequestMultiError(errors)
	}

	return nil
}

func (m *CreateUserRequest) _validateUuid(uuid string) error {
	if matched := _customer_user_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CreateUserRequestMultiError is an error wrapping multiple validation errors
// returned by CreateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserRequestMultiError) AllErrors() []error { return m }

// CreateUserRequestValidationError is the validation error returned by
// CreateUserRequest.Validate if the designated constraints aren't met.
type CreateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserRequestValidationError) ErrorName() string {
	return "CreateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserRequestValidationError{}

// Validate checks the field values on CreateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserResponseMultiError, or nil if none found.
func (m *CreateUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateUserResponseMultiError(errors)
	}

	return nil
}

// CreateUserResponseMultiError is an error wrapping multiple validation errors
// returned by CreateUserResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserResponseMultiError) AllErrors() []error { return m }

// CreateUserResponseValidationError is the validation error returned by
// CreateUserResponse.Validate if the designated constraints aren't met.
type CreateUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserResponseValidationError) ErrorName() string {
	return "CreateUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserResponseValidationError{}

// Validate checks the field values on GetUserListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserListRequestMultiError, or nil if none found.
func (m *GetUserListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetOfficeUuid()); err != nil {
		err = GetUserListRequestValidationError{
			field:  "OfficeUuid",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetUserListRequestMultiError(errors)
	}

	return nil
}

func (m *GetUserListRequest) _validateUuid(uuid string) error {
	if matched := _customer_user_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetUserListRequestMultiError is an error wrapping multiple validation errors
// returned by GetUserListRequest.ValidateAll() if the designated constraints
// aren't met.
type GetUserListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserListRequestMultiError) AllErrors() []error { return m }

// GetUserListRequestValidationError is the validation error returned by
// GetUserListRequest.Validate if the designated constraints aren't met.
type GetUserListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserListRequestValidationError) ErrorName() string {
	return "GetUserListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserListRequestValidationError{}

// Validate checks the field values on GetUserListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserListResponseMultiError, or nil if none found.
func (m *GetUserListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResult() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetUserListResponseValidationError{
						field:  fmt.Sprintf("Result[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetUserListResponseValidationError{
						field:  fmt.Sprintf("Result[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetUserListResponseValidationError{
					field:  fmt.Sprintf("Result[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetUserListResponseMultiError(errors)
	}

	return nil
}

// GetUserListResponseMultiError is an error wrapping multiple validation
// errors returned by GetUserListResponse.ValidateAll() if the designated
// constraints aren't met.
type GetUserListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserListResponseMultiError) AllErrors() []error { return m }

// GetUserListResponseValidationError is the validation error returned by
// GetUserListResponse.Validate if the designated constraints aren't met.
type GetUserListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserListResponseValidationError) ErrorName() string {
	return "GetUserListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserListResponseValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uuid

	// no validation rules for Name

	// no validation rules for OfficeUuid

	// no validation rules for OfficeName

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}