CREATE TABLE office (
                        uuid smallserial PRIMARY KEY NOT NULL,
                        name varchar(255),
                        address varchar,
                        created_at timestamp DEFAULT Now()
);