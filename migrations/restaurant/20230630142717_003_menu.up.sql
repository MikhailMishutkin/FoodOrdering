CREATE TABLE menu (
                        uuid smallserial PRIMARY KEY NOT NULL,
                        on_date timestamp,
                        opening_record_at timestamp,
                        closing_record_at timestamp,
                        product_id integer,
                        prod_type_m integer,
                        created_at timestamp DEFAULT Now(),
                        FOREIGN KEY (product_id) REFERENCES product (uuid) ON UPDATE CASCADE,
                        FOREIGN KEY (prod_type_m) REFERENCES product_type (id) ON UPDATE CASCADE
);