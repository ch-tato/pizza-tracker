# Joe's Pizza - Pizza Order Tracker

A fully-featured web application built with Go (Golang) and the Gin web framework. It allows customers to place custom pizza orders and track their delivery status in real-time. It also provides an admin dashboard for efficient order management.

---

## Key Features

### Customer Experience
* **Dynamic Order Placement**: An intuitive, modern UI built with Tailwind CSS allowing customers to fully customize their orders (Size, Flavor, Crust, Add-Ons, and Special Instructions).
* **Real-Time Order Tracking**: A live status stepper displaying exactly where the pizza is in the process (Received → Preparing → Baking → Out for Delivery → Delivered).
* **Instant Updates**: Powered by Server-Sent Events (SSE) to update the tracking page instantly without requiring the user to refresh the browser.

### Admin Dashboard
* **Secure Authentication**: Session-based secure login using `gin-contrib/sessions`.
* **Live Order Feed**: Admins see a live stream of incoming orders arriving in real-time via SSE.
* **Order Management**: Quick action buttons to update the status of any active order or cancel/delete an order directly from the dashboard.

---

## Tech Stack & Architecture

* **Backend**: Go (Golang) 1.25+
* **Web Framework**: [Gin Framework](https://github.com/gin-gonic/gin)
* **Database**: SQLite3 managed by [GORM](https://gorm.io/) (`gorm.io/gorm`)
* **Real-time Communication**: Server-Sent Events (SSE) using pure Go channels and Gin context streams.
* **Frontend**: HTML5 via Go Templates (`html/template`) and [Tailwind CSS](https://tailwindcss.com/) for a utility-first styling approach.
* **Live Reloading**: Supported natively using [Air](https://github.com/air-verse/air).
* **State Management:** GORM-backed persistent sessions (`github.com/gin-contrib/sessions`)
* **Security:** Bcrypt password hashing for admin authentication.

---

## Project Structure

```text
.
├── cmd/                # Entrypoint of the application (main.go, routes.go, handlers.go)
├── data/               # Local database storage (SQLite DB files)
├── internal/           # Private application code
│   └── models/         # Database models and GORM setup
├── templates/          # HTML Templates (Admin, Customer Tracking, Order Form)
│   └── static/         # Static assets (images, external CSS/JS)
├── Makefile            # Build and run commands
└── README.md           # Project documentation
```

---

## Getting Started

### Prerequisites
* Go 1.25 or higher installed
* `make` utility installed
* Air (for hot-reloading in development)

### Running the Application

1. **Clone the repository** (or navigate to the project directory).
2. **Install modules**:
   ```bash
   go mod download
   ```
3. **Run in Development Mode** (using `make` and `air`):
   ```bash
   make run-dev
   ```
4. **Access the application**:
   * **Customer Order Form**: [http://localhost:8080/](http://localhost:8080/)
   * **Admin Dashboard**: [http://localhost:8080/admin](http://localhost:8080/admin) (Requires authentication defined in environment configs)

---

## API Route Map

### Public Customer Routes

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/` | Serves the blank HTML order form. |
| `POST` | `/new-order` | Parses, validates, and saves the order. Redirects to tracker. |
| `GET` | `/customer/:id` | Serves the live tracking UI for a specific order. |
| `GET` | `/notifications` | The invisible SSE endpoint the customer's browser listens to. |

### Authentication Routes

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/login` | Serves the admin login form. |
| `POST` | `/login` | Validates credentials and generates a secure session cookie. |
| `POST` | `/logout` | Destroys the active session and clears the cookie. |

### Protected Admin Routes (Requires Middleware)

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/admin` | Serves the main dashboard listing all orders. |
| `POST` | `/admin/order/:id` | Updates an order's status and triggers the SSE broadcast. |
| `POST` | `/admin/order/:id/delete` | Executes a cascading delete for the order and its pizzas. |
| `GET` | `/admin/notifications` | A global SSE pipeline alerting admins of incoming new orders. |