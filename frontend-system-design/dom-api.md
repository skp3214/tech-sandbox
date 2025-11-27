## DOM API

## Table of Contents

- [1. DOM and Querying](#dom-and-querying)

### [DOM and Querying]()

The DOM API is a set of methods we can utilise in JavScript to manipulate the DOM.

**When to use:-**

- Building a low-level library (e.g virtualisation, DOM management etc)
- Creating Generic Component (like Video Player).

### Why Different Methods Exist

The DOM API provides multiple methods for similar tasks because they serve different purposes and use cases:

- **Performance vs. Convenience**: Methods like `getElementById()` are faster but limited to IDs, while `querySelector()` is more flexible but slightly slower
- **Single vs. Multiple Elements**: `querySelector()` returns one element, while `querySelectorAll()` returns all matches - choose based on your needs
- **Live vs. Static Collections**: `getElementsByClassName()` returns a live HTMLCollection that updates automatically, while `querySelectorAll()` returns a static NodeList - use live collections when DOM changes matter
- **Browser Compatibility**: Older methods like `getElementsByTagName()` have better legacy browser support, while newer methods like `querySelector()` offer modern CSS selector syntax
- **Specificity**: Specialized methods like `getElementById()` or `getElementsByClassName()` make code intent clearer than generic selectors
- **Use Case Optimization**: Different methods are optimized for different scenarios - use the most specific method for better performance and readability

**Rule of Thumb**: Use `querySelector()`/`querySelectorAll()` for modern projects with complex selectors, and specific methods like `getElementById()` for simple, performance-critical operations.

![alt text](image-12.png)

