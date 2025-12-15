# Web App Performances

## Table of Contents

- [1. Performance Optimization](#1-performance-optimization)
- [2. JavaScript Bundling and Loading](#2-javascript-bundling-and-loading)

## [1. Performance Optimization]()

### [Core Web Vitals]()

Core Web Vitals are Google's essential metrics for measuring real-world user experience on the web. These metrics impact SEO rankings and user satisfaction.

#### The Three Core Metrics

[**1. LCP (Largest Contentful Paint)**]()

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

[**2. INP (Interaction to Next Paint)**]()

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

[**3. CLS (Cumulative Layout Shift)**]()

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


**Tools for Testing:**
- Chrome DevTools (Lighthouse)
- PageSpeed Insights

---

### [HTTP Protocol Versions]()

Understanding HTTP versions is crucial for performance optimization as each version introduces improvements for faster, more efficient communication.

#### [HTTP/1.1 (1997)]()

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

#### [HTTP/2 (2015)]()

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

#### [HTTP/3 (2022)]()

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

## [2. JavaScript Bundling and Loading]()

### [What is Bundling?]()

Bundling is the process of combining multiple JavaScript files into fewer files to optimize loading performance. Modern bundlers also handle transpilation, minification, tree-shaking, and code splitting.

**Popular Bundlers:**
- **Webpack** - Most configurable, mature ecosystem
- **Vite** - Ultra-fast, ESM-based, great DX
- **Rollup** - Excellent for libraries, best tree-shaking
- **Parcel** - Zero config, automatic optimization
- **esbuild** - Extremely fast (Go-based)

---

### [Bundling Strategies]()

#### [1. Single Bundle (Simple but Slow)]()

```javascript
// All code in one file
bundle.js (500 KB)
├─ React
├─ App code
├─ Libraries
└─ Utilities

// Problems:
// - Large initial download
// - Download everything even if unused
// - No parallel loading
// - Cache invalidation on any change
```

**When to use:** Very small apps (< 100 KB)

---

#### [2. Code Splitting (Smart Loading)]()

Split code into smaller chunks loaded on-demand.

**Route-based splitting:**
```javascript
// React with lazy loading
import { lazy, Suspense } from 'react';

const Home = lazy(() => import('./pages/Home'));
const Dashboard = lazy(() => import('./pages/Dashboard'));
const Profile = lazy(() => import('./pages/Profile'));

function App() {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/profile" element={<Profile />} />
      </Routes>
    </Suspense>
  );
}

// Result:
// main.js (50 KB) - Core app + React
// home.chunk.js (20 KB) - Loaded on /
// dashboard.chunk.js (30 KB) - Loaded on /dashboard
// profile.chunk.js (15 KB) - Loaded on /profile
```

**Component-based splitting:**
```javascript
// Split heavy components
const HeavyChart = lazy(() => import('./HeavyChart'));
const VideoPlayer = lazy(() => import('./VideoPlayer'));

function Dashboard() {
  return (
    <div>
      <h1>Dashboard</h1>
      <Suspense fallback={<Spinner />}>
        <HeavyChart data={data} />
      </Suspense>
    </div>
  );
}
```

**Dynamic imports:**
```javascript
// Load on interaction
button.addEventListener('click', async () => {
  const module = await import('./heavyFeature.js');
  module.initialize();
});

// Conditional loading
if (user.isPremium) {
  const premium = await import('./premiumFeatures.js');
  premium.enable();
}
```


#### [3. Tree Shaking]()

Remove unused code during bundling.

```javascript
// library.js
export function used() { return 'I am used'; }
export function unused() { return 'I am never used'; }

// app.js
import { used } from './library';
console.log(used());

// Bundle output (with tree-shaking):
// Only 'used' function included, 'unused' removed
function used() { return 'I am used'; }
console.log(used());
```

**Best practices:**
```javascript
// ✅ Good: Named imports (tree-shakeable)
import { Button } from 'ui-library';

// ❌ Bad: Default imports (imports everything)
import * as UI from 'ui-library';

// ✅ Good: Import specific modules
import debounce from 'lodash/debounce';

// ❌ Bad: Import entire library
import _ from 'lodash';
```

---

### [Loading Strategies]()

#### [1. Script Tag Attributes]()

```html
<!-- Normal: Blocks parsing and rendering -->
<script src="app.js"></script>

<!-- Async: Downloads in parallel, executes when ready -->
<script src="app.js" async></script>
<!-- Use for: Analytics, ads, independent scripts -->

<!-- Defer: Downloads in parallel, executes after HTML parsed -->
<script src="app.js" defer></script>
<!-- Use for: Main application scripts -->

<!-- Module: Deferred by default, supports ES6 imports -->
<script type="module" src="app.js"></script>
```

**Visual Timeline:**
```
Normal:  HTML parsing ──[blocked]── JS download ──[blocked]── Execute ── Continue parsing
Async:   HTML parsing ──────────────────────────────── JS download ─[blocked]─ Execute
Defer:   HTML parsing ──────────────────────────────── Parsing done ─ JS Execute
```

---

#### [2. Preload & Prefetch]()

**Preload** - High priority, needed soon:
```html
<!-- Load critical resources early -->
<link rel="preload" as="script" href="critical.js">
<link rel="preload" as="style" href="critical.css">
<link rel="preload" as="font" href="font.woff2" crossorigin>

<!-- When to use: -->
<!-- - Above-the-fold resources -->
<!-- - Critical path scripts -->
<!-- - Web fonts -->
```

**Prefetch** - Low priority, might need later:
```html
<!-- Load for next navigation -->
<link rel="prefetch" as="script" href="next-page.js">
<link rel="prefetch" as="image" href="next-page-hero.jpg">

<!-- When to use: -->
<!-- - Next route resources -->
<!-- - Anticipated user actions -->
<!-- - Background prefetching -->
```

**Preconnect** - Establish early connection:
```html
<!-- Connect to third-party domains -->
<link rel="preconnect" href="https://api.example.com">
<link rel="dns-prefetch" href="https://cdn.example.com">

<!-- When to use: -->
<!-- - API endpoints -->
<!-- - CDN domains -->
<!-- - Third-party resources -->
```

---

#### [3. Module Preload]()

```html
<!-- Preload ES modules and their dependencies -->
<link rel="modulepreload" href="app.js">
<link rel="modulepreload" href="utils.js">

<!-- Then use -->
<script type="module" src="app.js"></script>
```

---

#### [4. Progressive Loading Pattern]()

```javascript
// 1. Load critical code immediately
import('./critical.js').then(module => {
  module.initApp();
});

// 2. Load important code after critical
setTimeout(() => {
  import('./important.js');
}, 0);

// 3. Load nice-to-have when idle
if ('requestIdleCallback' in window) {
  requestIdleCallback(() => {
    import('./optional.js');
  });
} else {
  setTimeout(() => import('./optional.js'), 2000);
}

// 4. Load on interaction
document.getElementById('btn').addEventListener('click', () => {
  import('./feature.js').then(module => {
    module.activate();
  });
}, { once: true });
```

---

### [Advanced Techniques]()

#### [1. Streaming SSR]()

Send HTML in chunks as it's generated.

```javascript
// Next.js App Router (React Server Components)
export default async function Page() {
  return (
    <div>
      <Header />
      <Suspense fallback={<Skeleton />}>
        <SlowComponent /> {/* Streamed when ready */}
      </Suspense>
      <Footer />
    </div>
  );
}

// Browser receives:
// 1. <Header /> immediately
// 2. <Skeleton /> placeholder
// 3. <SlowComponent /> when data ready (streamed)
// 4. <Footer /> immediately
```

---

#### [2. Service Worker Caching]()

```javascript
// Cache bundles for offline use
self.addEventListener('install', (event) => {
  event.waitUntil(
    caches.open('v1').then((cache) => {
      return cache.addAll([
        '/main.js',
        '/vendor.js',
        '/styles.css',
      ]);
    })
  );
});

self.addEventListener('fetch', (event) => {
  event.respondWith(
    caches.match(event.request).then((response) => {
      return response || fetch(event.request);
    })
  );
});
```
