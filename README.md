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
