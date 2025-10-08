# fetch-api

A small Go HTTP server that fetches user data from an external API and exposes it to clients as JSON. This project is a simple proxy/middleware example — useful for learning how to call upstream APIs from a Go server, add caching or authentication, and shape responses for frontends or mobile apps.

---

## Features

* Fetches users from `https://jsonplaceholder.typicode.com/users` (default)
* Exposes an endpoint `GET /users` that returns the fetched data as JSON
* Simple root health route `GET /` that returns a welcome message
* Built with Go standard library only (no external dependencies)
* Includes safe HTTP client usage with timeout and error logging

---

## Repository structure

```
fetch-api/
├── go.mod
└── main.go
```

---

## Requirements

* Go 1.20+ installed (the project uses modules)
* Internet access (so the server can reach the upstream API)

---

## Quick start

1. Clone the repository:

```bash
git clone https://github.com/EternalKnight002/fetch-api.git
cd fetch-api
```

2. Download dependencies (not strictly needed for this small project, but good practice):

```bash
go mod tidy
```

3. Run the server:

```bash
go run main.go
```

The server prints the address it is listening on; by default the example binds to port `8083`.

---

## Endpoints

### `GET /`

* Returns a simple welcome string (useful as a health check).

Example:

```bash
curl http://localhost:8083/
# -> Welcome to Fetch API Server 🌍
```

---

### `GET /users`

* Fetches users from the upstream API and returns them as JSON.
* By default the upstream URL is `https://jsonplaceholder.typicode.com/users`.

Example:

```bash
curl http://localhost:8083/users
```

A successful response returns a JSON array of user objects.

---

## Configuration & Customization

The example code hardcodes the upstream URL and the listen address. You can easily modify `main.go` to support environment variables or flags. Example suggestions:

* Add `UPSTREAM_URL` environment variable to point to the upstream API.
* Add `PORT` or `ADDR` environment variable to change the listening port.
* Add an API key header when calling private upstream services.

---

## Improvements you can add

If you want to extend this project, here are common enhancements:

* **Caching**: Add an in-memory cache (map with TTL) or Redis to avoid calling the upstream service on every request.
* **Structured types**: Replace `map[string]interface{}` with Go structs matching the upstream JSON for better type safety.
* **Retry/backoff**: Add retry logic with exponential backoff for transient network errors.
* **Rate limiting**: Protect your server from abuse and throttle calls to the upstream API.
* **Authentication**: Add JWT or API key protection to your endpoints.
* **Graceful shutdown**: Use `http.Server` with `Shutdown` to stop the server cleanly.
* **Tests**: Add unit tests for handlers and an integration test using a mocked upstream server.

---

## Troubleshooting

* If `/users` returns an error, check the server console — the program logs helpful error messages when upstream fetches fail.
* If you get timeouts or connection refused when contacting `jsonplaceholder.typicode.com`, verify your machine has outbound internet access and no firewall/proxy is blocking the request.
* If you started the server on a different port, make sure to use the correct port in requests.

---

## License
This project is licensed under the MIT License.


---

## Contact

If you want help extending this project (database integration, auth, deployment), open an issue or message me with what you want and I can provide example code and guidance.
