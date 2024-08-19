# The-Lazy-Chef
A simple recipe book application.

## How to Run the Backend

### Prerequisites

- **Docker**: Ensure Docker is installed and running on your machine.
- **Docker Compose**: Ensure Docker Compose is installed.

### Instructions

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/The-Lazy-Chef.git
   cd The-Lazy-Chef

2. **Build and Run the Containers**:

    Navigate to the backend folder

    ```bash
    cd backend
    ```

    Use Docker Compose to build and run the backend along with MongoDB:

    ```bash
    docker-compose up --build
    ```

3. **Backend Port Information**

    The backend service will be available on **port 8080** by default. Once the backend is running, you can access the API by navigating to the following URL in your browser or API client:

    ```http://localhost:8080```

## API Endpoints

The backend provides a set of RESTful API endpoints for managing recipes. Below are the available endpoints and their usage.

### 1. Get All Recipes
- **Endpoint**: `/recipes`
- **Method**: `GET`
- **Description**: Fetch all recipes stored in the database.
- **Response Example**:
  ```json
  [
    {
      "id": "66c2428bfd67acb043c553b7",
      "name": "Pasta",
      "category": "Dinner",
      "ingredients": ["Tomatoes", "Pasta", "Basil"],
      "steps": ["Boil water", "Cook pasta", "Prepare sauce"],
      "tags": ["Vegetarian"],
      "summary": "A simple and quick pasta recipe with tomatoes and basil for a satisfying dinner."
    },
    ...
  ]
  ```

### 2. Get Recipes by Category
- **Endpoint**: `/recipes/{category}`
- **Method**: `GET`
- **Description**: Fetch all recipes that belong to a specific category (e.g., Breakfast, Lunch, Dinner).
- **Path Parameter**:
  - `category`: The category of recipes to fetch (e.g., `Dinner`).
- **Response Example**:
  ```json
  [
    {
      "id": "66c2428bfd67acb043c553b7",
      "name": "Pasta",
      "category": "Dinner",
      "ingredients": ["Tomatoes", "Pasta", "Basil"],
      "steps": ["Boil water", "Cook pasta", "Prepare sauce"],
      "tags": ["Vegetarian"],
      "summary": "A simple and quick pasta recipe with tomatoes and basil for a satisfying dinner."
    },
    ...
  ]
  ```

### 3. Create a Recipe
- **Endpoint**: `/recipes`
- **Method**: `POST`
- **Description**: Add a new recipe to the database.
- **Request Body Example**:
  ```json
  {
    "name": "Chocolate Cake",
    "category": "Dessert",
    "ingredients": ["Flour", "Sugar", "Cocoa powder", "Eggs", "Butter"],
    "steps": ["Preheat oven to 350F", "Mix ingredients", "Bake for 30 minutes"],
    "tags": ["Sweet", "Baking"],
    "summary": "A rich and moist chocolate cake that's perfect for any occasion."
  }
  ```
- **Response Example**:
  ```json
  {
    "id": "66c248f1a5e6634150b8a794",
    "name": "Chocolate Cake",
    "category": "Dessert",
    "ingredients": ["Flour", "Sugar", "Cocoa powder", "Eggs", "Butter"],
    "steps": ["Preheat oven to 350F", "Mix ingredients", "Bake for 30 minutes"],
    "tags": ["Sweet", "Baking"],
    "summary": "A rich and moist chocolate cake that's perfect for any occasion."
  }
  ```

### 4. Get a Recipe by ID
- **Endpoint**: `/recipes/{id}`
- **Method**: `GET`
- **Description**: Fetch a single recipe by its unique ID.
- **Path Parameter**:
  - `id`: The unique identifier of the recipe.
- **Response Example**:
  ```json
  {
    "id": "66c2428bfd67acb043c553b7",
      "name": "Pasta",
      "category": "Dinner",
      "ingredients": ["Tomatoes", "Pasta", "Basil"],
      "steps": ["Boil water", "Cook pasta", "Prepare sauce"],
      "tags": ["Vegetarian"],
      "summary": "A simple and quick pasta recipe with tomatoes and basil for a satisfying dinner."
  }
  ```

### 5. Update a Recipe
- **Endpoint**: `/recipes/{id}`
- **Method**: `PUT`
- **Description**: Update the details of an existing recipe.
- **Path Parameter**:
  - `id`: The unique identifier of the recipe to update.
- **Request Body Example**:
  ```json
  {
    "name": "Spaghetti",
    "category": "Dinner",
    "ingredients": ["Tomatoes", "Spaghetti", "Basil"],
    "steps": ["Boil water", "Cook spaghetti", "Prepare sauce"],
    "tags": ["Vegetarian", "Italian"],
    "summary": "A simple and quick pasta recipe with tomatoes and basil for a satisfying dinner."
  }
  ```
- **Response Example**:
  ```json
  {
    "id": "66c2428bfd67acb043c553b7",
    "name": "Spaghetti",
    "category": "Dinner",
    "ingredients": ["Tomatoes", "Spaghetti", "Basil"],
    "steps": ["Boil water", "Cook spaghetti", "Prepare sauce"],
    "tags": ["Vegetarian", "Italian"],
    "summary": "A simple and quick pasta recipe with tomatoes and basil for a satisfying dinner."
  }
  ```

### 6. Delete a Recipe
- **Endpoint**: `/recipes/{id}`
- **Method**: `DELETE`
- **Description**: Delete a recipe by its unique ID.
- **Path Parameter**:
  - `id`: The unique identifier of the recipe to delete.
- **Response Example**:
  ```json
  {
    "message": "Recipe deleted successfully"
  }
  ```
