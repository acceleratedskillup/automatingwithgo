**Objective:**
You are part of a team managing a spacecraft exploring a distant planet. The spacecraft sends periodic data back to Earth, but due to the vast distance, there are often delays. Your task is to create a communication system that can handle data packets from the spacecraft and timeouts if the data isn't received within a certain period.

**Steps:**

1. Create a DataPacket struct with fields for the data type and its value.
2. Implement a function that sends random data packets to a buffered channel.
3. Use a select statement in the main function to either receive data packets or handle a timeout using the time.After function.
4. Print received data or a timeout message based on the select outcome.