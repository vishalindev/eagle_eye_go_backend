SHELL := /usr/bin/env bash
SERVICES := admin camera dashboard identity mobile notification platformuser reports rule zone
MODULES := shared $(addprefix services/,$(addsuffix -service,$(SERVICES)))

.PHONY: tidy test run-% docker-build-%

tidy:
	@for module in $(MODULES); do \
		echo "==> go mod tidy $$module"; \
		(cd $$module && go mod tidy); \
	done

test:
	@for module in $(MODULES); do \
		echo "==> go test $$module"; \
		(cd $$module && go test ./...); \
	done

run-%:
	cd services/$*-service && go run ./cmd

docker-build-%:
	docker build -t eagle-eye/$*-service:local services/$*-service
