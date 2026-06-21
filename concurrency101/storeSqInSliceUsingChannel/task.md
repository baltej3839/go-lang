Perfect. Since you've understood `WaitGroup` and `Mutex`, it's time to add **channels**, which are Go's primary communication mechanism.

# Task: Worker Pool with Channels

Write a program that computes the squares of numbers concurrently and collects the results using channels.

## Input

```go
nums := []int{1,2,3,4,5,6,7,8,9,10}
```

---

## Requirements

### Worker

Write a function:

```go
func GiveSquareOfNumber(
    num int,
    wg *sync.WaitGroup,
    results chan int,
)
```

Each worker should:

1. Compute `num*num`.
2. Print:

```text
square of 4 is 16
```

3. Send the square into the channel:

```go
results <- square
```

4. Call `wg.Done()`.

---

### Main

1. Create:

```go
var wg sync.WaitGroup
```

2. Create a channel:

```go
results := make(chan int)
```

3. Launch one goroutine per number.

4. Start another goroutine whose job is:

```go
wg.Wait()
close(results)
```

---

### Collect results

In main, receive values using:

```go
for square := range results {
    // store them in a slice
}
```

---

### Finally print

```go
fmt.Println(sliceOfSquares)
```

---

# Constraints

❌ No mutexes.

❌ No shared slice passed into workers.

✅ Use channels for communication.

---

# Bonus (Harder)

After finishing, modify the program so that:

* 3 worker goroutines are created.
* Numbers are sent into a `jobs` channel.
* Workers read from the `jobs` channel.
* Results are sent to a `results` channel.
* Main collects all results.

Architecture:

```text
main
 │
 ├── jobs channel -----> Worker1
 │                  └-> Worker2
 │                  └-> Worker3
 │
 └── results channel <---- workers
             │
             ▼
        sliceOfSquares
```

Try the first version yourself. If you get stuck, show me your code and we'll move to the full worker-pool version together.
