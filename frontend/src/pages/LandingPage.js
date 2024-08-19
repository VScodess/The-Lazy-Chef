import React from 'react';
import { useNavigate } from 'react-router-dom';
import './LandingPage.css';
const LandingPage = () => {
    const navigate = useNavigate();

    const handleButtonClick = (mealType) => {
        navigate(`/${mealType}`);
    };

    return (
        <div className='container'>
            
            <div className='buttonGrid'>
                <button className='meal-button' onClick={() => handleButtonClick('breakfast')}>
                    Breakfast
                </button>
                <button className='meal-button' onClick={() => handleButtonClick('lunch')}>
                    Lunch
                </button>
                <button className='meal-button' onClick={() => handleButtonClick('dinner')}>
                    Dinner
                </button>
                <button className='meal-button' onClick={() => handleButtonClick('snacks')}>
                    Snacks
                </button>
            </div>
        </div>
    );
};

export default LandingPage;