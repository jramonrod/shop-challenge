# Once the customer makes their selection, there should be an "add to cart" button. What happens when the customer clicks this button? 
Final Validation

Re-run the business rules on the backend to ensure the configuration is still valid (we need to guard against inventory changes, race conditions, invalid restriction selections).

If any violations remain, return a 422 Unprocessable Entity with details; the UI surfaces errors and aborts


# What is persisted in the database?
Cart Creation / Retrieval

If the user has no open cart, we create one. If they already have an open cart, load its ID. The make a POST/PUT request to update the cart. 
Between these 2 choices:
1. If we opt for a POST request instead of a PUT request we should send the whole cart and replace the cart in the database. 
2. If we opt instead for a PUT request we can update specific parts of the cart depending on the business rules, such as updating the quantity to buy

We can then either snapshot the current price_cents or alert a user that their price has been updated if the price returned by the DB is different from his cart during checkout.
1. With price snapshoting we ensure future price changes don’t retroactively affect the customer’s cart. This might be desirable in some shops but likely we prefer to keep the prices as up to date as possible
