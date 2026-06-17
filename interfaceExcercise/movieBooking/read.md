Since you're already writing Go code, working with interfaces, repositories, services, and testing, here's a **real Go developer exercise** that forces you to understand:

* Why interfaces exist
* Dependency inversion
* Mocking
* Polymorphism
* Interface composition
* Testing with mocks

without feeling like a toy example.

---

# Exercise: Movie Ticket Booking System

Build this package:

```text
booking/
├── booking.go
├── service.go
├── store.go
├── notifier.go
└── service_test.go
```

---

# Requirements

A user can book a seat.

When booking:

1. Check if seat exists
2. Check if seat already booked
3. Save booking
4. Send confirmation notification

---

## Domain

```go
type Booking struct {
    SeatID string
    UserID string
}
```

---

# Step 1

Create interface:

```go
type BookingStore interface {
    SeatExists(seatID string) bool
    IsBooked(seatID string) bool
    SaveBooking(b Booking) error
}
```

Question to think about:

> Why is Service accepting BookingStore instead of MemoryStore?

---

# Step 2

Create another interface

```go
type Notifier interface {
    Send(userID string, msg string) error
}
```

Implement:

```go
type EmailNotifier struct{}
```

Later implement:

```go
type SMSNotifier struct{}
```

without changing service code.

---

# Step 3

Service

```go
type Service struct {
    store    BookingStore
    notifier Notifier
}
```

Constructor:

```go
func NewService(
    store BookingStore,
    notifier Notifier,
) *Service
```

Booking function:

```go
func (s *Service) Book(
    seatID string,
    userID string,
) error
```

---

# Rule

```text
seat does not exist
      ↓
return error

seat already booked
      ↓
return error

save booking
      ↓
send notification
      ↓
success
```

---

# Step 4 (Hard Part)

Implement TWO stores.

### MemoryStore

```go
type MemoryStore struct {}
```

Uses map.

---

### SQLStore

```go
type SQLStore struct {}
```

No real database.

Just print:

```go
saving into postgres...
```

The point is:

```go
service := NewService(
    sqlStore,
    emailNotifier,
)
```

should work.

And:

```go
service := NewService(
    memoryStore,
    smsNotifier,
)
```

should also work.

without changing Service code.

---

# Step 5 (Harder)

Create interface composition.

```go
type Reader interface {
    SeatExists(string) bool
    IsBooked(string) bool
}

type Writer interface {
    SaveBooking(Booking) error
}

type BookingStore interface {
    Reader
    Writer
}
```

Question:

Why would someone split a big interface into smaller interfaces?

---

# Step 6 (Very Hard)

Create AuditLogger.

```go
type AuditLogger interface {
    Log(action string)
}
```

Implement:

```go
ConsoleLogger
FileLogger
```

Service now receives:

```go
type Service struct {
    store    BookingStore
    notifier Notifier
    logger   AuditLogger
}
```

Log:

```go
booking started
booking saved
notification sent
booking failed
```

---

# Step 7 (Senior Level)

Write tests.

Create fake implementations:

```go
type MockStore struct {}
type MockNotifier struct {}
type MockLogger struct {}
```

Your test should verify:

```go
service.Book(...)
```

actually called:

```go
SaveBooking()
Send()
Log()
```

without using real database,
without using real email.

This is exactly why Go developers love interfaces.

---

# Final Boss

Add payment processing.

```go
type PaymentGateway interface {
    Charge(
        userID string,
        amount int,
    ) error
}
```

Implement:

```go
StripeGateway
RazorpayGateway
FakeGateway
```

Service should work with all three.

---

If you finish all of this, you'll understand the question:

> "When should I accept an interface as a parameter?"

The answer becomes:

> Accept an interface when your code only cares about behavior, not a specific implementation.

That's the core idea behind most Go backend architectures (repositories, services, payment gateways, caches, message queues, loggers, etc.).
