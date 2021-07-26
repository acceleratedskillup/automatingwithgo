**Objective:**
Create a simple calculator that performs operations sent to it via channels. The calculator should be able to handle multiple operations concurrently.

**Steps:**

Define a struct for a mathematical operation with operands and an operator.
Implement a function that performs the operation and sends the result back via a channel.
In the main function, send multiple operations to the calculator and print the results.