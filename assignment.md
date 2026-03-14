# Refurbed – Senior Fullstack Engineer Home Assignment

**Backend:** Go  
**Frontend:** Vue + Tailwind **or** Vanilla JS/HTML/CSS  

This assignment simulates real squad work: 
- building product discovery features, aggregating data, adding filters, load more pagination and rendering product cards.

**Disclaimers:**
- We do **not** expect a production-ready system.  
- All snippets, data structures, are provided as examples. Feel free to modify as you see fit.
- We want to see **clarity, structure, and senior-level problem‑solving**.

---

# Overview

Build a small fullstack project consisting of:

### Go backend
- Reads product data from **two internal sources**
- Aggregates and merges the data
- Applies **search, filtering, load more**
- Exposes **one endpoint:** `GET /products`
- Includes **simple in-memory caching (30s TTL)**

### Frontend
- Generates a product listing UI
- Displays product cards (image, name, price, discount, colors, bestseller badge)
- Implements search, filters, load more type of pagination

You may choose between:

- **Vue + Tailwind** (starter provided)  
- **HTML/CSS + Vanilla JS** (starter provided)

**Figma design is provided here:**
- https://www.figma.com/design/98z5akfMhdgVvUvQIy2U77/Senior-Full-Stack-Engineer?node-id=0-1&p=f
- You may use product images from https://www.refurbed.at/

**Implementation of the design:**
Pixel-perfect accuracy isn’t required, but the final UI should be very close to the provided Figma designs (including key details like positioning, spacing, typography, colours, component states, etc).
Any vibe-coded / AI-assisted implementation that isn’t validated (e.g., basic QA + visual checks) and only follows the wireframe at a high level will be rejected.

---

# Backend Requirements (Go)

## 1. Two Internal Data Sources / APIs

You may accommodate the json structure and files provided as you wish. 

### **Metadata Source**
- `id`
- `name`
- `base_price`
- `image_url`

### **Details Source**
- `id`
- `discount_percent`
- `bestseller`
- `colors` (e.g., `["blue", "green", "red"]`)
- `stock`

**No database is required.**

---

## 2. Aggregator Endpoint

Expose:

```
GET /products
```

Sample:

```json
{
  "id": "p1",
  "name": "iPhone 15",
  "price": 610.99,
  "discount_percent": 25,
  "bestseller": true,
  "colors": ["blue", "red", "green"],
  "image_url": "",
  "stock": 34
}
```

---

## 3. Server side search, filters, load more capabilities

Support query params:

```
GET /products?search=iphone&color=blue&bestseller=true&minPrice=100&maxPrice=500
```

---

## 4. Server-side caching

Implement an **in-memory cache** for the full aggregated product list with TTL: **30 seconds**.

This does **not** need to be a production-ready cache.  
We want to see your **thought process and design**.

---

# Frontend Requirements

Choose one:

- `frontend-vue/`  
- `frontend-vanilla/`  

The environment is intentionally minimal so we can see **your** thought process.

### Focus on:
- Product Card Rendering
- Search, Filters & Load More UI
- A11y & SEO
- Responsivness

**Implementation of the design:**
Pixel-perfect accuracy isn’t required, but the final UI should be very close to the provided Figma designs (including key details like positioning, spacing, typography, colours, component states, etc).
Any vibe-coded / AI-assisted implementation that isn’t validated (e.g., basic QA + visual checks) and only follows the wireframe at a high level will be rejected.


### Bonus points:
- Sort by popularity ranking, provided by external source

---


## Deliverables

Send us back the zip file with:

- Your Go backend  
- Your frontend (Vue or vanilla)  
- A README including:
  - How to run backend
  - How to run frontend
  - What you would improve/change for production ready environment
  - Notes on architecture, decisions or other comments
  - Final thoughts


---
---

Happy coding,

Refurbed Engineering