CREATE TABLE orders (
                         on_date date DEFAULT CURRENT_DATE + INTERVAL '1 day',
                         user_uuid integer,
                         product_id integer,
                         count integer
)