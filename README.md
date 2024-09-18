# Golang Web Calculator
This is a simple web-based calculator built with Golang using the `net/http` package for handling HTTP requests. The web interface is developed using HTML and CSS, allowing users to perform basic arithmetic operations like addition, subtraction, multiplication, and division.

## Features
- Addition, subtraction, multiplication, and division.
- Simple and interactive web interface.
- Uses Go's `net/http` package to serve the web page.

## Prerequisites
Before running the project, make sure you have the following installed:
- [Go](https://golang.org/doc/install) (version 1.20+)
- A web browser

## Project Structure

```bash
-Go-Calculator
 ├── dockerfile # Docker file
 ├── main.go # Main Go file for handling requests 
 ├── go.mod # Go module file 
 └── templates/ 
    └── index.html # HTML template for the calculator UI
```

### Installation and Setup
1. **Clone Repository:**
```bash
git clone https://github.com/Gowtham-R-dev/Fo-Calculator.git
cd Go-Calculator
```

2. **Initialize Go Modules**
```bash
go mod tidy
go mod init <name>
```

3. **Run the Go Program**
```bash
go run main.go
```

4. **Access the Web Application**
```bash
http://localhost:8080
```

### Docker Build and Run Commands
1. **Build the Docker Image**
Run the following command in the root of your project directory:
```bash
docker build -t go-calculator .
```

2. **Run the Docker Container**
After the image is built, you can run the container using:
```bash
docker run -p 8080:8080 go-calculator .
```

3. **Access the Web Application**
```bash
http://localhost:8080
```
This setup will allow you to easily build and run your Go-based web calculator using Docker!
