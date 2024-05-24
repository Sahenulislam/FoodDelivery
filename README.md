API Documentation

    Restaurant Endpoints
    1. Get Top Restaurants
        - URL: http://localhost:8080/restaurants/top
        - Method: GET
        - Parameters:
            - number_of_dishes (optional): Number of dishes.
            - price_range (optional): Price range of dishes.
            - more_than (optional): Whether to return restaurants with more than the specified number of dishes.
        - Example: http://localhost:8080/restaurants/top?number_of_dishes=5&price_range=10-20&more_than=false 

    2. Get Open Restaurants
        - URL: http://localhost:8080/restaurants/open
        - Method: GET
        - Parameters:
                - datetime: Date and time in the format "YYYY-MM-DD HH:MM:SS".
        - Example: http://localhost:8080/restaurants/open?datetime=2024-05-24%2015:00:00
        
    3. Search Restaurants by Name
        - URL: http://localhost:8080/restaurants/search/restaurant
        - Method: GET
        - Parameters:
                - query: Restaurant name query.
        - Example: http://localhost:8080/restaurants/search/restaurant?query=pizza
        
    4. Search Dishes by Name
        - URL: http://localhost:8080/restaurants/search/dishes
        - Method: GET
        - Parameters:
                - query: Dish name query.
        - Example: http://localhost:8080/restaurants/search/dishes?query=tometo


    User Endpoint
    1. User Purchase
       - URL: http://localhost:8080/users/purchase
       - Method: POST
       - Body Parameters:
                - userId: ID of the user making the purchase.
                - menuItemId: ID of the menu item being purchased.
       - Example Body:
            {
  		        "userId": 5,
   		        "menuItemId": 1
            }


    ETL Script
       - To run the ETL script:
       - Ensure you have Go installed on your system.
       - Navigate to the directory containing the ETL script.
       - Run the script using the command : go run etl_script.go




    Server and Database Setup
       - Install MySQL on your system if not already installed.
       - Create a new MySQL database named food-delivery.
       - Update the database connection settings in the main package (main.go) to match your MySQL configuration.
       - Run the following command to set up the server and database: go run main.go
