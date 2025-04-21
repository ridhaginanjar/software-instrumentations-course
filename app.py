# Ini adalah contoh kode aplikasi yang membuat metrics secara manual

from prometheus_client import Counter, Histogram, generate_latest
from flask import Flask, Response, request
import time

app = Flask(__name__)

# METRICS
REQUEST_COUNT = Counter(
    'app_request_count', 
    'Total number of requests',
    ['method', 'endpoint']
)

REQUEST_LATENCY = Histogram(
    'app_request_latency_seconds', 
    'Latency of requests in seconds',
    ['endpoint']
)

# ROUTES
@app.route('/predict', methods=['POST'])
def predict():
    start_time = time.time()
    REQUEST_COUNT.labels(method='POST', endpoint='/predict').inc()

    # Simulasi proses prediksi
    time.sleep(0.2)
    latency = time.time() - start_time
    REQUEST_LATENCY.labels(endpoint='/predict').observe(latency)

    return {"result": "predicted_value"}

@app.route('/metrics')
def metrics():
    return Response(generate_latest(), mimetype='text/plain')

if __name__ == '__main__':
    app.run(debug=True)
