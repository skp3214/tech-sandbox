## Web APIs for Complex UI Patterns

## Table of Contents

- [1. Observer API](#observer-api)

## [Observer API]()


### Overview

The Observer APIs are a collection of modern browser APIs that allow you to efficiently observe changes to DOM elements and the viewport. These APIs are essential for building performant, complex UI patterns without relying on expensive polling or scroll event listeners.

![alt text](screenshots/image-14.png)

### Key Observer APIs

#### [1. **Intersection Observer API**](/frontend-system-design/infinite-scroll-demo.html)

Monitors when elements enter or exit the viewport or intersect with a parent element.

**Use Cases:**
- Lazy loading images and content
- Infinite scrolling
- Tracking ad visibility
- Triggering animations when elements come into view
- Analytics and user engagement tracking

**Basic Usage:**
```javascript
const observer = new IntersectionObserver((entries) => {
  entries.forEach(entry => {
    if (entry.isIntersecting) {
      console.log('Element is visible');
      // Load image, trigger animation, etc.
      entry.target.classList.add('visible');
    }
  });
}, {
  threshold: 0.5,  // Trigger when 50% visible
  rootMargin: '0px' // Offset from viewport
});

// Observe elements
const images = document.querySelectorAll('img[data-src]');
images.forEach(img => observer.observe(img));
```

**Options:**
- `root`: Element used as viewport (default: browser viewport)
- `rootMargin`: Margin around root (like CSS margin)
- `threshold`: Percentage of visibility to trigger callback (0 to 1)
