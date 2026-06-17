This is the right question. Patterns are useful only if you know **when to reach for them**.

The first-principles question behind Factory is:

> **Who should decide what concrete object gets created?**

If the answer is:

> "Not the caller"

then a factory is probably useful.

---

# 1. Multiple implementations exist

Suppose:

```go
type Store interface {
    Save()
}
```

Implementations:

```go
MemoryStore
RedisStore
PostgresStore
```

Without factory:

```go
store := &PostgresStore{}
```

The caller has to know which implementation to create.

With factory:

```go
store, _ := NewStore("postgres")
```

### Use Factory when:

```text
I have several implementations
and I want to centralize choosing one.
```

---

# 2. Object creation is complicated

Suppose creating a database requires:

```text
Open connection
↓
Ping database
↓
Configure connection pool
↓
Run migrations
↓
Return object
```

You don't want callers doing:

```go
db := &PostgresStore{
    ...
}
```

Instead:

```go
store, err := NewPostgresStore(connString)
```

### Use Factory when:

```text
Object creation itself contains logic.
```

---

# 3. Configuration decides behavior

Suppose:

```yaml
storage: redis
```

Then:

```go
store, _ := NewStore(cfg.Storage)
```

Tomorrow:

```yaml
storage: postgres
```

No application code changes.

### Use Factory when:

```text
Runtime configuration chooses implementation.
```

---

# 4. You want to hide concrete types

Suppose:

```go
type PaymentGateway interface {
    Pay()
}
```

Caller:

```go
gateway, _ := NewPaymentGateway("stripe")
```

Caller never sees:

```go
Stripe{}
```

### Use Factory when:

```text
I want callers to depend on interfaces,
not concrete structs.
```

---

# 5. Dependency Injection

Main:

```go
store, _ := NewStore("redis")

service := NewUserService(store)
```

Service doesn't know:

```text
Redis
Postgres
Memory
```

Only:

```go
Store
```

### Use Factory when:

```text
Main wires dependencies together.
```

Very common.

---

# 6. Future implementations are expected

Today:

```text
EmailNotifier
SMSNotifier
```

Tomorrow:

```text
SlackNotifier
PushNotifier
WhatsAppNotifier
```

Factory changes:

```go
NewNotifier()
```

Caller stays:

```go
notifier.Send()
```

### Use Factory when:

```text
New implementations are likely.
```

---

# When NOT to use Factory

Simple struct:

```go
type User struct {
    Name string
}
```

Don't do:

```go
NewUserFactory()
```

Just:

```go
user := User{
    Name: "Bob",
}
```

---

# Decision Algorithm

Whenever creating something, ask:

```text
Do I have multiple implementations?
        |
       yes
        ↓
    Factory

Do I have one object with many options?
        |
       yes
        ↓
 Builder / Functional Options

Do I simply need a struct?
        |
       yes
        ↓
 Composite literal

Do I need to change behavior at runtime?
        |
       yes
        ↓
 Strategy
```

---

# Real-world Go examples

### `sql.Open()`

```go
db, err := sql.Open("postgres", dsn)
```

Factory chooses the correct driver.

---

### `http.NewRequest()`

```go
req, err := http.NewRequest(...)
```

Complex object creation hidden.

---

### Your future booking project

```go
store, _ := NewStore(cfg.Storage)
```

returns:

```text
MemoryStore
RedisStore
PostgresStore
```

Your service only sees:

```go
BookingStore
```

---

## Rule of Thumb

If you ever write:

```go
if config.Type == "redis" {
    ...
}

if config.Type == "postgres" {
    ...
}
```

in multiple places,

that's usually a sign:

```text
Move this decision into a Factory.
```

A factory centralizes **object creation decisions**.
