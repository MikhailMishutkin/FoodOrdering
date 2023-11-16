CREATE TABLE menu_date (
                           uuid serial PRIMARY KEY NOT NULL,
                           date timestamptz,
                           opening_record_at timestamptz,
                           closing_record_at timestamptz,
                           created_at timestamptz DEFAULT Now()
);