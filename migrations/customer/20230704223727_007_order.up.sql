CREATE TABLE orders (
                       user_uuid int,
                       item_id int,
                       FOREIGN KEY (user_uuid) REFERENCES users (uuid),
                       FOREIGN KEY (item_id) REFERENCES order_item (id)
);