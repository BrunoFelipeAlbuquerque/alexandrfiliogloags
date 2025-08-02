# Imperial Fleet API (alexandrfiliogloags)

This is a simple RESTful API for managing a fleet of spaceships, built with Go and Gin, and backed by PostgreSQL using GORM.

---

## 🛠 Requirements

- **Go**: version **1.21+** (minimum required)
- **PostgreSQL**: running and accessible
- **Insomnia** (optional): for running test requests (file: `insomnia.yaml`)

---

## 📦 Installation & Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/imperial-fleet.git
   cd imperial-fleet
   ```

2. **Configure environment variables:**

   Create a `.env` file in the root directory:

   ```dotenv
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=yourpassword
   DB_NAME=imperial_fleet
   ```

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

4. **Run database migrations (if applicable):**

   You may need to apply schema migrations manually or via a tool like `golang-migrate` or GORM’s `AutoMigrate` (already included in the code).

5. **Run the project:**

   ```bash
   go run main.go
   ```

   The server will start at: `http://localhost:8080`

---

## 📥 Import Requests with Insomnia

1. Open **Insomnia**.
2. Go to the top menu → `File` → `Import`.
3. Choose `From File`.
4. Select the provided `insomnia.yaml` file in the project root.
5. Run and test the available routes.

---

## 📚 Available Endpoints

| Method | Route               | Description                  |
|--------|---------------------|------------------------------|
| GET    | `/spaceships`       | List spaceships with filters |
| GET    | `/spaceships/:id`   | Get spaceship by ID          |
| POST   | `/spaceships`       | Create a new spaceship       |
| PUT    | `/spaceships/:id`   | Update a spaceship           |
| DELETE | `/spaceships/:id`   | Delete a spaceship           |

All endpoints return JSON responses.

---

## ✅ Example Query (GET `/spaceships`)

You can pass query parameters such as:

- `name`
- `class`
- `status`
- `crew`
- `value`, `value_min`, `value_max`
- `created_at`, `created_at_min`, `created_at_max`
- `updated_at`, `updated_at_min`, `updated_at_max`
- `page`
- `limit`

Example:

```
/spaceships?name=Imperial&page=1&limit=10
```
