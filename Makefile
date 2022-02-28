prometheus:
	prometheus --config.file=prometheus.yml --web.enable-lifecycle

reload:
	curl -X POST http://127.0.0.1:9090/-/reload
