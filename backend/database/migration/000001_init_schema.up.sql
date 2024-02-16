CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "role" INT NOT NULL,
    "status" INT NOT NULL,
    CONSTRAINT "email_unique" UNIQUE ("email")
);

CREATE INDEX "idx_email" ON "users"("email");

CREATE TABLE IF NOT EXISTS "user_tags" (
    "id" SERIAL,
    "teacher_id" INT NOT NULL,
    "student_id" INT NOT NULL,
    CONSTRAINT "pk_user_tags" PRIMARY KEY ("teacher_id", "student_id")
); 

CREATE INDEX "idx_teacher_id" ON "user_tags"("teacher_id");