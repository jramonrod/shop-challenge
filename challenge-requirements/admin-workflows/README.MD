# Describe the main workflows for Marcus to manage his store.
There are largely 5 different workflows for Marcus to manage his store.
## 1. Product Lifecycle Management 
Marcus can create, update, and remove different types of products

Marcus can create a new product with attributes: name, slug , status, and optional image. This implies an upload of an image and an associations to a main product. This has implications for the product list in the front-end initially and allows for extensibility in the next sections.

## 2. Category Management
For each product type, Marcus defines the major customization categories. This will extend the steps for building the related product and add in a step in the builder per category.

## 3. Part/Customization Management
> [!WARNING]
> I feel like it's important to mention here that images and better display for parts should also be implemented and marcus should be able to customize the images he wants to use. But this was not a requirement so I leave this comment here only :).

Marcus manages individual customization options under each category.Marcus can create new parts for a category with attributes like name, inventory count, and price (in cents).

He can also update inventory levels as stock changes.

This has implications for the front-end in the sense that the user will be given new options for each category when building their product.

## 4. Constraint (Restrictions) Management
Marcus can define rules that prevent invalid combinations or enforce specific combinations.

For example:
> If selected Mountain Wheels ➔ must choose Full Suspension Frame”
> 
> This translates to a state in the database table restrictions with a row that references "Mountain Wheels" as a principal and "Full Suspension Frame" as a resource and the effect is set to "only".

This allows us to communicate with the user and our backend the rules that Marcus has about the products.

## 5. Inventory/Payments Monitoring
Keep track of parts that are running low and restock. Marcus should also be allowed to increment the amount of parts available for a currently existing part.
