-- unnecessary 
WITH bike_product AS (
    SELECT uuid FROM products WHERE slug = 'bikes'
)
INSERT INTO categories (name, product_id)
VALUES
    ('Frame type',      (SELECT uuid FROM bike_product)),
    ('Frame finish',    (SELECT uuid FROM bike_product)),
    ('Wheels',          (SELECT uuid FROM bike_product)),
    ('Rim Color',       (SELECT uuid FROM bike_product)),
    ('Chain',           (SELECT uuid FROM bike_product));
