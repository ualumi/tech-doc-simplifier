{/*import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './LoginForm.css'; // Подключаем CSS

const LoginForm = ({ onLoginSuccess, onClose }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [visible, setVisible] = useState(false);

  useEffect(() => {
    setTimeout(() => setVisible(true), 10);
  }, []);

  const handleLogin = async () => {
    try {
      setLoading(true);
      setError('');

      // Отправляем JSON-запрос
      const res = await axios.post('http://localhost:8081/login', {
        login: username,
        password: password
      }, {
        headers: {
          'Content-Type': 'application/json'
        },
        withCredentials: true
      });
      onLoginSuccess(res.data.token);
    } catch (err) {
      setError('Неверный логин или пароль');
    } finally {
      setLoading(false);
    }
  };

  const handleClose = () => {
    setVisible(false);
    setTimeout(onClose, 200);
  };

  return (
    <div className={`login-overlay ${visible ? 'visible' : ''}`}>
      <div className={`loginform ${visible ? 'visible' : ''}`}>
        <button onClick={handleClose} className="close">
          &times;
        </button>
        <h4>Login</h4>
        <form
          onSubmit={(e) => {
            e.preventDefault();
            handleLogin();
          }}
        >
          <input
            type="text"
            name="username"
            className="username form-control"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <input
            type="password"
            name="password"
            className="password form-control"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <input
            className="btn login"
            type="submit"
            value={loading ? 'Logging in...' : 'Login'}
            disabled={loading}
          />
        </form>
        {error && <p className="error">{error}</p>}
      </div>
    </div>
  );
};

export default LoginForm;*/}

import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './LoginForm.css';

const LoginForm = ({ onLoginSuccess, onClose }) => {
  const [mode, setMode] = useState('login'); // 'login' or 'signup'
  const [email, setEmail] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [visible, setVisible] = useState(false);

  useEffect(() => {
    setTimeout(() => setVisible(true), 10);
  }, []);

  const resetFields = () => {
    setEmail('');
    setUsername('');
    setPassword('');
    setError('');
  };

  const handleLogin = async () => {
    try {
      setLoading(true);
      setError('');

      const res = await axios.post('http://87.228.89.190/authorization/login', {
        login: username,
        password: password
      }, {
        headers: { 'Content-Type': 'application/json' },
        withCredentials: true
      });

      onLoginSuccess(res.data.token, username);
    } catch (err) {
      setError('Неверный логин или пароль');
    } finally {
      setLoading(false);
    }
  };

  const handleSignup = async () => {
    try {
      setLoading(true);
      setError('');

      const res = await axios.post('http://87.228.89.190/authorization/registration', {
        email,
        login: username,
        password
      }, {
        headers: { 'Content-Type': 'application/json' },
        withCredentials: true
      });

      // Если регистрация прошла успешно, сразу логиним
      await handleLogin(); // <--- авто-логин после регистрации

    } catch (err) {
      setError('Ошибка регистрации. Возможно, пользователь уже существует.');
    } finally {
      setLoading(false);
    }
  };


  const handleClose = () => {
    setVisible(false);
    setTimeout(() => {
      resetFields();
      onClose();
    }, 200);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (mode === 'login') {
      handleLogin();
    } else {
      handleSignup();
    }
  };

  return (
    <div className={`login-overlay ${visible ? 'visible' : ''}`}>
      <div className={`loginform ${visible ? 'visible' : ''}`}>
        <button onClick={handleClose} id='close1' className="close">&times;</button>

        <div className="tabs">
          
          <button
            id='loginbutton'
            className={mode === 'signup' ? 'active' : ''}
            onClick={() => {
              setMode('login');
              resetFields();
            }}
          >
            Log in
          </button>
          <button
            id='loginbutton'
            className={mode === 'login' ? 'active' : ''}
            onClick={() => {
              setMode('signup');
              resetFields();
            }}
          >
            Sign Up
          </button>
        </div>

        <form onSubmit={handleSubmit}>
          {mode === 'login' && (<div><p id='hithere'>Please, login or signup to use service</p></div>)}
          
          {mode === 'signup' && (
            <input
              type="email"
              name="email"
              className="form-control"
              placeholder="Email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          )}
          <input
            type="text"
            name="username"
            className="form-control"
            placeholder="Login"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
          <input
            type="password"
            name="password"
            className="form-control"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <input
            className="btn login"
            type="submit"
            value={loading ? (mode === 'login' ? 'Logging in...' : 'Signing up...') : (mode === 'login' ? 'Login' : 'Sign Up')}
            disabled={loading}
          />
        </form>

        {error && <p className="error">{error}</p>}
      </div>
    </div>
  );
};

export default LoginForm;

