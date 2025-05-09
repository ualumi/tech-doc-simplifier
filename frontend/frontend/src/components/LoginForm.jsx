import React, { useState, useEffect } from 'react';
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

export default LoginForm;
