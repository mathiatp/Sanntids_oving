Exercise 2 - Theory questions
-----------------------------

### What is an atomic operation?
> Atomic operations are program operations that run completely independently of any other processes.

### What is a critical section?
> Is a segment of code where shared variables can be accessed and can't be executed by more than one thread at the same time.

### What is the difference between race conditions and data races?
> Data race is when two instructions from different threads access the same memory location.
  Race condition occurs when a program depends on the timing of one mor multiple processes to function correcly, if a thread     finishes at an unexpected time, it might cause an error.

### What are the differences between semaphores, binary semaphores, and mutexes?
> Semaphore is a variable that is non-negative and shared between threads.
  Binary semaphore only have the value 0 or 1, which are used for implementing locks by signalling for acheving mutual           exclusion.
  Mutex provides that only one thread can enter into critical section at a time.
  
### What are the differences between channels (in Communicating Sequential Processes, or as used by Go, Rust), mailboxes (in the Actor model, or as used by Erlang, D, Akka), and queues (as used by Python)? 
> Channels Are a model for interprocess communication and synchronization via message passing.
  +++
  Queue is a linear data structure that stores items in FIFO manner,
### List some advantages of using message passing over lock-based synchronization primitives.
> Message passing algorithms tend to be simpler.
  Share data netweem threads or processes. 
 
### List some advantages of using lock-based synchronization primitives over message passing.
> *Your answer here*
