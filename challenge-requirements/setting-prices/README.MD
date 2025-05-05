# How can Marcus change the price of a specific part or specify particular pricing for combinations of choices? # How does the UI and database handle this?
Marcus can utilise the admin panel of the shop to modify specific parts including their prices. 

In order for him to provide sales or specific combo prices we can introduce "price_rules" and these price rules can, similarly to our restrictions, define a group of parts and an effect. 

My initial idea can be something like:

```sql 
CREATE TABLE price_rules (
  uuid            UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  principal_part  UUID        NOT NULL REFERENCES parts(uuid)
  resource_part   UUID        NOT NULL REFERENCES parts(uuid),
  override_effect INTEGER     NOT NULL,
  type            TEXT        NOT NULL CHECK(type IN ('absolute','delta', 'sale'))
);
```

Then when calculating the price in the backend we can apply either an absolute price, a reduction/delta in price or a price sale in the service-level code

This can be further expanded by introducting a `parts_price_rules` table if we need a combo of resources or a combo of principals to be met for the override_effect to apply.
