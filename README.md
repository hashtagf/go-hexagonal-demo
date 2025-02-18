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

## Create folder structure

- \_cmd
  - main.go
- internal
  - adapters
    - inbound
    - outbound
- cores
  - domain
  - ports
  - services
- utils
  - utils other all
- middlewares
