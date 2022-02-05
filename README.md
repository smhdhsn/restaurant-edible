# **Interview project**

## Food ordering system

Design a system that offers possible foods. User can place an order and the system will update the inventory after order's submittion.

## APIs

- As a user, we need an API that returns possible foods based on available raw materials.
- As a user, we need an API to register an order in the system.

## Required implementations

- Once the request is made, the `/menu` API will return a response in JSON format that includes possible foods based on available ingredients.
- User can order a food through `/order` API. After registering the order, the stock of raw materials for that food should be reduced.
- Inventory should have `expires_at` attribute which holds ingredients' expiration dates, foods with expired ingredients should not be recommended in `/menu` API.
- Inventory should have `best_before` attribute which holds ingredients' best usage dates, foods that their ingredients' `best-before` date have passed this time but not expired, should be at the end of the `/menu` API's response.
- Each ingredient has a stock that will be decreased when an order is submitted.
- If the stock of an ingredient is finished, foods with that ingredient should not be listed in `/menu` API's response.
- Every 15 minutes, ingredients that has been finished or expired should be deleted and new stocks from those finished or expired ingredients should be stored in the inventory.

## Optional implementations

- Implement `unit` and `integration` tests.
- Implement `README.md` file to guid the client in setup, use and testing the service.
- Implement `docker` and `docker-compose`.
