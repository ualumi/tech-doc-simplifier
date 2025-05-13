import React, { useState, useEffect } from 'react';
import LoginForm from './components/LoginForm.jsx';
import TextInputSender from './components/TextInputSender.jsx';
import NavBar from './components/NavBar.jsx';
import ResponseViewer from './components/ResponseViewer.jsx';
import HistoryList from './components/HistoryList.jsx';
import ChatHistory from './components/ChatHistory.jsx';
import ChatDetail from './components/ChatDetail.jsx'; 
import NewChatButton from './components/NewChatButton.jsx';
import { parseSimplifiedText } from './components/helpers/parseSimplified.js';

const App = () => {
  const [token, setToken] = useState(null);
  const [showLogin, setShowLogin] = useState(false);
  const [username, setUsername] = useState('');
  const [response, setResponse] = useState(null);
  const [messages, setMessages] = useState([]);
  const [selectedChat, setSelectedChat] = useState(null);
  const [historyRefreshToggle, setHistoryRefreshToggle] = useState(false);


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
    //if (response.original && response.simplified) {
    //  setMessages((prev) => [...prev, response]);
    //} else {
    //  setMessages((prev) => [...prev, { original: lastInput, simplified: response }]);
    //}

     // 👇 переключаем флаг, чтобы HistoryList знал, что нужно обновиться
    setHistoryRefreshToggle(prev => !prev);
  };


  return (
    
    <div className="main-content">
      <div className='history'>
      <NewChatButton onNewChat={() => setSelectedChat(null)} />
      {token && <HistoryList token={token} onSelect={setSelectedChat} refreshTrigger={historyRefreshToggle} />}
      </div>
      <div className='chat'>
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
            <div>
              {selectedChat ? (
                <ChatDetail
                  original={selectedChat.original}
                  simplified={selectedChat.simplified}
                />
              ) : (
                <ChatHistory messages={messages} />
              )}
            </div>
            <TextInputSender
              token={token}
              onTriggerLogin={() => {
                console.log('Запрос на логин');
                setShowLogin(true);
              }}
              onResponse={setResponse}
              onMessageAdd={handleNewMessage} 
              disabled={!!selectedChat}
            />
            
            {response && !selectedChat && (
              <ChatDetail 
                original={response.original || ''}
                simplified={response.simplified || ''}
                createdAt={response.createdAt || new Date().toISOString()}
              />
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default App;