import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    stages: [
        { duration: '2m', target: 200 },
        { duration: '5m', target: 200 },
        { duration: '1m', target: 0   },
    ],
    noConnectionReuse: true,
    noVUConnectionReuse: true,
};

// The function that defines VU logic.
//
// See https://grafana.com/docs/k6/latest/examples/get-started-with-k6/ to learn more
// about authoring k6 scripts.
//
export default function() {
    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const payload = JSON.stringify({
        test: 'hi',
        number: 10,
    });

    const res = http.post(`http://192.168.49.2:${__ENV.PORT}/`, payload, params);
    check(res, { 'Status is 200': (r) => r.status == 200 });
    sleep(1);
}
