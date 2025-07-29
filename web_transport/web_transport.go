//go:build js

package web_transport

import (
	"syscall/js"

	"github.com/sourcenetwork/goji"
	"github.com/sourcenetwork/goji/streams"
)

func init() {
	WebTransport = webTransportJS(js.Global().Get("WebTransport"))
}

type webTransportJS js.Value

// WebTransport is a wrapper for the WebTransport API.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport
var WebTransport webTransportJS

// WebTransportValue is an instance of a WebTransport.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport
type WebTransportValue js.Value

// New returns a new WebTransportValue.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/WebTransport
func (w webTransportJS) New(url string, opts ...webTransportOption) WebTransportValue {
	switch {
	case len(opts) > 0:
		options := js.ValueOf(map[string]any{})
		for _, opt := range opts {
			opt(options)
		}
		res := js.Value(w).New(url, options)
		return WebTransportValue(res)

	default:
		res := js.Value(w).New(url)
		return WebTransportValue(res)
	}
}

// Closed returns the WebTransport.closed property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/closed
func (w WebTransportValue) Closed() goji.PromiseValue {
	return goji.PromiseValue(js.Value(w).Get("closed"))
}

// Datagrams returns the WebTransport.datagrams property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/datagrams
func (w WebTransportValue) Datagrams() WebTransportDatagramDuplexStreamValue {
	return WebTransportDatagramDuplexStreamValue(js.Value(w).Get("datagrams"))
}

// Ready returns the WebTransport.ready property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/ready
func (w WebTransportValue) Ready() goji.PromiseValue {
	return goji.PromiseValue(js.Value(w).Get("ready"))
}

// IncomingBidirectionalStreams returns the WebTransport.incomingBidirectionalStreams property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/incomingBidirectionalStreams
func (w WebTransportValue) IncomingBidirectionalStreams() streams.ReadableStreamValue {
	return streams.ReadableStreamValue(js.Value(w).Get("incomingBidirectionalStreams"))
}

// IncomingUnidirectionalStreams returns the WebTransport.incomingUnidirectionalStreams property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/incomingUnidirectionalStreams
func (w WebTransportValue) IncomingUnidirectionalStreams() streams.ReadableStreamValue {
	return streams.ReadableStreamValue(js.Value(w).Get("incomingUnidirectionalStreams"))
}

// WebTransportCloseInfo provides additional info when closing a web transport.
type WebTransportCloseInfo struct {
	// CloseCode is a number representing the error code for the error.
	CloseCode int
	// Reason is a string representing the reason for closing the WebTransport.
	Reason string
}

// Close wraps the WebTransport.close method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/close
func (w WebTransportValue) Close(info *WebTransportCloseInfo) {
	if info != nil {
		js.Value(w).Call("close", map[string]any{
			"closeCode": info.CloseCode,
			"reason":    info.Reason,
		})
	} else {
		js.Value(w).Call("close")
	}
}

// WebTransportCreateStreamOptions is used to set stream creation options.
var WebTransportCreateStreamOptions = &webTransportCreateStreamOptions{}

type webTransportCreateStreamOptions struct{}

type webTransportCreateStreamOption func(opts js.Value)

// WithSendOrder sets the sendOrder option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/createBidirectionalStream#sendorder
func (e webTransportCreateStreamOptions) WithSendOrder(value int) webTransportCreateStreamOption {
	return func(opts js.Value) {
		opts.Set("sendOrder", value)
	}
}

// CreateBidirectionalStream wraps the WebTransport.createBidirectionalStream method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/createBidirectionalStream
func (w WebTransportValue) CreateBidirectionalStream(opts ...webTransportOption) goji.PromiseValue {
	switch {
	case len(opts) > 0:
		options := js.ValueOf(map[string]any{})
		for _, opt := range opts {
			opt(options)
		}
		res := js.Value(w).Call("createBidirectionalStream", options)
		return goji.PromiseValue(res)

	default:
		res := js.Value(w).Call("createBidirectionalStream")
		return goji.PromiseValue(res)
	}
}

