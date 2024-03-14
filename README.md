# gds-onecv-swe

# Documentation

The API is publicly available ~@https://dev.brianquek.live/api~

# ERD Diagram

ERD Diagram Link @https://drawsql.app/teams/astraxq/diagrams/gds-onecv-swe

# Setting Up Your Environment

### 1. Install Go 1.22 (https://go.dev/doc/install)

```sh
brew install go@1.22
```

### 2. Install Migrate Go 4.17.0 (https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

```sh
brew install golang-migrate@4.17.0
```

### 3. Ensure docker is installed (https://docs.docker.com/engine/install/)

### 4. Clone the repository and navigate to backend directory

```sh
git clone https://github.com/astraxq/gds-onecv-swe.git
cd gds-onecv-swe/backend
```

### 5. Run docker compose

```sh
docker-compose up -d
```

### 6. Run migration file

```sh
migrate -path database/migration -database "postgresql://root:secret@localhost:5432/class_db?sslmode=disable" -verbose up
```

### 7. Run Go Application

```sh
go run .
```

### 8. Feel free to test the endpoints via the Postman Collection or curl

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

### Others. Seed Data (QOL Endpoint)

```
# clears the table rows and populate the user table
curl --location --request POST 'http://localhost:8000/api/seed'


 id |     name      |           email           | role | status
----+---------------+---------------------------+------+--------+----------------------
 26 | Ken Doe       | teacherken@example.com    |    2 |      1
 27 | Brian Quek    | brianquek@example.com     |    2 |      1
 28 | John Tan      | johntan@example.com       |    3 |      1
 29 | Jane Smith    | jane.smith@example.com    |    3 |      1
 30 | Alice Johnson | alice.johnson@example.com |    3 |      1
 31 | James Lee     | james.lee@example.com     |    3 |      1
```
