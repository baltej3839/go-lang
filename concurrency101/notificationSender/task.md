Now that you've mastered the foundational worker pool, let's step it up to a classic production scenario. In the real world, systems don't just process batches of static data; they deal with **streaming data pipelines**, **rate limiting**, and **fan-in architecture**.

Here is your next challenge.

---

## Real-World Scenario: The Notification Dispatcher

Imagine you are building a notification service for an app like Uber or DoorDash. Your system receives a continuous stream of raw notification requests (SMS, Email, Push). You need to:

1. Parse and validate the incoming requests.
2. Hand them off to a worker pool that mimics sending them (e.g., calling an external API like Twilio).
3. Log the status of every single attempt into a central logging system.

---

## Task Requirements

### 1. The Request Generator (Producer)

Write a function that simulates a stream of incoming notification requests:

```go
type Notification struct {
    ID      int
    Type    string // "SMS", "Email", or "Push"
    UserID  int
    Message string
}

```

* This function should generate **20 distinct notifications** and send them into an `incomingRequests` channel.
* Once all 20 are sent, it should close the channel.

### 2. The Worker Pool (Processors)

Spin up **3 worker goroutines**. Each worker should:

* Read notifications from the `incomingRequests` channel.
* Simulate processing time by sleeping for a random duration (e.g., between 10 to 50 milliseconds).
* Create a `Receipt` struct for every job they process:

```go
type Receipt struct {
    NotificationID int
    WorkerID       int
    Status         string // "SUCCESS" or "FAILED"
}

```

* *Note:* Randomly make about 10% of the notifications "FAILED" to make it realistic.
* Send this `Receipt` into a shared `receipts` channel.

### 3. The Logger (The Consumer / Fan-In)

Instead of collecting results in a slice inside `main` like last time, you will create a **dedicated logger goroutine**.

* This goroutine listens to the `receipts` channel.
* Every time it receives a `Receipt`, it prints it nicely to the console:
`[Worker 2] Successfully sent Notification #12 (SMS)` or `[Worker 1] FAILED to send Notification #5 (Email)`.
* It tracks how many total successes and failures occurred.

### 4. Main Orchestration

Your `main` function should:

1. Initialize the necessary channels.
2. Launch the Logger goroutine.
3. Launch the 3 Workers.
4. Start the Generator (either inline or as a goroutine).
5. Cleanly shut down everything when done. When the logger finishes reading all receipts, it should print a final summary: `Total Processed: 20 | Success: 18 | Failed: 2`.

---

## Why this task is different:

* **Decoupled Logging:** The workers don't print to the console anymore. Printing to a console is slow; delegating it to a separate logger goroutine keeps your workers running at maximum speed.
* **Pipelines:** You are chaining channels together (`Generator -> incomingRequests -> Workers -> receipts -> Logger`).

Give this architecture a shot. Take your time setting up how the channels flow into each other! Where do you think you'll need `sync.WaitGroup` this time to ensure the Logger doesn't exit early?