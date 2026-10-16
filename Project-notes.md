Hello there!

1. Make the folder structure with this command:
mkdir -p cmd/api internal/config internal/database internal/handlers internal/middleware internal/models internal/repository migration

2. cmd/slappi/ – Main entry point of the application, containing main.go to start and run the server.

internal/config/ – Manages application configuration and loads environment variables.

internal/database/ – Handles database connection setup and connection pooling.

internal/handlers/ – Processes HTTP requests and formulates HTTP responses.

internal/middleware/ – Handles authentication and other request-processing tasks.

internal/models/ – Defines data structures using Go structs.

internal/repository/ – Handles database operations such as creating, reading, updating, and deleting data.

migrations/ – Contains database migration files for applying changes (up) and reverting changes (down).

API/main.go – Entry point that runs the application server and defines the initial server setup.

go.mod – Defines the Go module name and manages project dependencies.

3. - go mod init <module-name> – Initializes a new Go module, creates the go.mod file, and defines the module name (e.g., go mod init todo-api).

4. added gin framework: go get -u github.com/gin-gonic/gin

5. added postgres driver: go get github.com/jackc/pgx/v5

6. added connection pooler: go get -u  github.com/jackc/pgx/v5/pgxpool

7. added jwt auth: go get -u github.com/golang-jwt/jwt/v5   

8. added go-migrate: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest   

9. added password encrypter: go get golang.org/x/crypto/bcrypt   

10. added support for env: go get github.com/joho/godotenv   

11. added go air package for support server changes dynamically: go install github.com/air-verse/air@latest   

12. Here is the updated version of your notes, with the detailed explanations for the callback function, the context, and how `gin.H` maps work:
# Go & Gin Framework: Web Server Basics

## 1. The Router
*   **Definition:** The core component responsible for accepting client requests, differentiating between route types, and sending back responses.
*   **Role:** It is the first step in setting up a web server. It maps specific URLs to the code that should handle them.

## 2. Pointers & Performance in Go
*   **What they are:** Pointers (denoted by an asterisk `*`) store the memory address of an instance rather than copying the instance itself.
*   **Why use them:** Copying large objects (like a server engine) is slow and consumes memory. Pointers allow you to reference the existing object directly.
*   **Analogy:** Renting an existing apartment (using a pointer/address) is much faster and cheaper than building a brand new house from scratch (copying the object).

## 3. Initializing the Gin Router
*   **Gin Framework:** A fast web framework for Go used to handle the request/response cycle.
*   **`gin.Default()`:** A method that creates and returns a pointer to a new Gin `Engine` instance. It comes pre-attached with useful middleware (logger and recovery).
*   **Code Explanation:** `var router *gin.Engine = gin.Default()`
    *   This explicitly declares `router` as a pointer variable (`*gin.Engine`).
    *   **Strict Pointer Design:** The Gin framework is explicitly designed to require pointers for its core structures. If you try to declare it as a regular value (without the `*`), the compiler will throw an error. This enforces referencing the memory address rather than copying the heavy `Engine` object, ensuring your API remains fast.

## 4. Defining Routes & The Callback Function
*   Routes are defined by combining a URL path (e.g., `/`) with an HTTP method (e.g., GET).
*   **Callback Function:** When a client makes a request to a specific route, the server executes a callback function to handle that exact request.
*   **Code Example:** `router.GET("/", func(c *gin.Context) { ... })`
    *   The first argument (`"/"`) is the route path the user visits.
    *   The second argument (`func(c *gin.Context)`) is the callback function that dictates what happens when that route is hit.

## 5. The Context Object (`*gin.Context`)
*   The callback function takes a **Context** argument (usually named `c`).
*   **Purpose:** The context encapsulates *both* the incoming HTTP Request and the outgoing HTTP Response. It gives you the tools to read what the client sent and formulate what to send back.
*   Like the router, it is passed as a pointer (`*gin.Context`) so the server doesn't have to copy the heavy request/response data in memory.

## 6. Formulating JSON Responses & Maps (`gin.H`)
*   You can send JSON back to the client using `c.JSON()`, passing the HTTP status code (e.g., `200` for OK) and the data.
*   **`gin.H`:** A convenient shorthand provided by Gin to create JSON objects. 
*   **How it works under the hood:** It is essentially a Go map defined as `map[string]interface{}` (or `map[string]any` in modern Go).
    *   **`string` (The Key):** Represents the property names in your JSON object, such as `"message"`, `"status"`, or `"error"`.
    *   **`interface{}` / `any` (The Value):** Means the value assigned to that key can be of *any* data type (a string, a number, a boolean, etc.).
*   **Code Example:**
    ```go
    c.JSON(200, gin.H{
        "message": "todo API is running", // "message" is the string key, the text is the interface{} value
        "status": "success",              // "status" is the string key
    })
    ```

## 7. Starting the Server
*   To start listening for requests, call the `Run` method on your router and pass the desired port.
*   **Code Example:** `router.Run(":3000")`
*   **Tooling Note:** You can use a tool like `air` in your terminal to run the server with automatic live-reloading.

13. installed postgres
