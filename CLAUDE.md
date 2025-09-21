# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Testing
```bash
go test ./...                    # Run all tests
go test -v ./...                 # Run tests with verbose output
go test ./clients/datadog        # Run tests for specific package
```

### Building and Dependencies
```bash
go mod tidy                      # Clean up dependencies
go mod download                  # Download dependencies
go build ./...                   # Build all packages
```

### Mock Generation
```bash
make mocks                       # Generate all mocks using uber-go/mock
make generate                    # Alias for mocks target
```

The mock generation creates:
- `./mocks/metrics_impl.go` - Mock for MetricsImpl interface
- `./mocks/metrics_client.go` - Mock for Client interface
- `./mocks/stasd_client.go` - Mock for DataDog statsd ClientInterface

## Architecture Overview

This is a Go metrics library that provides a standardized interface for collecting application metrics with support for multiple backends (DataDog, Prometheus).

### Core Components

**Client Interface** (`client.go`): Defines the basic metrics operations (Histogram, Count, Gauge, Summary) that all metric backends must implement.

**MetricsImpl Interface** (`metrics_types.go`): High-level interface that wraps the Client interface and adds standard metric methods for common use cases.

**Standard Metrics** (`standard_metrics.go`): Defines 5 standardized metrics with specific label schemas:
- `resolver_request_duration_histogram_milliseconds` - GraphQL/API resolver performance
- `http_request_duration_histogram_milliseconds` - HTTP request performance
- `api_request_duration_histogram_milliseconds` - Service-to-service API calls
- `database_query_duration_histogram_milliseconds` - Database operation performance
- `call_duration_histogram_milliseconds` - Function call performance

Each standard metric has its own file (e.g., `resolver_metric.go`) with typed label structs.

**HTTP Middleware** (`middleware.go`): Provides automatic HTTP request duration tracking via middleware.

### Client Implementations

**DataDog Client** (`clients/datadog/`):
- Implements Client interface using DataDog's statsd library
- Supports custom histogram buckets via `CreateHistogram()`
- Auto-creates histograms with default buckets if not pre-created
- Maps Go labels to DataDog tags format

**Prometheus Client** (`clients/prometheus/`): Alternative implementation for Prometheus metrics.

### Usage Pattern

1. Create a client instance (DataDog or Prometheus)
2. Optionally pre-create histograms with custom buckets
3. Create a Metrics instance wrapping the client
4. Use either low-level methods (`HistogramMetric`) or high-level standard metrics (`ResolverMetric`, etc.)

The library enforces consistent metric naming and labeling across services while allowing flexibility in the underlying metrics backend.