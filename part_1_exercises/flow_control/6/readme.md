**Objective:** Create a function that simulates a panic and then recovers from it.

**Steps:**
1. Create a function named causePanic.
2. Inside the function, defer another function that calls recover and prints "Recovered from panic".
3. After the defer, simulate a panic by calling the panic function.
4. Test the causePanic function in the main function.