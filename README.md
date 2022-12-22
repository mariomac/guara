# Guara

My own personal Go toolset

## `pkg/cache`

Cache implementations.

## `pkg/rate`

Tools for rate calculation and limiting.

## `pkg/test`

Some tools used in testing.

* `eventually.go`: repeats a test until it finally succeeds (then the test is marked as succeeded),
  or until it fails (then the test is marked as failed).
* `ports.go`: find a free port for spawning test services.
