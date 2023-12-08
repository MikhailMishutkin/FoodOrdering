CREATE TABLE orders (
                        id serial primary key not null,
                        on_date date DEFAULT CURRENT_DATE + INTERVAL '1 day',
                        product_id integer,
                        count integer,
                        FOREIGN KEY (product_id) REFERENCES product (uuid)
);