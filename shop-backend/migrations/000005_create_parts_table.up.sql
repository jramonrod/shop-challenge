CREATE TABLE parts (
  uuid        UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  name        TEXT        NOT NULL,
  inventory   INTEGER     NOT NULL,
  price_cents INTEGER     NOT NULL,
  category_id UUID        NOT NULL REFERENCES categories(uuid) ON DELETE CASCADE,
  product_id  UUID        NOT NULL REFERENCES products(uuid) ON DELETE CASCADE,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT uq_parts_name_per_category_per_product UNIQUE (name, category_id, product_id),
  CONSTRAINT uq_parts_name_per_category UNIQUE (name, category_id)
);
