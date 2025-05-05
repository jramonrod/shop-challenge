-- unnecessary 
WITH category AS (
    SELECT categories.uuid, categories.product_id FROM categories 
    JOIN products ON categories.product_id = products.uuid 
    WHERE categories.name = 'Frame type' 
    AND products.slug = 'bikes'
)
INSERT INTO parts (name, category_id, product_id, inventory, price_cents)
VALUES
    ('Full-suspension', (SELECT uuid FROM category), (SELECT product_id FROM category), 2, 2230),
    ('Diamond',         (SELECT uuid FROM category), (SELECT product_id FROM category), 0, 3902),
    ('Step-through',    (SELECT uuid FROM category), (SELECT product_id FROM category), 3, 22032);

WITH category AS (
    SELECT categories.uuid, categories.product_id FROM categories 
    JOIN products ON categories.product_id = products.uuid 
    WHERE categories.name = 'Frame finish' 
    AND products.slug = 'bikes'
)
INSERT INTO parts (name, category_id, product_id, inventory, price_cents)
VALUES
    ('Matte',   (SELECT uuid FROM category), (SELECT product_id FROM category), 200, 32302),
    ('Shiny',   (SELECT uuid FROM category), (SELECT product_id FROM category), 2000, 33221);

WITH category AS (
    SELECT categories.uuid, categories.product_id FROM categories 
    JOIN products ON categories.product_id = products.uuid 
    WHERE categories.name = 'Wheels'
    AND products.slug = 'bikes'
)
INSERT INTO parts (name, category_id, product_id, inventory, price_cents)
VALUES
    ('Road wheels',       (SELECT uuid FROM category), (SELECT product_id FROM category), 3, 2332),
    ('Mountain wheels',   (SELECT uuid FROM category), (SELECT product_id FROM category), 0, 23112),
    ('Fat bike wheels',   (SELECT uuid FROM category), (SELECT product_id FROM category), 300, 233232);

WITH category AS (
    SELECT categories.uuid, categories.product_id FROM categories 
    JOIN products ON categories.product_id = products.uuid 
    WHERE categories.name = 'Rim Color'
    AND products.slug = 'bikes'
)
INSERT INTO parts (name, category_id, product_id, inventory, price_cents)
VALUES
    ('Red',     (SELECT uuid FROM category), (SELECT product_id FROM category), 2, 23232),
    ('Black',   (SELECT uuid FROM category), (SELECT product_id FROM category), 1, 323211),
    ('Blue',    (SELECT uuid FROM category), (SELECT product_id FROM category), 1, 78654);

WITH category AS (
    SELECT categories.uuid, categories.product_id FROM categories 
    JOIN products ON categories.product_id = products.uuid 
    WHERE categories.name = 'Chain'
    AND products.slug = 'bikes'
)
INSERT INTO parts (name, category_id, product_id, inventory, price_cents)
VALUES
    ('Single-speed chain', (SELECT uuid FROM category), (SELECT product_id FROM category), 1, 2356),
    ('8-speed chain',      (SELECT uuid FROM category), (SELECT product_id FROM category), 50, 32189);
