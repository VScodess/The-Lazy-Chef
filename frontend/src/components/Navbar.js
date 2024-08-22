import React from 'react';
import { useNavigate } from 'react-router-dom';
import './Navbar.css';

const Navbar = () => {
    const navigate = useNavigate();

    const handleLogoClick = () => {
        navigate('/');
    };

    return (
        <nav className="navbar bg-body-tertiary fixed-top">
            <div className="container-fluid">
                <div className="navbar-brand" onClick={handleLogoClick}>
                    <img src="/catCook.png" alt="Website Logo" />
                    <span>The Lazy Chef</span>
                </div>
                <form className="d-flex" role="search">
                    <input className="form-control me-2" type="search" placeholder="Search" aria-label="Search" />
                    <button className="btn btn-outline-success" type="submit">Search</button>
                </form>
            </div>
        </nav>
    );
};

export default Navbar;
