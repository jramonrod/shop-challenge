-- here we can implement other stuff like price snapshotting or quantity
CREATE TABLE cart_items (
  cart_id     UUID        NOT NULL REFERENCES carts(uuid) ON DELETE CASCADE,
  part_id     UUID        NOT NULL REFERENCES parts(uuid),
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  PRIMARY KEY (cart_id, part_id)
);
