# Expert Problems

1. lockfree_ringbuffer - Implement a lock-free MPMC ring buffer using atomic operations & generics.
2. custom_scheduler_sim - Simulate a minimal goroutine scheduler (work stealing conceptually).
3. zero_alloc_json - High-performance JSON encoder/decoder avoiding allocations (use sync.Pool, unsafe optional).
4. plugin_architecture - Dynamic plugin loading via Go plugins (build tags fallback when unsupported).
5. cgo_image_processing - Call a simple C image blur function from Go, manage memory correctly.
6. wasm_module - Compile a Go function to WebAssembly and call from JS (tiny example + fetch).
7. tracing_instrumentation - Add OpenTelemetry tracing + metrics to a layered service.
8. retry_circuitbreaker - Implement exponential backoff with jitter + circuit breaker state machine.
9. raft_like_log - Simplified replicated log (leader election via timeouts, no persistence needed).
10. grpc_interceptors - Unary & stream interceptors for metrics, auth, and retries.
11. advanced_generics_constraints - Type set for numeric ops + generic matrix operations.
12. data_race_detector_lab - Intentionally create + then fix data races, explaining -race output.
13. profiling_optimization - Provide naive code; profile (pprof), optimize CPU & memory, document gains.
14. build_tags_ci - Use build tags for different environments (dev, prod) and demonstrate go:build lines.
15. sharded_inmem_cache - Sharded concurrent cache with TTL, LFU eviction & metrics.

Each folder should include:
- README.md (problem statement, goals, hints, stretch ideas)
- main.go or package code stub
- *_test.go with TODO tests or failing tests guiding implementation

Stretch: Combine multiple problems into a mini distributed system (cache + tracing + gRPC interceptors).
