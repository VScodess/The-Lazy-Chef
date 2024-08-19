import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import './RecipesPage.css';
import RecipeTile from '../components/RecipeTile.js';

const recipesData = [
  {
    id: 1,
    name: 'Fluffy Pancakes',
    mealType: 'breakfast',
    tags: ['vegetarian'],
    summary: 'A classic breakfast treat, light and fluffy pancakes topped with your favorite syrup and fruit.',
  },
  {
    id: 2,
    name: 'Avocado Toast with Egg',
    mealType: 'breakfast',
    tags: ['vegetarian'],
    summary: 'A healthy and satisfying breakfast option, creamy avocado on toast with a perfectly cooked egg.',
  },
  {
    id: 3,
    name: 'Caprese Salad',
    mealType: 'lunch',
    tags: ['vegetarian', 'vegan'],
    summary: 'A refreshing and simple salad with juicy tomatoes, fresh mozzarella, and fragrant basil.',
  },
  {
    id: 4,
    name: 'Grilled Chicken Caesar Salad',
    mealType: 'lunch',
    tags: ['non-vegetarian'],
    summary: 'A classic Caesar salad with grilled chicken, crisp romaine lettuce, and a creamy dressing.',
  },
  {
    id: 5,
    name: 'Spaghetti Bolognese',
    mealType: 'dinner',
    tags: ['non-vegetarian'],
    summary: 'A hearty and comforting Italian dish with a rich meat sauce and spaghetti.',
  },
  {
    id: 6,
    name: 'Vegetable Curry',
    mealType: 'dinner',
    tags: ['vegetarian', 'vegan'],
    summary: 'A flavorful and aromatic curry with a variety of vegetables simmered in a creamy coconut sauce.',
  },
  {
    id: 7,
    name: 'Trail Mix',
    mealType: 'snacks',
    tags: ['vegetarian', 'vegan'],
    summary: 'A mix of nuts, seeds, and dried fruits for a quick and healthy energy boost.',
  },
  {
    id: 8,
    name: 'Hummus with Pita Bread',
    mealType: 'snacks',
    tags: ['vegetarian', 'vegan'],
    summary: 'A creamy and flavorful dip made from chickpeas, perfect for dipping with pita bread or vegetables.',
  },
];

const RecipeSummaryPage = () => {
  const { mealType } = useParams();
  const [recipes, setRecipes] = useState([]);

  useEffect(() => {
    const filteredRecipes = recipesData.filter(
      (recipe) => recipe.mealType === mealType
    );
    setRecipes(filteredRecipes);
  }, [mealType]);

  return (
    <div className="container">
        <div className='logo-container'>
                <img src='/catCook.png' alt='Website Logo' />
                <h1>The Lazy Chef</h1>
            </div>
      {recipes.map((recipe) => (
        <RecipeTile key={recipe.id} recipe={recipe} />
      ))}
    </div>
  );
};

export default RecipeSummaryPage;