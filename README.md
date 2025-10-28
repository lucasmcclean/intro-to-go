# Intro to Go

> Knight Hacks workshop on developing web servers with Go.

> 2025-10-28 16:30 EDT

## 1. Install Go

- Download and install from [https://go.dev/dl/](https://go.dev/dl/)
- Verify installation:

```bash
go version
```

- Recommended version: **1.25.x**

## 2. Clone this Repo

```bash
git clone https://github.com/lucasmcclean/intro-to-go
cd intro-to-go
go mod tidy
```

## 3. If Unable to Clone

- Create a new module:

```bash
mkdir my-go-workshop
cd my-go-workshop
go mod init my-go-workshop
```

- Copy the example files (main.go, server/, routes/) into your folder.

- Add dependencies:

```bash
go mod tidy
```

## 4. Run the Project

- To run the UI (optional):

```bash
cd frontend
npm run dev
```

- Visit `localhost:5173` in your browser

- To run the server:

```bash
go run .
```

- The endpoint is `localhost:8080` by default

## Encountering Issues?

Feel free to raise your hand or ask a neighbour for help. Worst-case scenario,
the code will be public after the workshop. You are also welcome to ask me any
questions afterwards.
