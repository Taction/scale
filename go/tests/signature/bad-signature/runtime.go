//go:build !tinygo && !js && !wasm
// +build !tinygo,!js,!wasm

// Code generated by scale-signature-go v0.0.1, DO NOT EDIT.
// source: signature/Bad/signature.proto

package bad

import (
	_ "embed"
	"github.com/loopholelabs/polyglot-go"
	"github.com/loopholelabs/scale-signature"
)

var _ signature.Signature = (*Context)(nil)
var _ signature.RuntimeContext = (*RuntimeContext)(nil)

type RuntimeContext Context

// Context is a context object for an incoming request. It is meant to be used
// inside the Scale function.
type Context struct {
	*BadContext
	buffer *polyglot.Buffer
}

// New creates a new empty Context
func New() *Context {
	return &Context{
		BadContext: NewBadContext(),
		buffer:     polyglot.NewBuffer(),
	}
}

// RuntimeContext converts a Context into a RuntimeContext.
func (x *Context) RuntimeContext() signature.RuntimeContext {
	return (*RuntimeContext)(x)
}

// Read reads the context from the given byte slice and returns an error if one occurred
//
// This method is meant to be used by the Scale Runtime to deserialize the Context
func (x *RuntimeContext) Read(b []byte) error {
	return x.internalDecode(b)
}

// Write writes the context into a byte slice and returns it
func (x *RuntimeContext) Write() []byte {
	x.buffer.Reset()
	x.internalEncode(x.buffer)
	return x.buffer.Bytes()
}

// Error writes the context into a byte slice and returns it
func (x *RuntimeContext) Error(err error) []byte {
	x.buffer.Reset()
	x.error(x.buffer, err)
	return x.buffer.Bytes()
}
