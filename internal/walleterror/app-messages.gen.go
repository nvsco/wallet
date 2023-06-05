package walleterror

import (
	"fmt"
	"strings"
)

/*
	Do not edit this file. This file was auto generated
	All the changes will be overwritten by the "make gen-app-messages" command
*/

func GetFileContent() []byte {
	return []byte(`{
	"messages": {
		"InternalAuthError": {
			"text": "internal auth error: {{err}}"
		},
		"InternalServiceError": {
			"text": "internal service error: {{err}}"
		},
		"UnableParseAmount": {
			"text": "unable to parse the field {{field}} with value {{value}} as amount"
		},
		"UnexpectedForeignType": {
			"text": "Unexpected foreign type {{foreignType}}"
		}
	}
}`)
}

const (
	InternalAuthError     = "InternalAuthError"
	InternalServiceError  = "InternalServiceError"
	UnableParseAmount     = "UnableParseAmount"
	UnexpectedForeignType = "UnexpectedForeignType"
)

type AppMessage struct {
	Code       string     `json:"code"`
	Message    string     `json:"message"`
	Attributes Attributes `json:"attributes"`
	ErrorStr   string     `json:"error"`
}

func (m AppMessage) String() string {
	text := m.Message

	for key, value := range m.Attributes {
		text = strings.ReplaceAll(text, fmt.Sprintf("{{%s}}", key), value)
	}

	return text
}

func (m AppMessage) Error() string {
	return m.String()
}

func (m AppMessage) Export() (code string, message string, attributes map[string]string) {
	return m.Code, m.Message, m.Attributes
}

type Attributes map[string]string

func toErrStr(errs []error) string {
	if len(errs) > 0 {
		err := errs[0]
		if err != nil {
			return errs[0].Error()
		}
	}
	return ""
}

// NewInternalAuthError generated from the code "InternalAuthError"
func NewInternalAuthError(err error, errs ...error) *AppMessage {
	if len(errs) > 0 && errs[0] == nil {
		return nil
	}

	return &AppMessage{
		Code:    InternalAuthError,
		Message: "internal auth error: {{err}}",
		Attributes: Attributes{
			"err": err.Error(),
		},
		ErrorStr: toErrStr(errs),
	}
}

// NewInternalServiceError generated from the code "InternalServiceError"
func NewInternalServiceError(err error, errs ...error) *AppMessage {
	if len(errs) > 0 && errs[0] == nil {
		return nil
	}

	return &AppMessage{
		Code:    InternalServiceError,
		Message: "internal service error: {{err}}",
		Attributes: Attributes{
			"err": err.Error(),
		},
		ErrorStr: toErrStr(errs),
	}
}

// NewUnableParseAmount generated from the code "UnableParseAmount"
func NewUnableParseAmount(field string, value string, errs ...error) *AppMessage {
	if len(errs) > 0 && errs[0] == nil {
		return nil
	}

	return &AppMessage{
		Code:    UnableParseAmount,
		Message: "unable to parse the field {{field}} with value {{value}} as amount",
		Attributes: Attributes{
			"field": field,
			"value": value,
		},
		ErrorStr: toErrStr(errs),
	}
}

// NewUnexpectedForeignType generated from the code "UnexpectedForeignType"
func NewUnexpectedForeignType(foreignType string, errs ...error) *AppMessage {
	if len(errs) > 0 && errs[0] == nil {
		return nil
	}

	return &AppMessage{
		Code:    UnexpectedForeignType,
		Message: "Unexpected foreign type {{foreignType}}",
		Attributes: Attributes{
			"foreignType": foreignType,
		},
		ErrorStr: toErrStr(errs),
	}
}

type AppMessageDesign struct {
	Code       string
	Message    string
	Attributes []string
}

func GetAppMessages() []AppMessageDesign {
	return []AppMessageDesign{
		{
			Code:       "InternalAuthError",
			Message:    "internal auth error: {{err}}",
			Attributes: []string{"err"},
		},
		{
			Code:       "InternalServiceError",
			Message:    "internal service error: {{err}}",
			Attributes: []string{"err"},
		},
		{
			Code:       "UnableParseAmount",
			Message:    "unable to parse the field {{field}} with value {{value}} as amount",
			Attributes: []string{"field", "value"},
		},
		{
			Code:       "UnexpectedForeignType",
			Message:    "Unexpected foreign type {{foreignType}}",
			Attributes: []string{"foreignType"},
		},
	}
}