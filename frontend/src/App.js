
import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import RecipesPage from './pages/RecipesPage'; 

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<LandingPage />} /> 
        <Route path="/:mealType" element={<RecipesPage/>} /> 
      </Routes>
    </BrowserRouter>
  );
}

export default App;