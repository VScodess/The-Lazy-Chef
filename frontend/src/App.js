import React from 'react';
import { BrowserRouter, Routes, Route, useNavigate } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import RecipesPage from './pages/RecipesPage';
import RecipeDetails from './components/RecipeDetails';
import './App.css';

function App() {
  const navigate = useNavigate();

  const handleLogoClick = () => {
    navigate('/');
  };

  return (
    <div className="App">
      <div className='logo-container' onClick={handleLogoClick} style={{ cursor: 'pointer' }}>
        <img src='/catCook.png' alt='Website Logo' />
        <h1>The Lazy Chef</h1>
      </div>

      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/:mealType" element={<RecipesPage />} />
        <Route path="/:mealType/:recipeId" element={<RecipeDetails />} />
      </Routes>
    </div>
  );
}

function AppWrapper() {
  return (
    <BrowserRouter>
      <App />
    </BrowserRouter>
  );
}

export default AppWrapper;
