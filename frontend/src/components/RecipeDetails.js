import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import './RecipeDetails.css';

const RecipeDetails = () => {
  const { mealType, recipeId } = useParams();
  const [recipe, setRecipe] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchRecipe = async () => {
      try {
        const response = await fetch(`http://localhost:8080/recipes/${mealType}/${recipeId}`);
        if (!response.ok) {
          throw new Error('Failed to fetch recipe');
        }
        const data = await response.json();
        setRecipe(data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchRecipe();
  }, [mealType, recipeId]);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div className="container">
      <h1 className="title">{recipe.name}</h1>

      <div className="section">
        <h2 className="section-title">Ingredients</h2>
        <ul className="list">
          {recipe.ingredients.map((ingredient, index) => (
            <li key={index} className="list-item">{ingredient}</li>
          ))}
        </ul>
      </div>

      <div className="section">
        <h2 className="section-title">Instructions</h2>
        <p className="instructions">{recipe.steps}</p>
      </div>
    </div>
  );
};

export default RecipeDetails;
