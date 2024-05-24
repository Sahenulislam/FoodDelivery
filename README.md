<<<<<<< HEAD
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

=======
# food-delivery



## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

## Add your files

- [ ] [Create](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#create-a-file) or [upload](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#upload-a-file) files
- [ ] [Add files using the command line](https://docs.gitlab.com/ee/gitlab-basics/add-file.html#add-a-file-using-the-command-line) or push an existing Git repository with the following command:

```
cd existing_repo
git remote add origin https://gitlab.com/personal4400321/food-delivery.git
git branch -M main
git push -uf origin main
```

## Integrate with your tools

- [ ] [Set up project integrations](https://gitlab.com/personal4400321/food-delivery/-/settings/integrations)

## Collaborate with your team

- [ ] [Invite team members and collaborators](https://docs.gitlab.com/ee/user/project/members/)
- [ ] [Create a new merge request](https://docs.gitlab.com/ee/user/project/merge_requests/creating_merge_requests.html)
- [ ] [Automatically close issues from merge requests](https://docs.gitlab.com/ee/user/project/issues/managing_issues.html#closing-issues-automatically)
- [ ] [Enable merge request approvals](https://docs.gitlab.com/ee/user/project/merge_requests/approvals/)
- [ ] [Set auto-merge](https://docs.gitlab.com/ee/user/project/merge_requests/merge_when_pipeline_succeeds.html)

## Test and Deploy

Use the built-in continuous integration in GitLab.

- [ ] [Get started with GitLab CI/CD](https://docs.gitlab.com/ee/ci/quick_start/index.html)
- [ ] [Analyze your code for known vulnerabilities with Static Application Security Testing (SAST)](https://docs.gitlab.com/ee/user/application_security/sast/)
- [ ] [Deploy to Kubernetes, Amazon EC2, or Amazon ECS using Auto Deploy](https://docs.gitlab.com/ee/topics/autodevops/requirements.html)
- [ ] [Use pull-based deployments for improved Kubernetes management](https://docs.gitlab.com/ee/user/clusters/agent/)
- [ ] [Set up protected environments](https://docs.gitlab.com/ee/ci/environments/protected_environments.html)

***

# Editing this README

When you're ready to make this README your own, just edit this file and use the handy template below (or feel free to structure it however you want - this is just a starting point!). Thanks to [makeareadme.com](https://www.makeareadme.com/) for this template.

## Suggestions for a good README

Every project is different, so consider which of these sections apply to yours. The sections used in the template are suggestions for most open source projects. Also keep in mind that while a README can be too long and detailed, too long is better than too short. If you think your README is too long, consider utilizing another form of documentation rather than cutting out information.

## Name
Choose a self-explaining name for your project.

## Description
Let people know what your project can do specifically. Provide context and add a link to any reference visitors might be unfamiliar with. A list of Features or a Background subsection can also be added here. If there are alternatives to your project, this is a good place to list differentiating factors.

## Badges
On some READMEs, you may see small images that convey metadata, such as whether or not all the tests are passing for the project. You can use Shields to add some to your README. Many services also have instructions for adding a badge.

## Visuals
Depending on what you are making, it can be a good idea to include screenshots or even a video (you'll frequently see GIFs rather than actual videos). Tools like ttygif can help, but check out Asciinema for a more sophisticated method.

## Installation
Within a particular ecosystem, there may be a common way of installing things, such as using Yarn, NuGet, or Homebrew. However, consider the possibility that whoever is reading your README is a novice and would like more guidance. Listing specific steps helps remove ambiguity and gets people to using your project as quickly as possible. If it only runs in a specific context like a particular programming language version or operating system or has dependencies that have to be installed manually, also add a Requirements subsection.

## Usage
Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

## Support
Tell people where they can go to for help. It can be any combination of an issue tracker, a chat room, an email address, etc.

## Roadmap
If you have ideas for releases in the future, it is a good idea to list them in the README.

## Contributing
State if you are open to contributions and what your requirements are for accepting them.

For people who want to make changes to your project, it's helpful to have some documentation on how to get started. Perhaps there is a script that they should run or some environment variables that they need to set. Make these steps explicit. These instructions could also be useful to your future self.

You can also document commands to lint the code or run tests. These steps help to ensure high code quality and reduce the likelihood that the changes inadvertently break something. Having instructions for running tests is especially helpful if it requires external setup, such as starting a Selenium server for testing in a browser.

## Authors and acknowledgment
Show your appreciation to those who have contributed to the project.

## License
For open source projects, say how it is licensed.

## Project status
If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers.
>>>>>>> 5d1b1fca6bfe9d191d1486baa017cec9897fffe7
