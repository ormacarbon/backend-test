# 🏭 Carbon Offset API

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/joaooliveira247/backend-test/dev/backend)

## 💻 Requirements:

### `Go >= 1.24.1`

### [`Docker`](https://www.docker.com/) & [`Docker compose`](https://docs.docker.com/compose/)

## 📜 Documentation:

`/users/`

<details>
<summary><code>POST /users/</code></summary>

- **Description**: Creates a new user.

- **Headers**:

    ```plaintext
    Content-Type: application/json

    ```

- **Request Body**:

    ```json
    {
        "name": "john",
        "email": "john@gmail.com",
        "phone": "+5519912345678"
    }
    ```

- **Success Response (201 Created)**:

    ```json
    {
         "affiliateCode": "1d47bbe5-c7d3-4580-ad2a-c4b192eeeb47"
    }
    ```

- **Errors**:

    - **400 Bad Request**: Invalid request body.

    - **500 Internal Server Error**: Failed to create the entity.

- **Example Request with cURL**:

```curl
curl -X POST localhost:8000/users/ \
-H "Content-Type: application/json" \
-d '{
    "name": "john",
    "email": "john@gmail.com",
    "phone": "+5519912345678"
}'
```
</details>

###

`/competitions/`

<details>
<summary><code>POST /competitions/</code></summary>

- **Description**: Creates a new competition.

- **Headers**:

    ```plaintext
    Content-Type: application/json

    ```

- **Request Body**:

    ```json
    ```

- **Success Response (201 Created)**:

    ```json
    {
         "id": "06ae5f86-46dd-42d3-8e6d-2abe26f6b07e"
    }
    ```

- **Errors**:

    - **409 Conflict**: competition already activated.

    - **500 Internal Server Error**: error create competition.

- **Example Request with cURL**:

```curl
curl -X POST localhost:8000/competitions/ \
-H "Content-Type: application/json" \
```
</details>

<details>
<summary><code>GET /competitions/</code></summary>

- **Description**: Get competition activated.

- **Headers**:

    ```plaintext
    Content-Type: application/json

    ```

- **Success Response (200 OK / 204 No Content)**:

    ```json
    {
        "id": "9fbae8ae-3ba9-4582-a931-04d2e3c6aa93",
        "createdAt": 1743096181,
        "status": true
    }
    ```

- **Example Request with cURL**:

```curl
curl -X GET localhost:8000/competitions/ \
-H "Content-Type: application/json" \
```
</details>

<details>
<summary><code>PUT /competitions/</code></summary>

- **Description**: Close Competition.

- **Headers**:

    ```plaintext
    Content-Type: application/json

    ```

- **Query Parameters**:

    **ID** (required, UUID): Competiton ID.

- **Request Body**:

    ```json
    ```

- **Success Response (204 No Content)**:

    ```json
    ```

- **Errors**:

    - **400 Bad Request**: invalid id.

    - **404 Not Found**: competition not found.

    - **409 Conflict**: competition already activated.

    - **500 Internal Server Error**: closed competition error.

- **Example Request with cURL**:

```curl
curl -X PUT localhost:8000/competitions/?ID=e4b5c0cc-2f47-4b29-9883-84f314600f71\
-H "Content-Type: application/json" \
```
</details>

<details>
<summary><code>GET /competitions/reports/</code></summary>

- **Description**: Get competition reports.

- **Headers**:

    ```plaintext
    Content-Type: application/json
    ```

- **Query Parameters**:

    **ID** (required, UUID): Competiton ID.

- **Success Response (200 OK)**:

    ```json
    [
        {
            "name": "User 6",
            "points": 10
        },
        {
            "name": "User 5",
            "points": 10
        },
        {
            "name": "User 13",
            "points": 9
        },
        {
            "name": "User 8",
            "points": 8
        },
        {
            "name": "User 9",
            "points": 8
        },
        {
            "name": "User 7",
            "points": 6
        },
        {
            "name": "User 2",
            "points": 4
        },
        {
            "name": "User 11",
            "points": 4
        },
        {
            "name": "User 15",
            "points": 4
        },
        {
            "name": "User 1",
            "points": 2
        }
    ]
    ```

- **Example Request with cURL**:

```curl
curl -X GET localhost:8000/competitions/reports/?ID=d3dd9c62-5cc0-4b57-b341-ea1876dadac6 \
-H "Content-Type: application/json" \
```
</details>

## 📦 Usage libraries:

- [gin](github.com/gin-gonic/gin)

- [gorm](https://gorm.io/)

- [gorm/postgres](https://github.com/go-gorm/postgres)

- [uuid](github.com/google/uuid)

- [gomail](https://pkg.go.dev/gopkg.in/gomail.v2?utm_source=godoc)