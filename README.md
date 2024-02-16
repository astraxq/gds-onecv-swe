# gds-onecv-swe

# Documentation

The API is publicly availale @https://dev.brianquek.live/api

# ERD Diagram

To be added.

# Setting Up Your Environment

### 1. Install Go 1.22

```sh
brew install go@1.22
```

### 2. Migrate Go 4.17.0

```sh
brew install migrate-go@4.17.0
```

### 3. Navigate to backend folder

```sh
cd backend
```

### 4. Run docker compose

```sh
docker-compose up -d
```

### 5. Run migration file

```sh
migrate -path database/migration -database "postgresql://root:secret@localhost:5432/class_db?sslmode=disable" -verbose up
```

### 5. Run Go Application

```sh
go run .
```

### 6. Feel free to test the endpoints via the Postman Collection or curl

```sh
# User Story 1: Register Students API
# public domain
curl --location 'https://dev.brianquek.live/api/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "teacher": "brianquek@example.com",
    "students": [
        "jane.smith@example.com",
        "alice.johnson@example.com"
    ]
}'
# localhost
curl --location 'http://localhost:8000/api/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "teacher": "brianquek@example.com",
    "students": [
        "jane.smith@example.com",
        "alice.johnson@example.com"
    ]
}'

# User Story 2: Common Students
# public domain
curl --location 'https://dev.brianquek.live/api/commonstudents?teacher=brianquek%40example.com'
# localhost
curl --location 'http://localhost:8000/api/commonstudents?teacher=brianquek%40example.com'

# User Story 3: Suspend Student
# public domain
curl --location 'https://dev.brianquek.live/api/suspend' \
--header 'Content-Type: application/json' \
--data-raw '{
    "student": "jane.smith@example.com"
}'
# localhost
curl --location 'http://localhost:8000/api/suspend' \
--header 'Content-Type: application/json' \
--data-raw '{
    "student": "jane.smith@example.com"
}'

# User Story 4: Retrieve Notification
# public domain
curl --location 'https://dev.brianquek.live/api/retrievefornotifications' \
--header 'Content-Type: application/json' \
--data-raw '{
    "teacher": "teacherken@example.com",
    "notification": "Hello students! studentagnes@gmail.com studentmiche@gmail.comwdadw diefneigna@sada.com safwaf safsaf@"
}'
# localhost
curl --location 'http://localhost:8000/api/retrievefornotifications' \
--header 'Content-Type: application/json' \
--data-raw '{
    "teacher": "teacherken@example.com",
    "notification": "Hello students! studentagnes@gmail.com studentmiche@gmail.comwdadw diefneigna@sada.com safwaf safsaf@"
}'
```

### 7. Seed Data (QOL Endpoint)

```
# clears the table rows and populate the user table
curl --location --request POST 'http://localhost:8000/api/seed'


 id |     name      |           email           | role | status | notification_allowed
----+---------------+---------------------------+------+--------+----------------------
 26 | Ken Doe       | teacherken@example.com    |    2 |      1 | t
 27 | Brian Quek    | brianquek@example.com     |    2 |      1 | t
 28 | John Tan      | johntan@example.com       |    3 |      1 | t
 29 | Jane Smith    | jane.smith@example.com    |    3 |      1 | t
 30 | Alice Johnson | alice.johnson@example.com |    3 |      1 | t
 31 | James Lee     | james.lee@example.com     |    3 |      1 | f
```
