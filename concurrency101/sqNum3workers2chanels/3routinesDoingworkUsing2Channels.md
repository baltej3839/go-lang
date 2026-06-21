Awesome job tackling the first part! Moving from "one goroutine per item" to a fixed pool of workers is a massive milestone in mastering Go concurrency. It changes how you think about data flow.

Here is the conceptual breakdown of how to limit your program to exactly 3 workers, without giving away the code.

---

## The Key Shift: Push vs. Pull

In your first solution, you spun up 10 goroutines and handed each one its specific number directly.

With a worker pool, you spin up **exactly 3 goroutines** at the very beginning. At the moment of creation, they don't have any data yet. Instead, they sit there waiting to **pull** work from a shared conveyor belt (the `jobs` channel).

---

## How to Guide Your Implementation

To make this happen, break it down into these four logical steps:

### 1. The Worker Loop

Instead of a worker function taking a single `num int`, it now needs to take the `jobs` channel (`<-chan int`) and the `results` channel (`chan<- int`).

Inside the worker, you need a mechanism that keeps it alive, constantly listening for new jobs until there are no more. You can do this elegantly using a `for ... range` loop over the `jobs` channel:

```go
for num := range jobs {
    // 1. Do the work (square the number)
    // 2. Send it to results
}
// 3. Signal that THIS worker is done out here!

```

*Crucial Detail:* Because a `range` loop over a channel only exits when the channel is **closed**, the worker will automatically block and wait if the channel is temporarily empty, and will cleanly exit once you close the `jobs` channel.

### 2. Spawning the Pool

In your `main` function, you will use a standard loop to launch your workers:

* Loop exactly 3 times.
* Inside the loop, spin up your worker goroutine using the `go` keyword.
* Don't forget to track these workers using your `sync.WaitGroup` (add to the WaitGroup *before* you start the loop or inside it).

### 3. Feeding the Conveyor Belt

Now that your 3 workers are standing by, you need to feed them.

* Loop through your `nums` slice and send each number into the `jobs` channel.
* **Immediately after** sending all 10 numbers, you **must close** the `jobs` channel. Closing it tells the 3 workers: *"Hey, no more numbers are coming, you can finish up and exit your loops."*

### 4. The Orchestration (The Tricky Part)

You have two separate channel closures to worry about here, and order matters:

1. **Close `jobs**`: Done in `main` right after sending the 10 numbers.
2. **Close `results**`: You still need a separate goroutine to wait for the workers to finish (`wg.Wait()`) and then close the `results` channel, just like you did in the first exercise.

---

## A Brain Teaser Before You Write the Code

Before you start typing, think about this:
Should your `jobs` channel be **buffered** or **unbuffered**? What happens if you try to send all 10 numbers into an unbuffered `jobs` channel before the workers start reading them?

Give the code a shot based on this structure. Let me know how it goes or if you hit a deadlock!