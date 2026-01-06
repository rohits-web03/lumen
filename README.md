# Lumen

**Lumen** is an **ESM-first dev server and build tool written in Go**.

Lumen is designed around a simple idea:  
**in development, the browser should load modules natively**.

Instead of bundling everything upfront, Lumen serves source files as ES modules, rewriting and resolving imports so the browser can execute them directly. Bundling and production optimizations are treated as a *separate, later step* — not the foundation.

---

## Philosophy

- **ESM-first**: ES modules are the canonical format
- **Dev-server centric**: fast startup, minimal work, on-demand transforms
- **Explicit modules**: one file maps to one module
- **Simple before powerful**: features are added progressively

For architectural principles and constraints, see  
→ [`ARCHITECTURE.md`](./ARCHITECTURE.md)

---

## v0.0 Scope

### Goal

> Serve a browser-runnable ESM application from source files **without bundling**.

If this goal is met, v0.0 is complete.

The authoritative definition of v0.0 lives in  
→ [`V0.md`](./V0.md)

---

### v0.0 Checklist

- [ ] HTTP dev server written in Go
- [ ] Serve `index.html`
- [ ] Serve `.js` files as native ES modules
- [ ] Support `<script type="module">`
- [ ] Rewrite bare imports (`react`, `lodash`, etc.) to browser-resolvable paths
- [ ] Resolve and serve modules from `node_modules`
- [ ] Prefer ESM entry points when resolving packages
- [ ] Minimal in-memory caching for transformed files
- [ ] Single command: `lumen dev`

---

## Decision Log

Major architectural decisions and their reasoning are recorded in  
→ [`DECISIONS.md`](./DECISIONS.md)

This file is append-only and captures *why* choices were made, not just what was implemented.

---

## Status

Lumen is currently in **early experimental development (v0.0)**.  
APIs, behavior, and structure are expected to evolve rapidly.

---
