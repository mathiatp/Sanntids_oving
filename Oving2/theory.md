Exercise 2 - Theory questions
-----------------------------

### What is an atomic operation?
> An operation that is running completly independent of other operations.

### What is a critical section?
> A part of the code that needs to be protected from concurrency

### What is the difference between race conditions and data races?
> A data race condition is a fault in the design of the interactions between two or more tasks whereby the result is unexpecte and critically dependent on the sequence of timing of accesses to shared data.
A race condition is that a device or system attempts to preform two or more operations at the same time, but because of the nature of the device or system, the operations must be done in the proper sequence to be done correctly.
The difference between the two is that data races is about the order of access to a variable/memory, while race condiotion is a more general settig, where the race is between operations.
### What are the differences between semaphores, binary semaphores, and mutexes?
> A semaphore is like an integere that after initialization can not be read, or changed in any other way than incrementing and decrementing it.
A binary semaphore is a semaphore that can only have the values 1 and 0. Compared to a normal one, binary is often easier to implmentand, but it is more limiting with how many resources that can runa the same time.
Mutexes is like a token that gives premission to run that is passed between threads resulting in only one thread running at the time. 
### What are the differences between channels (in Communicating Sequential Processes, or as used by Go, Rust), mailboxes (in the Actor model, or as used by Erlang, D, Akka), and queues (as used by Python)? 
> A mailbox in Actor model is like a channel in Go, but it is private to other processes, compared to channels where every thread can be given access to the channels.
queues in Python is how it keeps track of communitaction in multiprocessing. Often it is used as a FIFO queue, but can also be implemented as LIFO queue.
### List some advantages of using message passing over lock-based synchronization primitives.
> Message passing is generally safer, becuase it denies the programmer access to locking and shared resources, betting against the programmer. It is also more convinient
### List some advantages of using lock-based synchronization primitives over message passing.
> Some algorithms can be made much simpler. If there is no wait time in algorithms, then it will be faster than with message passing.
