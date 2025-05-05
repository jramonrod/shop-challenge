CREATE TABLE categories (
  uuid        UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  name        TEXT        NOT NULL,
  product_id  UUID        NOT NULL REFERENCES products(uuid) ON DELETE CASCADE,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT uq_categories_name_per_product UNIQUE (name, product_id)
);
