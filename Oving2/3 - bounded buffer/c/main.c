// compile with:  gcc -g main.c ringbuf.c -lpthread

#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <semaphore.h>
#include <time.h>

#include "ringbuf.h"

struct BoundedBuffer {
    struct RingBuffer*  buf;
    pthread_mutex_t     mtx;
    sem_t               empty;
    sem_t               full;
    
    
};

struct BoundedBuffer* buf_new(int size){
    struct BoundedBuffer* buf = malloc(sizeof(struct BoundedBuffer));
    buf->buf = rb_new(size);
    
    pthread_mutex_init(&buf->mtx, NULL);
    // TODO: initialize semaphores

    sem_init(&buf->full,  0, size);
	sem_init(&buf->empty, 0, 0);
    
    return buf;    
}

void buf_destroy(struct BoundedBuffer* buf){
    rb_destroy(buf->buf);
    pthread_mutex_destroy(&buf->mtx);
    sem_destroy(&buf->empty);
    sem_destroy(&buf->full);
    free(buf);
}




void buf_push(struct BoundedBuffer* buf, int val){    
    // TODO: wait for there to be room in the buffer
    sem_wait(&buf->full); //Decrement semaphore. If buffer is full, semaphore is 0. If semaphore is 0 it blocks. 
    sem_post(&buf->empty); //Increment semaphore

    // TODO: make sure there is no concurrent access to the buffer internals
    pthread_mutex_lock(&buf->mtx);
    rb_push(buf->buf, val);
    pthread_mutex_unlock(&buf->mtx);
    
    // TODO: signal that there are new elements in the buffer  
}

int buf_pop(struct BoundedBuffer* buf){
    sem_post(&buf->full); //Increment semaphore
    sem_wait(&buf->empty); //Decrement semaphore. If buffer is full, semaphore is 0. If semaphore <= 0 it blocks. 

    pthread_mutex_lock(&buf->mtx);
    int val = rb_pop(buf->buf);    
    pthread_mutex_unlock(&buf->mtx);

    return val;
}





void* producer(void* args){
    struct BoundedBuffer* buf = (struct BoundedBuffer*)(args);
    
    for(int i = 0; i < 10; i++){
        nanosleep(&(struct timespec){0, 100*1000*1000}, NULL);
        printf("[producer]: pushing %d\n", i);
        buf_push(buf, i);
    }
    return NULL;
}

void* consumer(void* args){
    struct BoundedBuffer* buf = (struct BoundedBuffer*)(args);
    
    // give the producer a 1-second head start
    nanosleep(&(struct timespec){1, 0}, NULL);
    while(1){
        int val = buf_pop(buf);
        printf("[consumer]: %d\n", val);
        nanosleep(&(struct timespec){0, 50*1000*1000}, NULL);
    }
}

int main(){ 
    
    struct BoundedBuffer* buf = buf_new(5);
    
    pthread_t producer_thr;
    pthread_t consumer_thr;
    pthread_create(&producer_thr, NULL, producer, buf);
    pthread_create(&consumer_thr, NULL, consumer, buf);
    
    pthread_join(producer_thr, NULL);
    pthread_cancel(consumer_thr);
    
    buf_destroy(buf);
    
    return 0;
}