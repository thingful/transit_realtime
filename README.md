# transit_realtime

Golang binding generated from the
[GTFS-realtime](https://developers.google.com/transit/gtfs-realtime/) protocol
buffer spec.

## Introduction

[GTFS-realtime](https://developers.google.com/transit/gtfs-realtime/) is a data
format for communicating real-time information about public transit systems.
GTFS-realtime data is encoded and decoded using [Protocol
Buffers](https://developers.google.com/protocol-buffers/), a compact binary
representation designed for fast and efficient processing.  The data schema
itself is defined in
[gtfs-realtime.proto](https://developers.google.com/transit/gtfs-realtime/gtfs-realtime-proto).

To work with GTFS-realtime data, a developer would typically use the
`gtfs-realtime.proto` schema to generate classes in the programming language of
their choice.  These classes can then be used for constructing GTFS-realtime
data model objects and serializing them as binary data or, in the reverse
direction, parsing binary data into data model objects.

This library simply contains a generated binding for Go, generated from the
version of the GTFS-realtime protocol buffer spec included here.

A number of other pre-generated bindings are provided by Google here:
https://github.com/google/gtfs-realtime-bindings

## Updating

To update the bindings if the protocol buffer spec changes, you need to:

1. Install the protocol buffer compiler (`protoc`). This binding was generated using version `2.6.1`.
2. Regenerate the language bindings using protoc:

    ```bash
    protoc --go_out=. --proto_path=. ./gtfs-realtime.proto
    ```

3. Run the tests:

    ```bash
    go test
    ```

## Build status

[![wercker status](https://app.wercker.com/status/b3cc161ff9667edc85ca273f0286cf3c/s/master "wercker status")](https://app.wercker.com/project/bykey/b3cc161ff9667edc85ca273f0286cf3c)
