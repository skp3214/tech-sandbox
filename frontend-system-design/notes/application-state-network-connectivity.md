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

## [3. Server-Sent Events (SSE)]()

### What is SSE?

Server-Sent Events is a standard for **one-way, server-to-client** real-time communication over HTTP. The server can push updates to the client over a single, long-lived HTTP connection.

**Key Characteristics:**
- Unidirectional (server → client only)
- Built on standard HTTP/HTTPS
- Automatic reconnection on disconnect
- Simple text-based protocol
- Native browser support via `EventSource` API

### How It Works

```javascript
// Client-side
const eventSource = new EventSource('/api/stream');

eventSource.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log('Received:', data);
};

// Listen to specific event types
eventSource.addEventListener('update', (event) => {
  console.log('Update:', event.data);
});

eventSource.onerror = (error) => {
  console.error('SSE Error:', error);
};

// Close connection when done
eventSource.close();
```

```javascript
// Server-side (Node.js/Express)
app.get('/api/stream', (req, res) => {
  res.setHeader('Content-Type', 'text/event-stream');
  res.setHeader('Cache-Control', 'no-cache');
  res.setHeader('Connection', 'keep-alive');
  
  // Send data every 2 seconds
  const interval = setInterval(() => {
    res.write(`data: ${JSON.stringify({ time: new Date() })}\n\n`);
  }, 2000);
  
  // Cleanup on close
  req.on('close', () => {
    clearInterval(interval);
  });
});
```

### When to Use SSE

**Good For:**
- **Live Notifications:** Real-time alerts, news feeds
- **Stock Tickers:** Continuous price updates
- **Social Media Feeds:** New posts, likes, comments
- **Progress Updates:** File uploads, background tasks
- **Live Dashboards:** Metrics, analytics, monitoring
- **Chat Applications:** Read-only message streams

**Not Good For:**
- Bidirectional communication (use WebSockets)
- Binary data transfer (text-only)
- IE/Edge legacy support (no native support)
