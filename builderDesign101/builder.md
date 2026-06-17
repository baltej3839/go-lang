Since you like learning from first principles, let's build up the **Builder Pattern** from a problem rather than a definition.

# Problem

Imagine you're creating a server configuration.

```go
type Server struct {
    Host     string
    Port     int
    TLS      bool
    Timeout  int
    MaxConns int
}
```

Without a builder, you might write:

```go
server := Server{
    Host:     "localhost",
    Port:     8080,
    TLS:      true,
    Timeout:  30,
    MaxConns: 100,
}
```

This is okay for 5 fields.

But imagine:

```go
type Server struct {
    Host string
    Port int

    TLS bool
    Timeout int
    MaxConns int

    CacheEnabled bool
    CacheSize int

    Compression bool

    LogLevel string

    RetryCount int
}
```

Now creating objects becomes messy.

---

# First Attempt: Constructor

```go
func NewServer(
    host string,
    port int,
    tls bool,
    timeout int,
    maxConns int,
) *Server
```

Usage:

```go
server := NewServer(
    "localhost",
    8080,
    true,
    30,
    100,
)
```

Problem:

Can you immediately tell what `30` means?

```go
NewServer(
    "localhost",
    8080,
    true,
    30,
    100,
)
```

Not really.

---

# Builder Idea

Instead of supplying everything at once:

```text
Create empty object
      ↓
Configure it step by step
      ↓
Build final object
```

---

# Builder Structure

```go
type ServerBuilder struct {
    server Server
}
```

The builder's job is:

```text
Collect configuration
      ↓
Produce Server
```

---

# Step 1: Constructor

```go
func NewServerBuilder() *ServerBuilder {
    return &ServerBuilder{}
}
```

---

# Step 2: Configuration Methods

```go
func (b *ServerBuilder) Host(host string) *ServerBuilder {
    b.server.Host = host
    return b
}
```

Why return `*ServerBuilder`?

So you can chain:

```go
builder.
    Host("localhost").
    Port(8080)
```

---

Another method:

```go
func (b *ServerBuilder) Port(port int) *ServerBuilder {
    b.server.Port = port
    return b
}
```

---

And:

```go
func (b *ServerBuilder) TLS(enabled bool) *ServerBuilder {
    b.server.TLS = enabled
    return b
}
```

---

# Build Method

Finally:

```go
func (b *ServerBuilder) Build() Server {
    return b.server
}
```

---

# Usage

```go
server := NewServerBuilder().
    Host("localhost").
    Port(8080).
    TLS(true).
    Build()
```

Now it reads almost like English:

```text
Create Builder
    ↓
Host localhost
    ↓
Port 8080
    ↓
Enable TLS
    ↓
Build
```

---

# What Problem Did We Solve?

Instead of:

```go
NewServer(
    "localhost",
    8080,
    true,
    30,
    100,
)
```

we now have:

```go
NewServerBuilder().
    Host("localhost").
    Port(8080).
    TLS(true).
    Timeout(30).
    Build()
```

Much more readable.

---

# Real Go Example: HTTP Request

Imagine building an HTTP request.

Without builder:

```go
req := Request{
    Method: "POST",
    URL: "/users",
    Headers: headers,
    Body: body,
}
```

Builder:

```go
req := NewRequestBuilder().
    Method("POST").
    URL("/users").
    Header("Authorization", token).
    Body(data).
    Build()
```

This is very common.

---

# First Principles Algorithm

Whenever you see:

```text
Object has many optional fields
```

consider:

```text
Builder
```

Algorithm:

```text
1. Create Builder
2. Store partial configuration
3. Each method modifies state
4. Return builder
5. Build final object
```

---

# Builder vs Factory

Many developers confuse these.

## Factory

```go
db := NewPostgresDB()
```

Purpose:

```text
Hide object creation
```

---

## Builder

```go
db := NewDBBuilder().
    Host("localhost").
    Port(5432).
    SSL(true).
    Build()
```

Purpose:

```text
Construct complex objects step-by-step
```

---

# Real Exercise

Build a `PizzaBuilder`.

```go
type Pizza struct {
    Size string

    Cheese bool
    Olives bool
    Mushroom bool
    Paneer bool

    ExtraSauce bool
}
```

Expected usage:

```go
pizza := NewPizzaBuilder().
    Size("Large").
    Cheese().
    Mushroom().
    Paneer().
    ExtraSauce().
    Build()
```

Try implementing that yourself.

You'll discover the three core pieces of Builder:

```text
Builder Struct
      ↓
Chainable Methods
      ↓
Build()
```

Once you can build `PizzaBuilder` without looking at notes, you've understood the Builder pattern.
