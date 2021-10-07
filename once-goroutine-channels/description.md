# Running a goroutine inside a `once` construct and make it communicate with the outer caller through a channel. 
**Description**

Experiment to find out how goroutines work inside a `once` construct and communicate with the outer caller through a channel.
**Findings**
1. The `once.Do(f)` returns once the `f` function returns. 
1. `f` must be *niladic*
1. `once.Do(f)` returning doesn't terminate the `goroutine` triggered inside it. 
