The Food Delivery API project is a comprehensive backend solution designed to manage and facilitate a food delivery service. It provides robust functionality for handling restaurant operations, menu management, user transactions, and more. Built using the Go programming language and the Gin web framework, this project ensures high performance and scalability. The database interactions are managed using GORM, providing a smooth ORM experience.

    Key Features
        1. Restaurant Management:
            - Load restaurant data from JSON files into the database.
            - Retrieve open restaurants based on a specific date and time.
            - Retrieve top restaurants based on the number of dishes and price range.
            - Search for restaurants by name.

        2. Menu Management:
            - Load menu items into the database.
            - Search for dishes by name.

        3. User Management:
            - Load user data and purchase history from JSON files into the database.
            - Handle user purchase transactions, ensuring balance checks and updates.

        4. Transaction Management:
            - Record and manage user transactions.
            - Ensure data consistency and atomicity during purchase operations.

    Key Modules
        Controllers
            RestaurantController: Handles HTTP requests related to restaurant operations.
                - GetOpenRestaurants: Retrieves a list of open restaurants based on a provided datetime.
                - GetTopRestaurants: Retrieves a list of top restaurants based on various criteria such as the number of dishes and price range.
                - SearchRestaurants: Searches for restaurants by name.
                - SearchDishes: Searches for dishes by name.

            UserController: Handles HTTP requests related to user operations.
                - Purchase: Processes a user's purchase transaction, updating their balance and recording the transaction.

        Services
            RestaurantService: Provides business logic for restaurant-related operations.
                - LoadData: Loads restaurant data and menu items from a JSON file into the database.
                - GetOpenRestaurants: Retrieves open restaurants based on a specific datetime.
                - GetTopRestaurants: Retrieves top restaurants based on the number of dishes and price range.
                - SearchRestaurants: Searches for restaurants by name.
                - SearchDishes: Searches for dishes by name.

            UserService: Provides business logic for user-related operations.
                - LoadData: Loads user data and purchase history from a JSON file into the database.
                - Purchase: Processes a user's purchase transaction, ensuring balance checks and recording the transaction.

        Models
            - Restaurant: Represents a restaurant entity with fields such as name, cash balance, opening hours, and associated menu items.
            - MenuItem: Represents a menu item with fields such as name, price, and the associated restaurant ID.
            - User: Represents a user entity with fields such as name, cash balance, and purchase history.
            - Transaction: Represents a transaction entity with fields such as user ID, menu item ID, amount, and purchase date.

        Main
            main.go: The entry point of the application.
                - Configures the database connection using GORM.
                - Initializes the services and controllers.
                - Defines the routes for the API endpoints.
                - Starts the HTTP server.
        
    Technologies Used:
        - Backend Framework: Go (Golang)
        - Web Framework: Gin
        - Database: MySQL
        - ORM: GORM


API DOCUMENTATION

    Restaurant Endpoints
    1. Get Open Restaurants
        - URL: http://localhost:8080/restaurants/open
        - Method: GET
        - Parameters:
                - datetime: day and time in the format "Monday 10:30 am".
        - Example: http://localhost:8080/restaurants/open?datetime=Monday 10:16 am
    
    2. Get Top Restaurants
        - URL: http://localhost:8080/restaurants/top
        - Method: GET
        - Parameters:
            - number_of_dishes (optional): Number of dishes.
            - price_range (optional): Price range of dishes.
            - more_than (optional): Whether to return restaurants with more than the specified number of dishes.
        - Example: http://localhost:8080/restaurants/top?number_of_dishes=5&price_range=10-20&more_than=false 

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

