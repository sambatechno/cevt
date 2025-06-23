# Cata Event

Centralized protobuf definition fot events.


# Development setup

## Prerequisites

- Go 1.21
- `buf` cli (https://buf.build/docs/cli/installation/)

## Setup

1. `buf dep update`

## Generating code

1. After doing change to the proto files, run `buf generate`.
2. Sometime `gopls` would have errors after generate, just reload the IDE window.