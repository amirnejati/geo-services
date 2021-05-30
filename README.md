# Ride-Hailing Dynamic Pricing System
This project addresses the problem mentioned [here](https://github.com/AliKarami/interview-tasks/tree/master/surge). A proof of concept for any location-based service (LBS) which needs to manage prices dynamically based on supply/demand relation.

___
## Narrative
In a ride-hailing system, based on changes in demand rate, prices should be able to fluctuate as incentives to cover market requirements.

A very first solution is to count requests per municipal districts, obtained from each request's origin-point. So, there should be a system which has geo-data and gets districtNo based on point. Demand values per district must be checked against a reference table to change coefficients in realtime, if thresholds met.

___
## Technologies Used
- [Go](https://golang.org) programming language
- [gokit](https://gokit.io) A toolkit for microservices in Go
- RESTful API & [gRPC](https://grpc.io) + [protobuf](https://github.com/protocolbuffers/protobuf) for communications
- [MongoDB](https://www.mongodb.com) for CRUD & geospatial queries
- [Redis](https://redis.io) for caching
- [NATS](https://nats.io) as messaging system
- [Docker](https://www.docker.com) & Compose for containerization & orchestration
- [Nginx](https://nginx.org) as HTTP-server & traffic mirroring tool

___
## Terminology & Parts
* **RP**: A *reverse-proxy* able to make mirrored(shadowed) traffic & pass it to some additional backend
* **MQ**: A message queue with fast pub/sub support
* **Cache**: An in-memory address/db, remotely accessible to multiple clients
* **Mediator**: Service which accepts HTTP traffics and puts them into the *MQ*
* **CoreSystem(CrS)**: The main service responsible to respond users requests & calculate trip price based on the most recent data received from *surge-supervisor*
* **SurgeSupervisor(SrS)**: Service that monitors demand rates in realtime & informs new pricing coefficients to *core-system* if any changes happen to its value
* **MapServices(MpS)**: Service responsible for geo queries including the mapping of any point(x, y) to district-number

___
## Project Design
Components orchestration as shown in the image below:

![surge-pricing-architecture](https://user-images.githubusercontent.com/8034519/120118898-0eb0c900-c1aa-11eb-96be-746646cced01.jpg)

___
## System Workflow
The sequence diagram below shows flows in the system:

![surge-pricing-sequences-diagram](https://user-images.githubusercontent.com/8034519/120118912-20926c00-c1aa-11eb-936f-39ee27c35f43.jpg)

> * request<sub>client</sub>
>   1. RP accepts a user request & passes the traffic to both CrS & SrS via request<sub>main</sub> & request<sub>mirror</sub> respectively.

> * request<sub>main</sub>
>   1. CrS accepts request & asks MpS which district the request's origin point belongs to.
>   1. CrS queries its cache for the coefficient value of the district number.  
>   1. CrS calculates price, and returns the results to RP.  

> * request<sub>mirror</sub>
>   1. Mediator accepts request & pushes it into the MQ asap.
>   1. SrS pulls request data from MQ & asks MpS which district the request's origin point belongs to.
>   1. SrS inserts a record into its DB for the district number plus timestamp.
>   1. SrS counts DB records for the district number during a specific time-window and calculates coefficient value.
>   1. SrS informs CrS in case of any fluctuations in coefficient value based on its previously cached value; And when receives ack response, updates its cache based on new calculated value.
