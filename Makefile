NAME = $(notdir $(PWD))

VERSION = $(shell printf "%s.%s" \
		$$(git rev-list --count HEAD) \
		$$(git rev-parse --short HEAD))

build:
	CGO_ENABLED=0 GOOS=linux go build -o bin/app \
		-ldflags "-X main.version=$(VERSION)" \
		-gcflags "-trimpath $(GOPATH)/src" \
		./cmd/$(NAME)


image: build
	docker build -t $(NAME):$(VERSION) -f Dockerfile .

push@%:
	$(eval VERSION ?= latest)
	$(eval TAG = $*/$(NAME):$(VERSION))
	docker tag $(NAME):$(VERSION) $(TAG)
	docker push $(TAG)

	@if [[ "$(tag-file)" ]]; then echo "$(TAG)" > "$(tag-file)"; fi
	@if [[ "$(version-file)" ]]; then echo "$(VERSION)" > "$(version-file)"; fi
