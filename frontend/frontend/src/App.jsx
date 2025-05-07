import React, { useState, useEffect } from 'react';
import LoginForm from './components/LoginForm.jsx';
import TextInputSender from './components/TextInputSender.jsx';

const App = () => {
  const [token, setToken] = useState(null);
  const [showLogin, setShowLogin] = useState(false);

  useEffect(() => {
    const storedToken = localStorage.getItem('token');
    if (storedToken) {
      setToken(storedToken);
    }
  }, []);

  const handleLoginSuccess = (newToken) => {
    localStorage.setItem('token', newToken);
    setToken(newToken);
    setShowLogin(false);
  };

  return (
    <div className="p-4 max-w-md mx-auto">
      {!token && showLogin && <LoginForm onLoginSuccess={handleLoginSuccess} />}
      <TextInputSender
        token={token}
        onTriggerLogin={() => {
          console.log('Запрос на логин');
           setShowLogin(true);
        }}
      />
    </div>
  );
};

export default App;
