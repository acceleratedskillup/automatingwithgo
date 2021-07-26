**Database Initialization with init Function**

**Objective:** Understand the use of the init function in Go by simulating the initialization of a database connection.

**Background:** The init function in Go is designed to handle initial setup tasks before the main logic of the program runs. In this exercise, we'll simulate the initialization of a database connection.

**Scenario:** You are building a simple application that interacts with a mock database. Before any database operation can be performed, a connection to the database must be established. Use the init function to simulate this connection setup.

**Steps:**

1. Create two global string variables: databaseStatus and queryMessage.
2. Use the init function to set databaseStatus with the message "Database connected!" and queryMessage with the message "Querying the database...".
3. Define a function displayStatus that prints the value of databaseStatus.
4. Define a function queryDatabase that prints the value of queryMessage.
5. In the main function, call displayStatus and then queryDatabase.