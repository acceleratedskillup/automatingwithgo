**Objective:** Implement a function that uses defer, panic, and recover to handle errors gracefully.

**Steps:**
1. Create a function named safeDivision.
2. Inside the function, defer another function that calls recover and prints "Division by zero error" if a panic occurs.
3. If the denominator is zero, trigger a panic.
4. Otherwise, perform the division and return the result.
5. Test the safeDivision function in the main function with various numbers.s