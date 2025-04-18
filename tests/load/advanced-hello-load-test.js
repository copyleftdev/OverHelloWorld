import http from 'k6/http';
import { check, sleep } from 'k6';
import { Trend } from 'k6/metrics';

export let responseTime = new Trend('custom_response_time');

export const options = {
  scenarios: {
    ramping_users: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '30s', target: 10 },
        { duration: '1m', target: 50 },
        { duration: '1m', target: 100 },
        { duration: '30s', target: 0 },
      ],
      gracefulRampDown: '30s',
    },
    constant_rate: {
      executor: 'constant-arrival-rate',
      rate: 40, // 40 iterations per second
      timeUnit: '1s',
      duration: '2m',
      preAllocatedVUs: 20,
      maxVUs: 100,
      exec: 'apiFlow',
    },
  },
  thresholds: {
    http_req_duration: ['p(95)<500'], // 95% of requests < 500ms
    'custom_response_time': ['avg<300'],
    http_req_failed: ['rate<0.01'], // <1% errors
  },
  tags: { project: 'OverHelloWorld' },
};

export function setup() {
  // Could create initial data if needed
  return { testUser: 'performance-tester' };
}

export default function (data) {
  // Simple GET root endpoint
  let res = http.get('http://localhost:8080/', { tags: { name: 'RootEndpoint' } });
  responseTime.add(res.timings.duration);
  check(res, {
    'GET / status 200': (r) => r.status === 200,
    'body contains OverEngineered': (r) => r.body && r.body.includes('OverEngineered'),
  });
  sleep(Math.random() * 2);
}

export function apiFlow() {
  // POST /hello with random message
  const msg = `Hello from VU ${__VU} at ${Date.now()}`;
  let res1 = http.post('http://localhost:8080/hello', JSON.stringify({ message: msg }), {
    headers: { 'Content-Type': 'application/json' },
    tags: { name: 'PostHello' },
  });
  check(res1, { 'POST /hello status 202': (r) => r.status === 202 });

  // GET /hello
  let res2 = http.get('http://localhost:8080/hello', { tags: { name: 'GetHellos' } });
  check(res2, { 'GET /hello status 200': (r) => r.status === 200 });
  responseTime.add(res2.timings.duration);
  sleep(Math.random() * 2);
}

export function teardown(data) {
  // Could clean up test data if needed
}
