import React from 'react';
import { useNavigate } from 'react-router-dom';
import './RecipeTile.css';

const RecipeTile = ({ recipe }) => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate(`/${recipe.mealType}/${recipe.id}`); 
  };

  return (
    <div className="tile-container" onClick={handleClick}>
      <img className="image" src='/pancake.png' alt={recipe.name} />
      <div className="content">
        <h3 className="title">{recipe.name}</h3>
        <p className="summary">{recipe.summary}</p>
        <div className="tags">
          {recipe.tags.map((tag) => (
            <span key={tag} className="tag">{tag}</span>
          ))}
        </div>
      </div>
    </div>
  );
};

export default RecipeTile;