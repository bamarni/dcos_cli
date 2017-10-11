deps:
	pip install -e .[dev]
.PHONY: deps

test:
	pytest --cov=dcos_cli --cov-fail-under=100 tests/
.PHONY: test
