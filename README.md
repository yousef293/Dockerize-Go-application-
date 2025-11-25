# Dockerize-Go-application-

Features

## RPC-based communication using net/rpc

Thread-safe message handling with sync.Mutex

In-memory chat history

Supports multiple concurrent clients

Listens on TCP port 1234

## How It Works

The server stores all messages in memory. Each RPC call:

Locks the message slice.

Appends the new formatted message.

Returns the full chat history.