// CreateUnidirectionalStream wraps the WebTransport.createUnidirectionalStream method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/createUnidirectionalStream
func (w WebTransportValue) CreateUnidirectionalStream(opts ...webTransportOption) goji.PromiseValue {
	switch {
	case len(opts) > 0:
		options := js.ValueOf(map[string]any{})
		for _, opt := range opts {
			opt(options)
		}
		res := js.Value(w).Call("createUnidirectionalStream", options)
		return goji.PromiseValue(res)

	default:
		res := js.Value(w).Call("createUnidirectionalStream")
		return goji.PromiseValue(res)
	}
}

// CongestionControl specifies the available congestion control algorithms.
type CongestionControl string

var (
	// CongestionControlDefault is the default congestion control tuning for the transport.
	CongestionControlDefault = CongestionControl("default")
	// CongestionControlThroughput prefers congestion control to be tuned for throughput.
	CongestionControlThroughput = CongestionControl("throughput")
	// CongestionControlLowLatency prefers congestion control to be tuned for low-latency.
	CongestionControlLowLatency = CongestionControl("low-latency")
)

// WebTransportOptions is used to set web transport options.
var WebTransportOptions = &webTransportOptions{}

type webTransportOptions struct{}

type webTransportOption func(opts js.Value)

// WithAllowPooling sets the allow pooling option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/WebTransport#allowpooling
func (e webTransportOptions) WithAllowPooling(enable bool) webTransportOption {
	return func(opts js.Value) {
		opts.Set("allowPooling", enable)
	}
}

// WithCongestionControl sets the congestion control option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/WebTransport#congestioncontrol
func (e webTransportOptions) WithCongestionControl(control CongestionControl) webTransportOption {
	return func(opts js.Value) {
		opts.Set("congestionControl", string(control))
	}
}

// WithRequireUnreliable sets the require unreliable option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/WebTransport#requireunreliable
func (e webTransportOptions) WithRequireUnreliable(enable bool) webTransportOption {
	return func(opts js.Value) {
		opts.Set("requireUnreliable", enable)
	}
}

// WithServerCertificateHashes sets the server certificate hashes option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/WebTransport#servercertificatehashes
func (e webTransportOptions) WithServerCertificateHashes(hashes ...CertificateHashValue) webTransportOption {
	return func(opts js.Value) {
		val := make([]any, len(hashes))
		for i, h := range hashes {
			val[i] = any(map[string]any{
				"algorithm": "sha-256",
				"value":     js.Value(h).Get("value"),
			})
		}
		opts.Set("serverCertificateHashes", val)
	}
}

// WebTransportBidirectionalStreamValue is an instance of WebTransportBidirectionalStream.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportBidirectionalStream
type WebTransportBidirectionalStreamValue js.Value

// Readable returns the WebTransportBidirectionalStream.readable property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportBidirectionalStream/readable
func (v WebTransportBidirectionalStreamValue) Readable() streams.ReadableStreamValue {
	return streams.ReadableStreamValue(js.Value(v).Get("readable"))
}

// Writable returns the WebTransportBidirectionalStream.writable property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportBidirectionalStream/writable
func (v WebTransportBidirectionalStreamValue) Writable() streams.WritableStreamValue {
	return streams.WritableStreamValue(js.Value(v).Get("writable"))
}

// WebTransportDatagramDuplexStreamValue is an instance of WebTransportDatagramDuplexStream.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportDatagramDuplexStream
type WebTransportDatagramDuplexStreamValue js.Value

// MaxDatagramSize returns the WebTransportDatagramDuplexStream.maxDatagramSize property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportDatagramDuplexStream/maxDatagramSize
func (v WebTransportDatagramDuplexStreamValue) MaxDatagramSize() int {
	return js.Value(v).Get("maxDatagramSize").Int()
}

