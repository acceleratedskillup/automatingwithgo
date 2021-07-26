**Objective:**
You are tasked with creating a downloader that fetches multiple files concurrently. Use wait groups to ensure all downloads are complete before the program exits.

**Steps:**

1. Create a function that simulates downloading a file.
2. Use a wait group to track the completion of each download.
3. In the main function, spawn multiple go routines to start the downloads and wait for all of them to complete.