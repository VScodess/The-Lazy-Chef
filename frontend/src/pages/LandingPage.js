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
      // ... append other fields (ingredients, steps, tags) as needed
      formData.append('image', newRecipe.image);

      const response = await fetch('http://localhost:8080/recipes', { 
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        throw new Error('Failed to add recipe');
      }

      setShowAddRecipeForm(false);
      // ... you might want to reset the newRecipe state here
      alert('Recipe added successfully!');
    } catch (err) {
      console.error('Error adding recipe:', err);
      alert('Failed to add recipe. Please try again.');
    }
  };

  return (
    <div className="container">
      <button onClick={handleAddRecipeClick}>Add New Recipe</button>

      {showAddRecipeForm && (
        <div className="form-container"> {/* Add a container for the form */}
          <h2>Add New Recipe</h2>
          <form onSubmit={handleAddRecipeSubmit}>
            <div>
              <label htmlFor="name">Name:</label>
              <input
                type="text"
                id="name"
                name="name"
                value={newRecipe.name}
                onChange={handleInputChange}
                required
              />
            </div>
            <div>
              <label htmlFor="category">Category:</label>   

              <select
                id="category"
                name="category"
                value={newRecipe.category}
                onChange={handleInputChange}   

                required
              >
                <option value="">Select Category</option> {/* Add a default option */}
                <option value="breakfast">Breakfast</option>
                <option value="lunch">Lunch</option>
                <option value="dinner">Dinner</option>
                <option   
 value="snacks">Snacks</option>
              </select>
            </div>
            <div>   

              <label htmlFor="summary">Summary:</label>
              <textarea
                id="summary"
                name="summary"
                value={newRecipe.summary}
                onChange={handleInputChange}
                required
              />
            </div>
            <div>
              <label htmlFor="ingredients">Ingredients (comma-separated):</label>
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
              />
            </div>
            <div>
              <label htmlFor="steps">Steps (one step per line):</label>
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
              />
            </div>
            <div>
              <label htmlFor="tags">Tags (comma-separated):</label>
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
              />
            </div>
            <div>
              <label htmlFor="image">Image:</label>
              <input
                type="file"
                id="image"
                name="image"
                accept="image/*"
                onChange={handleImageUpload}
              />
            </div>
            <button type="submit">Submit</button>
          </form>
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
            </div>
        </div>
    );
};

export default LandingPage;
