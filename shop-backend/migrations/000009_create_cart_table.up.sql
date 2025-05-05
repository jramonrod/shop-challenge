CREATE TABLE carts (
  uuid        UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  status      TEXT        NOT NULL CHECK(status IN ('active','completed')) DEFAULT 'active',
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
