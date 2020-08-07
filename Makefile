.PHONY: tests test-coverage

tests:
	@./_scripts/tests_unit.sh

test-coverage:
	@./_scripts/tests_coverage.sh