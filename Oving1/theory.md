Exercise 1 - Theory questions
-----------------------------
 
 ### What is the difference between concurrency and parallelism?
 > Concurrency is that two or more threads can start, run and complete in an overlapping time periode. They dont have to do something at the same time. So they split the resources between them.

 Parallelism  is that two threads run at the literally same time.
 
 ### Why have machines become increasingly multicore in the past decade?
 > Becuase we are beginneing to hit a physical limit to how fast a single core can operate. The expectations for computers to become faster and more powerful has not stopped, and a solution to this is to use multicore processors. This move leads to a distribution of load over several cores, increasing the max capacity, even with only minimal increases in single core preformance.
 
 ### Why do we divide software (programs) into concurrently executing parts (like threads or processes)?
 (Or phrased differently: What problems do concurrency help in solving?)
 > Concurrency solves the problem that a single thread can hinder all other threads from running, even when it is "waiting for a message" and not preforming any computations. If we had used concurrency instead, the thread waiting for a message could open up for other threads to run while it is waiting, and then when it recives a message, it could ask for premission to continue its computation. 

 What this does is that we can divide a limited resource on several threads, so more work can be done in total.
 
 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
 > Yes, because we can optimize the usage of the computers resources. At the same time, i think it also will include another layer of problems that need solving. Right now I dont know enough about concurrency to sway either way.
 
 ### What is the conceptual difference between threads and processes?
 > A thread is a segment of a process, while a process is program that is running.
 A thread contains the information the scheduler needs to swap to the next task, while a proscess contains multiple threads and some more data.
 
 ### Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they?
 > Threads can consist of several fibers. However they can not run at the same time; they are co-operative.
 
 ### What is the Go-language's "goroutine"? A C/POSIX "pthread"?
 > A goroutine is a lightweight thread managed by the Go runtime.
 pthread is a way i C to creat multithreading. So in a way it is the same as goroutines in GO.
 
 ### In Go, what does `func GOMAXPROCS(n int) int` change? 
 >GOMAXPROCS limits the number of operating system threads that can execute user-levle Go code.



 
 
 
 
