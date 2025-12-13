## Memory Management

Efficient memory management is critical for frontend applications, especially those dealing with large datasets.

### [Memory Offloading]()

**Technique:** Move data from RAM to persistent storage when not immediately needed.


#### Browser Storage Solutions

Different storage mechanisms for different use cases based on data size, persistence requirements, and access patterns.

#### Comparison Table

| Storage | Size Limit | Persistence | Synchronous | Use Case |
|---------|-----------|-------------|-------------|----------|
| **LocalStorage** | ~5-10 MB | Permanent | Yes | Small persistent data, user preferences |
| **SessionStorage** | ~5-10 MB | Session only | Yes | Temporary session data, form drafts |
| **IndexedDB** | ~50% of disk | Permanent | No (async) | Large datasets, offline-first apps |


#### [LocalStorage]()

**Best For:**
- User preferences and settings
- Small persistent data (< 1 MB)
- Simple key-value pairs
- Data that needs to survive browser restarts

**Example:**
```javascript
// Store user preferences
const preferences = {
  theme: 'dark',
  language: 'en',
  fontSize: 14
};

localStorage.setItem('userPreferences', JSON.stringify(preferences));

// Retrieve
const stored = JSON.parse(localStorage.getItem('userPreferences'));
```

**Limitations:**
- Synchronous (blocks main thread)
- String values only (need JSON serialization)
- Small size limit (~5-10 MB)
- No expiration mechanism

#### [SessionStorage]()

**Best For:**
- Temporary session-specific data
- Form data during multi-step processes
- Data that should not persist across tabs
- Wizard/stepper state

**Example:**
```javascript
// Multi-step form
class FormWizard {
  saveStep(stepNumber, data) {
    const formData = this.getFormData() || {};
    formData[`step${stepNumber}`] = data;
    sessionStorage.setItem('wizardData', JSON.stringify(formData));
  }
  
  getFormData() {
    const data = sessionStorage.getItem('wizardData');
    return data ? JSON.parse(data) : null;
  }
  
  clearForm() {
    sessionStorage.removeItem('wizardData');
  }
}

// Usage
const wizard = new FormWizard();
wizard.saveStep(1, { name: 'Alice', email: 'alice@example.com' });
wizard.saveStep(2, { address: '123 Main St', city: 'NYC' });
```

**Use Cases:**
- Shopping cart for single session
- Form auto-save during navigation
- Temporary authentication tokens
- Tab-specific UI state

#### [IndexedDB]()

**Best For:**
- Large datasets (> 10 MB)
- Structured data with complex queries
- Offline-first applications
- Client-side database operations

**Example:**
```javascript
// Open database
const dbPromise = indexedDB.open('myDatabase', 1);

dbPromise.onupgradeneeded = (event) => {
  const db = event.target.result;
  // Create object store
  const store = db.createObjectStore('posts', { keyPath: 'id' });
  store.createIndex('category', 'category', { unique: false });
};

dbPromise.onsuccess = (event) => {
  const db = event.target.result;
  
  // Add data
  const tx = db.transaction('posts', 'readwrite');
  const store = tx.objectStore('posts');
  store.add({ id: 1, title: 'Post 1', category: 'tech' });
  
  // Get data
  const getRequest = store.get(1);
  getRequest.onsuccess = () => console.log(getRequest.result);
  
  // Query by index
  const index = store.index('category');
  const query = index.getAll('tech');
  query.onsuccess = () => console.log(query.result);
};
```

**Use Cases:**
- Caching API responses for offline use
- Storing user-generated content (drafts, notes)
- Large product catalogs
- Email clients (storing messages)
- Media files and blobs


#### Summary: Choosing the Right Storage

**Decision Tree:**

1. **Data Size < 1 MB + Simple Key-Value + Persistent?**
   → Use **LocalStorage**

2. **Data Size < 1 MB + Temporary (Session Only)?**
   → Use **SessionStorage**

3. **Data Size > 1 MB OR Complex Queries + Persistent?**
   → Use **IndexedDB**


### [Sharding]()

**Technique:** Split data across multiple stores to reduce the size of individual data structures.

**Use Cases:**
- Very large datasets that don't fit efficiently in memory
- Isolating data by user, tenant, or time period
- Distributing load across different storage mechanisms

**Example: Time-based Sharding**
```javascript
const state = {
  // Shard orders by year-month
  orders: {
    "2024-01": { /* orders from Jan 2024 */ },
    "2024-02": { /* orders from Feb 2024 */ },
    "2024-03": { /* orders from Mar 2024 */ }
  },
  
  // Only keep recent shards in memory
  activeShards: ["2024-02", "2024-03"],
  
  // Older shards in IndexedDB
  archivedShards: ["2024-01", "2023-12", "2023-11"]
};

async function getOrder(orderId, date) {
  const shard = formatYearMonth(date); // "2024-03"
  
  // Check if shard is in memory
  if (state.activeShards.includes(shard)) {
    return state.orders[shard][orderId];
  }
  
  // Load from IndexedDB
  return await loadFromArchive(shard, orderId);
}
```

**Example: User-based Sharding**
```javascript
// For multi-tenant applications
const state = {
  userDataShards: {
    "user-123": { /* user 123's data */ },
    "user-456": { /* user 456's data */ }
  },
  
  currentUserId: "user-123"
};

// Only load current user's data into memory
function switchUser(userId) {
  // Offload previous user's data
  offloadToIndexedDB(state.currentUserId, state.userDataShards[state.currentUserId]);
  delete state.userDataShards[state.currentUserId];
  
  // Load new user's data
  state.userDataShards[userId] = await loadFromIndexedDB(userId);
  state.currentUserId = userId;
}
```