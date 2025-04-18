# OverHelloWorld

A robust, over-engineered "Hello World" Go application demonstrating:
- DDD, CQRS, and event sourcing
- Redis and file-backed event bus
- Plugin system (ASCII, TTS, LED, etc)
- Prometheus & OpenTelemetry observability
- Property-based and integration tests
- Automated test reporting and coverage

---

## 🚀 Quick Start

```bash
# Run all tests with coverage and generate a timestamped report folder
test_reports/YYYY-MM-DDTHH-MM-SS/
./run_tests.sh
```

## 🧪 Test Suite & Coverage

- **Unit, integration, and property-based tests**
- **Detailed coverage report** in each test run folder
- **Property-based tests** for:
  - Event sourcing (file store)
  - Plugins
  - API roundtrip
  - Redis event bus (if REDIS_ADDR set)

### Run All Tests
```bash
./run_tests.sh
```

### View Coverage
- Open the generated HTML report, e.g.:
  ```
  xdg-open test_reports/YYYY-MM-DDTHH-MM-SS/coverage.html
  ```

### Test Status & Coverage Badges

If using GitHub Actions, add this to your README:

```
![Go](https://github.com/<your-username>/OverHelloWorld/actions/workflows/go.yml/badge.svg)
![Coverage](https://img.shields.io/badge/coverage-dynamic-brightgreen)
```

(Replace `<your-username>` with your GitHub username)

---

![CI](https://github.com/copyleftdev/OverHelloWorld/actions/workflows/ci.yml/badge.svg)
![Coverage](https://img.shields.io/badge/coverage-dynamic-brightgreen?style=for-the-badge)

---

# 📊 Benchmark

See [BENCHMARK.md](./BENCHMARK.md) for the latest load test results using k6.

---

# 📄 Documentation

- [OVERENGINEERING_ARTICLE.md](./OVERENGINEERING_ARTICLE.md): Satirical deep-dive on overengineering.
- [BENCHMARK.md](./BENCHMARK.md): Performance and load test results.
- [CONTRIBUTING.md](./CONTRIBUTING.md): How to contribute, code style, PR guidelines.
- [CODE_OF_CONDUCT.md](./CODE_OF_CONDUCT.md): Community standards and behavior.
- [SECURITY.md](./SECURITY.md): How to report vulnerabilities.
- [SUPPORT.md](./SUPPORT.md): Where and how to get help.
- [LICENSE](./LICENSE): Project license.

---

## 🔬 Property-Based Testing

- Uses [gopter](https://github.com/leanovate/gopter) for hypothesis-style tests.
- See `tests/integration/property_test.go`, `property_redis_test.go`, `property_file_event_store_test.go`.

## 📦 Test Reports
- All test runs output to `test_reports/<timestamp>/`:
  - `report.json`: Full JSON test log
  - `coverage.out`: Go coverage profile
  - `coverage.html`: Visual HTML coverage report

---

## 🛠️ CI/CD
- Recommended: Add a GitHub Actions workflow for Go tests and coverage.
- Artifacts can be uploaded for each run.

## 📝 How to Add More Tests
- Add new property-based or integration tests in `tests/integration/`.
- Extend the plugin system or event store, and verify with new properties!

---

## 📚 Documentation
- See inline comments and test files for usage and extension examples.
- For questions, open an issue or PR!

---

## 🏆 Project Status
- **Feature-complete, robustly tested, and ready for extension or production demo!**

---

![Overengineering](https://img.shields.io/badge/overengineering-100%25-red?style=for-the-badge&logo=github)
![Microservices](https://img.shields.io/badge/microservices-42-in-production-blueviolet?style=for-the-badge)
![Unnecessary Patterns](https://img.shields.io/badge/patterns-Duck%20Factory%20CQRS-orange?style=for-the-badge)
![Job Security](https://img.shields.io/badge/job%20security-guaranteed-success?style=for-the-badge)
![Agility](https://img.shields.io/badge/agility-0%25-critical?style=for-the-badge)
![CI/CD](https://img.shields.io/badge/CI%2FCD-Just%20for%20Show-lightgrey?style=for-the-badge)
![Ego Driven](https://img.shields.io/badge/ego-driven-important?style=for-the-badge)
![YAGNI](https://img.shields.io/badge/YAGNI-violated-red?style=for-the-badge)
![KISS](https://img.shields.io/badge/KISS-ignored-yellow?style=for-the-badge)
