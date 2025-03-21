Interview
---

## Question 1:

Implement Workqueue, a task queue designed to execute tasks in a First-In-First-Out (FIFO) order. If a task encounters an error during execution, Workqueue will automatically requeue it for retries. Each task can be executed up to a maximum of three times. If it still fails after the third attempt, it will be discarded. Additionally, retries follow an incremental delay strategy: the first retry occurs after a 2-second wait, while the second retry takes place after a 4-second wait.

You shall write your code in `workqueue.go` file. You can test the implementation with `make test`. You may checkout `workqueue_test.go` for more details.

<!--
## Other Questions

- Your resume states that you are proficient in Golang and C++. Could you elaborate on your in-depth experience with these languages?
- When faced with a serious performance bottleneck in a high-concurrency system, what is your approach to troubleshooting and resolving the issue?
- You contributed to the implementation of Kubernetes GPU sharing in Alibaba Cloud PAI. Could you provide a detailed explanation of the technical solution used for GPU sharing?
- Can you elaborate on the optimizations you made to the vector recall components?
- Could you share insights into your customized Kubernetes scheduling strategies?
- What has been the most challenging task you have ever encountered, and how did you tackle it?
-->
