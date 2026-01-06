# Lumen Architecture

## Core Idea

Lumen is an **ESM-first dev server**.

In development, the browser is the module loader.
Lumen exists to make source files browser-runnable without bundling.

---

## Design Principles

1. **ESM is canonical**
   - All code served to the browser is ES modules
   - Other formats are transformed into ESM

2. **Dev server first**
   - Development experience drives architecture
   - Production build is a derivative concern

3. **One file = one module**
   - No bundling in dev
   - Files map directly to URLs

4. **Explicit over implicit**
   - Import resolution is visible and debuggable
   - No hidden magic or global state

---

## What Lumen Will Not Do

- Execute CommonJS at runtime
- Bundle during development
- Hide module boundaries
- Optimize for legacy browsers in dev
