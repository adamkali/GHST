# Ghost
**GHST stack (Gin - HTMX - SurrealDB - and Tailwindcss)** is a BLAZINGLY fast framework that sends down minimal JavaScript to the client.

```
       _,.,
     /'    '\
    / ()  () \
 \""          ""/    
  '-_    o     /    
    |        ./
  ,;______,-'
```

## Overview

Ghost is a spooky high-performance web framework that leverages the GHST stack to deliver fast and responsive web applications. It combines the power of Gin for server-side logic, HTMX for dynamic HTML updates, SurrealDB for futeristic and surreal data storage, and Tailwind CSS for stylish UI components. Ghost is designed to minimize the use of client-side JavaScript, ensuring your web apps load quickly and run efficiently.

_As a side note: Ghost does not solve a problem. It solves the inconvienience of setting up these services and then packages them into a tool; however, in a personal opinion, this is the best stack for making fullstack apps._

## Features

- **Blazing Fast**: Ghost is optimized for speed, delivering near-instantaneous page loads and responsive interactions.
- **Minimal JavaScript**: Ghost minimizes the use of client-side JavaScript, reducing the burden on the client's device.
- **Modern Stack**: Utilizes the GHST stack, which combines the best technologies for building modern web applications.
- **SurrealDB Integration**: Easily connect to SurrealDB to store and manage your data with speed and reliability.
- **Tailwind CSS**: Create beautiful and customizable user interfaces with the power of Tailwind CSS.

## Requirements

Before you can use Ghost, please ensure that the following requirements are met:

### Golang
Ghost is built with Go (Golang). Make sure you have Go installed on your machine. You can download it from the [official website](https://golang.org/dl/).

### Tailwind CSS Cli
Ghost relies on Tailwind CSS for styling. To use Ghost effectively, you should have the Tailwind CSS CLI installed and runnable. Visit the [home page](https://tailwindcss.com/blog/standalone-cli) and make sure that the cli is runnable in a shell and the executable directory is sourced to the Path environment variables for your operating system. You can verify the installation with: 

```bash
tailwindcss --help
```

### SurrealDB
Ghost does not yet manage your installation of SurrealDB. So you need to have an installation of SurrealDB that is open to listening to `Port 8000`. See the [SurrealDB Installation Guide](https://surrealdb.com/docs/installation) and verify the installation by: 

```bash
surreal help
```

## Installation 
You can use `go install` to download the installation: 
```bash
go install github.com/adamkali/ghost
```

## Quickstart
`ghost` is made to work as a wrapper for these utilities. To start:

```bash
ghost new <project name> 
```

## Commands 
The following is a combined list of commands which ghost uses:
```bash
# To create a new project
ghost new <project name>

# Run that project
ghost run

# Build the production project
ghost build

# To checkhealth
ghost checkhealth
```

Use `-h` on any command in to see the entire list of options or `ghost help` for
more top level information 


