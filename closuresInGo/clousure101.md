    Closures are one of those things that seem complicated until you realize:

> A closure is just a function that remembers variables from where it was created.

Let's start without Functional Options.

---

# Example 1

```go
func MakeGreeter(name string) func() {
    return func() {
        fmt.Println("Hello", name)
    }
}
```

Usage:

```go
greetAlice := MakeGreeter("Alice")
greetBob := MakeGreeter("Bob")

greetAlice()
greetBob()
```

Output:

```text
Hello Alice
Hello Bob
```

---

## What's happening?

When you call:

```go
greetAlice := MakeGreeter("Alice")
```

the outer function finishes and returns.

Normally you'd think:

```text
name is gone now
```

But it isn't.

The returned function "remembers":

```text
name = Alice
```

That's a closure.

---

# Visualize It

```text
MakeGreeter("Alice")
       │
       ▼
name = Alice
       │
       ▼
return function
       │
       ▼
function remembers:
    name = Alice
```

Later:

```go
greetAlice()
```

still knows:

```text
name = Alice
```

---

# Example 2: Counter

```go
func NewCounter() func() int {
    count := 0

    return func() int {
        count++
        return count
    }
}
```

Usage:

```go
counter := NewCounter()

fmt.Println(counter())
fmt.Println(counter())
fmt.Println(counter())
```

Output:

```text
1
2
3
```

---

## Why?

The closure remembers:

```go
count := 0
```

Even after `NewCounter()` returns.

Think:

```text
counter
  remembers
      count
```

Each call updates the same remembered variable.

---

# Two Different Closures

```go
c1 := NewCounter()
c2 := NewCounter()
```

Memory:

```text
c1 → count = 0

c2 → count = 0
```

Now:

```go
c1()
c1()
c2()
```

Output:

```text
1
2
1
```

Each closure has its own remembered state.

---

# Now Back to Functional Options

You wrote:

```go
func WithHost(newHost string) Option {
    return func(db *Database) {
        db.Host = newHost
    }
}
```

Let's expand it mentally.

When you call:

```go
WithHost("prod-db")
```

Go creates:

```go
func(db *Database) {
    db.Host = "prod-db"
}
```

That returned function remembers:

```text
newHost = prod-db
```

That's the closure.

---

# Constructor View

Suppose:

```go
db := NewDatabase(
    WithHost("prod-db"),
)
```

Execution:

### Step 1

```go
WithHost("prod-db")
```

returns:

```go
func(db *Database) {
    db.Host = "prod-db"
}
```

### Step 2

Stored in:

```go
opts []Option
```

### Step 3

Constructor runs:

```go
for _, opt := range opts {
    opt(db)
}
```

which executes:

```go
db.Host = "prod-db"
```

---

# Why Closures Are Useful

Imagine this:

```go
func WithRetry(count int) Option {
    return func(db *Database) {
        db.Retries = count
    }
}
```

Each option remembers a different value:

```go
WithRetry(3)
```

remembers:

```text
count = 3
```

while:

```go
WithRetry(10)
```

remembers:

```text
count = 10
```

Same code.

Different remembered state.

---

# Real Go Usage

Closures show up everywhere:

## HTTP Handlers

```go
func CreateHandler(service UserService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        service.GetUser()
    }
}
```

The returned handler remembers:

```text
service
```

---

## Goroutines

```go
name := "Alice"

go func() {
    fmt.Println(name)
}()
```

The anonymous function closes over:

```text
name
```

---

## Middleware

```go
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("request received")
        next.ServeHTTP(w, r)
    })
}
```

The returned function remembers:

```text
next
```

---

# Mental Model

Whenever you see:

```go
func (...) {
    ...
}
```

ask:

```text
Is this function using variables
that were declared outside it?
```

If yes:

```text
It's a closure.
```

Example:

```go
func WithHost(host string) Option {
    return func(db *Database) {
        db.Host = host
    }
}
```

The inner function uses:

```go
host
```

which belongs to the outer function.

Therefore:

```text
Closure detected.
```

---

# Exercise

Predict the output before running:

```go
func Multiplier(x int) func(int) int {
    return func(y int) int {
        return x * y
    }
}

func main() {
    double := Multiplier(2)
    triple := Multiplier(3)

    fmt.Println(double(5))
    fmt.Println(triple(5))
}
```

Draw the memory like this:

```text
double remembers:
    x = ?

triple remembers:
    x = ?
```

If you can explain why the output is what it is, you've understood closures well enough to understand Functional Options, middleware, and many Go web frameworks.
