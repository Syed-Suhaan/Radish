# Radish 🥕

Radish is a lightweight, high-performance, in-memory key-value store built from scratch in **Go**. It implements the **Redis Serialization Protocol (RESP)**, making it compatible with standard Redis clients and tools.

> **Status:** Active Development. Core TCP networking and RESP parsing are operational. Currently implementing storage primitives and command dispatching.

---

## ⚡ Technical Overview

Radish is engineered to handle concurrent connections efficiently without relying on external frameworks. It interacts directly with raw TCP sockets and manages state using Go's native concurrency primitives.

### Core Features
* **Custom Protocol Implementation:** Features a hand-written **RESP parser** capable of deserializing arrays, bulk strings, and simple strings from raw byte streams.
* **Concurrent Architecture:** Utilizes a "goroutine-per-connection" model to handle multiple simultaneous clients non-blockingly.
* **Zero Dependencies:** Built entirely using the Go Standard Library (`net`, `sync`, `bufio`) to minimize overhead.

---

## 🗺️ Development Roadmap

### ✅ v0.1: Networking & Parsing (Completed)
* [x] **TCP Server:** Custom implementation using `net.Listen` to accept raw socket connections.
* [x] **Connection Handling:** Asynchronous client management using Goroutines.
* [x] **Protocol Parser:** Deserializer for the RESP format (reading `+OK`, `$BulkString`, `*Arrays`).
* [x] **Response Serializer:** Writers to format Go structures back into RESP bytes for the client.

### 🚧 v0.2: Storage Engine (In Progress)
* [x] **Command Dispatcher:** Routing logic to map text commands to internal Go functions.
* [x] **Basic Commands:** `PING`, `ECHO`.
* [ ] **In-Memory Store:** Implementing the global hash map implementation for `SET` and `GET`.
* [ ] **Thread Safety:** Integrating `sync.RWMutex` to ensure atomic operations on shared data structures.

### 🔮 v0.3: Advanced Features (Planned)
* [ ] **TTL & Eviction:** Background workers (`time.Ticker`) to handle key expiration.
* [ ] **Persistence:** Append-Only File (AOF) logging to ensure data durability across restarts.

---

## 🛠 Architecture

Radish follows a clean, modular architecture:

1.  **The Listener:** Accepts TCP connections on port `6379`.
2.  **The Handler:** Spawns a Goroutine for each connection to ensure isolation.
3.  **The Parser:** Reads from the socket buffer and constructs a Command object.
4.  **The Executor:** Executes the command against the data store and returns the result.

---

## 🚀 Getting Started

### Prerequisites
* Go 1.18+
* `redis-cli` (optional, for testing)

### Installation
```bash
git clone [https://github.com/Syed-Suhaan/Radish.git](https://github.com/Syed-Suhaan/Radish.git)
cd Radish
go run main.go
