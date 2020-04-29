// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v2/cluster/circuit_breaker.proto

package envoy_api_v2_cluster

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

	"github.com/gogo/protobuf/types"

	envoy_api_v2_core "github.com/datawire/ambassador/pkg/api/envoy/api/v2/core"
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
	_ = types.DynamicAny{}

	_ = envoy_api_v2_core.RoutingPriority(0)
)

// define the regex for a UUID once up-front
var _circuit_breaker_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CircuitBreakers with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CircuitBreakers) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetThresholds() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return CircuitBreakersValidationError{
						field:  fmt.Sprintf("Thresholds[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// CircuitBreakersValidationError is the validation error returned by
// CircuitBreakers.Validate if the designated constraints aren't met.
type CircuitBreakersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CircuitBreakersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CircuitBreakersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CircuitBreakersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CircuitBreakersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CircuitBreakersValidationError) ErrorName() string { return "CircuitBreakersValidationError" }

// Error satisfies the builtin error interface
func (e CircuitBreakersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitBreakers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CircuitBreakersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CircuitBreakersValidationError{}

// Validate checks the field values on CircuitBreakers_Thresholds with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CircuitBreakers_Thresholds) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := envoy_api_v2_core.RoutingPriority_name[int32(m.GetPriority())]; !ok {
		return CircuitBreakers_ThresholdsValidationError{
			field:  "Priority",
			reason: "value must be one of the defined enum values",
		}
	}

	{
		tmp := m.GetMaxConnections()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CircuitBreakers_ThresholdsValidationError{
					field:  "MaxConnections",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetMaxPendingRequests()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CircuitBreakers_ThresholdsValidationError{
					field:  "MaxPendingRequests",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetMaxRequests()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CircuitBreakers_ThresholdsValidationError{
					field:  "MaxRequests",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetMaxRetries()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CircuitBreakers_ThresholdsValidationError{
					field:  "MaxRetries",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetRetryBudget()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CircuitBreakers_ThresholdsValidationError{
					field:  "RetryBudget",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for TrackRemaining

	{
		tmp := m.GetMaxConnectionPools()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CircuitBreakers_ThresholdsValidationError{
					field:  "MaxConnectionPools",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// CircuitBreakers_ThresholdsValidationError is the validation error returned
// by CircuitBreakers_Thresholds.Validate if the designated constraints aren't met.
type CircuitBreakers_ThresholdsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CircuitBreakers_ThresholdsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CircuitBreakers_ThresholdsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CircuitBreakers_ThresholdsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CircuitBreakers_ThresholdsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CircuitBreakers_ThresholdsValidationError) ErrorName() string {
	return "CircuitBreakers_ThresholdsValidationError"
}

// Error satisfies the builtin error interface
func (e CircuitBreakers_ThresholdsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitBreakers_Thresholds.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CircuitBreakers_ThresholdsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CircuitBreakers_ThresholdsValidationError{}

// Validate checks the field values on CircuitBreakers_Thresholds_RetryBudget
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *CircuitBreakers_Thresholds_RetryBudget) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetBudgetPercent()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CircuitBreakers_Thresholds_RetryBudgetValidationError{
					field:  "BudgetPercent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetMinRetryConcurrency()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CircuitBreakers_Thresholds_RetryBudgetValidationError{
					field:  "MinRetryConcurrency",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// CircuitBreakers_Thresholds_RetryBudgetValidationError is the validation
// error returned by CircuitBreakers_Thresholds_RetryBudget.Validate if the
// designated constraints aren't met.
type CircuitBreakers_Thresholds_RetryBudgetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CircuitBreakers_Thresholds_RetryBudgetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CircuitBreakers_Thresholds_RetryBudgetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CircuitBreakers_Thresholds_RetryBudgetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CircuitBreakers_Thresholds_RetryBudgetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CircuitBreakers_Thresholds_RetryBudgetValidationError) ErrorName() string {
	return "CircuitBreakers_Thresholds_RetryBudgetValidationError"
}

// Error satisfies the builtin error interface
func (e CircuitBreakers_Thresholds_RetryBudgetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitBreakers_Thresholds_RetryBudget.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CircuitBreakers_Thresholds_RetryBudgetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CircuitBreakers_Thresholds_RetryBudgetValidationError{}
