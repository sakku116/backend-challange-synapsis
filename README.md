# Backend Challenge Synapsis

## CLI Commands
- `seed-superuser`
    ```
    go run . seed-superuser
    ```
    or
    ```
    go run . seed-superuser superuser1  superuserpass
    ```
    - args (optional):
        args should be empty for default seed (superuser;superuser). Or args must be containing 2 strings for custom username, and passwords.


- `seed-data`
    used to seed all required data for mandatory minimum requirements.
    ```
    go run . seed-data
    ```

## Swagger
go to `/swagger/index.html` route path to see API documentation via Swagger.

## API Mandatory Requirements
- Customer can view product list by product category
- Customer can add product to shopping cart
- Customers can see a list of products that have been added to the shopping cart
- Customer can delete product list in shopping cart
- Customers can checkout and make payment transactions
- Login and register customers
