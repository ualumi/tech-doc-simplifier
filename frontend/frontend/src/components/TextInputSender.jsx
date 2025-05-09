import React, { useState } from 'react';
import axios from 'axios';

const TextInputSender = ({ token, onTriggerLogin }) => {
  const [text, setText] = useState('');
  const [response, setResponse] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleInputClick = () => {
    if (!token) {
      console.log('Клик по input без токена');
      onTriggerLogin();
    }
  };

  const handleSend = async () => {
    if (!token) return;

    try {
      setLoading(true);
      const res = await axios.post(
        'http://localhost:8080/simplify',
        { text },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
      setResponse(res.data);
      setText('');
    } catch (error) {
      console.error('Ошибка отправки:', error);
      setResponse('Ошибка при отправке текста');
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
      {response && <div className="mt-3 p-2 bg-gray-100">{JSON.stringify(response)}</div>}
    </div>
  );
};

export default TextInputSender;
