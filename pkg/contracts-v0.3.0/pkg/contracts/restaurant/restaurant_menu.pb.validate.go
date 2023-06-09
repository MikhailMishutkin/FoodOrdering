// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: restaurant_menu.proto

package restaurant

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

// Validate checks the field values on CreateMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateMenuRequestMultiError, or nil if none found.
func (m *CreateMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOnDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateMenuRequestValidationError{
					field:  "OnDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateMenuRequestValidationError{
					field:  "OnDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOnDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateMenuRequestValidationError{
				field:  "OnDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOpeningRecordAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateMenuRequestValidationError{
					field:  "OpeningRecordAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateMenuRequestValidationError{
					field:  "OpeningRecordAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpeningRecordAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateMenuRequestValidationError{
				field:  "OpeningRecordAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetClosingRecordAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateMenuRequestValidationError{
					field:  "ClosingRecordAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateMenuRequestValidationError{
					field:  "ClosingRecordAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetClosingRecordAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateMenuRequestValidationError{
				field:  "ClosingRecordAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateMenuRequestMultiError(errors)
	}

	return nil
}

// CreateMenuRequestMultiError is an error wrapping multiple validation errors
// returned by CreateMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateMenuRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateMenuRequestMultiError) AllErrors() []error { return m }

// CreateMenuRequestValidationError is the validation error returned by
// CreateMenuRequest.Validate if the designated constraints aren't met.
type CreateMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMenuRequestValidationError) ErrorName() string {
	return "CreateMenuRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMenuRequestValidationError{}

// Validate checks the field values on CreateMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateMenuResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateMenuResponseMultiError, or nil if none found.
func (m *CreateMenuResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateMenuResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateMenuResponseMultiError(errors)
	}

	return nil
}

// CreateMenuResponseMultiError is an error wrapping multiple validation errors
// returned by CreateMenuResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateMenuResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateMenuResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateMenuResponseMultiError) AllErrors() []error { return m }

// CreateMenuResponseValidationError is the validation error returned by
// CreateMenuResponse.Validate if the designated constraints aren't met.
type CreateMenuResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMenuResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMenuResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMenuResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMenuResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMenuResponseValidationError) ErrorName() string {
	return "CreateMenuResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMenuResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMenuResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMenuResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMenuResponseValidationError{}

// Validate checks the field values on GetMenuRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetMenuRequestMultiError,
// or nil if none found.
func (m *GetMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOnDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetMenuRequestValidationError{
					field:  "OnDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetMenuRequestValidationError{
					field:  "OnDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOnDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetMenuRequestValidationError{
				field:  "OnDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetMenuRequestMultiError(errors)
	}

	return nil
}

// GetMenuRequestMultiError is an error wrapping multiple validation errors
// returned by GetMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type GetMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMenuRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMenuRequestMultiError) AllErrors() []error { return m }

// GetMenuRequestValidationError is the validation error returned by
// GetMenuRequest.Validate if the designated constraints aren't met.
type GetMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMenuRequestValidationError) ErrorName() string { return "GetMenuRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMenuRequestValidationError{}

// Validate checks the field values on GetMenuResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetMenuResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetMenuResponseMultiError, or nil if none found.
func (m *GetMenuResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMenuResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMenu()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetMenuResponseValidationError{
					field:  "Menu",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetMenuResponseValidationError{
					field:  "Menu",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMenu()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetMenuResponseValidationError{
				field:  "Menu",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetMenuResponseMultiError(errors)
	}

	return nil
}

// GetMenuResponseMultiError is an error wrapping multiple validation errors
// returned by GetMenuResponse.ValidateAll() if the designated constraints
// aren't met.
type GetMenuResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMenuResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMenuResponseMultiError) AllErrors() []error { return m }

// GetMenuResponseValidationError is the validation error returned by
// GetMenuResponse.Validate if the designated constraints aren't met.
type GetMenuResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMenuResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMenuResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMenuResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMenuResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMenuResponseValidationError) ErrorName() string { return "GetMenuResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetMenuResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMenuResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMenuResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMenuResponseValidationError{}

// Validate checks the field values on Menu with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Menu) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Menu with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MenuMultiError, or nil if none found.
func (m *Menu) ValidateAll() error {
	return m.validate(true)
}

func (m *Menu) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uuid

	if all {
		switch v := interface{}(m.GetOnDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MenuValidationError{
					field:  "OnDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MenuValidationError{
					field:  "OnDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOnDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MenuValidationError{
				field:  "OnDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOpeningRecordAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MenuValidationError{
					field:  "OpeningRecordAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MenuValidationError{
					field:  "OpeningRecordAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpeningRecordAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MenuValidationError{
				field:  "OpeningRecordAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetClosingRecordAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MenuValidationError{
					field:  "ClosingRecordAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MenuValidationError{
					field:  "ClosingRecordAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetClosingRecordAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MenuValidationError{
				field:  "ClosingRecordAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSalads() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Salads[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Salads[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MenuValidationError{
					field:  fmt.Sprintf("Salads[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetGarnishes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Garnishes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Garnishes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MenuValidationError{
					field:  fmt.Sprintf("Garnishes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetMeats() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Meats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Meats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MenuValidationError{
					field:  fmt.Sprintf("Meats[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetSoups() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Soups[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Soups[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MenuValidationError{
					field:  fmt.Sprintf("Soups[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetDrinks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Drinks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Drinks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MenuValidationError{
					field:  fmt.Sprintf("Drinks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetDesserts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Desserts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Desserts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MenuValidationError{
					field:  fmt.Sprintf("Desserts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MenuValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MenuValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MenuValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return MenuMultiError(errors)
	}

	return nil
}

// MenuMultiError is an error wrapping multiple validation errors returned by
// Menu.ValidateAll() if the designated constraints aren't met.
type MenuMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MenuMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MenuMultiError) AllErrors() []error { return m }

// MenuValidationError is the validation error returned by Menu.Validate if the
// designated constraints aren't met.
type MenuValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MenuValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MenuValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MenuValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MenuValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MenuValidationError) ErrorName() string { return "MenuValidationError" }

// Error satisfies the builtin error interface
func (e MenuValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMenu.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MenuValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MenuValidationError{}
