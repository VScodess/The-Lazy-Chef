import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import RecipesPage from './pages/RecipesPage'; 
import RecipeDetails from './components/RecipeDetails'; 
import './App.css'; // Make sure you import your App.css

function App() {
  return (
    <div className="App"> {/* Main container for your app */}
      <div className='logo-container'>
        <img src='/catCook.png' alt='Website Logo' />
        <h1>The Lazy Chef</h1>
      </div>

      <BrowserRouter>
        <Routes>
          <Route path="/" element={<LandingPage />} /> 
          <Route path="/:mealType" element={<RecipesPage/>} /> 
          <Route path="/:mealType/:recipeId" element={<RecipeDetails />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;