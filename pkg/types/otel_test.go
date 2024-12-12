package types

import (
	"github.com/stretchr/testify/assert"
	"testing"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func TestSetupSamplerAlwaysOn(t *testing.T) {
	sampler := OTelSampler{Type: otelSamplerAlwaysOnType}
	result, err := sampler.setupSampler()
	assert.NoError(t, err)
	assert.Equal(t, sdktrace.AlwaysSample(), result)
}

func TestSetupSamplerAlwaysOff(t *testing.T) {
	sampler := OTelSampler{Type: otelSamplerAlwaysOffType}
	result, err := sampler.setupSampler()
	assert.NoError(t, err)
	assert.Equal(t, sdktrace.NeverSample(), result)
}

func TestSetupSamplerTraceIDRatio(t *testing.T) {
	sampler := OTelSampler{Type: otelSamplerTraceIDRatioType, Arguments: "0.5"}
	result, err := sampler.setupSampler()
	assert.NoError(t, err)
	assert.IsType(t, sdktrace.TraceIDRatioBased(0.5), result)
}

func TestSetupSamplerTraceIDRatioInvalidArgument(t *testing.T) {
	sampler := OTelSampler{Type: otelSamplerTraceIDRatioType, Arguments: "invalid"}
	_, err := sampler.setupSampler()
	assert.Error(t, err)
}

func TestSetupSamplerParentBasedAlwaysOn(t *testing.T) {
	sampler := OTelSampler{Type: otelSamplerParentBasedAlwaysOnType}
	result, err := sampler.setupSampler()
	assert.NoError(t, err)
	assert.IsType(t, sdktrace.ParentBased(sdktrace.AlwaysSample()), result)
}

func TestSetupSamplerParentBasedAlwaysOff(t *testing.T) {
	sampler := OTelSampler{Type: otelSamplerParentBasedAlwaysOffType}
	result, err := sampler.setupSampler()
	assert.NoError(t, err)
	assert.IsType(t, sdktrace.ParentBased(sdktrace.NeverSample()), result)
}

func TestSetupSamplerParentBasedTraceIDRatio(t *testing.T) {
	sampler := OTelSampler{Type: otelSamplerParentBasedTraceIDRatioType, Arguments: "0.5"}
	result, err := sampler.setupSampler()
	assert.NoError(t, err)
	assert.IsType(t, sdktrace.ParentBased(sdktrace.TraceIDRatioBased(0.5)), result)
}

func TestSetupSamplerParentBasedTraceIDRatioInvalidArgument(t *testing.T) {
	sampler := OTelSampler{Type: otelSamplerParentBasedTraceIDRatioType, Arguments: "invalid"}
	_, err := sampler.setupSampler()
	assert.Error(t, err)
}

func TestSetupSamplerUnsupportedType(t *testing.T) {
	sampler := OTelSampler{Type: "unsupported"}
	_, err := sampler.setupSampler()
	assert.Error(t, err)
}
