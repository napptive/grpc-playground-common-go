// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: playground-common/entities.proto

package grpc_playground_common_go

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on EmptyRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EmptyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmptyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EmptyRequestMultiError, or
// nil if none found.
func (m *EmptyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *EmptyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyRequestMultiError(errors)
	}
	return nil
}

// EmptyRequestMultiError is an error wrapping multiple validation errors
// returned by EmptyRequest.ValidateAll() if the designated constraints aren't met.
type EmptyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyRequestMultiError) AllErrors() []error { return m }

// EmptyRequestValidationError is the validation error returned by
// EmptyRequest.Validate if the designated constraints aren't met.
type EmptyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyRequestValidationError) ErrorName() string { return "EmptyRequestValidationError" }

// Error satisfies the builtin error interface
func (e EmptyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmptyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyRequestValidationError{}

// Validate checks the field values on OpResponse with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OpResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OpResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OpResponseMultiError, or
// nil if none found.
func (m *OpResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OpResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for StatusName

	// no validation rules for UserInfo

	if len(errors) > 0 {
		return OpResponseMultiError(errors)
	}
	return nil
}

// OpResponseMultiError is an error wrapping multiple validation errors
// returned by OpResponse.ValidateAll() if the designated constraints aren't met.
type OpResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OpResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OpResponseMultiError) AllErrors() []error { return m }

// OpResponseValidationError is the validation error returned by
// OpResponse.Validate if the designated constraints aren't met.
type OpResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpResponseValidationError) ErrorName() string { return "OpResponseValidationError" }

// Error satisfies the builtin error interface
func (e OpResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpResponseValidationError{}

// Validate checks the field values on TokenResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TokenResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TokenResponseMultiError, or
// nil if none found.
func (m *TokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *TokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for RefreshToken

	// no validation rules for AccountName

	// no validation rules for EnvironmentName

	// no validation rules for PlatformUsername

	if len(errors) > 0 {
		return TokenResponseMultiError(errors)
	}
	return nil
}

// TokenResponseMultiError is an error wrapping multiple validation errors
// returned by TokenResponse.ValidateAll() if the designated constraints
// aren't met.
type TokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TokenResponseMultiError) AllErrors() []error { return m }

// TokenResponseValidationError is the validation error returned by
// TokenResponse.Validate if the designated constraints aren't met.
type TokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokenResponseValidationError) ErrorName() string { return "TokenResponseValidationError" }

// Error satisfies the builtin error interface
func (e TokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokenResponseValidationError{}

// Validate checks the field values on GenericListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GenericListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GenericListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GenericListRequestMultiError, or nil if none found.
func (m *GenericListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GenericListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IncludeYamlConversion

	// no validation rules for IncludeJsonConversion

	if len(errors) > 0 {
		return GenericListRequestMultiError(errors)
	}
	return nil
}

// GenericListRequestMultiError is an error wrapping multiple validation errors
// returned by GenericListRequest.ValidateAll() if the designated constraints
// aren't met.
type GenericListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GenericListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GenericListRequestMultiError) AllErrors() []error { return m }

// GenericListRequestValidationError is the validation error returned by
// GenericListRequest.Validate if the designated constraints aren't met.
type GenericListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenericListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenericListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenericListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenericListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenericListRequestValidationError) ErrorName() string {
	return "GenericListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GenericListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenericListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenericListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenericListRequestValidationError{}

// Validate checks the field values on RemoveGenericRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveGenericRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveGenericRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveGenericRequestMultiError, or nil if none found.
func (m *RemoveGenericRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveGenericRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := RemoveGenericRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RemoveGenericRequestMultiError(errors)
	}
	return nil
}

// RemoveGenericRequestMultiError is an error wrapping multiple validation
// errors returned by RemoveGenericRequest.ValidateAll() if the designated
// constraints aren't met.
type RemoveGenericRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveGenericRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveGenericRequestMultiError) AllErrors() []error { return m }

// RemoveGenericRequestValidationError is the validation error returned by
// RemoveGenericRequest.Validate if the designated constraints aren't met.
type RemoveGenericRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveGenericRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveGenericRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveGenericRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveGenericRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveGenericRequestValidationError) ErrorName() string {
	return "RemoveGenericRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveGenericRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveGenericRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveGenericRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveGenericRequestValidationError{}

// Validate checks the field values on GenericGetRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GenericGetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GenericGetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GenericGetRequestMultiError, or nil if none found.
func (m *GenericGetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GenericGetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for EnvironmentQualifiedName

	// no validation rules for AccountId

	// no validation rules for EnvironmentId

	// no validation rules for Name

	// no validation rules for IncludeYamlConversion

	// no validation rules for IncludeJsonConversion

	if len(errors) > 0 {
		return GenericGetRequestMultiError(errors)
	}
	return nil
}

// GenericGetRequestMultiError is an error wrapping multiple validation errors
// returned by GenericGetRequest.ValidateAll() if the designated constraints
// aren't met.
type GenericGetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GenericGetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GenericGetRequestMultiError) AllErrors() []error { return m }

// GenericGetRequestValidationError is the validation error returned by
// GenericGetRequest.Validate if the designated constraints aren't met.
type GenericGetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenericGetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenericGetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenericGetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenericGetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenericGetRequestValidationError) ErrorName() string {
	return "GenericGetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GenericGetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenericGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenericGetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenericGetRequestValidationError{}
