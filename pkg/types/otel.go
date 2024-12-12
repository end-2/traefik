package types

import (
	"fmt"
	"strconv"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// OTelGRPC provides configuration settings for the gRPC open-telemetry.
type OTelGRPC struct {
	Endpoint string            `description:"Sets the gRPC endpoint (host:port) of the collector." json:"endpoint,omitempty" toml:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Insecure bool              `description:"Disables client transport security for the exporter." json:"insecure,omitempty" toml:"insecure,omitempty" yaml:"insecure,omitempty" export:"true"`
	TLS      *ClientTLS        `description:"Defines client transport security parameters." json:"tls,omitempty" toml:"tls,omitempty" yaml:"tls,omitempty" export:"true"`
	Headers  map[string]string `description:"Headers sent with payload." json:"headers,omitempty" toml:"headers,omitempty" yaml:"headers,omitempty"`
}

// SetDefaults sets the default values.
func (o *OTelGRPC) SetDefaults() {
	o.Endpoint = "localhost:4317"
}

// OTelHTTP provides configuration settings for the HTTP open-telemetry.
type OTelHTTP struct {
	Endpoint string            `description:"Sets the HTTP endpoint (scheme://host:port/path) of the collector." json:"endpoint,omitempty" toml:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	TLS      *ClientTLS        `description:"Defines client transport security parameters." json:"tls,omitempty" toml:"tls,omitempty" yaml:"tls,omitempty" export:"true"`
	Headers  map[string]string `description:"Headers sent with payload." json:"headers,omitempty" toml:"headers,omitempty" yaml:"headers,omitempty"`
}

// SetDefaults sets the default values.
func (o *OTelHTTP) SetDefaults() {
	o.Endpoint = "https://localhost:4318"
}

// https://opentelemetry.io/docs/languages/sdk-configuration/general/#otel_traces_sampler
const (
	// otelSamplerAlwaysOnType is the always on sampler type.
	otelSamplerAlwaysOnType = "always_on"
	// otelSamplerAlwaysOffType is the always off sampler type.
	otelSamplerAlwaysOffType = "always_off"
	// otelSamplerTraceIDRatioType is the trace id ratio sampler type.
	otelSamplerTraceIDRatioType = "traceidratio"
	// otelSamplerParentBasedAlwaysOnType is the parent based always on sampler type.
	otelSamplerParentBasedAlwaysOnType = "parentbased_always_on"
	// otelSamplerParentBasedAlwaysOffType is the parent based always off sampler type.
	otelSamplerParentBasedAlwaysOffType = "parentbased_always_off"
	// otelSamplerParentBasedTraceIDRatioType is the parent based trace id ratio sampler type.
	otelSamplerParentBasedTraceIDRatioType = "parentbased_traceidratio"
)

// OtelSampler provides configuration settings for the open-telemetry sampler.
type OTelSampler struct {
	Type      string `description:"The sampler type to use." json:"type,omitempty" toml:"type,omitempty" yaml:"type,omitempty" export:"true"`
	Arguments string `description:"The sampler arguments." json:"arguments,omitempty" toml:"arguments,omitempty" yaml:"arguments,omitempty" export:"true"`
}

// SetDefaults sets the default(otelSamplerTraceIDRatioType) values.
func (o *OTelSampler) SetDefaults() {
	o.Type = otelSamplerTraceIDRatioType
}

// setupSampler sets up the sampler.
func (o *OTelSampler) setupSampler() (sdktrace.Sampler, error) {
	switch o.Type {
	case otelSamplerAlwaysOnType:
		return sdktrace.AlwaysSample(), nil
	case otelSamplerAlwaysOffType:
		return sdktrace.NeverSample(), nil
	case otelSamplerTraceIDRatioType:
		ratio, err := strconv.ParseFloat(o.Arguments, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse trace id ratio: %w", err)
		}
		return sdktrace.TraceIDRatioBased(ratio), nil
	case otelSamplerParentBasedAlwaysOnType:
		return sdktrace.ParentBased(sdktrace.AlwaysSample()), nil
	case otelSamplerParentBasedAlwaysOffType:
		return sdktrace.ParentBased(sdktrace.NeverSample()), nil
	case otelSamplerParentBasedTraceIDRatioType:
		ratio, err := strconv.ParseFloat(o.Arguments, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse trace id ratio: %w", err)
		}
		return sdktrace.ParentBased(sdktrace.TraceIDRatioBased(ratio)), nil
	default:
		return nil, fmt.Errorf("unsupported sampler type: %s", o.Type)
	}
}
