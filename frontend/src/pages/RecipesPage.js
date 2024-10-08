import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import './RecipesPage.css';
import RecipeTile from '../components/RecipeTile.js';

const RecipeSummaryPage = () => {
   const { mealType } = useParams();
   const [recipes, setRecipes] = useState([]);
   const [loading, setLoading] = useState(true);
   const [error, setError] = useState(null);

   useEffect(() => {
      const fetchRecipes = async () => {
         try {
            const response = await fetch(`http://localhost:8080/recipes/${mealType}`);
            if (!response.ok) {
               if (response.status === 404) {
                  throw new Error('No recipes found for the given category.');
               } else {
                  throw new Error('Failed to fetch recipes');
               }
            }
            const data = await response.json();
            setRecipes(data);
         } catch (err) {
            setError(err.message);
         } finally {
            setLoading(false);
         }
      };

      fetchRecipes();
   }, [mealType]);

   if (loading) {
      return <p>Loading...</p>;
   }

   if (error) {
      return <p>Error: {error}</p>;
   }

   const handleDelete = async (recipeId) => {
      try {
         const response = await fetch(`http://localhost:8080/recipes/${mealType}/${recipeId}`, {
            method: 'DELETE',
         });

         if (!response.ok) {
            throw new Error('Failed to delete recipe');
         }
         setRecipes(recipes.filter((recipe) => recipe.id !== recipeId));
      } catch (err) {
         setError(err.message);
      }
   };

   return (
      <div className="recipe-wrapper">
      <div className="recipe-grid-container">
        {recipes.map((recipe) => (
          <div className="recipe-tile-grid" key={recipe.id}>
            <div className="recipe-tile-container">
              <RecipeTile recipe={recipe} />
            </div>
            <div className="delete-button-container">
              <button onClick={() => handleDelete(recipe.id)} className="delete-button">Delete</button>
            </div>
          </div>
        ))}
      </div>
    </div>
   );
};

export default RecipeSummaryPage;
