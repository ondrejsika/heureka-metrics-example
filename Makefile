prometheus:
	prometheus --config.file=prometheus.yml --web.enable-lifecycle

reload:
	curl -X POST http://127.0.0.1:9090/-/reload

alertmanager:
	docker run --name alertmanager -p 127.0.0.1:9093:9093 -v $(shell pwd):/etc/alertmanager quay.io/prometheus/alertmanager
