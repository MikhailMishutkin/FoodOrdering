CREATE TABLE orders (
                         id serial primary key not null,
                         on_date date DEFAULT CURRENT_DATE + INTERVAL '1 day',
                         user_uuid integer,
                         product_id integer,
                         count integer,
                         FOREIGN KEY (product_id) REFERENCES product (uuid),
                         FOREIGN KEY (user_uuid) REFERENCES users (id)
);