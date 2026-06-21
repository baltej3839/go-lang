Write a program that processes a single customer order through three distinct stages in a strict, sequential pipeline: Validate Order $\rightarrow$ Charge Card $\rightarrow$ Ship Item.Each stage must be handled by its own dedicated goroutine, passing the order data forward using individual unbuffered channels


Write a program that processes a single customer order through three distinct stages in a strict, sequential pipeline: 
Validate Order $\rightarrow$ 
Charge Card $\rightarrow$ 
Ship Item.

Each stage must be handled by its own dedicated goroutine, passing the order data forward using individual unbuffered channels