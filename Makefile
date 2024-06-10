.SILENT:
run:
	find ./question -name $(shell printf '%03d' $(num))* -print0 | xargs -0 -I {}  go run  {}
