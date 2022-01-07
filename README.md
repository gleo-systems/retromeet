# What is RetroMeet?
RetroMeet is a web browser application for software developers.<br />
The tool improves development process by facilitating retrospective ceremonies.

## Running application
Set required OS environmental variables.<br />
Start application server by executing command `go run main.go server`.

## Technology stack
* `Go` programming language,
* `Amazon AWS` cloud provider,
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
* improve DB connection,
* add Docker descriptor,
* add Kubernetes descriptor,
* integrate with SES AWS email service,
* `gocraft/health` implement health check feature.
