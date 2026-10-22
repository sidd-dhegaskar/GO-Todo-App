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

command for logging into postgres:
psql -U postgres

command to check the DBs:
\l

NOTE: we always write SQL queries in uppercase
";" terminate the line

command to create database:
CREATE DATABASE todo_api;

command to select database:
\c todo_api;

command to exit postgres:
\q

14. Need to add postgres connection string to the application and make ENV file

15. made a config file with this command:
mkdir -p internal/config && touch internal/config/config.go

16. Configure the confing file.
why is it package config and not config main?
Here is the new information formatted to match your existing study notes:

# Go Project Structure & Packages

## 1. The `package` Declaration
*   Every Go file starts with a `package` declaration, which names the package the file belongs to.
*   **Convention:** The package name should match the name of the directory it resides in (e.g., files in the `internal/config` directory should use `package config`).
*   **Usage:** The package name is what you type before the dot when calling its functions from another file (e.g., `config.Load()`).

## 2. `package main` (The Executable)
*   **Special Purpose:** `package main` is unique. It tells the Go compiler to build an executable program rather than a shared library.
*   **Entry Point:** It *must* contain a `func main()` function, which serves as the starting point of your application.
*   **Where to use it:** Only in the entry-point directory of your project (e.g., `cmd/api/main.go`).
*   **Important Rule:** You *cannot* import a `package main` into another file. Go will reject it with an error.

## 3. Library Packages
*   Any directory that isn't the main entry point is considered a library package.
*   These are not standalone programs; they are designed to be imported and used by other parts of your application (e.g., `import "your-module/internal/config"`).

## 4. The `internal/` Directory Rule
*   **Go-Specific Feature:** Any code placed inside a directory named `internal/` can *only* be imported by other code within your own module.
*   **Why use it:** It acts as a privacy boundary, preventing external projects from importing your app-specific code (like database configurations or internal business logic).

## 5. Example: Structuring a Config Package
*   **File Path:** `internal/config/config.go`
*   **Declaration:** Starts with `package config`.
*   **Contents:** Typically contains a `Config` struct (e.g., holding a `DatabaseURL` string) and a `Load()` function to read environment variables.
*   **Execution:** Your `main.go` file imports this package and calls `config.Load()` at startup to initialize the app's configuration.

17. in the config file, create a config struct and create a function to load it
Here is the detailed explanation of your `config.go` file, formatted to match your existing study notes:

# Configuration Package & Environment Variables

## 1. Structs in Go (`type Config struct`)
*   **What they are:** Structs in Go are used to define custom data structures. You can think of them as similar to classes or objects in JavaScript/TypeScript. They hold properties (key-value pairs).
*   **In this code:** The `Config` struct acts as a blueprint to hold our application's configuration data. It contains two properties: `DatabaseURL` and `Port`, both of which are strictly typed as `string`.

## 2. Exported Functions (Capitalization Rule)
*   **The `Load()` Function:** Notice that the function name starts with a capital `L`. 
*   **Why it matters:** In Go, if a variable, struct, or function starts with a **capital letter**, it is "exported" (public). This means it can be imported and used by other packages (like your `main.go` file). If it started with a lowercase `l`, it would be private to the `config` package.

## 3. Multiple Return Types & Pointers
*   **Signature:** `func Load() (*Config, error)`
*   **Multiple Returns:** Go functions can return more than one value. Here, it returns the configuration data *and* an error status.
*   **Pointer Return (`*Config`):** Instead of returning a heavy copy of the `Config` struct, the function returns a pointer to it. This is consistent with Go's focus on performance—passing the memory address rather than copying the whole object.

## 4. Loading the `.env` File & Error Handling
*   **`godotenv.Load()`:** This function from the `joho/godotenv` package looks for a `.env` file in your project and loads its variables into the system's environment.
*   **Error Handling:** It returns an error if it fails (e.g., the file is missing). 
*   **The `if err != nil` block:** In Go, `nil` means "null" or "nothing". If `err` is *not* `nil`, it means an error occurred. The code catches this and logs a message (`"env file not found"`).

## 5. Security & Environment Variables
*   **Best Practice:** Real-world projects never hardcode sensitive information (like database passwords or API keys) directly into the code. Instead, they are stored in a `.env` file.
*   **`os.Getenv("VARIABLE_NAME")`:** This method from Go's standard `os` package reaches into the operating system's environment variables and grabs the value associated with the key (e.g., `"DATABASE_URL"`). The keys must exactly match what is written in your `.env` file.

## 6. Initializing the Struct Pointer (`&`)
*   **Code:** `var config *Config = &Config{ ... }`
*   **The Ampersand (`&`):** While the asterisk (`*`) is used to *declare* a pointer variable, the ampersand (`&`) is used to *get the memory address* of a value. 
*   **What is happening:** This block creates a new instance of the `Config` struct, populates it with the values fetched from `os.Getenv()`, grabs its memory address using `&`, and stores that address in the `config` pointer variable.
*   **Final Step:** `return config, nil` sends the pointer back to whoever called the function, along with `nil` for the error (since it succeeded).

Here is a deeper dive into how the ampersand (`&`) works in Go, specifically answering your question about whether we are writing to or fetching from that address.

