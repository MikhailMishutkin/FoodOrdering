CREATE TABLE users (
                        id integer PRIMARY KEY NOT NULL,
                        user_name varchar(255),
                        office_uuid integer,
                        FOREIGN KEY (office_uuid) REFERENCES office (id)
);