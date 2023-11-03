CREATE TABLE IF NOT EXISTS products(
   id           serial PRIMARY KEY,
   name         VARCHAR (50) UNIQUE NOT NULL,
   description  text,
   price        real,
   weight       int
);