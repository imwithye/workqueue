package workqueue

type Task interface {
	Run() error
}

// Implement the logic to enqueue tasks.
//   When the queue is running, tasks should be executed in a First-In, First-Out (FIFO) order.
//   If a task encounters an error, it should be re-enqueued for a retry.
//   Each task can be run up to a maximum of three times. If it fails more than three times, it should be considered as failed (no more retries).
//   Retries should follow a backoff mechanism:
//   The first retry waits 2 second.
//   The second retry waits 4 seconds.

type WorkQueue struct {
	/* Add your code here */
}

// NewWorkQueue creates a new WorkQueue.
func NewWorkQueue() *WorkQueue {
	return &WorkQueue{}
}

// Enqueue will add a task to the queue.
func (w *WorkQueue) Enqueue(task Task) {
	/* Add your code here */
}

// Run will block the thread and execute the tasks in a FIFO order.
// It exists when there are no more tasks to execute.
func (w *WorkQueue) Run() {
	/* Add your code here */
}
