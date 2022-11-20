# Udacity_Golang

This project is a part of Udacity's Golang Course. The main ask here is to build the backend for the CRM project.
The backend has the following features:

1. The API can be accessed via localhost.

2. The application handles the following 5 operations for customers in the "database":

    Getting a single customer through a /customers/{id} path
    Getting all customers through a the /customers path
    Creating a customer through a /customers path
    Updating a customer through a /customers/{id} path
    Deleting a customer through a /customers/{id} path
    
    Each RESTful route is associated with the correct HTTP verb.
    
3. The application leverages the encoding/json package to return JSON values (i.e., not text, etc.) to the user.

4. The home route is a client API endpoint, and includes a brief overview of the API (e.g., available endpoints). Note: This is the only route that does not return a JSON response.

5. The application uses a router (e.g., gorilla/mux, http.ServeMux, etc.) that supports HTTP method-based routing and variables in URL paths.

6. The Handler interface is used to handle HTTP requests sent to defined paths. There are five routes that return a JSON response, and are each is registered to a dedicated handler:

    getCustomers()
    getCustomer()
    addCustomer()
    updateCustomer()
    deleteCustomer()
    
7. If the user queries for a customer that doesn't exist (i.e., when getting a customer, updating a customer, or deleting a customer), the server response includes:

    A 404 status code in the header
    null or an empty JSON object literal or an error message
    
8.  An appropriate Content-Type header is sent in server responses.
9.  The application leverages the io/ioutil package to read I/O (e.g., request) data.
10. The applications leverages the encoding/json package to parse JSON data.
