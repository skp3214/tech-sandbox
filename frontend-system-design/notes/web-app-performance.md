# Web App Performances

## Table of Contents

- [1. Performance Optimization](#1-performance-optimization)
- [2. JavaScript Bundling and Loading](#2-javascript-bundling-and-loading)
- [3. CSS, Images and Rendering](#3-css-image-and-rendering)

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
## [3. CSS, Images, Fonts, and Rendering Optimization]()

### [CSS Optimization]()

#### [1. Minification and Compression]()

**Minification** removes whitespace, comments, and shortens code:

```css
/* Before minification (10 KB) */
.button {
  background-color: #007bff;
  padding: 10px 20px;
  border-radius: 4px;
  color: white;
}

/* After minification (2 KB) */
.button{background-color:#007bff;padding:10px 20px;border-radius:4px;color:#fff}
```

**Tools:**
- **cssnano** - PostCSS plugin
- **clean-css** - Fast minifier
- **PurgeCSS** - Removes unused CSS

```javascript
// Build configuration
module.exports = {
  plugins: [
    require('cssnano')({
      preset: ['default', {
        discardComments: { removeAll: true },
      }]
    })
  ]
}
```

---

#### [2. Critical CSS]()

Inline critical above-the-fold CSS to avoid render-blocking:

```html
<!DOCTYPE html>
<html>
<head>
  <!-- Inline critical CSS -->
  <style>
    .header{background:#000;padding:20px}
    .hero{min-height:400px;background:#f5f5f5}
  </style>
  
  <!-- Load full CSS async -->
  <link rel="preload" href="styles.css" as="style" onload="this.onload=null;this.rel='stylesheet'">
  <noscript><link rel="stylesheet" href="styles.css"></noscript>
</head>
```

**Tools:**
- **Critical** - Extract & inline critical CSS
- **Critters** - Webpack plugin

---

#### [3. CSS-in-JS vs Traditional CSS]()

**Traditional CSS (Better for performance):**
```css
/* styles.css - Single file, cached, processed at build time */
.button { color: blue; }
```

**CSS-in-JS (Runtime overhead):**
```javascript
// Generates styles at runtime
const Button = styled.div`
  color: blue;
`;
```

**Performance Tips:**
- Use CSS Modules for scoping
- Avoid runtime CSS-in-JS in performance-critical apps
- If using CSS-in-JS, use static extraction (Linaria, vanilla-extract)

---

#### [4. Utility-First CSS (Tailwind)]()

**Benefits:**
- Smaller bundle size (remove unused classes)
- No naming conventions needed
- Consistent design tokens

```html
<!-- Before: Custom CSS -->
<style>
  .custom-button { @apply bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded; }
</style>
<button class="custom-button">Click</button>

<!-- After: Utility classes -->
<button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
  Click
</button>
```

**Optimization with PurgeCSS:**
```javascript
// tailwind.config.js
module.exports = {
  content: [
    './src/**/*.{js,jsx,ts,tsx,html}',
  ],
  // Before purge: 3 MB
  // After purge: ~10 KB (only used classes)
}
```

**Best Practices:**
- Enable JIT mode for development
- Purge unused classes in production
- Use `@apply` sparingly (reduces utility benefits)

---

### [Image Optimization]()

#### [1. Image Formats Comparison]()

| Format | Type | Compression | Transparency | Animation | Best For | Size |
|--------|------|-------------|--------------|-----------|----------|------|
| **JPEG** | Raster | Lossy | No | No | Photos | Medium |
| **PNG** | Raster | Lossless | Yes | No | Graphics with transparency | Large |
| **WebP** | Raster | Both | Yes | Yes | General purpose (modern) | Small |
| **AVIF** | Raster | Both | Yes | Yes | Next-gen format | Smallest |
| **SVG** | Vector | Lossless | Yes | Yes | Icons, logos | Tiny (scalable) |
| **GIF** | Raster | Lossless | Yes | Yes | Legacy animations | Large |

---

#### [2. Modern Raster Formats]()

**WebP** (30% smaller than JPEG):
```html
<picture>
  <source srcset="image.webp" type="image/webp">
  <img src="image.jpg" alt="Fallback">
</picture>
```

**AVIF** (50% smaller than JPEG, newest format):
```html
<picture>
  <source srcset="image.avif" type="image/avif">
  <source srcset="image.webp" type="image/webp">
  <img src="image.jpg" alt="Fallback">
</picture>
```

**Browser Support (2025):**
- WebP: ~97% (all modern browsers)
- AVIF: ~90% (Chrome, Firefox, Safari 16+)

**Conversion:**
```bash
# Convert to WebP
cwebp image.jpg -q 80 -o image.webp

# Convert to AVIF
avifenc image.jpg image.avif --quality 80

# Batch convert
for f in *.jpg; do cwebp "$f" -o "${f%.jpg}.webp"; done
```

---

#### [3. Animations: GIF vs MP4 vs WebP]()

**Problem with GIF:**
- Large file size (no modern compression)
- Limited colors (256 colors)
- No audio support

**Solution: Use Video Formats**

| Format | Size | Quality | Browser Support | Best For |
|--------|------|---------|----------------|----------|
| **GIF** | 3.2 MB | Poor | Universal | Legacy only |
| **MP4** | 500 KB | Excellent | Universal | General animations |
| **WebM** | 400 KB | Excellent | 95% | Modern browsers |
| **Animated WebP** | 450 KB | Good | 97% | Simple animations |

**Convert GIF to Video:**
```bash
# GIF to MP4 (80-95% size reduction)
ffmpeg -i animation.gif -movflags faststart -pix_fmt yuv420p -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" animation.mp4

# GIF to WebM
ffmpeg -i animation.gif -c vp9 -b:v 0 -crf 30 animation.webm

# GIF to Animated WebP
gif2webp animation.gif -o animation.webp
```

**Implementation:**
```html
<!-- Replace GIF with video -->
<video autoplay loop muted playsinline>
  <source src="animation.webm" type="video/webm">
  <source src="animation.mp4" type="video/mp4">
  <!-- Fallback -->
  <img src="animation.gif" alt="Animation">
</video>

<!-- Or use animated WebP -->
<picture>
  <source srcset="animation.webp" type="image/webp">
  <img src="animation.gif" alt="Animation">
</picture>
```

**Results:**
```
GIF:           3,200 KB  (baseline)
MP4:             500 KB  (84% smaller)
WebM:            400 KB  (87% smaller)
Animated WebP:   450 KB  (86% smaller)
```

---

#### [4. Responsive Images]()

**Use `srcset` for different screen sizes:**
```html
<img 
  src="image-800.jpg"
  srcset="
    image-400.jpg 400w,
    image-800.jpg 800w,
    image-1200.jpg 1200w,
    image-1600.jpg 1600w
  "
  sizes="(max-width: 600px) 400px,
         (max-width: 1000px) 800px,
         1200px"
  alt="Responsive image"
  loading="lazy"
  decoding="async"
>
```

**Art direction with `<picture>`:**
```html
<picture>
  <!-- Mobile: Portrait crop -->
  <source media="(max-width: 600px)" srcset="image-portrait.jpg">
  
  <!-- Tablet: Square crop -->
  <source media="(max-width: 1000px)" srcset="image-square.jpg">
  
  <!-- Desktop: Landscape -->
  <img src="image-landscape.jpg" alt="Adaptive image">
</picture>
```

---

#### [5. Lazy Loading]()

```html
<!-- Native lazy loading -->
<img src="image.jpg" loading="lazy" alt="Lazy loaded">

<!-- Eager load for above-the-fold -->
<img src="hero.jpg" loading="eager" fetchpriority="high" alt="Hero">

<!-- JavaScript fallback -->
<img 
  data-src="image.jpg" 
  class="lazy"
  alt="Image"
>

<script>
  const images = document.querySelectorAll('img.lazy');
  
  const imageObserver = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        const img = entry.target;
        img.src = img.dataset.src;
        img.classList.remove('lazy');
        imageObserver.unobserve(img);
      }
    });
  });
  
  images.forEach(img => imageObserver.observe(img));
</script>
```

---

#### [6. Image CDN & Optimization Services]()

```html
<!-- Cloudinary -->
<img src="https://res.cloudinary.com/demo/image/upload/
  w_800,          <!-- Width -->
  f_auto,         <!-- Auto format (WebP/AVIF) -->
  q_auto,         <!-- Auto quality -->
  c_fill,         <!-- Crop fill -->
  g_auto          <!-- Auto focus -->
/sample.jpg" alt="Optimized">

<!-- imgix -->
<img src="https://demo.imgix.net/image.jpg?
  w=800&
  auto=format,compress&
  fit=crop
" alt="Optimized">
```

---

### [Font Optimization]()

#### [1. Font Loading Strategies]()

**font-display options:**
```css
@font-face {
  font-family: 'CustomFont';
  src: url('font.woff2') format('woff2');
  font-display: swap; /* Show fallback immediately, swap when loaded */
}

/* Options:
   - auto: Browser default
   - block: Hide text (max 3s), then show (FOIT - Flash of Invisible Text)
   - swap: Show fallback immediately, swap when loaded (FOUT - Flash of Unstyled Text)
   - fallback: Brief hide (100ms), swap if loads within 3s, else use fallback
   - optional: Brief hide, use only if cached (best for performance)
*/
```

**Recommendation:** Use `swap` for better UX, `optional` for best performance

---

#### [2. Font Format Priorities]()

```css
@font-face {
  font-family: 'CustomFont';
  src: url('font.woff2') format('woff2'),      /* Modern (best compression) */
       url('font.woff') format('woff'),        /* Fallback */
       url('font.ttf') format('truetype');     /* Legacy fallback */
  font-display: swap;
}

/* Format comparison:
   WOFF2: 30% smaller than WOFF (use this)
   WOFF:  Older format
   TTF:   Uncompressed (avoid)
*/
```

---

#### [3. Preload Critical Fonts]()

```html
<head>
  <!-- Preload fonts used above the fold -->
  <link 
    rel="preload" 
    href="fonts/main-font.woff2" 
    as="font" 
    type="font/woff2"
    crossorigin
  >
  
  <!-- Preconnect to Google Fonts -->
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
</head>
```

---

#### [4. Variable Fonts]()

Single file contains multiple weights/styles:

```css
/* Traditional: Multiple files */
@font-face { font-family: 'Font'; font-weight: 400; src: url('regular.woff2'); }
@font-face { font-family: 'Font'; font-weight: 700; src: url('bold.woff2'); }
/* Total: 200 KB */

/* Variable font: One file for all weights */
@font-face {
  font-family: 'FontVariable';
  src: url('font-variable.woff2') format('woff2-variations');
  font-weight: 100 900; /* Supports all weights */
}
/* Total: 120 KB (40% smaller) */

.text {
  font-family: 'FontVariable';
  font-weight: 456; /* Any weight between 100-900 */
}
```


### [Rendering Optimization]()

#### [1. Critical Rendering Path]()

```
1. HTML Parsing → DOM Tree
2. CSS Parsing  → CSSOM Tree
3. DOM + CSSOM  → Render Tree
4. Layout       → Calculate positions
5. Paint        → Draw pixels
6. Composite    → Layer composition
```

**Bottlenecks:**
- Render-blocking CSS
- Parser-blocking JavaScript
- Large DOM trees
- Complex CSS selectors

---

#### [2. Reduce Reflows and Repaints]()

**Reflow** (expensive) - Recalculate layout:
- Changing: width, height, position, display
- Adding/removing elements
- Font changes

**Repaint** (less expensive) - Redraw pixels:
- Changing: color, background, visibility
- No layout change

```javascript
// ❌ Bad: Multiple reflows (Layout Thrashing)
for (let i = 0; i < elements.length; i++) {
  elements[i].style.width = elements[i].offsetWidth + 10 + 'px';
  // Read (offsetWidth) → Write (style.width) → Reflow!
}

// ✅ Good: Batch reads and writes
const widths = [];
for (let i = 0; i < elements.length; i++) {
  widths[i] = elements[i].offsetWidth; // Read all
}
for (let i = 0; i < elements.length; i++) {
  elements[i].style.width = widths[i] + 10 + 'px'; // Write all
}

// ✅ Better: Use CSS transforms (no reflow)
element.style.transform = 'translateX(100px)'; // Only composite
```

---

#### [3. Use CSS Transform and Opacity]()

**These properties only trigger composite (cheapest):**
```css
/* ✅ Performant - GPU accelerated */
.animated {
  transform: translateX(100px);
  opacity: 0.5;
  will-change: transform, opacity;
}

/* ❌ Expensive - Triggers layout/paint */
.animated {
  left: 100px;      /* Reflow */
  width: 200px;     /* Reflow */
  background: red;  /* Repaint */
}
```

**CSS Triggers Reference:**
- `transform`, `opacity` → Composite only (best)
- `color`, `background` → Paint + Composite
- `width`, `height`, `position` → Layout + Paint + Composite (worst)

---

#### [4. Layer Promotion]()

Create separate compositing layers:

```css
/* Promote to own layer */
.animated-element {
  will-change: transform, opacity;
  /* Or */
  transform: translateZ(0); /* Force GPU */
}

/* Warning: Don't overuse! */
/* Each layer uses memory (~10MB for 1000px x 1000px) */
```

**When to use:**
- Elements with frequent animations
- Fixed/sticky positioned elements
- Canvas/video elements

**When NOT to use:**
- Static elements
- Too many layers (memory issues)

---
