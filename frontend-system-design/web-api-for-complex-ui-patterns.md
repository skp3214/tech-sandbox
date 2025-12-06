## Web APIs for Complex UI Patterns

## Table of Contents

- [1. Observer API](#observer-api)
  - [1.1. Intersection Observer API](#1-intersection-observer-api)
  - [1.2. Mutation Observer API](#2-mutation-observer-api)
  - [1.3. Resize Observer API](#3-resize-observer-api)

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

#### [2. **Mutation Observer API**](/mutation-observer-demo.html)

Watches for changes to the DOM tree (additions, removals, attribute changes).

**Use Cases:**
- Detecting dynamic content changes
- Third-party script integration
- Form validation on DOM changes
- Tracking user-generated content
- Debugging DOM modifications

**Basic Usage:**
```javascript
const mutationObserver = new MutationObserver((mutations) => {
  mutations.forEach(mutation => {
    if (mutation.type === 'childList') {
      console.log('Nodes added or removed');
      mutation.addedNodes.forEach(node => {
        console.log('Added:', node);
      });
    } else if (mutation.type === 'attributes') {
      console.log(`Attribute ${mutation.attributeName} changed`);
    }
  });
});

mutationObserver.observe(document.querySelector('#container'), {
  childList: true,      // Watch for child additions/removals
  attributes: true,     // Watch for attribute changes
  subtree: true,        // Watch descendants too
  characterData: true,  // Watch text content changes
  attributeFilter:['one','two'],
  attributeOldValue: true,  // Record old attribute values
  characterDataOldValue: true  // Record old text values
});
```

**Configuration Options:**
- `childList`: Watch for child node changes
- `attributes`: Watch for attribute changes
- `characterData`: Watch for text content changes
- `subtree`: Apply to entire subtree
- `attributeFilter`: Array of specific attributes to watch

#### [3. **Resize Observer API**](/resize-observer-demo.html)

Monitors changes to element dimensions (width, height) without relying on window resize events.

**Use Cases:**
- Responsive component layouts
- Dynamic chart/graph resizing
- Textarea auto-resize
- Container query-like behavior
- Adjusting content based on element size changes

**Basic Usage:**
```javascript
const resizeObserver = new ResizeObserver((entries) => {
  entries.forEach(entry => {
    const { width, height } = entry.contentRect;
    console.log(`Element resized to ${width}x${height}`);
    
    // Adjust layout based on size
    if (width < 500) {
      entry.target.classList.add('compact');
    } else {
      entry.target.classList.remove('compact');
    }
  });
});

// Observe element
const container = document.querySelector('#responsive-container');
resizeObserver.observe(container);
```

**Key Properties:**
- `contentRect`: Element's dimensions (width, height, top, left)
- `borderBoxSize`: Size including padding and border
- `contentBoxSize`: Size of content box only
- `devicePixelContentBoxSize`: Size in device pixels

