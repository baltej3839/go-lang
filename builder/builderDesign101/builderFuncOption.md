Functional Options are basically Go's answer to:

> "I need a configurable object, but I don't want a constructor with 15 parameters and I don't want a Builder."

---

# The Problem

Suppose you start with:

```go
func NewServer(
    host string,
    port int,
) *Server
```

Life is good.

Later you add:

```go
TLS
Timeout
Retries
MaxConns
LogLevel
```

Now you get:

```go
func NewServer(
    host string,
    port int,
    tls bool,
    timeout int,
    retries int,
    maxConns int,
    logLevel string,
) *Server
```

When you call it:

```go
server := NewServer(
    "localhost",
    8080,
    true,
    30,
    3,
    100,
    "debug",
)
```

Nobody knows what those numbers mean.

---

# Builder Solution

You already learned:

```go
server := NewServerBuilder().
    Host("localhost").
    Port(8080).
    TLS(true).
    Timeout(30).
    Build()
```

Works.

---

# Functional Options Idea

Instead of:

```go
NewServer(host, port, tls, timeout)
```

think:

```text
Create server
    ↓
Apply modifications
    ↓
Return server
```

Each option becomes a function.

---

# First Principles

Suppose:

```go
type Server struct {
    Host    string
    Port    int
    TLS     bool
    Timeout int
}
```

Create a default server:

```go
func NewServer() *Server {
    return &Server{
        Host: "localhost",
        Port: 8080,
        TLS: false,
        Timeout: 10,
    }
}
```

---

Now ask:

> How can somebody modify this during creation?

Answer:

Pass functions.

---

# Option Type

Define:

```go
type Option func(*Server)
```

Read it slowly:

```text
Option
    =
a function

that receives a Server pointer

and modifies it
```

---

# Example Option

```go
func WithTLS(enabled bool) Option {
    return func(s *Server) {
        s.TLS = enabled
    }
}
```

This function returns another function.

Think:

```text
WithTLS(true)

creates

a function that knows
TLS should become true
```

---

# Another Option

```go
func WithTimeout(timeout int) Option {
    return func(s *Server) {
        s.Timeout = timeout
    }
}
```

---

# Constructor

Now:

```go
func NewServer(
    opts ...Option,
) *Server {
    s := &Server{
        Host: "localhost",
        Port: 8080,
        TLS: false,
        Timeout: 10,
    }

    for _, opt := range opts {
        opt(s)
    }

    return s
}
```

---

# Usage

```go
server := NewServer(
    WithTLS(true),
    WithTimeout(30),
)
```

Let's mentally execute it.

---

## Step 1

Create default server

```go
s := &Server{
    Host: "localhost",
    Port: 8080,
    TLS: false,
    Timeout: 10,
}
```

---

## Step 2

Apply:

```go
WithTLS(true)
```

returns:

```go
func(s *Server) {
    s.TLS = true
}
```

Execute:

```go
opt(s)
```

Now:

```go
TLS = true
```

---

## Step 3

Apply:

```go
WithTimeout(30)
```

returns:

```go
func(s *Server) {
    s.Timeout = 30
}
```

Execute.

Now:

```go
Timeout = 30
```

---

## Final Result

```go
Server{
    Host: "localhost",
    Port: 8080,
    TLS: true,
    Timeout: 30,
}
```

---

# Why Go Developers Like This

Imagine next month you add:

```go
MaxConns
```

You don't change:

```go
NewServer()
```

at all.

Just add:

```go
func WithMaxConns(n int) Option
```

Usage:

```go
server := NewServer(
    WithTLS(true),
    WithTimeout(30),
    WithMaxConns(100),
)
```

No constructor explosion.

---

# Compare Builder vs Functional Options

Builder:

```go
server := NewServerBuilder().
    TLS(true).
    Timeout(30).
    Build()
```

Functional Options:

```go
server := NewServer(
    WithTLS(true),
    WithTimeout(30),
)
```

Both solve the same problem.

---

# Mental Model

Most beginners see:

```go
type Option func(*Server)
```

and panic.

Instead think:

```text
Server starts with defaults

Each Option is a tiny function

that changes one thing

Constructor runs all those functions
```

Like:

```text
New Server
     ↓
Apply TLS option
     ↓
Apply Timeout option
     ↓
Apply Retry option
     ↓
Done
```

---

# Real World Examples

You'll see this pattern in many Go libraries, including:

* [gRPC Go](https://grpc.io/docs/languages/go/?utm_source=chatgpt.com)
* [Zap Logger](https://go.uber.org/zap?utm_source=chatgpt.com)
* [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/?utm_source=chatgpt.com)

---

# Exercise

Implement a `Database` struct:

```go
type Database struct {
    Host string
    Port int
    SSL bool
    MaxConnections int
}
```

Requirements:

```go
db := NewDatabase(
    WithSSL(true),
    WithMaxConnections(50),
)
```

Default values:

```text
Host = localhost
Port = 5432
SSL = false
MaxConnections = 10
```

Write:

1. `type Option func(*Database)`
2. `WithSSL`
3. `WithMaxConnections`
4. `NewDatabase(opts ...Option)`

If you can implement that without looking back, you've understood the core of Functional Options. The only remaining concept afterward is understanding **why returning a function can capture values (`enabled`, `timeout`)**, which leads into closures.
