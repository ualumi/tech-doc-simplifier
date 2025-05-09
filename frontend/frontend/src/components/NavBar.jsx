import React from 'react';


const NavBar = ({ isAuthenticated, username, onLoginClick }) => {
  const getUserInitial = (name) => {
    if (!name) return '?';
    return name.charAt(0).toUpperCase();
  };

  return (
    <nav className="navbar">
      <div className="logo">
        {/* Место под ваш логотип */}
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

export default NavBar;
