import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import './RecipeDetails.css';

const RecipeDetails = () => {
    const { mealType, recipeId } = useParams();
    const [recipe, setRecipe] = useState(null);
    const recipesData = [
        {
          id: 1,
          name: 'Fluffy Pancakes',
          mealType: 'breakfast',
          tags: ['vegetarian'],
          summary: 'A classic breakfast treat, light and fluffy pancakes topped with your favorite syrup and fruit.',
          image: 'https://example.com/pancakes.jpg', // Replace with actual image URL
          ingredients: [
            '1 cup all-purpose flour',
            '2 tablespoons granulated sugar',
            '2 teaspoons baking powder',
            '1/2 teaspoon baking soda',
            '1/4 teaspoon salt',
            '1 cup milk',
            '1 egg',
            '2 tablespoons unsalted butter, melted',
          ],
          instructions: `
      1. In a large bowl, whisk together the flour, sugar, baking powder, baking soda, and salt.
      2. In a separate bowl, whisk together the milk, egg, and melted butter.
      3. Pour the wet ingredients into the dry ingredients and whisk until just combined (some lumps are okay).
      4. Heat a lightly greased griddle or frying pan over medium heat.
      5. Pour 1/4 cup of batter onto the hot griddle for each pancake.
      6. Cook until bubbles form on the surface and the edges look set, then flip and cook until golden brown on the other side.
      7. Serve warm with your favorite toppings.
          `,
        },
        {
          id: 2,
          name: 'Avocado Toast with Egg',
          mealType: 'breakfast',
          tags: ['vegetarian'],
          summary: 'A healthy and satisfying breakfast option, creamy avocado on toast with a perfectly cooked egg.',
          image: 'https://example.com/avocado-toast.jpg',
          ingredients: [
            '1 slice of bread',
            '1/2 ripe avocado',
            '1 egg',
            'Salt and pepper to taste',
            'Optional toppings: red pepper flakes, Everything but the Bagel seasoning, lemon juice',
          ],
          instructions: `
      1. Toast the bread to your desired level of crispiness.
      2. While the bread is toasting, mash the avocado in a small bowl and season with salt and pepper.
      3. Heat a small skillet over medium heat and add a drizzle of oil.
      4. Crack the egg into the skillet and cook to your desired doneness (sunny-side up, over-easy, etc.).
      5. Spread the mashed avocado on the toasted bread.
      6. Top with the cooked egg and any additional toppings.
      7. Enjoy!
          `,
        },
        // ... add details for other recipes in a similar format
      ];
      
    useEffect(() => {
      // Fetch recipe data based on mealType and recipeId (replace with actual API call)
      const foundRecipe = recipesData.find(
        (r) => r.mealType === mealType && r.id === parseInt(recipeId)
      );
      setRecipe(foundRecipe);
    }, [mealType, recipeId]);
  
    if (!recipe) {
      return <div>Loading...</div>; // Or display an error message
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
        <p className="instructions">{recipe.instructions}</p>
      </div>
    </div>
  );
};

export default RecipeDetails;