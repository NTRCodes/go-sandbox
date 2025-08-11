# go-sandbox

**Purpose:**  
This repository is my personal Go lab — a dedicated space for building small, focused exercises ("katas") and documenting idioms that sharpen my skills as a backend/infrastructure engineer.

The work here directly supports my long-term goal: delivering secure, observable, production-grade services in Go.

---

## 📂 Structure

go-sandbox/
├── katas/ # Self-contained exercises with tests
│ ├── lru_cache/
│ │ ├── lru.go
│ │ └── lru_test.go
│ └── worker_pool/
│ ├── pool.go
│ └── pool_test.go
└── notes/ # Written learnings
├── idioms.md # Go language patterns worth keeping
└── lessons.md # Mistakes, insights, and "gotchas"

---

## 🛠 Practices

This repo follows the same principles I use for production work:

- **Table-driven tests** for clear, repeatable coverage
- **TDD-first** development on each kata
- `go test ./... -race` clean before commit
- Idiomatic Go conventions (interfaces, zero-values, error handling)
- Incremental commits with descriptive messages

---

## 🎯 Goals

- Build muscle memory for Go syntax, idioms, and concurrency patterns
- Maintain a living reference (`idioms.md`) of patterns to use in production
- Apply test-driven workflows to both small katas and larger services
- Demonstrate continuous improvement and professional discipline in code

---

## 📈 Progression

Katas increase in complexity over time:

1. Core data structures (LRU cache, ring buffer)
2. Concurrency tools (worker pools, pipelines, cancellation)
3. Networked examples (basic HTTP services, request handling)
4. Integrations (databases, metrics, tracing — in mini form)

---

## 🧪 Running Tests

```bash
go test ./... -v