// IncomingHighWaterMark returns the WebTransportDatagramDuplexStream.incomingHighWaterMark property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportDatagramDuplexStream/incomingHighWaterMark
func (v WebTransportDatagramDuplexStreamValue) IncomingHighWaterMark() int {
	return js.Value(v).Get("incomingHighWaterMark").Int()
}

// SetIncomingHighWaterMark sets the WebTransportDatagramDuplexStream.incomingHighWaterMark property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportDatagramDuplexStream/incomingHighWaterMark
func (v WebTransportDatagramDuplexStreamValue) SetIncomingHighWaterMark(value int) {
	js.Value(v).Set("incomingHighWaterMark", value)
}

// OutgoingHighWaterMark returns the WebTransportDatagramDuplexStream.outgoingHighWaterMark property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportDatagramDuplexStream/outgoingHighWaterMark
func (v WebTransportDatagramDuplexStreamValue) OutgoingHighWaterMark() int {
	return js.Value(v).Get("outgoingHighWaterMark").Int()
}

// SetOutgoingHighWaterMark sets the WebTransportDatagramDuplexStream.outgoingHighWaterMark property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportDatagramDuplexStream/outgoingHighWaterMark
func (v WebTransportDatagramDuplexStreamValue) SetOutgoingHighWaterMark(value int) {
	js.Value(v).Set("outgoingHighWaterMark", value)
}

// IncomingMaxAge returns the WebTransportDatagramDuplexStream.incomingMaxAge property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportDatagramDuplexStream/incomingMaxAge
func (v WebTransportDatagramDuplexStreamValue) IncomingMaxAge() *int {
	res := js.Value(v).Get("incomingMaxAge")
	if res.IsNull() {
		return nil
	}
	val := res.Int()
	return &val
}

// SetIncomingMaxAge sets the WebTransportDatagramDuplexStream.incomingMaxAge property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportDatagramDuplexStream/incomingMaxAge
func (v WebTransportDatagramDuplexStreamValue) SetIncomingMaxAge(value *int) {
	if value == nil {
		js.Value(v).Set("incomingMaxAge", js.Null())
	} else {
		js.Value(v).Set("incomingMaxAge", *value)
	}
}

// OutgoingMaxAge returns the WebTransportDatagramDuplexStream.outgoingMaxAge property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportDatagramDuplexStream/outgoingMaxAge
func (v WebTransportDatagramDuplexStreamValue) OutgoingMaxAge() *int {
	res := js.Value(v).Get("outgoingMaxAge")
	if res.IsNull() {
		return nil
	}
	val := res.Int()
	return &val
}

// SetOutgoingMaxAge sets the WebTransportDatagramDuplexStream.outgoingMaxAge property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportDatagramDuplexStream/outgoingMaxAge
func (v WebTransportDatagramDuplexStreamValue) SetOutgoingMaxAge(value *int) {
	if value == nil {
		js.Value(v).Set("outgoingMaxAge", js.Null())
	} else {
		js.Value(v).Set("outgoingMaxAge", *value)
	}
}

// Readable returns the WebTransportDatagramDuplexStream.readable property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebTransportDatagramDuplexStream/readable
func (v WebTransportDatagramDuplexStreamValue) Readable() streams.ReadableStreamValue {
	res := js.Value(v).Get("readable")
	return streams.ReadableStreamValue(res)
}

// CertificateHashAlgorithm specifies the available certificate hash algorithms
type CertificateHashAlgorithm string

// CertificateHashAlgorithmSHA256 uses the sha-256 hash algorithm
var CertificateHashAlgorithmSHA256 = CertificateHashAlgorithm("sha-256")

// CertificateHashValue is used to configure server certificate hashes.
type CertificateHashValue js.Value

// CertificateHash returns a new certificate hash object.
func CertificateHash(algorithm CertificateHashAlgorithm, hash []byte) CertificateHashValue {
	arr := goji.Uint8ArrayFromBytes(hash)
	res := js.ValueOf(map[string]any{
		"algorithm": string(algorithm),
		"value":     js.Value(arr),
	})
	return CertificateHashValue(res)
}
