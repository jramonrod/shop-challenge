CREATE TABLE products (
  uuid        UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  name        TEXT        NOT NULL,
  status      TEXT        NOT NULL CHECK(status IN ('unavailable','available')),
  slug        TEXT        NOT NULL,
  location    TEXT        DEFAULT NULL,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT uq_products_slug UNIQUE (slug)
);
