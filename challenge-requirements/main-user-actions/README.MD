# Explain the main actions users would take on this e-commerce website in detail.
* Browse and Discover Products

Users land on the home page and browse available products, each representing a customizable item (e.g., a bicycle).

* View a Product Detail Page

Clicking on a product navigates users to the product's landing page, this is where they can meet a catalogue of different options offered by the e-commerce site

*  Customize their Product

A user can now chose to build their own product. Users step through each category, selecting exactly one part per category.

* Validate Configuration

As the building process happens users verify that their custom build has no compatibility issues. Calls in the background upon part selection happen to the backend to verify the business rules set in the admin panel. Also a call is made when submission is done to double check the business rules and prices once again as asynchronously a change could have happened. Minimizing our chances of a race condition. By the end of the building process a user can create a cart.

* Create a Cart

If the configuration is valid, the user can add the build to their cart. This gets linked to their user/session in the backend (though not in this implementation for reasons state previously)
Users in their cart can see prices, products selected, totals, and make final adjustments.

* Proceed to Checkout
* 
From the cart, users initiate the checkout process. While this project may not implement a real checkout, in production, this would trigger workflows involving inventory validation, payment tokenization, order creation, and handoff to fulfillment services.
