CREATE TABLE order_item (
                            id smallserial PRIMARY KEY NOT NULL,
                            count integer,
                            product_uuid integer,
                            FOREIGN KEY (product_uuid) REFERENCES product (uuid)
);