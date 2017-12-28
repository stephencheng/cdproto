package runtime

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"github.com/mailru/easyjson"
)

// EventConsoleAPICalled issued when console API was called.
type EventConsoleAPICalled struct {
	Type               APIType            `json:"type"`                 // Type of the call.
	Args               []*RemoteObject    `json:"args"`                 // Call arguments.
	ExecutionContextID ExecutionContextID `json:"executionContextId"`   // Identifier of the context where the call was made.
	Timestamp          *Timestamp         `json:"timestamp"`            // Call timestamp.
	StackTrace         *StackTrace        `json:"stackTrace,omitempty"` // Stack trace captured when the call was made.
	Context            string             `json:"context,omitempty"`    // Console context descriptor for calls on non-default console context (not console.*): 'anonymous#unique-logger-id' for call on unnamed context, 'name#unique-logger-id' for call on named context.
}

// EventExceptionRevoked issued when unhandled exception was revoked.
type EventExceptionRevoked struct {
	Reason      string `json:"reason"`      // Reason describing why exception was revoked.
	ExceptionID int64  `json:"exceptionId"` // The id of revoked exception, as reported in exceptionThrown.
}

// EventExceptionThrown issued when exception was thrown and unhandled.
type EventExceptionThrown struct {
	Timestamp        *Timestamp        `json:"timestamp"` // Timestamp of the exception.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails"`
}

// EventExecutionContextCreated issued when new execution context is created.
type EventExecutionContextCreated struct {
	Context *ExecutionContextDescription `json:"context"` // A newly created execution context.
}

// EventExecutionContextDestroyed issued when execution context is destroyed.
type EventExecutionContextDestroyed struct {
	ExecutionContextID ExecutionContextID `json:"executionContextId"` // Id of the destroyed context
}

// EventExecutionContextsCleared issued when all executionContexts were
// cleared in browser.
type EventExecutionContextsCleared struct{}

// EventInspectRequested issued when object should be inspected (for example,
// as a result of inspect() command line API call).
type EventInspectRequested struct {
	Object *RemoteObject       `json:"object"`
	Hints  easyjson.RawMessage `json:"hints"`
}
