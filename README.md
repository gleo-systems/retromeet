# What is RetroMeet?
RetroMeet is a web browser application for software developers.<br />
The tool improves development process by facilitating retrospective ceremonies.

## Running application server
Start server with `go run main.go` or `retromeet run -p 8080`.

## Technology stack
* `Go` language,
* `Amazon Web Services` cloud,
* `Kubernetes` containers,
* `Postgres` database.

### Project structure
* `build/package` contains Docker descriptors,
* `cmd` contains application commands code,
* `deployments` contains Kubernetes descriptors,
* `internal` contains application code,
* `scripts/db` contains DB init scripts,
* `test` contains application tests code.

## Future improvements
* add Docker descriptor,
* add Kubernetes descriptor,
* integrate with SES AWS email service. 
