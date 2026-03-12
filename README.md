# URL-Shortener

----


A minimal HTTP server built in Go designed to shorten long URLs into manageable, Base62-encoded links.

----

## Installation & Setup

1. ##### Clone the repository
   ```bash
   git clone https://github.com/mmmion/url-shortener.git
   cd url-shortener
   ```

2. ##### Build the project
    This command creates a build/ directory and compiles the binary.
    ```bash
   make build
   ```
   
   
3. ##### Running the Application
    *Option A*: Development Mode 
    ```bash
    make run
    ```
    *Option B*: Binary execution (Windows only)
    ```bash
    # Windows
    ./build/url_shortener.exe
    ```
