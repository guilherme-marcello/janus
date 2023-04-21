# Janus WebRTC Client

## Overview

The Janus Client is a Go library that provides a convenient way to interact with the Janus WebRTC Gateway API. It is designed to be used in Go projects for building applications that require real-time communication using the Janus Gateway.

The library is organized into different packages based on the functionality they provide. The main packages in the Janus Client are:

- `elements`: Contains Go files related to Janus recording and streaming, including structs and functions for handling recording and streaming operations.
- `plugin`: Contains plugin-agnostic structures and functions for attaching plugins to a session. It also includes specific implementations for Janus plugins like record and play, and streaming, with their corresponding structs and functions.
- `requests`: Contains Go files for handling HTTP requests and parsing response payloads related to sessions and plugins.
- `session`: Provides generic session-related structures and methods for creating, destroying, and keeping alive sessions using the HTTP requests available in the `requests` package.
- `janus.go`: Defines the `Http` struct, which specifies the Janus HTTP API endpoint and is used by other structs in the library for making API requests.

Please note that currently, there are no interfaces available in the Janus Client. The only way to interact with the Janus API is through the HTTP transport. However, future developments in the open-source project may include the creation of interfaces to create a more generic and Janus-transport agnostic code.

## Package details

### `elements` 

The `elements` package contains structs related to Janus recording and streaming. It includes the following files:

- `elements.go`: This file contains common structs used in the Janus Client.
- `recording.go`: This file contains structs related to Janus recording, such as the `Recording` struct.
- `streaming.go`: This file contains structs related to Janus streaming, such as the `Mountpoint` struct.

### `plugin` 

The `plugin` package contains plugin-agnostic structures and functions, as well as specific implementations for Janus plugins, such as janus.plugin.recordplay and janus.plugin.streaming. It includes the following files:

- `plugin.go`: This file contains plugin-agnostic structures and functions, such as attaching a plugin to a session.
- `recordnplay.go`: This file contains the `RecordPlay` struct, which is a specific implementation for the janus.plugin.recordplay plugin. It is composed of the `Plugin` struct from `plugin.go` and specifies functions for janus.plugin.recordplay specific usages, such as listing available recordings.
- `streaming.go`: This file contains the `Streaming` struct, which is a specific implementation for the janus.plugin.streaming plugin. It is composed of the `Plugin` struct from `plugin.go` and specifies functions for janus.plugin.streaming specific usages, such as listing available mountpoints to watch.

### `requests`

The `requests` package contains utility functions for sending HTTP requests and parsing responses. It includes the following files:

- `session.go`: This file contains generic session-related answer JSON models and request payloads, such as keeping a session alive, attaching a plugin to a session, creating and destroying a session.
- `plugin_streaming.go`: This file contains specific requests for the janus.plugin.streaming plugin, such as starting and stopping a stream, creating and destroying a mountpoint.
- `plugin_recordnplay.go`: This file contains specific requests for the janus.plugin.recordplay plugin, such as starting and stopping a recording, getting recording status, and setting recording options.

### `session`

The `session` package contains generic session-related structures and provides methods to create, destroy, and keep alive sessions using the requests available in the `requests` package.

### `janus`

The `janus` package contains the definition of the `Http` struct, which specifies the Janus HTTP API endpoint. All subsequent structs in the library use this `Http` struct, as all requests are made to a specific Janus API endpoint.
