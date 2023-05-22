// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: customer_order.proto

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

// Validate checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderRequestMultiError, or nil if none found.
func (m *CreateOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserUuid

	for idx, item := range m.GetSalads() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Salads[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Salads[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateOrderRequestValidationError{
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
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Garnishes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Garnishes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateOrderRequestValidationError{
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
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Meats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Meats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateOrderRequestValidationError{
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
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Soups[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Soups[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateOrderRequestValidationError{
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
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Drinks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Drinks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateOrderRequestValidationError{
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
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Desserts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Desserts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateOrderRequestValidationError{
					field:  fmt.Sprintf("Desserts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CreateOrderRequestMultiError(errors)
	}

	return nil
}

// CreateOrderRequestMultiError is an error wrapping multiple validation errors
// returned by CreateOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderRequestMultiError) AllErrors() []error { return m }

// CreateOrderRequestValidationError is the validation error returned by
// CreateOrderRequest.Validate if the designated constraints aren't met.
type CreateOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderRequestValidationError) ErrorName() string {
	return "CreateOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderRequestValidationError{}

// Validate checks the field values on CreateOrderResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderResponseMultiError, or nil if none found.
func (m *CreateOrderResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateOrderResponseMultiError(errors)
	}

	return nil
}

// CreateOrderResponseMultiError is an error wrapping multiple validation
// errors returned by CreateOrderResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateOrderResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderResponseMultiError) AllErrors() []error { return m }

// CreateOrderResponseValidationError is the validation error returned by
// CreateOrderResponse.Validate if the designated constraints aren't met.
type CreateOrderResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderResponseValidationError) ErrorName() string {
	return "CreateOrderResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderResponseValidationError{}

// Validate checks the field values on GetActualMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetActualMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetActualMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetActualMenuRequestMultiError, or nil if none found.
func (m *GetActualMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetActualMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetSalads() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetActualMenuRequestValidationError{
						field:  fmt.Sprintf("Salads[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetActualMenuRequestValidationError{
						field:  fmt.Sprintf("Salads[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetActualMenuRequestValidationError{
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
					errors = append(errors, GetActualMenuRequestValidationError{
						field:  fmt.Sprintf("Garnishes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetActualMenuRequestValidationError{
						field:  fmt.Sprintf("Garnishes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetActualMenuRequestValidationError{
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
					errors = append(errors, GetActualMenuRequestValidationError{
						field:  fmt.Sprintf("Meats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetActualMenuRequestValidationError{
						field:  fmt.Sprintf("Meats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetActualMenuRequestValidationError{
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
					errors = append(errors, GetActualMenuRequestValidationError{
						field:  fmt.Sprintf("Soups[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetActualMenuRequestValidationError{
						field:  fmt.Sprintf("Soups[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetActualMenuRequestValidationError{
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
					errors = append(errors, GetActualMenuRequestValidationError{
						field:  fmt.Sprintf("Drinks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetActualMenuRequestValidationError{
						field:  fmt.Sprintf("Drinks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetActualMenuRequestValidationError{
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
					errors = append(errors, GetActualMenuRequestValidationError{
						field:  fmt.Sprintf("Desserts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetActualMenuRequestValidationError{
						field:  fmt.Sprintf("Desserts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetActualMenuRequestValidationError{
					field:  fmt.Sprintf("Desserts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetActualMenuRequestMultiError(errors)
	}

	return nil
}

// GetActualMenuRequestMultiError is an error wrapping multiple validation
// errors returned by GetActualMenuRequest.ValidateAll() if the designated
// constraints aren't met.
type GetActualMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetActualMenuRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetActualMenuRequestMultiError) AllErrors() []error { return m }

// GetActualMenuRequestValidationError is the validation error returned by
// GetActualMenuRequest.Validate if the designated constraints aren't met.
type GetActualMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetActualMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetActualMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetActualMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetActualMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetActualMenuRequestValidationError) ErrorName() string {
	return "GetActualMenuRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetActualMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetActualMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetActualMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetActualMenuRequestValidationError{}

// Validate checks the field values on GetActualMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetActualMenuResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetActualMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetActualMenuResponseMultiError, or nil if none found.
func (m *GetActualMenuResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetActualMenuResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetActualMenuResponseMultiError(errors)
	}

	return nil
}

// GetActualMenuResponseMultiError is an error wrapping multiple validation
// errors returned by GetActualMenuResponse.ValidateAll() if the designated
// constraints aren't met.
type GetActualMenuResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetActualMenuResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetActualMenuResponseMultiError) AllErrors() []error { return m }

// GetActualMenuResponseValidationError is the validation error returned by
// GetActualMenuResponse.Validate if the designated constraints aren't met.
type GetActualMenuResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetActualMenuResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetActualMenuResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetActualMenuResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetActualMenuResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetActualMenuResponseValidationError) ErrorName() string {
	return "GetActualMenuResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetActualMenuResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetActualMenuResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetActualMenuResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetActualMenuResponseValidationError{}

// Validate checks the field values on OrderItem with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderItem with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderItemMultiError, or nil
// if none found.
func (m *OrderItem) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Count

	// no validation rules for ProductUuid

	if len(errors) > 0 {
		return OrderItemMultiError(errors)
	}

	return nil
}

// OrderItemMultiError is an error wrapping multiple validation errors returned
// by OrderItem.ValidateAll() if the designated constraints aren't met.
type OrderItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderItemMultiError) AllErrors() []error { return m }

// OrderItemValidationError is the validation error returned by
// OrderItem.Validate if the designated constraints aren't met.
type OrderItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderItemValidationError) ErrorName() string { return "OrderItemValidationError" }

// Error satisfies the builtin error interface
func (e OrderItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderItemValidationError{}

// Validate checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Product) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ProductMultiError, or nil if none found.
func (m *Product) ValidateAll() error {
	return m.validate(true)
}

func (m *Product) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uuid

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Type

	// no validation rules for Weight

	// no validation rules for Price

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProductValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProductValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProductMultiError(errors)
	}

	return nil
}

// ProductMultiError is an error wrapping multiple validation errors returned
// by Product.ValidateAll() if the designated constraints aren't met.
type ProductMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductMultiError) AllErrors() []error { return m }

// ProductValidationError is the validation error returned by Product.Validate
// if the designated constraints aren't met.
type ProductValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductValidationError) ErrorName() string { return "ProductValidationError" }

// Error satisfies the builtin error interface
func (e ProductValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProduct.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductValidationError{}
