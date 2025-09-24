# Event Service - Clean Architecture with Go and RabbitMQ 🐇

This project is a simple example of **clean architecture** applied to **event and queue** management with **RabbitMQ** in Go.  

The main idea is to separate responsibilities into layers so that the infrastructure (RabbitMQ) does not contaminate the business logic.  

---

## 🏗️ Project Architecture

```bash
├── cmd
│   ├── consumer
│   │   └── main.go            # Consumer entry point
│   └── publisher
│       └── main.go            # Publisher entry point
│
├── internal
│   ├── domain
│   │   ├── event_repository.go # Event repository contract (interface)
│   │   └── event.go            # Domain entity (Event)
│   │
│   ├── infrastructure
│   │   └── rabbitmq
│   │       ├── connection.go   # Configuration and connection with RabbitMQ
│   │       └── publisher.go    # Implementation of the repository using RabbitMQ
│   │
│   └── usecase
│       └── event_usecase.go    # Use case: publish event
│
├── go.mod
├── go.sum
└── README.md
```

## 📚 Layers explained

- Domain → Defines the entities and contracts (interfaces) that represent the core of the business.
- Use case → Implements business logic. Depends only on domain.
- Infrastructure → Implements specific adapters (e.g., RabbitMQ). Depends on domain.
- cmd → Entry points for running the program (publisher/consumer).

## 🚀 How to launch the project
1. Using the run.sh script (recommended)

```sh
chmod +x run.sh
./run.sh

```

### This will:

- Build Docker images.

- Set up RabbitMQ + Publisher + Consumer.

- RabbitMQ UI available at 👉 http://localhost:15672


## 🧩 System flow

- The Publisher creates an event and sends it to a queue in RabbitMQ.
- The Consumer listens to that queue and processes the received event.
- The domain layer defines what an Event is.
- The use case layer orchestrates the publication of events.
- The infrastructure (RabbitMQ) is just a technical detail, easily replaceable.