# Deep Dive: Pointers, `&`, and Memory Addresses

To answer your question directly: **At the exact moment we use the `&`, we are neither writing to the address nor fetching from it. We are simply *finding out what the address is* so we can hand it to someone else.**

Let's break down exactly what happens in this line of code:
`var config *Config = &Config{ DatabaseURL: "...", Port: "..." }`

## 1. The Creation Phase: `Config{ ... }`
Before the ampersand does anything, Go looks at the right side of the code: `Config{ DatabaseURL: "...", Port: "..." }`.
*   **What happens:** Go acts like a construction worker. It finds an empty plot of land in your computer's memory, builds a new `Config` struct there, and **writes** the database URL and port into that space.
*   *Analogy:* You just built a new apartment and put furniture inside it.

## 2. The Ampersand Phase: `&` (The "Address-Of" Operator)
Now that the struct exists in memory, the ampersand (`&`) steps in. 
*   **What happens:** The `&` operator looks at that newly built struct and asks, *"Where exactly was this built?"* It grabs the physical memory address (e.g., `0x14000123456`). 
*   *Analogy:* The `&` is you writing down the street address of that new apartment on a piece of paper. You aren't changing the apartment, and you aren't looking at the furniture inside it yet. You are just writing down *where* it is.

## 3. The Assignment Phase: `var config *Config =`
Finally, we store that address.
*   **What happens:** We take that memory address (the piece of paper) and save it inside the `config` variable. The asterisk (`*`) in `*Config` just tells Go: *"Hey, this variable isn't holding a struct; it's holding a memory address that points to a struct."*

---

## Summary: Writing vs. Fetching

*   **When do we write?** We write to memory the moment we declare `Config{...}`. Go allocates the space and fills it with the environment variables.
*   **What does `&` do?** It just generates the "map" or "address" to that space.
*   **When do we fetch?** We fetch *later*. When this function finishes, it returns this address to your `main.go` file. Later on, when your main file wants to know the Database URL, it will use this address to go to that exact spot in memory and **fetch** the value. 

By returning the address (`&`) instead of the whole struct, you are handing your `main.go` file a lightweight piece of paper with an address on it, rather than forcing Go to pick up the entire apartment and copy it over to `main.go`.


17. Create a postgres.go file in the database folder
This is to centralise the database logic and conection pooling in a single place.

context is going to be used to control the database operations
Here are your study notes for the database connection logic, formatted to match your previous notes:

# Database Connection & Pooling

## 1. Connection Pooling (The "Why")
*   **The Problem:** Opening a new database connection for every single API request is very slow (taking 20-50 milliseconds each). If you get 100 requests per second, your API will bottleneck and crash.
*   **The Solution (Pooling):** A connection pool opens a set number of connections (e.g., 10-20) and keeps them alive. 
*   **How it works:** When an API request comes in, it "borrows" a connection from the pool. When the database query finishes, the connection is returned to the pool to be reused by the next request. 
*   **Pointers again:** Because we are sharing and borrowing this pool across the whole app, we pass it around as a pointer (`*pgxpool.Pool`).

## 2. The `Connect` Function Signature
*   **Code:** `func Connect(databaseURL string) (*pgxpool.Pool, error)`
*   **Capitalization:** Starts with a capital `C` so it is exported and can be called from `main.go`.
*   **Returns:** It returns a pointer to the connection pool (`*pgxpool.Pool`) and an `error` (if something fails).

## 3. The `context` Package
*   **Code:** `var ctx context.Context = context.Background()`
*   **What it is:** In Go, a `context` is used to carry deadlines, cancellation signals, and request-scoped values across API boundaries.
*   **Database Usage:** We pass `ctx` to database functions so that if a query takes too long (e.g., the database hangs), we can cancel the operation instead of freezing the whole server. `context.Background()` is the standard starting point for an empty context.

## 4. Parsing the Database URL
*   **Code:** `config, err = pgxpool.ParseConfig(databaseURL)`
*   **Purpose:** The `databaseURL` is just a long string (containing the user, password, port, and DB name). `ParseConfig` breaks this string down into a structured `Config` object that the `pgxpool` package can actually understand.

## 5. Creating the Connection Pool
*   **Code:** `pool, err = pgxpool.NewWithConfig(ctx, config)`
*   **Purpose:** This actually establishes the pool of connections using the parsed configuration and the context. It returns the pointer to the pool (`*pgxpool.Pool`).

## 6. Pinging the Database (`pool.Ping`)
*   **Code:** `err = pool.Ping(ctx)`
*   **Why do this?** Just because the pool was created doesn't mean the database is actually turned on or reachable. `Ping` sends a tiny test signal to the database.
*   **Error Handling:** If the ping fails, we log the error, immediately close the broken pool (`pool.Close()`), and return the error so the app knows it failed to connect.

## 7. Integration & Cleanup in `main.go` (From Transcript)
*   When you call this `Connect` function in your `main.go` file, you need to ensure the pool eventually closes when the server shuts down.
*   **The `defer` keyword:** In `main.go`, you will use `defer pool.Close()`. `defer` tells Go: *"Schedule this function to run at the very end, right before the current function (`main`) exits."* This ensures clean memory management and prevents ghost connections to your database.