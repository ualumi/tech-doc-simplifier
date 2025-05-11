{/*import React, { useState } from 'react';
import axios from 'axios';

const TextInputSender = ({ token, onTriggerLogin, onResponse }) => {
  const [text, setText] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSend = async () => {
    if (!token || !text.trim()) return;

    const userMessage = { type: 'user', text };
    onResponse(userMessage); // показать введённый текст сразу
    const originalText = text;
    setText(''); // очистка поля ввода

    try {
      setLoading(true);
      const res = await axios.post(
        'http://localhost:8080/simplify',
        { text: originalText },
        {
          headers: { Authorization: `Bearer ${token}` },
          validateStatus: () => true,
        }
      );

      if (res.status === 401) {
        console.warn('⛔ Unauthorized – вызываем форму входа');
        onTriggerLogin();
        return;
      }

      if (res.status >= 200 && res.status < 300) {
        const simplified = res.data?.simplified || 'Нет упрощённого текста';
        onResponse({ type: 'bot', text: simplified });
      } else {
        onResponse({ type: 'bot', text: 'Ошибка при отправке текста' });
      }
    } catch (error) {
      console.error('❌ Ошибка отправки:', error);
      onResponse({ type: 'bot', text: 'Ошибка при отправке текста' });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <input
        type="text"
        className="InputText"
        placeholder="Введите текст..."
        value={text}
        readOnly={!token}
        onClick={() => !token && onTriggerLogin()}
        onChange={(e) => setText(e.target.value)}
      />
      <button
        onClick={handleSend}
        disabled={!token || loading}
        className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 disabled:opacity-50"
      >
        {loading ? 'Отправка...' : 'Отправить'}
      </button>
    </div>
  );
};

export default TextInputSender;*/}

import React, { useState } from 'react';
import axios from 'axios';

const TextInputSender = ({ token, onTriggerLogin, onResponse }) => {
  const [text, setText] = useState('');
  const [loading, setLoading] = useState(false);

  const handleInputClick = () => {
    if (!token) {
      console.log('Клик по input без токена');
      onTriggerLogin();
    }
  };

  const handleSend = async () => {
    if (!token || !text.trim()) return;

    const userInput = text;
    setText('');
    onResponse({ original: userInput, simplified: '...' }); // Показать до ответа

    try {
      setLoading(true);
      const res = await axios.post(
        'http://localhost:8080/simplify',
        { text: userInput },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
          validateStatus: () => true,
        }
      );

      if (res.status === 401) {
        console.warn('⛔ Unauthorized – вызываем форму входа');
        onTriggerLogin();
        return;
      }

      if (res.status >= 200 && res.status < 300) {
        onResponse({ original: userInput, simplified: res.data.simplified });
      } else {
        onResponse({ original: userInput, simplified: 'Ошибка при отправке текста' });
      }
    } catch (error) {
      console.error('❌ Ошибка отправки:', error);
      onResponse({ original: userInput, simplified: 'Ошибка при отправке текста' });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="mt-4 flex gap-2">
      <input
        type="text"
        className="InputText flex-1 p-2 border rounded"
        placeholder="Введите текст..."
        value={text}
        onClick={handleInputClick}
        readOnly={!token}
        onChange={(e) => setText(e.target.value)}
      />
      <button
        onClick={handleSend}
        disabled={!token || loading}
        className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 disabled:opacity-50"
      >
        {loading ? 'Отправка...' : 'Отправить'}
      </button>
    </div>
  );
};

export default TextInputSender;

