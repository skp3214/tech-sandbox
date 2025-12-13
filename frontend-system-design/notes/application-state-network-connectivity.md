# Application State and Network Connectivity

## Table of Contents

- [1. Application State Design](#1-application-state-design)



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
