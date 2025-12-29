# 🚰 Sink

> **Lightweight observability & logging for edge containers and infrastructure**  
> Configurable for any development workflow. Built in Go.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

---

## 🎯 What is Sink?

**Sink** is a fast, edge-optimized logging library and agent designed for modern distributed systems. It handles observability where traditional solutions fall short: edge nodes, intermittent connectivity, resource-constrained environments, and high-throughput services.

### Built For

- **Edge Computing** - ARM devices, IoT gateways, CDN nodes
- **Cloud Infrastructure** - Kubernetes, Docker, serverless
- **CI/CD Pipelines** - Build agents, test runners
- **Microservices** - High-performance logging with minimal overhead

---

## ✨ Key Features

- ⚡ **Blazing Fast** - Async logging with ring buffers, minimal allocations
- 🧱 **Dual Mode** - Use as a Go library or standalone CLI agent
- 🌍 **Edge-Friendly** - Local buffering, works offline, low memory footprint
- 📦 **Zero Heavy Dependencies** - Pure Go, stdlib only
- 🔌 **Pluggable Sinks** - Console, file, HTTP, or build your own
- 🎨 **Structured Logging** - Type-safe fields, JSON output
- 🛡️ **Production Ready** - Thread-safe, graceful shutdown, log rotation

---

## 🚀 Quick Start

### As a Library

```go
package main

import (
    "github.com/stawuah/sink/v2/pkg/logger"
    "github.com/stawuah/sink/v2/pkg/sinks"
)

func main() {
    log := logger.New(
        logger.Config{
            Service:  "my-service",
            Env:      "production",
            MinLevel: logger.InfoLevel,
        },
        sinks.NewConsole(true),
    )
    defer log.Shutdown()

    log.Info("service started", logger.Fields{
        "version": "1.0.0",
        "port":    8080,
    })

    log.Error("connection failed", logger.Fields{
        "error":    "timeout",
        "retry":    3,
        "endpoint": "api.example.com",
    })
}
```

### As a CLI Agent

```bash
# Run the test demo
go run cmd/sink/main.go test

# Output:
# [15:04:05] INFO | test-service | sink initialized
# [15:04:05] DEBUG | test-service | starting test sequence | version=0.1.0 mode=test
# [15:04:05] WARN | test-service | this is a warning | retry_count=3
# [15:04:05] ERROR | test-service | simulated error | error_code=E001 user_id=u123
```

---

## 📦 Installation

```bash
go get github.com/stawuah/sink/v2
```

Or clone and build:

```bash
git clone https://github.com/stawuah/sink.git
cd sink
go build -o sink cmd/sink/main.go
```

---

## 🏗️ Architecture

```
┌─────────────────────────────────────────────────────┐
│                    Your Application                  │
└─────────────────────┬───────────────────────────────┘
                      │
                      ▼
           ┌──────────────────────┐
           │  Sink Logger Core    │
           │  - Level Filtering   │
           │  - Field Validation  │
           │  - Thread Safety     │
           └──────────┬───────────┘
                      │
        ┌─────────────┼─────────────┐
        ▼             ▼             ▼
   ┌─────────┐  ┌─────────┐  ┌──────────┐
   │ Console │  │  File   │  │   HTTP   │
   │  Sink   │  │  Sink   │  │   Sink   │
   └─────────┘  └─────────┘  └──────────┘
        │             │             │
        ▼             ▼             ▼
     stdout       local disk    remote API
```

---

## 🎨 Output Formats

### Pretty Mode (Development)
```
[15:04:05] INFO | payments-api | payment processed | user_id=u123 amount=99.99
[15:04:06] ERROR | payments-api | transaction failed | error=insufficient_funds
```

### JSON Mode (Production)
```json
{
  "time": "2025-12-29T15:04:05Z",
  "level": "INFO",
  "service": "payments-api",
  "env": "production",
  "message": "payment processed",
  "fields": {
    "user_id": "u123",
    "amount": 99.99
  }
}
```

---

## 🔧 Configuration

Sink adapts to your workflow:

```go
logger.Config{
    Service:  "edge-node-01",    // Service identifier
    Env:      "edge",             // Environment (dev/staging/prod)
    MinLevel: logger.InfoLevel,  // Filter logs below this level
    Async:    true,               // Enable async buffering (coming soon)
}
```

---

## 🌍 Why Edge-Optimized?

Traditional logging solutions struggle at the edge:

| Challenge | Sink's Solution |
|-----------|-----------------|
| Intermittent connectivity | Local buffering, flush when online |
| Limited memory | Ring buffer with backpressure |
| ARM/low-power CPUs | Zero-allocation hot path |
| No Docker/K8s | Single binary, no dependencies |
| High log volume | Level filtering, async writes |

---

## 📚 Use Cases

### Edge IoT Gateway
```go
log := logger.New(
    logger.Config{Service: "iot-gateway", Env: "edge"},
    sinks.NewFile("/var/log/gateway.log"),
)
log.Info("sensor reading", logger.Fields{"temp": 23.5, "humidity": 65})
```

### CI/CD Build Agent
```go
log := logger.New(
    logger.Config{Service: "build-agent", Env: "ci"},
    sinks.NewConsole(false), // JSON for parsing
)
log.Info("test passed", logger.Fields{"suite": "integration", "duration_ms": 1240})
```

### Kubernetes Microservice
```go
log := logger.New(
    logger.Config{Service: os.Getenv("SERVICE_NAME"), Env: "prod"},
    sinks.NewConsole(false),
    sinks.NewHTTP("https://logs.example.com/ingest"), // coming soon
)
```

---

## 🛣️ Roadmap

### Phase 1 - Foundation ✅
- [x] Core logger implementation
- [x] Console sink (pretty & JSON)
- [x] Structured fields
- [x] Log levels & filtering

### Phase 2 - Edge Power 🚧
- [ ] Async ring buffer with backpressure
- [ ] File sink with rotation
- [ ] YAML configuration
- [ ] Benchmarks & performance tests

### Phase 3 - Production 📋
- [ ] HTTP exporter (Loki, Grafana Cloud)
- [ ] Unix socket ingestion (sidecar mode)
- [ ] WAL (write-ahead log) for durability
- [ ] Compression & batching
- [ ] Full CLI with commands

---

## 🤝 Contributing

Contributions welcome! This project is being built in public as a learning resource for Go infrastructure projects.

```bash
# Get started
git clone https://github.com/stawuah/sink.git
cd sink
go run cmd/sink/main.go test
```

---

## 📄 License

MIT License - see [LICENSE](LICENSE) for details

---

## 🙏 Acknowledgments

Inspired by observability challenges in edge computing and the need for lightweight, dependency-free logging solutions for distributed systems.

**Built with ❤️ for the edge computing community**

---

## 📬 Contact

- **GitHub**: [@stawuah](https://github.com/stawuah)
- **Issues**: [github.com/stawuah/sink/issues](https://github.com/stawuah/sink/issues)

---

<p align="center">
  <strong>⭐ Star this repo if you find it useful!</strong>
</p>
