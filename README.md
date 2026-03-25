# LRU Cache in Go

## Benchmarks

### Pre Doubly Linked List

> `goos: darwin` · `goarch: arm64` · `cpu: Apple M5`

| Benchmark | Iterations | ns/op | B/op | allocs/op |
|---|---|---|---|---|
| Get/Empty Cache | 599,566,063 | 1.834 | 0 | 0 |
| Get/Non Empty Cache Hit | 22,574,494 | 53.02 | 0 | 0 |
| Get/Non Empty Cache Miss | 190,246,208 | 6.304 | 0 | 0 |
| Put/Empty Cache | 40,485,033 | 29.70 | 0 | 0 |
| Put/Non Empty Cache | 40,274,199 | 29.85 | 0 | 0 |
| Put/Full Cache | 40,606,789 | 29.75 | 0 | 0 |
| removeLRU/Full Cache | 755,883,090 | 1.579 | 0 | 0 |
| removeLRU/Empty Cache | 759,861,380 | 1.579 | 0 | 0 |

### Post Doubly Linked List

> `goos: darwin` · `goarch: arm64` · `cpu: Apple M5`

| Benchmark | Iterations | ns/op | B/op | allocs/op |
|---|---|---|---|---|
| Get/Empty Cache | 226,472,496 | 5.183 | 0 | 0 |
| Get/Non Empty Cache Hit | 168,641,706 | 7.117 | 0 | 0 |
| Get/Non Empty Cache Miss | 157,311,098 | 7.561 | 0 | 0 |
| Put/Empty Cache | 144,611,583 | 8.290 | 0 | 0 |
| Put/Non Empty Cache | 145,636,326 | 8.221 | 0 | 0 |
| Put/Non Empty Cache (Already Exists) | 143,289,498 | 8.364 | 0 | 0 |
| Put/Full Cache | 146,371,555 | 8.201 | 0 | 0 |

### Comparison

| Benchmark | Before (ns/op) | After (ns/op) | Improvement |
|---|---|---|---|
| Get/Empty Cache | 1.834 | 5.183 | -2.8x |
| Get/Non Empty Cache Hit | 53.02 | 7.117 | **7.5x faster** |
| Get/Non Empty Cache Miss | 6.304 | 7.561 | ~same |
| Put/Empty Cache | 29.70 | 8.290 | **3.6x faster** |
| Put/Non Empty Cache | 29.85 | 8.221 | **3.6x faster** |
| Put/Full Cache | 29.75 | 8.201 | **3.6x faster** |