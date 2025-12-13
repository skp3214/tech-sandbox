## Indexing and Search Optimization

Indexes are additional data structures that improve search performance at the cost of extra memory and maintenance.

#### Types of Indexes

[**1. Primary Index (Direct Access by ID)**]()
```javascript
const usersById = {
  1: { id: 1, name: "Alice", email: "alice@example.com" },
  2: { id: 2, name: "Bob", email: "bob@example.com" }
};
```

[**2. Secondary Indexes (Search by Other Properties)**]()
```javascript
const state = {
  users: {
    1: { id: 1, name: "Alice", email: "alice@example.com" },
    2: { id: 2, name: "Bob", email: "bob@example.com" }
  },
  
  // Index by email for fast email lookup
  usersByEmail: {
    "alice@example.com": 1,
    "bob@example.com": 2
  },
  
  // Index by name (if names can be duplicated, store arrays)
  usersByName: {
    "Alice": [1],
    "Bob": [2]
  }
};

// Fast email search - O(1)
const userId = state.usersByEmail["alice@example.com"];
const user = state.users[userId];
```

[**3. Composite Indexes (Multiple Properties)**]()
```javascript
const state = {
  products: {
    501: { id: 501, name: "Laptop", categoryId: 10, inStock: true },
    502: { id: 502, name: "Mouse", categoryId: 10, inStock: false }
  },
  
  // Composite index: category + stock status
  productsByCategoryAndStock: {
    "10_true": [501],   // Category 10, in stock
    "10_false": [502],  // Category 10, out of stock
    "20_true": [503]
  }
};

// Fast query: Get in-stock products from category 10
const key = `10_true`;
const productIds = state.productsByCategoryAndStock[key];
```

[**4. Full-Text Search Indexes**]()
```javascript
// Inverted index for text search
const state = {
  products: {
    501: { id: 501, name: "Gaming Laptop", description: "High-performance laptop" },
    502: { id: 502, name: "Wireless Mouse", description: "Ergonomic mouse" }
  },
  
  // Inverted index: word -> product IDs
  searchIndex: {
    "gaming": [501],
    "laptop": [501],
    "high": [501],
    "performance": [501],
    "wireless": [502],
    "mouse": [502],
    "ergonomic": [502]
  }
};

// Search for "laptop"
const results = state.searchIndex["laptop"].map(id => state.products[id]);
```
