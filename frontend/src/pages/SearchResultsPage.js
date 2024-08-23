import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import './RecipesPage.css'; // Reusing the same CSS
import RecipeTile from '../components/RecipeTile.js';

const SearchResultsPage = () => {
   const [recipes, setRecipes] = useState([]);
   const [loading, setLoading] = useState(true);
   const [error, setError] = useState(null);
   const location = useLocation();

   useEffect(() => {
      const fetchSearchResults = async () => {
         const query = new URLSearchParams(location.search).get('q');
         const category = new URLSearchParams(location.search).get('category');
         try {
            const response = await fetch(`http://localhost:8080/recipes/search?q=${encodeURIComponent(query)}&category=${encodeURIComponent(category)}`);
            if (!response.ok) {
               if (response.status === 404) {
                  throw new Error('No recipes found for the given search.');
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

      fetchSearchResults();
   }, [location.search]);

   if (loading) {
      return <p>Loading...</p>;
   }

   if (error) {
      return <p>Error: {error}</p>;
   }

   const handleDelete = async (recipeId) => {
      try {
         const response = await fetch(`http://localhost:8080/recipes/${recipeId}`, { // Adjusted the URL for deletion
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

export default SearchResultsPage;
