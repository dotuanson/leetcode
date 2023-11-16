gen_tests:
	gotests -all -w $(PROBLEM)

run_tests:
	grc gow test $(PROBLEM)/main.go  $(PROBLEM)/main_test.go

.PHONY: gen_tests run_tests