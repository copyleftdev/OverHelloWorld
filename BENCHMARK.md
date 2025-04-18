# OverHelloWorld Benchmark Results

_This document summarizes the results of advanced load testing performed with [k6](https://k6.io/) on the OverHelloWorld project._

## Test Environment
- **Date:** 2025-04-17
- **Host:** Local Docker Compose (app + Redis)
- **Load Tool:** k6 v0.49+
- **Script:** `tests/load/advanced-hello-load-test.js`
- **App Version:** Initial public release

## Scenarios
- **Ramping Users:** Up to 100 VUs over 3 minutes
- **Constant Arrival Rate:** 40 iterations/sec for 2 minutes (max 100 VUs)
- **Mixed GET/POST traffic** to `/`, `/hello`, and `/hello` (POST)

## Thresholds (all passed)
- 95% of requests < 500ms
- Avg custom response time < 300ms
- Error rate < 1%

## Key Results

| Metric                     | Value                |
|---------------------------|----------------------|
| Total iterations          | 12,665               |
| Total requests            | 17,434               |
| Avg response time (custom)| 0.83 ms              |
| 95th percentile duration  | 2.61 ms              |
| Error rate                | 0%                   |
| Data received             | 1.5 GB               |
| Data sent                 | 2.0 MB               |
| Max VUs                   | 150                  |

## Detailed Results
```
custom_response_time (avg): 0.83 ms
http_req_duration (p95): 2.61 ms
http_req_failed: 0%
checks_total: 25,330
checks_succeeded: 100%
iterations: 12,665
http_reqs: 17,434
vus_max: 150
```

## Scenario Checks
- ✓ POST /hello status 202
- ✓ GET /hello status 200
- ✓ GET / status 200
- ✓ body contains OverEngineered

## Interpretation
- The OverHelloWorld app demonstrates extreme overengineering *and* extreme performance.
- System easily handled heavy, mixed traffic with zero errors and sub-millisecond latencies.
- Further scaling or chaos testing is likely to be limited by network or Docker host, not the app itself.

---

*For more details, see the k6 script and raw output in `tests/load/`.*
