CREATE TABLE menu (
                        uuid smallserial PRIMARY KEY NOT NULL,
                        on_date timestamptz,
                        opening_record_at timestamptz,
                        closing_record_at timestamptz,
                        product_id integer,
                        prod_type_m integer,
                        created_at timestamptz DEFAULT Now(),
                        FOREIGN KEY (product_id) REFERENCES product (uuid) ON UPDATE CASCADE,
                        FOREIGN KEY (prod_type_m) REFERENCES product_type (id) ON UPDATE CASCADE
);