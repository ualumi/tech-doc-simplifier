import React, { useState, useEffect } from 'react';
import LoginForm from './components/LoginForm.jsx';
import TextInputSender from './components/TextInputSender.jsx';
import NavBar from './components/NavBar.jsx';
import FloatingCircles from './components/FloatingCircles.jsx';


const App = () => {
  const [token, setToken] = useState(null);
  const [showLogin, setShowLogin] = useState(false);
  const [username, setUsername] = useState('');

  useEffect(() => {
    const storedToken = localStorage.getItem('token');
    if (storedToken) {
      setToken(storedToken);
    }
  }, []);

  const handleLoginSuccess = (newToken) => {
    localStorage.setItem('token', newToken);
    setToken(newToken);
  
    // Раскодируй JWT или запроси имя пользователя с сервера — временно можно заглушку:
    const mockUsername = 'Ivan'; // ← замени на реальное имя, полученное с сервера
    setUsername(mockUsername);
  
    setShowLogin(false);
  };

  return (
    
    <div className="main-content">
      <FloatingCircles />
      <div>
        <NavBar
          isAuthenticated={!!token}
          username={username}
          onLoginClick={() => setShowLogin(true)}
        />
      </div>

      <div className='content'>
        <div className='text-field'>
          <div>
            <p className='HighText'>How can I help you?</p>
          </div>
          <div>
            <p className='RegularText'>Lorem ipsum dolor sit amet consectetur adipisicing elit. Debitis illum, aperiam eum laudantium impedit rerum?</p>
          </div>
          
        </div>

        <div>
          {showLogin && (
            <LoginForm
              onLoginSuccess={handleLoginSuccess}
              onClose={() => setShowLogin(false)}
            />
          )}
        </div>
        
        <div className="input-field">
          <TextInputSender
            token={token}
            onTriggerLogin={() => {
              console.log('Запрос на логин');
              setShowLogin(true);
            }}
          />
        </div>
      </div>
      
      <div className='footer-elem'>
        <div className='RegularText'>Lorem ipsum dolor sit amet consectetur adipisicing elit.</div>
        <div className='RegularText'>Lorem ipsum, dolor sit amet consectetur adipisicing elit. Iusto consequatur hic alias pariatur, officiis culp.</div>
      </div>
    </div>
  );
};

export default App;
