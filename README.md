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

#### LRU Cache

| Benchmark | Iterations | ns/op | B/op | allocs/op |
|---|---|---|---|---|
| Get/Empty Cache | 227,664,852 | 5.129 | 0 | 0 |
| Get/Non Empty Cache Hit | 167,288,635 | 7.165 | 0 | 0 |
| Get/Non Empty Cache Miss | 157,714,860 | 7.584 | 0 | 0 |
| Put/Empty Cache | 147,354,675 | 8.145 | 0 | 0 |
| Put/Non Empty Cache | 144,633,174 | 8.296 | 0 | 0 |
| Put/Non Empty Cache (Already Exists) | 143,765,977 | 8.344 | 0 | 0 |
| Put/Full Cache | 144,393,660 | 8.307 | 0 | 0 |

#### Doubly Linked List

| Benchmark | Iterations | ns/op | B/op | allocs/op |
|---|---|---|---|---|
| MoveToFront | 473,908,225 | 2.381 | 0 | 0 |
| Back | 755,061,663 | 1.578 | 0 | 0 |
| Remove | 750,716,698 | 1.595 | 0 | 0 |

### Comparison

| Benchmark | Before (ns/op) | After (ns/op) | Improvement |
|---|---|---|---|
| Get/Empty Cache | 1.834 | 5.129 | -2.8x |
| Get/Non Empty Cache Hit | 53.02 | 7.165 | **7.4x faster** |
| Get/Non Empty Cache Miss | 6.304 | 7.584 | ~same |
| Put/Empty Cache | 29.70 | 8.145 | **3.6x faster** |
| Put/Non Empty Cache | 29.85 | 8.296 | **3.6x faster** |
| Put/Full Cache | 29.75 | 8.307 | **3.6x faster** |