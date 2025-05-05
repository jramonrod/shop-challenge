-- unnecessary 
WITH parts AS (
    SELECT uuid, name FROM parts WHERE name = 'Mountain wheels' OR name = 'Full-suspension'
)
INSERT INTO restrictions (principal, resource, effect)
VALUES
    ((SELECT uuid FROM parts WHERE name = 'Mountain wheels'), (SELECT uuid FROM parts WHERE name = 'Full-suspension'), 'only');

WITH parts AS (
    SELECT uuid, name FROM parts WHERE name = 'Fat bike wheels' OR name = 'Red'
)
INSERT INTO restrictions (principal, resource, effect)
VALUES
    ((SELECT uuid FROM parts WHERE name = 'Fat bike wheels'), (SELECT uuid FROM parts WHERE name = 'Red'), 'forbid');

