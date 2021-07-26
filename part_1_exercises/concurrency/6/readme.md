**Objective:**
You are tasked with creating an online classroom system where students join and leave classes. Use wait groups to ensure all students have joined before starting the class.

**Steps:**

1. Create a function that simulates a student joining a class.
2. Use a wait group to track each student joining.
3. In the main function, spawn multiple go routines for students joining and wait for all of them before starting the class.