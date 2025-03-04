# go-hexagonal-demo

## Hexagonal Architecture Components

### Core

- services
- domain
- ports

### Services

- business logic

### Domain

- models
- rules

### Ports

- Primary Ports (Inbound)
- Secondary Ports (Outbound)

### Adapters

- Primary Adapters (Inbound)
- Secondary Adapters (Outbound)

### Create Diagram Folder Structure

```mermaid
graph TD
    Root[go-hexagonal-demo] --> Internal[internal/]
    Root --> CMD[cmd/]
    Root --> Utils[utils/]
    Root --> Middlewares[middlewares/]
    Root --> Config[go.mod/go.sum]

    Internal --> Adapters[adapters/]
    Internal --> Core[core/]

    Adapters --> Inbound[inbound/]
    Adapters --> Outbound[outbound/]

    Core --> Domain[domains/]
    Core --> Ports[ports/]
    Core --> Services[services/]

    CMD --> Main[main.go]

    %% Styling
    classDef default fill:#f9f9f9,stroke:#333,stroke-width:1px
    classDef root fill:#e1f5fe,stroke:#01579b,stroke-width:2px
    classDef major fill:#fff3e0,stroke:#ff6f00,stroke-width:2px
    classDef module fill:#f3e5f5,stroke:#4a148c,stroke-width:1px

    class Root root
    class Internal,Adapters,Core major
    class CMD,Utils,Middlewares,Config module
```

### Hexagonal Architecture Diagram

```mermaid
graph TB
    subgraph External World
        REST[REST API]
        DB[(Database)]
    end

    subgraph Hexagonal Architecture
        subgraph Adapters
            direction TB
            subgraph "Inbound Adapters"
                REST --> |HTTP Requests| HTTP[HTTP Handler]
            end

            subgraph "Outbound Adapters"
                REPO[Repository Implementation] --> DB
            end
        end

        subgraph Core
            direction TB
            subgraph "Ports"
                IN[Inbound Ports/Interfaces]
                OUT[Outbound Ports/Interfaces]
            end

            subgraph "Domain"
                DOM[Domain Models]
            end

            subgraph "Services"
                SVC[Business Logic/Services]
            end
        end

        %% Connections
        HTTP --> IN
        IN --> SVC
        SVC --> DOM
        SVC --> OUT
        OUT --> REPO
    end

    classDef core fill:#f9f,stroke:#333,stroke-width:2px
    classDef adapter fill:#bbf,stroke:#333,stroke-width:2px
    classDef external fill:#ddd,stroke:#333,stroke-width:2px

    class DOM,SVC core
    class HTTP,REPO adapter
    class REST,DB external
```
