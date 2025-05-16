{/*import React from 'react';


const NavBar = ({ isAuthenticated, username, onLoginClick }) => {
  const getUserInitial = (name) => {
    if (!name) return '?';
    return name.charAt(0).toUpperCase();
  };

  return (
    <nav className="navbar">
      <div className="logo">
        
        <span className='RegularText' >logo</span>
      </div>

      <ul className="nav-links">
        <li><a className='RegularText' href="/about">О нас</a></li>
        <li><a className='RegularText' href="/simplifier">ИИ</a></li>
        <li><a className='RegularText' href="/api">API</a></li>
        <li><a className='RegularText' href="/api">Прочее</a></li>
      </ul>

      <div className="auth-section">
        {!isAuthenticated ? (
          <button onClick={onLoginClick}>Login</button>
        ) : (
          <div className="user-avatar">
            {getUserInitial(username)}
          </div>
        )}
      </div>
    </nav>
  );
};

export default NavBar;*/}

import React, { useState, useRef, useEffect } from 'react';
import './NavBar.css'; // Добавь стили или CSS-in-JS

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
        {/*<li><a className="RegularText" href="/about">О нас</a></li>
        <li><a className="RegularText" href="/simplifier">ИИ</a></li>
        <li><a className="RegularText" href="/api">Прочее</a></li>
        <li><a className="RegularText" href="/misc">Прочее</a></li>*/}
      </ul>

      <div className="auth-section" ref={menuRef}>
        {!isAuthenticated ? (
          <button onClick={onLoginClick}>Login</button>
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

