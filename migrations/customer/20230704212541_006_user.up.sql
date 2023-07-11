CREATE TABLE IF NOT EXISTS users (
                      uuid smallserial PRIMARY KEY NOT NULL,
                      name varchar(255),
                      office_uuid integer,
                      created_at timestamp DEFAULT Now(),
                      FOREIGN KEY (office_uuid) REFERENCES office (uuid) ON DELETE CASCADE
);