Factory Method is much easier to understand if you already know interfaces.

In fact, Factory Method answers one question:

> "Who should decide which concrete implementation to create?"

---

# Problem

Suppose you have:

```go id="j9xq56"
type PaymentGateway interface {
    Pay(amount int) error
}
```

Implementations:

```go id="wlyi8j"
type Stripe struct{}
type Razorpay struct{}
```

Both implement:

```go id="gg8t5i"
Pay(amount int) error
```

---

Without a factory:

```go id="h18h7e"
stripe := &Stripe{}
```

or

```go id="xwok6r"
razorpay := &Razorpay{}
```

Now imagine configuration says:

```text id="dtvw0x"
PAYMENT_PROVIDER=stripe
```

How does the application know what to create?

---

# First Principles

Think:

```text id="4h9gsw"
Input:
    "stripe"

Output:
    Stripe object
```

or

```text id="t4wlsm"
Input:
    "razorpay"

Output:
    Razorpay object
```

This is object creation logic.

Instead of spreading it everywhere:

```go id="3a1r7h"
if provider == "stripe" {
    ...
}
```

we centralize it.

---

# Factory Function

```go id="qg5x0q"
func NewPaymentGateway(
    provider string,
) PaymentGateway {

    switch provider {

    case "stripe":
        return &Stripe{}

    case "razorpay":
        return &Razorpay{}

    default:
        panic("unknown provider")
    }
}
```

Usage:

```go id="k0p46o"
gateway := NewPaymentGateway(
    "stripe",
)
```

Result:

```text id="0h44e0"
gateway is PaymentGateway

actual object is Stripe
```

---

# Why Is This Useful?

Your service doesn't care.

```go id="5i04rm"
type Service struct {
    gateway PaymentGateway
}
```

Later:

```go id="57v5qf"
service.gateway.Pay(100)
```

No knowledge of Stripe or Razorpay.

Only interface.

---

# Visual Diagram

Without Factory:

```text id="6w7u7g"
Service
   │
   ├── Stripe
   ├── Razorpay
   ├── Paypal
   └── ...
```

Service knows too much.

---

With Factory:

```text id="f8mn3x"
Service
   │
   ▼
PaymentGateway

Factory
   │
   ├── Stripe
   ├── Razorpay
   └── Paypal
```

Creation logic is isolated.

---

# Is This a Design Pattern?

Yes.

The Factory Method pattern is:

```text id="7klx7z"
Hide object creation behind a method.
```

The caller asks:

```text id="xot8kv"
Give me a PaymentGateway
```

instead of:

```text id="qvzn5n"
Give me a Stripe
```

---

# Real Go Examples

You already use factories constantly.

Example:

```go id="95d0g5"
file, err := os.Open("data.txt")
```

You don't do:

```go id="fgg8v4"
file := &os.File{
    ...
}
```

Instead:

```go id="r6x64l"
os.Open()
```

creates it.

That's factory-like behavior.

---

# Builder vs Factory

This confuses many developers.

## Factory

Focus:

```text id="gq1v0v"
WHICH object?
```

Example:

```go id="buzg5c"
NewPaymentGateway("stripe")
```

Decision:

```text id="4s1jof"
Stripe or Razorpay?
```

---

## Builder

Focus:

```text id="x2clnh"
HOW should object be configured?
```

Example:

```go id="u1fjlwm"
NewServerBuilder().
    TLS(true).
    Timeout(30).
    Build()
```

Decision:

```text id="2wxvri"
What settings should Server have?
```

---

# A More Go-like Example

Imagine your booking project.

Interface:

```go id="tx75hf"
type BookingStore interface {
    SaveBooking(...)
}
```

Implementations:

```go id="g7k4x5"
MemoryStore
PostgresStore
RedisStore
```

Factory:

```go id="3k3d9j"
func NewStore(
    storeType string,
) BookingStore
```

Usage:

```go id="8l6n4t"
store := NewStore("postgres")
```

Application doesn't know:

```text id="vtajv9"
How PostgresStore is created
```

Factory does.

---

# The Factory Method Mental Model

Whenever you see:

```text id="xl7khe"
I have multiple implementations
and I must choose one.
```

think:

```text id="d3t6v4"
Factory
```

Whenever you see:

```text id="4vq6bo"
I have one object
with lots of configuration.
```

think:

```text id="flquaq"
Builder / Functional Options
```

---

### Exercise

Design a notification system:

```go id="1j2avf"
type Notifier interface {
    Send(msg string) error
}
```

Implement:

```go id="4h7g9p"
EmailNotifier
SMSNotifier
PushNotifier
```

Write a factory:

```go id="1hjlwm"
func NewNotifier(
    notifierType string,
) Notifier
```

Then use it like:

```go id="jlwm8v"
notifier := NewNotifier("sms")
notifier.Send("hello")
```

Try it yourself before looking up any examples. You'll see why interfaces and factories fit together naturally in Go.
