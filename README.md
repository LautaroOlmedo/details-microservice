# Microservice Details API

## This API provides the necessary product details to the frontend.

**Version:** 1.0.0

### How to Run

### Prerequisites
* Go 1.22 or higher
* Docker (optional, for running the service in a container)
* Make (optional, for running commands)




**Technology Stack:** Go 1.24.2. Standard Library, mockery for testing.

### Architecture

This microservice follows a Hexagonal Architecture (also known as Ports and Adapters) to ensure a clear separation of concerns and make the application independent of external services.

**Domain Layer:** This layer contains the core business logic and entities.

**Application Layer (Services):** This layer contains the use cases and interfaces that define the input and output ports of the system.

**Infrastructure Layer:** This layer holds the implementations of the output ports, such as the Details repository. This is also where you would add implementations for message queues, middlewares, and other external integrations.


### Implemented Strategy

This service assumes that the requested ID is shared between the product and its details. It uses CQRS (Command Query Responsibility Segregation) to separate read and write operations. The Repository Pattern is used to abstract data access logic, with in-memory Map and Slice data structures serving as the current data persistence layer. Use injection dependency to decouple the service implementation and the repository implementation


### Calling endpoints

Get Details



![img.png](img.png)



### API Endpoint Design

We assume that all incoming request contain a detail identifier in the Details-ID header.

 * Endpoint ```GET /api/v1/details/```
 * Header 
```
Details-ID: "550e8400-e29b-41d4-a716-446655440000"
```
* Success Response

```
{
  "ID": "550e8400-e29b-41d4-a716-446655440000",
  "Description": "Sample description",
  "CreatedAt": "2025-08-08T09:25:31.095446-03:00",
  "UpdatedAt": "2025-08-08T09:25:31.095446-03:00"
}
```

Response Code Errors
```
200 Ok
400 Bad Request
404 Not Found
500 Internal Server Error
```


**/ping**

Method: GET
Description: A health check endpoint that returns "pong" to verify the microservice is running.

### Future Enhancements

**Communication with the Products Service:** A validation could be implemented to check if the product exists. This could be done by making a call to the products microservice using a pattern like the Circuit Breaker.

**Caching:** A cache could be implemented to temporarily store requested resources, reducing the number of database queries.

**Metrics (CPU, RAM, Requests per Second):** A metrics collection system, such as Prometheus, could be integrated to monitor the microservice's performance, including CPU usage, RAM, and requests per second.
