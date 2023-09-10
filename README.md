### Geo-Port service

This service has been built following several architecture paradigms:
- [Hexagonal](https://alistair.cockburn.us/hexagonal-architecture/)
- [Domain-Driven Design](https://martinfowler.com/bliki/DomainDrivenDesign.html) (DDD)
- [Command and Query Separation](https://martinfowler.com/bliki/CommandQuerySeparation.html) (CQS)
- [SOLID](https://en.wikipedia.org/wiki/SOLID) principles

### TODO: Architecture Diagram (Architecture as Code https://plantuml.com/component-diagram)

### Setup

- Go >= 1.20
- Docker >= 24.0.5

### Requirements

Install dependencies by running:

```bash
$ make deps
```

Install golangci-lint locally for Go linting (https://golangci-lint.run/)

To install 
```bash
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2

golangci-lint --version
```

### Run unit tests

```bash
$ make test
```

### Start the application

1. Build Docker image

```bash
$ make docker-build
```

2. Run the container

```bash
docker run -v ./data:/data geo-port aup -f /data/ports.json
```

### What is missing ? Unordered response

- Feature test to verify e2e workflow, inserting ports data through CLI and verify data is correctly stored.

- Integration test at the component level to test infrastructure store adapter and database together to verify that port store methods store data correctly.

- Unit tests:
    - service layer tests should cover unhappy paths

- Proper logging system in a shared infrastructure.

- Engineering / Product effort to figure if service could be used by concurrent users and what would be the impact on service (especially at application level workflow).