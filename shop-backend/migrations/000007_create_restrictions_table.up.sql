CREATE TABLE restrictions (
  uuid            UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  principal       UUID NOT NULL REFERENCES parts(uuid),
  resource        UUID NOT NULL REFERENCES parts(uuid),
  effect          TEXT NOT NULL CHECK(effect IN ('allow','forbid', 'only'))
);
