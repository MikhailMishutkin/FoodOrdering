CREATE TABLE product (
                         uuid smallserial PRIMARY KEY NOT NULL,
                         name varchar(255),
                         description varchar,
                         type_id integer,
                         weight integer,
                         price numeric(6, 2),
                         created_at timestamp DEFAULT Now(),
                         FOREIGN KEY (type_id) REFERENCES product_type (id) ON UPDATE CASCADE
);