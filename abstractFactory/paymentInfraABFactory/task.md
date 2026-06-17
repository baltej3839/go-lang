Excellent. Here's one that a Go backend developer might actually encounter.

---

# Exercise: Payment Infrastructure

You're building an application that runs in two modes:

```text
Sandbox Environment
Production Environment
```

Each environment needs a family of related services.

---

## Interfaces

### Payment Gateway

```go
type PaymentGateway interface {
    Charge(amount int) error
}
```

---

### Logger

```go
type Logger interface {
    Log(msg string)
}
```

---

### Notification Service

```go
type Notifier interface {
    Send(msg string) error
}
```

---

# Sandbox Implementations

### Payment Gateway

```text
FakeGateway
```

Print:

```text
Sandbox charge successful
```

---

### Logger

```text
ConsoleLogger
```

Print:

```text
[Sandbox] <message>
```

---

### Notifier

```text
FakeNotifier
```

Print:

```text
Pretending to send notification
```

---

# Production Implementations

### Payment Gateway

```text
StripeGateway
```

Print:

```text
Charging customer via Stripe
```

---

### Logger

```text
ZapLogger
```

Print:

```text
JSON log written
```

---

### Notifier

```text
EmailNotifier
```

Print:

```text
Email sent successfully
```

---

# Abstract Factory Interface

You should design something like:

```go
type PaymentInfraFactory interface {
    CreateGateway() PaymentGateway
    CreateLogger() Logger
    CreateNotifier() Notifier
}
```

---

# Concrete Factories

Implement:

```go
type SandboxFactory struct {}
```

and

```go
type ProductionFactory struct {}
```

---

# Main

This should work:

```go
var factory PaymentInfraFactory

factory = &ProductionFactory{}

gateway := factory.CreateGateway()
logger := factory.CreateLogger()
notifier := factory.CreateNotifier()

logger.Log("Charging user")

gateway.Charge(100)

notifier.Send("Payment successful")
```

Expected output:

```text
JSON log written
Charging customer via Stripe
Email sent successfully
```

---

# Bonus (Real-life)

Write a normal Factory:

```go
func NewFactory(env string) PaymentInfraFactory
```

so that:

```go
factory := NewFactory("prod")
```

returns:

```text
ProductionFactory
```

and

```go
factory := NewFactory("sandbox")
```

returns:

```text
SandboxFactory
```

---

# Super Bonus (Very Real)

Add one more interface:

```go
type Cache interface {
    Set(key string, value any)
}
```

Implement:

### Sandbox

```text
MemoryCache
```

### Production

```text
RedisCache
```

and add:

```go
CreateCache() Cache
```

to the abstract factory.

---

## Why this exercise is good

Notice that:

```text
Sandbox
---------
FakeGateway
ConsoleLogger
FakeNotifier
MemoryCache

Production
-----------
StripeGateway
ZapLogger
EmailNotifier
RedisCache
```

These objects belong together.

If tomorrow you add:

```text
Staging Environment
```

you only create:

```go
type StagingFactory struct {}
```

and the application code below:

```go
gateway := factory.CreateGateway()
logger := factory.CreateLogger()
notifier := factory.CreateNotifier()
cache := factory.CreateCache()
```

never changes.

That's the essence of Abstract Factory.

---

Once you've finished it, paste your code, and I'll review it like a Go code review and point out what a senior Go developer would improve.
