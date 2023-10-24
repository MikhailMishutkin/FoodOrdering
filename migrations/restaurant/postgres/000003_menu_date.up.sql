CREATE TABLE menu_date (
                           uuid serial PRIMARY KEY NOT NULL,
                           date timestamptz,
                           created_at timestamptz DEFAULT Now()
);