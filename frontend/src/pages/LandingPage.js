import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './LandingPage.css';

const LandingPage = () => {
  const
    navigate = useNavigate();
  const
    [showAddRecipeForm, setShowAddRecipeForm] = useState(false);
  const [newRecipe, setNewRecipe] = useState({
    name: '',
    category: '',
    ingredients: [],
    steps: [],
    tags: [],
    summary: '',
    image: null,
  });

  const handleButtonClick = (mealType) => {
    navigate(`/${mealType}`);
  };

  const handleAddRecipeClick = () => {
    setShowAddRecipeForm(true);
  };

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setNewRecipe({ ...newRecipe, [name]: value });
  };

  const handleImageUpload = (event) => {
    const file = event.target.files[0];
    setNewRecipe({ ...newRecipe, image: file });
  };

  const handleAddRecipeSubmit = async (event) => {
    event.preventDefault();

    try {
      const formData = new FormData();
      formData.append('name', newRecipe.name);
      formData.append('category', newRecipe.category);
      formData.append('ingredients', newRecipe.ingredients);
      formData.append('summary', newRecipe.summary);
      formData.append('steps', newRecipe.steps);
      formData.append('tags', newRecipe.tags);
      formData.append('image', newRecipe.image);

      const response = await fetch('http://localhost:8080/recipes', {
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        throw new Error('Failed to add recipe');
      }

      setShowAddRecipeForm(false);

      alert('Recipe added successfully!');
    } catch (err) {
      console.error('Error adding recipe:', err);
      alert('Failed to add recipe. Please try again.');
    }
  };

  return (
    <div className="landing-container">
      {showAddRecipeForm && (
        <div
          className="form-overlay"
          onClick={() => setShowAddRecipeForm(false)}
        >
          <div
            className="form-container"
            onClick={(e) => e.stopPropagation()}
          >
            <h2>Add New Recipe</h2>
            <form onSubmit={handleAddRecipeSubmit}>

              <div>
                <label htmlFor="name">Name</label>
                <input
                  type="text"
                  id="name"
                  name="name"
                  value={newRecipe.name}
                  onChange={handleInputChange}
                  required
                  placeholder="What should we call this masterpiece?"
                />
              </div>

              <div>
                <label htmlFor="category">Category</label>
                <select
                  id="category"
                  name="category"
                  value={newRecipe.category}
                  onChange={handleInputChange}
                  required
                >
                  <option value="">Choose a category...</option>
                  <option value="breakfast">Breakfast</option>
                  <option value="lunch">Lunch</option>
                  <option value="dinner">Dinner</option>
                  <option value="snacks">Snacks</option>
                </select>
              </div>

              <div>
                <label htmlFor="summary">Summary</label>
                <textarea
                  id="summary"
                  name="summary"
                  value={newRecipe.summary}
                  onChange={handleInputChange}
                  required
                  placeholder="Describe your recipe in a delicious one-liner!"
                />
              </div>

              <div>
                <label htmlFor="ingredients">Ingredients</label>
                <textarea
                  id="ingredients"
                  name="ingredients"
                  value={newRecipe.ingredients.join(', ')}
                  onChange={(e) =>
                    setNewRecipe({
                      ...newRecipe,
                      ingredients: e.target.value
                        .split(',')
                        .map((item) => item.trim()),
                    })
                  }
                  required
                  placeholder="What do we need? List ingredients, separated by commas"
                />
              </div>

              <div>
                <label htmlFor="steps">Steps</label>
                <textarea
                  id="steps"
                  name="steps"
                  value={newRecipe.steps.join('\n')}
                  onChange={(e) =>
                    setNewRecipe({
                      ...newRecipe,
                      steps: e.target.value.split('\n'),
                    })
                  }
                  required
                  placeholder="Step-by-step: One instruction per line, please!"
                />
              </div>

              <div>
                <label htmlFor="tags">Tags</label>
                <input
                  type="text"
                  id="tags"
                  name="tags"
                  value={newRecipe.tags.join(', ')}
                  onChange={(e) =>
                    setNewRecipe({
                      ...newRecipe,
                      tags: e.target.value
                        .split(',')
                        .map((item) => item.trim()),
                    })
                  }
                  required
                  placeholder="How would you tag this? (e.g., spicy, vegan)"
                />
              </div>

              <div className="file-upload-container">
                <label htmlFor="image">Image:</label>
                <input
                  type="file"
                  id="image"
                  name="image"
                  accept="image/*"
                  onChange={handleImageUpload}
                  required
                />
              </div>

              <button type="submit">Submit</button>
            </form>
          </div>
        </div>
      )}


      <div className="buttonGrid">

        <button className='meal-button' onClick={() => handleButtonClick('breakfast')}>
          Breakfast
        </button>
        <button className='meal-button' onClick={() => handleButtonClick('lunch')}>
          Lunch
        </button>
        <button className='meal-button' onClick={() => handleButtonClick('dinner')}>
          Dinner
        </button>
        <button className='meal-button' onClick={() => handleButtonClick('snacks')}>
          Snacks
        </button>
        <button className='meal-button add-recipe-button' onClick={handleAddRecipeClick}>Add your recipe</button>
      </div>
    </div>
  );
};

export default LandingPage;
