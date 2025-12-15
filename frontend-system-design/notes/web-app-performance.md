# Web App Performances

## Table of Contents

- [1. Performance Optimization](#1-performance-optimization)
  - [1.1. Core Web Vitals](#core-web-vitals)
  - [1.2. HTTP Protocol Versions](#http-protocol-versions)

## [1. Performance Optimization]()

### Core Web Vitals

Core Web Vitals are Google's essential metrics for measuring real-world user experience on the web. These metrics impact SEO rankings and user satisfaction.

#### The Three Core Metrics

**1. LCP (Largest Contentful Paint)**

Measures **loading performance** - how long it takes for the largest content element to become visible.

**What it measures:**
- Time from page start to largest image/text block rendered
- Focuses on perceived load speed

**Good Thresholds:**
- ✅ **Good:** < 2.5 seconds
- 🟡 **Needs Improvement:** 2.5 - 4.0 seconds
- ❌ **Poor:** > 4.0 seconds

**Common LCP Elements:**
- Hero images
- Video thumbnails
- Large text blocks
- Background images

**How to Improve:**
```javascript
// Optimize images
<img src="hero.jpg" 
     loading="eager" 
     fetchpriority="high"
     width="1200" 
     height="600" 
     alt="Hero" />

// Preload critical resources
<link rel="preload" as="image" href="hero.jpg" />

// Use modern image formats
<picture>
  <source srcset="hero.webp" type="image/webp">
  <img src="hero.jpg" alt="Hero">
</picture>
```

**Optimization Strategies:**
- Use CDN for images
- Compress images (WebP, AVIF)
- Lazy load below-the-fold content
- Remove render-blocking resources
- Use server-side rendering (SSR)

---

**2. INP (Interaction to Next Paint)**

Measures **responsiveness** - how quickly the page responds to user interactions (clicks, taps, keyboard).

**What it measures:**
- Delay from user interaction to visual feedback
- Replaces FID (First Input Delay) as of 2024
- Considers all interactions, not just first

**Good Thresholds:**
- ✅ **Good:** < 200 milliseconds
- 🟡 **Needs Improvement:** 200 - 500 milliseconds
- ❌ **Poor:** > 500 milliseconds

**How to Improve:**
```javascript
// Bad: Long-running task blocks interaction
function processData() {
  for (let i = 0; i < 1000000; i++) {
    // Heavy computation
  }
}

// Good: Break into smaller tasks
async function processDataAsync() {
  const chunks = 100;
  const chunkSize = 10000;
  
  for (let i = 0; i < chunks; i++) {
    // Process chunk
    processChunk(i * chunkSize, chunkSize);
    
    // Yield to browser
    await new Promise(r => setTimeout(r, 0));
  }
}

// Use Web Workers for heavy tasks
const worker = new Worker('heavy-task.js');
worker.postMessage({ data: largeDataset });
worker.onmessage = (e) => updateUI(e.data);

// Debounce frequent events
const debouncedSearch = debounce((query) => {
  fetchResults(query);
}, 300);

input.addEventListener('input', (e) => {
  debouncedSearch(e.target.value);
});
```

**Optimization Strategies:**
- Break up long JavaScript tasks
- Use Web Workers for CPU-intensive work
- Debounce/throttle event handlers
- Minimize DOM manipulation
- Use `requestIdleCallback` for non-critical work

---

**3. CLS (Cumulative Layout Shift)**

Measures **visual stability** - unexpected layout shifts that cause elements to move while page loads.

**What it measures:**
- Sum of all unexpected layout shift scores
- Tracks entire page lifetime

**Good Thresholds:**
- ✅ **Good:** < 0.1
- 🟡 **Needs Improvement:** 0.1 - 0.25
- ❌ **Poor:** > 0.25

**Common Causes:**
- Images without dimensions
- Ads/embeds/iframes without reserved space
- Dynamically injected content
- Web fonts causing FOIT/FOUT
- Animations using layout properties

**How to Improve:**
```html
<!-- Always include width/height -->
<img src="image.jpg" width="800" height="600" alt="Photo">

<!-- Reserve space for ads -->
<div style="min-height: 250px;">
  <!-- Ad slot -->
</div>

<!-- Use CSS aspect ratio -->
<style>
  .video-container {
    aspect-ratio: 16 / 9;
    width: 100%;
  }
</style>

<!-- Preload fonts -->
<link rel="preload" href="font.woff2" as="font" type="font/woff2" crossorigin>

<!-- Use font-display -->
<style>
  @font-face {
    font-family: 'CustomFont';
    src: url('font.woff2');
    font-display: swap; /* or optional */
  }
</style>
```

**Optimization Strategies:**
- Set explicit dimensions for images/videos
- Reserve space for dynamic content
- Avoid inserting content above existing content
- Use CSS transforms for animations (not top/left)
- Preload fonts and use `font-display: swap`

---

#### Measuring Web Vitals

```javascript
// Using web-vitals library
import { onLCP, onINP, onCLS } from 'web-vitals';

onLCP(metric => {
  console.log('LCP:', metric.value);
  // Send to analytics
  analytics.send({ metric: 'LCP', value: metric.value });
});

onINP(metric => {
  console.log('INP:', metric.value);
  analytics.send({ metric: 'INP', value: metric.value });
});

onCLS(metric => {
  console.log('CLS:', metric.value);
  analytics.send({ metric: 'CLS', value: metric.value });
});
```

**Tools for Testing:**
- Chrome DevTools (Lighthouse)
- PageSpeed Insights
- Web Vitals Chrome Extension
- Search Console (Core Web Vitals report)

---

### HTTP Protocol Versions

Understanding HTTP versions is crucial for performance optimization as each version introduces improvements for faster, more efficient communication.

#### HTTP/1.1 (1997)

**Characteristics:**
- One request per TCP connection (or sequential with keep-alive)
- Text-based protocol
- Uncompressed headers
- Head-of-line blocking

**Problems:**
```
Request 1 → Wait → Response 1
Request 2 → Wait → Response 2
Request 3 → Wait → Response 3

// Each request blocks the next (even with keep-alive)
```

**Workarounds Used:**
- Domain sharding (multiple subdomains)
- CSS sprites (combine images)
- Inlining assets (base64)
- Concatenating JS/CSS files

**Performance:**
- ❌ 6-8 connections per domain limit
- ❌ Header overhead (500-800 bytes per request)
- ❌ No request prioritization
- ❌ Sequential processing

---

#### HTTP/2 (2015)

**Key Improvements:**

**1. Multiplexing:**
```
Connection 1:
├─ Request 1 → Response 1 (streaming)
├─ Request 2 → Response 2 (streaming)
└─ Request 3 → Response 3 (streaming)

// All requests/responses in parallel over one connection
```

**2. Header Compression (HPACK):**
- Compresses headers (reduces overhead by ~80%)
- Maintains header state between requests

**3. Server Push:**
```html
<!-- Server can push resources before requested -->
<html>
<head>
  <!-- Server pushes style.css before browser asks -->
  <link rel="stylesheet" href="style.css">
</head>
```

**4. Stream Prioritization:**
- Critical resources loaded first
- Dependencies specified

**Benefits:**
- ✅ Multiple requests over single connection
- ✅ Reduced latency
- ✅ No need for domain sharding
- ✅ Better compression
- ✅ Binary protocol (faster parsing)

**Limitations:**
- ❌ Still has head-of-line blocking at TCP level
- ❌ Server push often misused (cache issues)
- ❌ Requires HTTPS (TLS)

**Best Practices:**
```javascript
// No need for:
// - Domain sharding
// - CSS sprites (can serve individual images)
// - Concatenating all JS/CSS

// Instead:
// - Serve individual, cacheable resources
// - Use code splitting
// - Let HTTP/2 handle parallel loading
```

---

#### HTTP/3 (2022)

**Revolutionary Change: Built on QUIC (UDP instead of TCP)**

**Key Improvements:**

**1. No Head-of-Line Blocking:**
```
Stream 1: ████████░░ (packet lost, only affects Stream 1)
Stream 2: ██████████ (continues unaffected)
Stream 3: ██████████ (continues unaffected)

// TCP would block all streams if one packet lost
```

**2. Faster Connection Setup:**
```
HTTP/2 (TCP + TLS):
TCP handshake (1 RTT) + TLS handshake (1-2 RTT) = 2-3 RTT

HTTP/3 (QUIC):
Combined handshake = 0-1 RTT (with 0-RTT on reconnection)
```

**3. Connection Migration:**
```javascript
// Switching from WiFi to Cellular?
// HTTP/2: Connection drops, reconnect required
// HTTP/3: Connection survives, no interruption
```

**4. Improved Congestion Control:**
- Better loss recovery
- Faster retransmission
- More efficient on lossy networks

**Performance:**
- ✅ ~10-30% faster than HTTP/2 on average
- ✅ ~50% faster on poor networks (mobile)
- ✅ Survives network changes
- ✅ No TCP head-of-line blocking
- ✅ Faster connection setup

**Limitations:**
- ❌ Some networks block UDP
- ❌ Less mature than HTTP/2
- ❌ Higher CPU usage (QUIC in userspace)
- ❌ Not all CDNs support it yet

---

#### HTTP Version Comparison

| Feature | HTTP/1.1 | HTTP/2 | HTTP/3 |
|---------|----------|--------|--------|
| **Protocol** | TCP | TCP | QUIC (UDP) |
| **Multiplexing** | No | Yes | Yes |
| **Header Compression** | No | HPACK | QPACK |
| **Head-of-line Blocking** | Yes (request level) | Yes (TCP level) | No |
| **Connection Setup** | 2-3 RTT | 2-3 RTT | 0-1 RTT |
| **Connection Migration** | No | No | Yes |
| **Server Push** | No | Yes | Yes |
| **Performance** | Baseline | 15-50% faster | 10-30% faster than HTTP/2 |
| **Mobile Performance** | Poor | Good | Excellent |
| **Browser Support** | Universal | Universal | ~95% (2025) |

---

#### How to Enable

**HTTP/2:**
```nginx
# Nginx
server {
    listen 443 ssl http2;
    ssl_certificate cert.pem;
    ssl_certificate_key key.pem;
}
```

**HTTP/3:**
```nginx
# Nginx (with QUIC module)
server {
    listen 443 quic reuseport;
    listen 443 ssl http2;
    
    ssl_protocols TLSv1.3;
    add_header Alt-Svc 'h3=":443"; ma=86400';
}
```

**Detection:**
```javascript
// Check protocol used
fetch('/api/data')
  .then(response => {
    console.log('Protocol:', response.headers.get('x-protocol'));
  });

// Browser typically tries: HTTP/3 → HTTP/2 → HTTP/1.1
```

---

#### Optimization Strategy by Version

**HTTP/1.1:**
- Minimize requests (sprites, concatenation)
- Domain sharding (4-6 domains)
- Inline critical CSS
- Use CDN

**HTTP/2:**
- Avoid concatenation (use code splitting)
- Remove domain sharding
- Serve many small, cacheable files
- Use server push carefully

**HTTP/3:**
- Same as HTTP/2
- Especially beneficial for mobile users
- Consider fallback to HTTP/2 for enterprise networks
