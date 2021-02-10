#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t lock; 
//Chose mutex since we want to give exclusive access to one thread at the time. Semaphore is better to use when we have to wait for another thread to finish a operation before we continue.

// Note the return type: void*
void* incrementingThreadFunction(){
  for (int j = 0; j < 1000001; j++){
    pthread_mutex_lock(&lock);
	  i++;
    pthread_mutex_unlock(&lock); 
  }
  return NULL;
}

void* decrementingThreadFunction(){
  for (int j = 0; j < 1000000; j++){
    pthread_mutex_lock(&lock);
    i--;
    pthread_mutex_unlock(&lock); 
  }
  return NULL;
}


int main(){
    pthread_t incrementingThread, decrementingThread;

    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);

    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);

    printf("The magic number is: %d\n", i);
    return 0;
}
