## Database Normalization

Database normalization principles apply to frontend state design. Understanding **Normal Forms (NF)** helps create efficient state structures.

#### [First Normal Form (1NF)]()
- **Rule:** Each field contains atomic (indivisible) values
- **Rule:** No repeating groups or arrays within a record

**Example:**
```javascript
// Violates 1NF - phone numbers in array
const user1NF_Bad = {
  id: 1,
  name: "Alice",
  phones: ["123-456-7890", "098-765-4321"] // Array of phones
};

// Follows 1NF - separate records for each phone
const users1NF = {
  1: { id: 1, name: "Alice" }
};
const userPhones = {
  1: { userId: 1, phone: "123-456-7890", type: "mobile" },
  2: { userId: 1, phone: "098-765-4321", type: "home" }
};
```

#### [Second Normal Form (2NF)]()
- **Rule:** Must be in 1NF
- **Rule:** All non-key attributes must depend on the entire primary key

**Example:**
```javascript
// Violates 2NF - category name depends only on categoryId, not on product id
const products2NF_Bad = {
  501: { 
    id: 501, 
    name: "Laptop", 
    categoryId: 10, 
    categoryName: "Electronics" // Redundant
  }
};

// Follows 2NF - separate categories
const products2NF = {
  501: { id: 501, name: "Laptop", categoryId: 10 }
};
const categories = {
  10: { id: 10, name: "Electronics" }
};
```

#### [Third Normal Form (3NF)]()
- **Rule:** Must be in 2NF
- **Rule:** No transitive dependencies (non-key attributes should not depend on other non-key attributes)

**Example:**
```javascript
// Violates 3NF - country depends on city, not on user
const users3NF_Bad = {
  1: { 
    id: 1, 
    name: "Alice", 
    city: "New York", 
    country: "USA" // Depends on city, not user
  }
};

// Follows 3NF - separate cities
const users3NF = {
  1: { id: 1, name: "Alice", cityId: 100 }
};
const cities = {
  100: { id: 100, name: "New York", countryId: 1 }
};
const countries = {
  1: { id: 1, name: "USA" }
};
```

**Benefits of Normalization:**
- Reduces data redundancy
- Improves data consistency
- Makes updates easier (update once, not in multiple places)
- Reduces memory usage
