// Package domdebugger provides the Chrome Debugging Protocol
// commands, types, and events for the DOMDebugger domain.
//
// DOM debugging allows setting breakpoints on particular DOM operations and
// events. JavaScript execution will stop on these operations as if there was a
// regular breakpoint set.
//
// Generated by the chromedp-gen command.
package domdebugger

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/cdp/runtime"
)

// SetDOMBreakpointParams sets breakpoint on particular operation with DOM.
type SetDOMBreakpointParams struct {
	NodeID cdp.NodeID        `json:"nodeId"` // Identifier of the node to set breakpoint on.
	Type   DOMBreakpointType `json:"type"`   // Type of the operation to stop upon.
}

// SetDOMBreakpoint sets breakpoint on particular operation with DOM.
//
// parameters:
//   nodeID - Identifier of the node to set breakpoint on.
//   type - Type of the operation to stop upon.
func SetDOMBreakpoint(nodeID cdp.NodeID, typeVal DOMBreakpointType) *SetDOMBreakpointParams {
	return &SetDOMBreakpointParams{
		NodeID: nodeID,
		Type:   typeVal,
	}
}

// Do executes DOMDebugger.setDOMBreakpoint against the provided context and
// target handler.
func (p *SetDOMBreakpointParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDOMDebuggerSetDOMBreakpoint, p, nil)
}

// RemoveDOMBreakpointParams removes DOM breakpoint that was set using
// setDOMBreakpoint.
type RemoveDOMBreakpointParams struct {
	NodeID cdp.NodeID        `json:"nodeId"` // Identifier of the node to remove breakpoint from.
	Type   DOMBreakpointType `json:"type"`   // Type of the breakpoint to remove.
}

// RemoveDOMBreakpoint removes DOM breakpoint that was set using
// setDOMBreakpoint.
//
// parameters:
//   nodeID - Identifier of the node to remove breakpoint from.
//   type - Type of the breakpoint to remove.
func RemoveDOMBreakpoint(nodeID cdp.NodeID, typeVal DOMBreakpointType) *RemoveDOMBreakpointParams {
	return &RemoveDOMBreakpointParams{
		NodeID: nodeID,
		Type:   typeVal,
	}
}

// Do executes DOMDebugger.removeDOMBreakpoint against the provided context and
// target handler.
func (p *RemoveDOMBreakpointParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDOMDebuggerRemoveDOMBreakpoint, p, nil)
}

// SetEventListenerBreakpointParams sets breakpoint on particular DOM event.
type SetEventListenerBreakpointParams struct {
	EventName  string `json:"eventName"`            // DOM Event name to stop on (any DOM event will do).
	TargetName string `json:"targetName,omitempty"` // EventTarget interface name to stop on. If equal to "*" or not provided, will stop on any EventTarget.
}

// SetEventListenerBreakpoint sets breakpoint on particular DOM event.
//
// parameters:
//   eventName - DOM Event name to stop on (any DOM event will do).
func SetEventListenerBreakpoint(eventName string) *SetEventListenerBreakpointParams {
	return &SetEventListenerBreakpointParams{
		EventName: eventName,
	}
}

// WithTargetName eventTarget interface name to stop on. If equal to "*" or
// not provided, will stop on any EventTarget.
func (p SetEventListenerBreakpointParams) WithTargetName(targetName string) *SetEventListenerBreakpointParams {
	p.TargetName = targetName
	return &p
}

// Do executes DOMDebugger.setEventListenerBreakpoint against the provided context and
// target handler.
func (p *SetEventListenerBreakpointParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDOMDebuggerSetEventListenerBreakpoint, p, nil)
}

// RemoveEventListenerBreakpointParams removes breakpoint on particular DOM
// event.
type RemoveEventListenerBreakpointParams struct {
	EventName  string `json:"eventName"`            // Event name.
	TargetName string `json:"targetName,omitempty"` // EventTarget interface name.
}

// RemoveEventListenerBreakpoint removes breakpoint on particular DOM event.
//
// parameters:
//   eventName - Event name.
func RemoveEventListenerBreakpoint(eventName string) *RemoveEventListenerBreakpointParams {
	return &RemoveEventListenerBreakpointParams{
		EventName: eventName,
	}
}

// WithTargetName eventTarget interface name.
func (p RemoveEventListenerBreakpointParams) WithTargetName(targetName string) *RemoveEventListenerBreakpointParams {
	p.TargetName = targetName
	return &p
}

// Do executes DOMDebugger.removeEventListenerBreakpoint against the provided context and
// target handler.
func (p *RemoveEventListenerBreakpointParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDOMDebuggerRemoveEventListenerBreakpoint, p, nil)
}

