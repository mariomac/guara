# Guara

My own personal Go toolset

## Package `cache`

Cache implementations.

## Package `casing`

String casing conversion tools. For example, from `CamelCase` or `DromedaryCase` to `dot.case` or `snake_case`.

## Package `err`

Error handling tools.

## Package `maps`

Esoteric data structures based on maps.

## Package `rate`

Tools for rate calculation and limiting.

## Package `test`

Some tools used in testing.

* `eventually.go`: repeats a test until it finally succeeds (then the test is marked as succeeded),
  or until it fails (then the test is marked as failed).
* `ports.go`: find a free port for spawning test services.

