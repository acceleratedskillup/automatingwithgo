**Objective:** Create a function that simulates a basic calculator with a twist. The calculator should be able to perform addition, subtraction, multiplication, and division. However, instead of using the typical arithmetic symbols, the function should take in words like "add", "subtract", "multiply", and "divide". If the operation word is "exit", the calculator should stop processing and return a message "Calculator exited". Use labels and for loops to achieve this.

**Steps:**
1. Define a function named loopingCalculator that takes a slice of numbers and a slice of operations in the form of words.
2. Use a labeled for loop to iterate over the operations.
3. Inside the loop, use a switch-case structure to determine the operation.
4. For each operation, perform the corresponding arithmetic on the numbers slice and update the slice.
5. If the operation word is "exit", break out of the loop using the label.
6. Return the modified numbers slice and any messages.
7. In the main function, test the loopingCalculator function with various numbers and operations.