CREATE TYPE statuses AS ENUM (
  'active',
  'banned',
  'deleted'
);

CREATE TYPE genders AS ENUM (
  'man',
  'woman'
);

CREATE TABLE users (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "surname" varchar,
  "gender" genders,
  "status" statuses,
  "birth_date" date,
  "creat_data" date DEFAULT now()
);