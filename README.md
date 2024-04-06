# go-fan-out-fan-in-pipeline

The fan-out/fan-in pattern is a concurrency design pattern commonly used in Go for parallelizing and coordinating concurrent tasks. It is particularly useful when you have a time-consuming task that can be divided into smaller subtasks that can be executed concurrently.

## How it works

### The pattern consists of two main stages: fan-out and fan-in.

Fan-out: In the fan-out stage, a single task is divided into multiple smaller subtasks, which are then executed concurrently. Each subtask can be assigned to a separate goroutine (lightweight concurrent thread in Go) to run in parallel. This stage distributes the workload across multiple goroutines, allowing for parallel processing.

Fan-in: In the fan-in stage, the results or outputs from all the concurrently executing subtasks are collected and combined into a single result. This stage waits for all the subtasks to complete and aggregates their results. The fan-in stage can also handle synchronization and coordination between the goroutines to ensure that all results are collected before proceeding.

<img width="953" alt="Screenshot 2024-04-07 at 1 42 53 AM" src="https://github.com/AmitrajitDas/go-fan-out-fan-in-pipeline/assets/60144525/6760b5e3-a509-4f7e-9cd0-91b992f15895">

## Prerequisites

- Go installed on your machine. You can download and install it from [here](https://golang.org/dl/).

## Usage

1. Clone the repository to your local machine:

```bash
git clone https://github.com/AmitrajitDas/go-fan-out-fan-in-pipeline.git
```

2. Execute

```bash
go run main.go
```
