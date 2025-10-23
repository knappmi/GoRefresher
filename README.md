# GoRefresher
A structured refresher to practice Go from fundamentals through expert-level concepts. All problem implementations are placeholders so you can branch off main and create your own solutions.

## Running Tests Manually

Basic (run everything):
```
go test ./...
```

Single problem (example – fundamental/fizzbuzz_variant):
```
go test ./problems/fundamental/fizzbuzz_variant
```

All problems in a category:
```
go test ./problems/fundamental/...
```

Only test problems changed since main (local diff example):
```
for d in $(git diff --name-only origin/main...HEAD | grep '^problems/' | cut -d/ -f1-3 | sort -u); do 
  echo "Testing ./$d"; go test "./$d" || exit 1; 
done
```
(Adjust origin/main if your remote or base branch differs.)

Run benchmarks in a specific package:
```
go test -bench . -benchmem ./problems/intermediate/benchmarking
```

Run race detector (recommended for concurrency problems):
```
go test -race ./problems/intermediate/goroutines_basics
```

## Learning Path Overview

### Fundamentals
Core language building blocks and standard library essentials.
Focus Areas:
- Control flow & basic algorithms (fizzbuzz_variant, fibonacci)
- Unicode & text handling (reverse_unicode, word_count)
- Data structures & interfaces (shape_interface)
- JSON encoding/decoding & tags (json_struct)
- Error modeling (custom_error)
- File I/O & performance basics (file_linecount)
Concepts Reinforced: slices vs arrays, rune handling, structs & tags, interfaces, error wrapping, bufio efficiency, benchmarking fundamentals.

### Intermediate
Applied concurrency patterns, generics, testing strategies, and modular design.
Focus Areas:
- Goroutines & synchronization (goroutines_basics)
- Channel patterns (channel_patterns)
- Context for cancellation & timeouts (context_cancellation)
- HTTP server + resilient client (http_rest_client_server)
- Generics collections & constraints (generics_collections)
- Streaming large data (json_streaming_decoder)
- Reflection for dynamic inspection (interface_reflection)
- Module versioning (go_modules_versioning)
- Table & property-based testing (testing_table_property)
- Benchmarking & allocation awareness (benchmarking)
- Error wrapping strategies (error_wrapping)
- File watching & debounce (file_watcher)
Concepts Reinforced: race detection, channel design patterns (fan-in/fan-out, pipelines), context propagation, exponential backoff ideas, generics syntax & type parameters, reflection API basics, property-based testing, performance profiling entry points.

### Expert
Advanced concurrency, performance optimization, systems design, and platform integration.
Focus Areas:
- Lock-free data structures (lockfree_ringbuffer)
- Scheduler simulation & work stealing concepts (custom_scheduler_sim)
- Allocation minimization & pooling (zero_alloc_json)
- Plugin architecture & dynamic loading (plugin_architecture)
- Interop with C via cgo (cgo_image_processing)
- WebAssembly targets (wasm_module)
- Observability (tracing_instrumentation)
- Resilience patterns: retry & circuit breaker (retry_circuitbreaker)
- Distributed consensus concepts (raft_like_log)
- gRPC interceptor chain design (grpc_interceptors)
- Advanced generics & numeric type sets (advanced_generics_constraints)
- Data race diagnosis & mitigation (data_race_detector_lab)
- Profiling & optimization workflow (profiling_optimization)
- Build tags for environment-specific behavior (build_tags_ci)
- Scalable concurrent cache design (sharded_inmem_cache)
Concepts Reinforced: atomics & memory ordering intuition, scheduling principles, unsafe & pooling trade-offs, dynamic linking constraints, ABI considerations with cgo, cross-compilation (WASM), telemetry instrumentation, failure state machines, simplified consensus workflow, middleware layering, generic constraints & type sets, race detector output analysis, CPU & memory profiling (pprof), conditional compilation, sharding & hashing strategies (FNV / consistent hashing potential).

## Suggested Workflow
1. Create a feature branch (e.g., `git checkout -b solve/fizzbuzz`).
2. Implement solution replacing TODOs.
3. Add/activate tests (remove t.Skip). Write benchmarks where relevant.
4. Run lint/static analysis (optionally add `golangci-lint` locally).
5. Merge via PR (manual review) back into main.

## Stretch Integrations
- Add a Makefile (make test, make bench, make race).
- Add coverage reporting: `go test -coverprofile=coverage.out ./...`.
- Integrate tooling: golangci-lint, code generation for mocks.
- Introduce real tracing via OpenTelemetry SDK.

## Contributing
All stubs are intentionally minimal; prefer incremental commits per problem. Keep solutions self-contained within their problem directory.

Happy coding!
