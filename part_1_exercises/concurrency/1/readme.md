**Objective:**
You are tasked with creating a simple bank system where multiple go routines will attempt to deposit and withdraw money. Your goal is to ensure that the bank's balance remains consistent and free from data races using mutexes.

**Steps:**

1. Create a Bank struct with a balance and a mutex.
2. Implement Deposit and Withdraw methods for the Bank struct.
3. Spawn multiple go routines to simulate concurrent deposits and withdrawals.
4. Print the final balance after all operations are done.