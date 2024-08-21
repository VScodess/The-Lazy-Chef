import React from 'react';
import { useNavigate } from 'react-router-dom';
import './RecipeTile.css';

const RecipeTile = ({ recipe }) => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate(`/${recipe.category}/${recipe.id}`);
  };

  return (
    <div className="tile-container" onClick={handleClick}>
      <img
        src={`data:image/jpeg;base64,${recipe.image}`}
        alt="Recipe Image"
        className="recipe-image"
      />
      <div className="content">
        <h3 className="tile-title">{recipe.name}</h3>
        <p className="summary">{recipe.summary}</p>
        <div className="tags">
          {recipe.tags.map((tag, index) => (
            <span key={index} className="tag">{tag}</span>
          ))}
        </div>
      </div>
    </div>
  );
};

export default RecipeTile;
