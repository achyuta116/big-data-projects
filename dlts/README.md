## Big Data - Distributed Load Testing System
### Goal
Design and build a distributed load-testing system that co-ordinates between
multiple driver nodes to run a highly concurrent, high-throughput load test on a
web server. This system will use Kafka as a communication service.

### File Structure
```bash
.
├── client
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── docker-compose.nodes.yml
├── docker-compose.yml
├── driver
│   ├── Dockerfile
│   └── main.go
├── go.mod
├── go.sum
├── lib
│   └── dlts.go
├── orchestrator
│   ├── Dockerfile
│   └── main.go
└── server
    ├── Dockerfile
    └── main.go
```
- `client` (Test - Client) -> Monitors Heartbeat sent out by all nodes, verifying the functionality offered by segmentio/kafka-go
- `driver` (Driver Node) -> Responsible for conducting specific load tests on the `server` - Details of which are mentioned in the project description
- `orchestrator` (Orchestrator Node) -> Responsible for managing Driver Nodes and triggering tests via sending messages through Kafka
- `server` (Test - Server) -> The server on which load tests are being conducted on

### Requirements
- Install `docker` for your respective Operating System
- Make sure ports 8000 and 8001 are not used by other processes running on your system
> [!NOTE]
> Check this by running `lsof -i:tcp` on a Unix based or Mac OS system
### Setup
Execute the following steps **in order** from the root of this repository
1. `docker compose up` -> sets up Kafka and Zookeeper running on your system and exposes relevant ports 9093 and 2181 respectively
2. `docker compose -f docker-compose.nodes.yml up` -> sets up Driver and Orchestrator nodes as docker images and runs them on the same network as the Kafka and Zookeeper containers

You will now be able to send HTTP requests to the Orchestrator Node endpoints mapped to port 8001 on your system  
Heartbeats will now be picked up by the Orchestrator Node, reflecting in `docker logs`
  
**Example**  
- Test Server
```sh 
$ curl 127.0.0.1:8000 # returns `Hello World` from the server to be tested
> Hello World
$ curl 127.0.0.1:8000/metric # returns the number of times the server has been sent a request
> 1
```
- Run Avalanche Test
```sh
$ curl -X POST 127.0.0.1:8001/avalanche
# requests the orchestrator to send an avalanche test config to the driver nodes
> {"TestId":"XVlBzgbaiCMRAjWwhTHc"}
$ curl -X POST 127.0.0.1:8001/trigger/XVlBzgbaiCMRAjWwhTHc
# trigger the actual avalanche test specifying the test ID
```
- Run Tsunami Test
```sh
$ curl -X POST 127.0.0.1:8001/tsunami/2
# requests the orchestrator to send an tsunami test config to the driver nodes
> {"TestId":"tcuAxhxKQFDaFpLSjFbc"}
# The route parameter passed in dictates the number of seconds
# the tsunami intervals are configured to be between subsequent loads
$ curl -X POST 127.0.0.1:8001/trigger/tcuAxhxKQFDaFpLSjFbc
# trigger the actual tsunami test specifying the test ID
```
On successful submission of these commands, the docker containers should log out metrics of the ongoing test, updating each second.  
> [!NOTE]
> - Each log message is a JSON object in the format specified by the project reference
> - The message formats are shared between the Orchestrator and Driver nodes through `lib/dlts.go`
### References
- [This](https://hackmd.io/@pesu-bigdata/S1nvSXAza) wonderful project description written by our TAs are Uni
- [Learn](https://rmoff.net/2018/08/02/kafka-listeners-explained/) how Kafka listeners work and how to configure them
- [Setup](https://www.baeldung.com/ops/kafka-docker-setup) Kafka and Zookeeper on `docker` containers
