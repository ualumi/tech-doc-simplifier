import React, { useState, useEffect } from 'react';
import LoginForm from './components/LoginForm.jsx';
import TextInputSender from './components/TextInputSender.jsx';
import NavBar from './components/NavBar.jsx';
import ResponseViewer from './components/ResponseViewer.jsx';
import HistoryList from './components/HistoryList.jsx';
import ChatHistory from './components/ChatHistory.jsx'

const App = () => {
  const [token, setToken] = useState(null);
  const [showLogin, setShowLogin] = useState(false);
  const [username, setUsername] = useState('');
  const [response, setResponse] = useState(null);
  const [messages, setMessages] = useState([]);

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

  const handleLogout = () => {
    localStorage.removeItem('token');
    setToken(null);
    setUsername('');
    setShowLogin(true);
  };

  const handleNewMessage = (response) => {
    if (response.original && response.simplified) {
      setMessages((prev) => [...prev, response]);
    } else {
      setMessages((prev) => [...prev, { original: lastInput, simplified: response }]);
    }
  };

  return (
    
    <div className="main-content">
      <div>
        {token && <HistoryList token={token} />}
      </div>
      <div>
      <div>
        <NavBar
          isAuthenticated={!!token}
          username={username}
          onLoginClick={() => setShowLogin(true)}
          onLogout={handleLogout}
        />
        </div>

        <div className='content'>
          <div className='text-field'>
            <div>
              <p className='HighText'>Чем могу помочь?</p>
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
            <ChatHistory messages={messages} />
            <TextInputSender
              token={token}
              onTriggerLogin={() => {
                console.log('Запрос на логин');
                setShowLogin(true);
              }}
              onResponse={setResponse}
            />
            <ResponseViewer response={response} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default App;
