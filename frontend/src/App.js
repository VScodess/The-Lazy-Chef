import React from 'react';
import { BrowserRouter, Routes, Route, useNavigate } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import RecipesPage from './pages/RecipesPage';
import SearchResultsPage from './pages/SearchResultsPage';
import RecipeDetails from './components/RecipeDetails';
import Navbar from './components/Navbar';
import './App.css';

function App() {
  const navigate = useNavigate();

  const handleLogoClick = () => {
    navigate('/');
  };

  return (
    <div className="App">
      <Navbar />

      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/:mealType" element={<RecipesPage />} />
        <Route path="/:mealType/:recipeId" element={<RecipeDetails />} />
        <Route path="/search" element={<SearchResultsPage />} />
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
