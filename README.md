## Golang test - NodeArt

# By Adailson

## Installation & Run

First, running this API

1 - Create environment

```bash
cp env.example .env
```

2 - Docker running this API (my recommenation)

```bash
docker compose up --build
```

## Decision

## Technical Decisions and Justifications

1. Choosing Fiber as the Framework:
   We chose Fiber due to its high performance and simplicity. Built on Fasthttp, one of the fastest HTTP libraries for Go, Fiber provides a lightweight and efficient framework for web development. Its intuitive APIs, ease of use, and native support for middlewares were crucial factors in meeting the project's performance and productivity requirements.

2. Using Clean Architecture:
   The decision to adopt Clean Architecture was driven by the need to create a scalable, maintainable, and decoupled system. This approach divides system responsibilities into well-defined layers (such as domain, use cases, interfaces, and infrastructure), promoting:

Flexibility: easy to adapt and evolve the system without affecting unrelated parts.
Testability: with decoupled responsibilities, each layer can be tested independently.
Sustainability: ensures the system remains robust and ready for future enhancements.

3. Multistage Docker Configuration:
   We implemented two separate Dockerfiles for different purposes:

Multistage for Production: reduces the final image size by excluding unnecessary dependencies, improving performance and security in production. This ensures that only the essential binaries and resources are included, resulting in a lightweight and efficient image.
Standard Docker for Development: includes all dependencies and tools needed to facilitate debugging and updates during development.
This approach balances efficiency in production with flexibility during development.

4. Absence of Tests:
   Although we understand the importance of testing for ensuring software quality, tests were not implemented due to time constraints and limited experience with testing in Go. However, prior experience with other languages like Node.js and PHP, where unit and integration tests were successfully developed, reflects a commitment to code quality. Moving forward, creating tests will be prioritized as part of continuous learning and adopting best practices for testing in Go.
