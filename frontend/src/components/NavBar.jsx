

import React, { useState, useRef, useEffect } from 'react';
import './NavBar.css'; 

const NavBar = ({ isAuthenticated, username, onLoginClick, onLogout }) => {
  const [menuOpen, setMenuOpen] = useState(false);
  const menuRef = useRef();

  const getUserInitial = (name) => {
    if (!name) return '?';
    return name.charAt(0).toUpperCase();
  };

  const toggleMenu = () => setMenuOpen((prev) => !prev);

  const handleClickOutside = (event) => {
    if (menuRef.current && !menuRef.current.contains(event.target)) {
      setMenuOpen(false);
    }
  };

  useEffect(() => {
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  return (
    <nav className="navbar">
      <div className="logo">
        <span className="RegularText">logo</span>
      </div>

      <ul className="nav-links">

      </ul>

      <div className="auth-section" ref={menuRef}>
        {!isAuthenticated ? (
          <button id='loginbutton' onClick={onLoginClick}>Log in</button>
        ) : (
          <div className="avatar-wrapper" onClick={toggleMenu}>
            <div className="user-avatar">
              {getUserInitial(username)}
            </div>
            {menuOpen && (
              <div className="dropdown-menu">
                <p className="dropdown-username">{username}</p>
                <button className="dropdown-logout" onClick={onLogout}>
                  Выйти
                </button>
              </div>
            )}
          </div>
        )}
      </div>
    </nav>
  );
};

export default NavBar;

