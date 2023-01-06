
CREATE TABLE IF NOT EXISTS expenses (
    id SERIAL PRIMARY KEY,
    title TEXT,
    amount FLOAT,
    note TEXT,
    tags TEXT[]
);


INSERT INTO "expenses" ("id", "title", "amount", "note", "tags") VALUES (1, 'Golang', 200, 'simple', ARRAY['banana']);