Interview
---

## Question 1:

Implement Workqueue, a task queue designed to execute tasks in a First-In-First-Out (FIFO) order. If a task encounters an error during execution, Workqueue will automatically requeue it for retries. Each task can be executed up to a maximum of three times. If it still fails after the third attempt, it will be discarded. Additionally, retries follow an incremental delay strategy: the first retry occurs after a 2-second wait, while the second retry takes place after a 4-second wait.

You shall write your code in `workqueue.go` file. You can test the implementation with `make test`. You may checkout `workqueue_test.go` for more details.
