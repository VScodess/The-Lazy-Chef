import React, { useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import './Navbar.css';

const Navbar = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const [searchQuery, setSearchQuery] = useState('');

    const handleLogoClick = () => {
        navigate('/');
    };

    const handleSearchInputChange = (e) => {
        setSearchQuery(e.target.value);
    };

    const handleSearchSubmit = (e) => {
        e.preventDefault();
        if (searchQuery.trim() !== '') {
            const searchParams = new URLSearchParams(location.search);
            let category = searchParams.get('category');

            if (!category) {
                category = location.pathname.split('/')[1];
            }

            if (category) {
                navigate(`/search?q=${encodeURIComponent(searchQuery)}&category=${encodeURIComponent(category)}`);
            } else {
                alert("Pleas select a category before searching.")
            }
        }
    };

    const shouldShowSearch = location.pathname !== '/';


    return (
        <nav className="navbar fixed-top navbar-custom">
            <div className="container-fluid">
                <div className="navbar-brand" onClick={handleLogoClick}>
                    <img src="/catCook.png" alt="Website Logo" />
                    <span>The Lazy Chef</span>
                </div>
                {shouldShowSearch && (
                    <form className="d-flex" role="search" onSubmit={handleSearchSubmit}>
                        <input
                            className="form-control me-2"
                            type="search"
                            placeholder="Look up a recipe"
                            aria-label="Search"
                            value={searchQuery}
                            onChange={handleSearchInputChange}
                        />
                        <button className="btn custom-search-button" type="submit">Search</button>
                    </form>
                )}
            </div>
        </nav>
    );
};

export default Navbar;