// SetInstrumentationBreakpointParams sets breakpoint on particular native
// event.
type SetInstrumentationBreakpointParams struct {
	EventName string `json:"eventName"` // Instrumentation name to stop on.
}

// SetInstrumentationBreakpoint sets breakpoint on particular native event.
//
// parameters:
//   eventName - Instrumentation name to stop on.
func SetInstrumentationBreakpoint(eventName string) *SetInstrumentationBreakpointParams {
	return &SetInstrumentationBreakpointParams{
		EventName: eventName,
	}
}

// Do executes DOMDebugger.setInstrumentationBreakpoint against the provided context and
// target handler.
func (p *SetInstrumentationBreakpointParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDOMDebuggerSetInstrumentationBreakpoint, p, nil)
}

// RemoveInstrumentationBreakpointParams removes breakpoint on particular
// native event.
type RemoveInstrumentationBreakpointParams struct {
	EventName string `json:"eventName"` // Instrumentation name to stop on.
}

// RemoveInstrumentationBreakpoint removes breakpoint on particular native
// event.
//
// parameters:
//   eventName - Instrumentation name to stop on.
func RemoveInstrumentationBreakpoint(eventName string) *RemoveInstrumentationBreakpointParams {
	return &RemoveInstrumentationBreakpointParams{
		EventName: eventName,
	}
}

// Do executes DOMDebugger.removeInstrumentationBreakpoint against the provided context and
// target handler.
func (p *RemoveInstrumentationBreakpointParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDOMDebuggerRemoveInstrumentationBreakpoint, p, nil)
}

// SetXHRBreakpointParams sets breakpoint on XMLHttpRequest.
type SetXHRBreakpointParams struct {
	URL string `json:"url"` // Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
}

// SetXHRBreakpoint sets breakpoint on XMLHttpRequest.
//
// parameters:
//   url - Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
func SetXHRBreakpoint(url string) *SetXHRBreakpointParams {
	return &SetXHRBreakpointParams{
		URL: url,
	}
}

// Do executes DOMDebugger.setXHRBreakpoint against the provided context and
// target handler.
func (p *SetXHRBreakpointParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDOMDebuggerSetXHRBreakpoint, p, nil)
}

// RemoveXHRBreakpointParams removes breakpoint from XMLHttpRequest.
type RemoveXHRBreakpointParams struct {
	URL string `json:"url"` // Resource URL substring.
}

// RemoveXHRBreakpoint removes breakpoint from XMLHttpRequest.
//
// parameters:
//   url - Resource URL substring.
func RemoveXHRBreakpoint(url string) *RemoveXHRBreakpointParams {
	return &RemoveXHRBreakpointParams{
		URL: url,
	}
}

// Do executes DOMDebugger.removeXHRBreakpoint against the provided context and
// target handler.
func (p *RemoveXHRBreakpointParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDOMDebuggerRemoveXHRBreakpoint, p, nil)
}

// GetEventListenersParams returns event listeners of the given object.
type GetEventListenersParams struct {
	ObjectID runtime.RemoteObjectID `json:"objectId"`         // Identifier of the object to return listeners for.
	Depth    int64                  `json:"depth,omitempty"`  // The maximum depth at which Node children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Pierce   bool                   `json:"pierce,omitempty"` // Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false). Reports listeners for all contexts if pierce is enabled.
}

// GetEventListeners returns event listeners of the given object.
//
// parameters:
//   objectID - Identifier of the object to return listeners for.
func GetEventListeners(objectID runtime.RemoteObjectID) *GetEventListenersParams {
	return &GetEventListenersParams{
		ObjectID: objectID,
	}
}

// WithDepth the maximum depth at which Node children should be retrieved,
// defaults to 1. Use -1 for the entire subtree or provide an integer larger
// than 0.
func (p GetEventListenersParams) WithDepth(depth int64) *GetEventListenersParams {
	p.Depth = depth
	return &p
}

// WithPierce whether or not iframes and shadow roots should be traversed
// when returning the subtree (default is false). Reports listeners for all
// contexts if pierce is enabled.
func (p GetEventListenersParams) WithPierce(pierce bool) *GetEventListenersParams {
	p.Pierce = pierce
	return &p
}

// GetEventListenersReturns return values.
type GetEventListenersReturns struct {
	Listeners []*EventListener `json:"listeners,omitempty"` // Array of relevant listeners.
}

// Do executes DOMDebugger.getEventListeners against the provided context and
// target handler.
//
// returns:
//   listeners - Array of relevant listeners.
func (p *GetEventListenersParams) Do(ctxt context.Context, h cdp.Handler) (listeners []*EventListener, err error) {
	// execute
	var res GetEventListenersReturns
	err = h.Execute(ctxt, cdp.CommandDOMDebuggerGetEventListeners, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Listeners, nil
}
