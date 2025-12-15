# Application State and Network Connectivity

## Table of Contents

- [1. Application State Design](#1-application-state-design)
- [2. Network Connectivity](#2-network-connectivity)
- [3. Server-Sent Events (SSE)](#3-server-sent-events-sse)
- [4. Web Sockets](#4-websockets)
- [5. Classic REST and GraphQL](#5-classic-rest-and-graphql)



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

### [How It Works]()

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

### [When to Use SSE]()

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

## [4. WebSockets]()

### What are WebSockets?

WebSockets provide **full-duplex, bidirectional** communication between client and server over a single, persistent TCP connection. Unlike HTTP's request-response model, both sides can send messages independently at any time.

**Key Characteristics:**
- Bidirectional (client ↔ server)
- Persistent connection
- Low latency (~1-2ms overhead)
- Supports text and binary data
- Uses `ws://` or `wss://` protocol
- Native browser support via `WebSocket` API

### [How It Works]()

**Connection Upgrade:**
1. Client initiates HTTP request with `Upgrade: websocket` header
2. Server responds with `101 Switching Protocols`
3. Connection upgraded to WebSocket protocol
4. Both sides can now send/receive messages freely

### [When to Use WebSockets]()

**Perfect For:**
- **Real-time Chat:** Instant messaging, group chats
- **Multiplayer Games:** Low-latency game state sync
- **Collaborative Editing:** Google Docs-style collaboration
- **Live Trading:** Stock/crypto trading platforms
- **IoT Dashboards:** Real-time sensor data
- **Live Sports Scores:** Instant score updates
- **Video Conferencing:** Signaling for WebRTC

**Not Ideal For:**
- One-way server updates (use SSE instead)
- Occasional updates (use polling or SSE)
- Static content fetching (use HTTP)

## [5. Classic REST and GraphQL]()

### REST (Representational State Transfer)

REST is an architectural style for building APIs using standard HTTP methods and resource-based URLs. It's the traditional approach for client-server communication on the web.

**Core Principles:**
- **Resources:** Everything is a resource with a unique URI
- **HTTP Methods:** GET, POST, PUT, PATCH, DELETE
- **Stateless:** Each request contains all needed information
- **Standard Status Codes:** 200, 404, 500, etc.

**Example:**
```javascript
// Fetch user
GET /api/users/123
Response: { id: 123, name: "Alice", email: "alice@example.com" }

// Fetch user's posts
GET /api/users/123/posts
Response: [{ id: 1, title: "Post 1" }, { id: 2, title: "Post 2" }]

// Fetch post comments
GET /api/posts/1/comments
Response: [{ id: 1, text: "Great!" }, { id: 2, text: "Nice!" }]

// Create new post
POST /api/posts
Body: { title: "New Post", content: "..." }
Response: { id: 3, title: "New Post", ... }

// Update user
PUT /api/users/123
Body: { name: "Alice Smith", email: "alice@example.com" }

// Partial update
PATCH /api/users/123
Body: { name: "Alice Smith" }

// Delete post
DELETE /api/posts/1
Response: 204 No Content
```

**Advantages:**
✅ Simple and well-understood
✅ Cacheable (HTTP caching)
✅ Stateless and scalable
✅ Language/platform agnostic
✅ Wide tooling support

**Problems:**
❌ **Over-fetching:** Get more data than needed
❌ **Under-fetching:** Need multiple requests for related data
❌ **N+1 Problem:** Fetch user → fetch posts → fetch comments (waterfall)
❌ **Versioning:** Breaking changes require API versioning
❌ **Fixed Structure:** Server decides response shape

### GraphQL

GraphQL is a query language that lets clients request exactly the data they need in a single request, with a flexible, strongly-typed schema.

**Core Concepts:**
- **Single Endpoint:** All queries go to `/graphql`
- **Client-Specified Queries:** Client controls response structure
- **Strongly Typed:** Schema defines all types and operations
- **Introspection:** API is self-documenting

**Example:**
```javascript
// Single query for user, posts, and comments
POST /graphql
Body: {
  query: `
    query {
      user(id: 123) {
        name
        email
        posts {
          title
          comments {
            text
            author {
              name
            }
          }
        }
      }
    }
  `
}

Response: {
  data: {
    user: {
      name: "Alice",
      email: "alice@example.com",
      posts: [
        {
          title: "Post 1",
          comments: [
            { text: "Great!", author: { name: "Bob" } },
            { text: "Nice!", author: { name: "Charlie" } }
          ]
        }
      ]
    }
  }
}

// Mutation (create/update/delete)
mutation {
  createPost(title: "New Post", content: "...") {
    id
    title
  }
}

// Request only specific fields
query {
  user(id: 123) {
    name  # Only fetch name, skip email
  }
}
```

**Schema Definition (Server-side):**
```graphql
type User {
  id: ID!
  name: String!
  email: String!
  posts: [Post!]!
}

type Post {
  id: ID!
  title: String!
  content: String!
  author: User!
  comments: [Comment!]!
}

type Comment {
  id: ID!
  text: String!
  author: User!
}

type Query {
  user(id: ID!): User
  posts: [Post!]!
}

type Mutation {
  createPost(title: String!, content: String!): Post!
  updateUser(id: ID!, name: String): User!
}
```

**Advantages:**
✅ **No Over/Under-fetching:** Get exactly what you need
✅ **Single Request:** Fetch related data in one query
✅ **Strong Typing:** Type safety and validation
✅ **Self-Documenting:** Schema serves as documentation
✅ **Flexible:** Client controls response shape
✅ **Versioning-Free:** Add fields without breaking changes

**Challenges:**
❌ **Complexity:** Steeper learning curve
❌ **Caching:** HTTP caching is harder (all POST to /graphql)
❌ **Query Cost:** Complex queries can be expensive
❌ **Over-Querying:** Malicious/accidental expensive queries
❌ **File Uploads:** Not built-in, needs workarounds
❌ **Backend Complexity:** Resolver implementation can be complex

### REST vs GraphQL Comparison

| Feature | REST | GraphQL |
|---------|------|---------|
| **Endpoints** | Multiple (`/users`, `/posts`) | Single (`/graphql`) |
| **Data Fetching** | Fixed by server | Client-specified |
| **Over-fetching** | Common | No |
| **Under-fetching** | Common (N+1) | No |
| **Caching** | HTTP caching (easy) | Custom caching (complex) |
| **Learning Curve** | Easy | Moderate |
| **Versioning** | Required (v1, v2) | Not needed |
| **Real-time** | Requires WebSocket | Subscriptions built-in |
| **Tooling** | Mature | Growing |
| **Best For** | Simple CRUD, public APIs | Complex data, mobile apps |

### When to Use Each

**Use REST when:**
- Building simple CRUD APIs
- Need standard HTTP caching
- Team unfamiliar with GraphQL
- Public API for third parties
- Simple data relationships

**Use GraphQL when:**
- Mobile apps (reduce data transfer)
- Complex, nested data requirements
- Multiple clients with different needs
- Rapid frontend iteration
- Need real-time updates (subscriptions)

### Hybrid Approach

Many teams use both:
```javascript
// REST for simple operations
GET /api/auth/login
POST /api/files/upload

// GraphQL for complex data fetching
POST /graphql
query { dashboard { user { ... }, stats { ... } } }
```

**Best Practice:** Start with REST, migrate to GraphQL when:
- Multiple related API calls slow down app
- Over-fetching wastes bandwidth
- Different clients need different data shapes
- Frontend team wants more control