# Architectural Decisions

This file records **why** decisions were made.
It is append-only.

---

## Decision 001: ESM-first Architecture

**Choice**
- Lumen is ESM-first, not bundler-first

**Alternatives Considered**
- Bundle-first dev workflow (Webpack-style)
- Hybrid ESM + CJS runtime

**Reason**
- ESM enables instant startup and fine-grained invalidation
- Browser-native loading simplifies the mental model
- Bundling complicates development unnecessarily

---

## Decision 002: Go as the Implementation Language

**Choice**
- Core of Lumen is written in Go

**Alternatives Considered**
- TypeScript / Node.js
- Rust

**Reason**
- Strong filesystem and networking primitives
- Fast startup and single binary distribution
- Better suited for systems-level tooling
- Aligns with learning goals of understanding internals
- I don't know Rust :(
