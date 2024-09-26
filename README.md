Project Title: Monthly Expense Tracker

Description: A web application designed to help users manage their monthly expenses effectively 
by providing intuitive tools to add, track, update, and delete expenses, ensuring that users stay within their budget.

- Technologies Used: Go, PostgreSQL, Gin, GORM

Key Features:
 - Salary and Expense Management: Users can submit their monthly salary, and expenses are validated in real-time to ensure they don't exceed the remaining salary. 
The app provides an overview of all expenses, including the total amount spent and the remaining balance.
 - CRUD Operations: Full support for creating, reading, updating, and deleting expenses using a RESTful API designed with the Gin framework, 
integrated with GORM for database interaction.
 - Error Handling and Validation: Error handling and data validation to ensure accurate and correct data entry,
providing clear feedback for invalid inputs using status codes and responding with a JSON response.

My role:
 - Backend Development: Designed and implemented the backend API using the Gin web framework in Go,
ensuring efficient routing and integration with the database for salary and expense management features.
 - Database Integration: Integrated GORM for handling database operations, including salary and expense CRUD functionality, ensuring smooth and reliable data storage and retrieval.

Challenges and Solutions:
 - Challenge 1: Ensuring that expenses do not exceed the remaining salary, and handling cases where users input incorrect data.
 - Solution 1: Implemented data validation using Gin’s bind JSON and custom error handling to provide clear error messages and maintain data consistency.

 - Challenge 2: Users must input their monthly salary before entering expenses to prevent errors.
 - Solution 2: Implemented a check for salary requirements before inputting expenses. An error will be displayed if expenses are entered first.

Results and Impact:
Achieved a 95% satisfaction rate from 5 testing users.

Links: 
[Live Demo](https://drive.google.com/file/d/1hWUrDZC_jcpqRa5A08EuZS134KdjixD9/view?usp=drive_link)
[GitHub Repository](https://github.com/Zmohamed6991/Expense-Tracker)

Lessons Learned: Improved skills in backend optimisation and gained experience using RESTful API with GORM for database interaction. For example, retrieving data from the database using gorm first and find.

Lessons Learned: Improved skills in backend optimisation and gained experience using RESTful API with GORM for database interaction. 

**Full API Documentation**


# Setup Steps:
1. Go installed (version 1.16 or higher): https://golang.org/dl/
Gin framework: Install using the terminal:

go get -u github.com/gin-gonic/gin

GORM ORM (for database interaction): Install using the terminal
	go get -u gorm.io/gorm
2. PostgreSQL: Download and install PostgreSQL to manage the database. You can also install pgAdmin 4 to manage the database via GUI - https://www.postgresql.org/download/ 
3. Postman: Download Postman to test API endpoints: https://www.postman.com/downloads/
4. Clone the repository: git clone https://github.com/Zmohamed6991/Expense-Tracker.git
5. Go Modules (for dependency management): Run this command in the terminal: 
go mod init example/expense-tracker

# Running the Application:
To start the application use:
go run main.go

The application will start on the web server on https://localhost:8080


# Testing Endpoints using Postman:

1. **Add Monthly Salary**
Adds user’s monthly salary
Method: POST
URL: https://localhost:8080/salary

Request body example: 
{
    "monthly_salary": 500
}

Success response with status code 200 OK.
{
    "ID": 1,
    "CreatedAt": "2024-09-24T18:46:49.153877+01:00",
    "UpdatedAt": "2024-09-24T18:46:49.153877+01:00",
    "DeletedAt": null,
    "monthly_salary": 500,
    "remaining_salary": 500
}

Error responses:
- 400 Bad Request: If invalid data is inputted.
- 404 Not found: If there is no salary input.

-----

2. **Add An Expense**
User can start adding their expenses
Method: POST
URL: https://localhost:8080/add
Request Body example:
{
    "expense_name": "Netflix",
    "amount": 9.99,
    "category": "Entertainment"
}

Success response with status code 200 OK. It also responds with the remaining salary after adding the expense.
 

   "Expense added": {
        "ID": 2,
        "CreatedAt": "2024-09-24T18:47:12.8118176+01:00",
        "UpdatedAt": "2024-09-24T18:47:12.8118176+01:00",
        "DeletedAt": null,
        "expense_name": "Netflix",
        "amount": 9.99,
        "category": "Entertainment"
    },
    "Remaining salary": 90.01,
    "Total amount of expenses": [
        9.99
    ]


Error responses:
404 Not Found: 
- if the user adds an expense before inputting a monthly salary,
- if no remaining salary is available.

400 Bad Request: 
- If the user inputs invalid data,
- If the expense amount is 0 or less.
- If there is no input for expense name or category.
- If the total amount exceeds the remaining salary.

500 Internal Server Error:
- If there is an error creating a total table in the database.
- If there is an error updating the remaining salary table in the database.

-----


3. **Get All Expenses**
Retrieves all expenses for the user
Method: GET
URL: https://localhost:8080/all
No request body, click send and it will retrieve all expenses in the body. Including the remaining salary. 

		
[
    {
        "ID": 1,
        "CreatedAt": "2024-09-24T18:46:53.983546+01:00",
        "UpdatedAt": "2024-09-24T18:47:33.613354+01:00",
        "DeletedAt": null,
        "expense_name": "rent",
        "amount": 450,
        "category": "House Bills"
    }
]
{
    "ID": 1,
    "CreatedAt": "2024-09-24T18:46:49.153877+01:00",
    "UpdatedAt": "2024-09-24T18:47:47.283083+01:00",
    "DeletedAt": null,
    "monthly_salary": 500,
    "remaining_salary": 50.00000000000001
}

Success response with status code 200 OK. It also responds with the remaining salary after adding the expense. 

Error response:
404 Not Found: 
- If there are no expenses in the database.
- If it cannot retrieve the monthly salary and remaining salary.

------

4. **Update an Expense**
Updates an expense’s amount by its ID
Method: PUT
URL: https://localhost:8080/update/:id
Enter the ID of the expense in the URL then use the request body to update the amount of the expense. 

Request Body example: 

{    
"amount": 450.00
}

Success response with status code 200 OK. With the updated remaining salary, update amount and the name of the expense.

{
    "remaining salary": 40.010000000000005,
    "updated expense": 450,
    "updated expense name": "rent"
}


Error response:
400 Bad Request: 
- If there is an invalid data input when binding.
- If the updated amount exceeds the remaining salary.

404 Status Not Found: 
- If there are no expenses in the database.
- If the requested expense is not in the database.
- If there is no salary input.

500 Internal Server Error:
- If it fails to update the remaining salary with the updated amount.
- If it fails to update the expense table in the database

-------

5. **Delete an expense**
Deletes an expense by its ID 
Method: DELETE
URL: https://localhost:8080/delete/:id
Enter the ID of the expense in the URL then send.

Success response with status code 200 OK. Also with the expense ID, a message to confirm the expense is deleted and the remaining salary is for the user to keep track of.

{
    "expense id": 2,
    "message": "Expense deleted",
    "updated remaining salary": 50.00000000000001
}

Error response:
404 Status Not Found:
- If the expense is not found in the database.
- If the salary is not found in the database.

500 Internal Server Error:
- If there is an error updating the remaining salary table in the database.



