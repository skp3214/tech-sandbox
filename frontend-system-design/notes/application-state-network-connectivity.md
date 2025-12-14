# Application State and Network Connectivity

## Table of Contents

- [1. Application State Design](#1-application-state-design)
- [2. Network Connectivity](#2-network-connectivity)



## [1. Application State Design]()

Application state design is crucial for building performant, scalable, and maintainable frontend applications. The key is to structure your data efficiently while following core principles that optimize access patterns, search operations, and memory usage.

### [General Principles]()

Three fundamental principles guide efficient state design:

#### [1. Minimize Access Cost]()

**Goal:** Reduce the time and computational steps needed to retrieve data.

**Strategies:**
- **Use normalization:** Store entities separately and use references (IDs) instead of embedding.

`CLICK HERE` -  [Normalization](/frontend-system-design/notes/normalization.md)




#### [2. Minimize Search Cost]()

**Goal:** Enable fast querying and filtering of data.

**Strategies:**
- **Create indexes:** Maintain additional data structures for common queries
- **Pre-sort data:** Keep data sorted if order matters for queries

`CLICK HERE` - [Indexing](/frontend-system-design/notes/indexing.md)
#### [3. Minimize RAM Usage]()

**Goal:** Keep memory footprint small, especially for large datasets.

**Strategies:**
- **Offload to persistent storage:** Use IndexedDB for large datasets
- **Implement data pruning:** Remove stale or unused data periodically

`Click Here` - [Memory Management](/frontend-system-design/notes/memory-browser.md)

## [2. Network Connectivity]()

### Transport Protocols: TCP vs UDP

[**TCP (Transmission Control Protocol):**]()
- Connection-oriented, reliable delivery
- Guarantees packet order and delivery
- Higher overhead due to acknowledgments
- Used by: HTTP, WebSockets
- Best for: Data integrity matters (web pages, APIs)

[**UDP (User Datagram Protocol):**]()
- Connectionless, no delivery guarantee
- Faster with lower latency
- No packet order guarantee
- Used by: WebRTC, live streaming
- Best for: Speed matters more than reliability (video calls, gaming)

### [Long Polling: Problems and Limitations]()

**How It Works:**
Client sends HTTP request → Server holds connection open → Server responds when data available → Client immediately sends new request

**Problems:**

[**1. High Latency:**]()
- Request/response cycle delay (100-500ms overhead)
- Connection establishment time for each request
- Server processing time adds up
- Not suitable for real-time applications

[**2. Energy Consumption:**]()
- Constant HTTP requests drain mobile battery
- Keeps radio awake continuously on mobile devices
- CPU cycles wasted on frequent connections
- Network interface stays active

**Better Alternatives:**
- **WebSockets:** Full-duplex, persistent connection, low latency
- **Server-Sent Events (SSE):** Unidirectional, simple, efficient for server-to-client
- **WebRTC:** Peer-to-peer, ultra-low latency for real-time communication