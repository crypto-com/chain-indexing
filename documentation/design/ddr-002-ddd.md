# Domain-Driven Development

## Changelog

Nov 1st, 2020: Created the record template

## Design

### System

The domain-driven system is divided two major systems:
1. Command Execution
2. Projection

#### Command Execution

Command execution concerns about the command and outcomes. It is a flow of command creation, business logic execution to yield the events and events persistance.

#### Projection

Projection is the accumulated representation of the concerned events. It listens to interested events from command execution and produce a state based on it.
