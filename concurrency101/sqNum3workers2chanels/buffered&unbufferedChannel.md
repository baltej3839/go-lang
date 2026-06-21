Unbuffered Channels: The High-Five ✋
An unbuffered channel (make(chan int)) has a capacity of 0. It cannot hold data; it only passes it.

Think of it like a high-five. A high-five requires two people to be there at the exact same instant.

If a sender wants to send a job, it blocks (freezes) and waits until a worker is ready to receive it.

If a worker is ready to receive a job, it blocks and waits until the sender sends one.

If jobs is unbuffered, main will send the 1st number, and then main will freeze right there until one of your 3 workers wakes up and takes it. Once the worker takes it, main unfreezes, loops, sends the 2nd number, and freezes again.

For this specific task, both will work, but they change the behavior:

If you use Unbuffered: main and the workers actively "hand off" data. main can only send as fast as the 3 workers can process.

If you use Buffered (e.g., size 10): main unloads all the work instantly into memory and moves on, leaving the workers to clean up the queue.

tell me how to close a channel ?
and if a channel is closed and the data inside it , has to be yet processed , what happens to that ?
