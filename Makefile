deps:
	pip install -e .[dev]
.PHONY: deps

lint:
	flake8 --verbose dcos_cli/ tests/
	pylint --disable=missing-docstring dcos_cli/ tests/
.PHONY: lint

test: lint
	pytest --cov=dcos_cli --cov-fail-under=100 tests/
.PHONY: test
