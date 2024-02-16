# gds-onecv-swe

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
curl --location 'https://dev.brianquek.live/commonstudents?teacher=brianquek%40example.com'
curl --location 'http://localhost:8000/commonstudents?teacher=brianquek%40example.com'

curl --location 'https://dev.brianquek.live/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "teacher": "brianquek@example.com",
    "students": [
        "jane.smith@example.com",
        "alice.johnson@example.com"
    ]

}'

curl --location 'http://localhost:8000/suspend' \
--header 'Content-Type: application/json' \
--data-raw '{
    "student": "jane.smith@example.com"
}'

    curl --location 'http://localhost:8000/retrievefornotifications' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "teacher": "teacherken@example.com",
        "notification": "Hello students! studentagnes@gmail.com studentmiche@gmail.comwdadw diefneigna@sada.com safwaf safsaf@"

    }'
```

### 7. Seed Data (I've created seed endpoint that drops the users & student-teacher table, where we subsequently repopulate the users table as given below)

```
curl --location --request POST 'http://localhost:8000/seed'


 id |     name      |           email           | role | status | notification_allowed
----+---------------+---------------------------+------+--------+----------------------
 26 | Ken Doe       | teacherken@example.com    |    2 |      1 | t
 27 | Brian Quek    | brianquek@example.com     |    2 |      1 | t
 28 | John Tan      | johntan@example.com       |    3 |      1 | t
 29 | Jane Smith    | jane.smith@example.com    |    3 |      1 | t
 30 | Alice Johnson | alice.johnson@example.com |    3 |      1 | t
 31 | James Lee     | james.lee@example.com     |    3 |      1 | f
```
