CREATE TABLE menu (
                        uuid smallserial PRIMARY KEY NOT NULL,
                        menu_id integer,
                        product_id integer,
                        prod_type_m integer,
                        FOREIGN KEY (menu_id) REFERENCES menu_date (uuid) ON DELETE CASCADE,
                        FOREIGN KEY (product_id) REFERENCES product (uuid) ON UPDATE CASCADE,
                        FOREIGN KEY (prod_type_m) REFERENCES product_type (id) ON UPDATE CASCADE
);