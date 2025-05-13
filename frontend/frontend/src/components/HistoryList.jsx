{/*import React, { useEffect, useState } from 'react'; 
import axios from 'axios';
import HistoryItem from './HistoryItem';
import './HistoryList.css';

const groupByDate = (items) => {
  const today = new Date();
  const yesterday = new Date();
  yesterday.setDate(today.getDate() - 1);

  const formatDateOnly = (dateStr) => new Date(dateStr).toDateString();

  const groups = {
    Сегодня: [],
    Вчера: [],
    Ранее: [],
  };

  items.forEach((item) => {
    const itemDate = new Date(item.created_at);
    const itemDay = formatDateOnly(itemDate);

    if (itemDay === formatDateOnly(today)) {
      groups.Сегодня.push(item);
    } else if (itemDay === formatDateOnly(yesterday)) {
      groups.Вчера.push(item);
    } else {
      groups.Ранее.push(item);
    }
  });

  return groups;
};

const HistoryList = ({ token, onSelect }) => {
  const [historyGroups, setHistoryGroups] = useState(null);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchHistory = async () => {
      try {
        const res = await axios.get('http://localhost:8080/history', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
          validateStatus: () => true,
        });

        if (res.status === 500) {
          setError('Ошибка при загрузке истории');
          return;
        }

        if (res.status >= 400) {
          setError('Ошибка при загрузке истории');
          return;
        }

        const sorted = [...res.data.results].sort(
          (a, b) => new Date(b.created_at) - new Date(a.created_at)
        );

        const grouped = groupByDate(sorted);
        setHistoryGroups(grouped);
      } catch (err) {
        console.error('Ошибка загрузки истории:', err);
        setError('Ошибка при загрузке истории');
      }
    };

    if (token) {
      fetchHistory();
    }
  }, [token]);

  return (
    <div className="historycontent">
      <h2 className="text-xl font-bold mb-4">История упрощений</h2>

      {error ? (
        <p className="text-red-500">{error}</p>
      ) : !historyGroups ? (
        <p className="text-gray-500">Загрузка...</p>
      ) : (
        Object.entries(historyGroups).map(([label, items]) =>
          items.length > 0 ? (
            <div key={label} className="mb-6">
              <h3 className="text-lg font-semibold mb-2">{label}</h3>
              {items.map((item, index) => (
                <HistoryItem
                  key={index}
                  original={item.original}
                  simplified={item.simplified}
                  createdAt={item.created_at}
                  onClick={() =>
                    onSelect({
                      original: item.original,
                      simplified: item.simplified,
                      createdAt: item.created_at
                    })
                  }
                />
              ))}
            </div>
          ) : null
        )
      )}
    </div>
  );
};

export default HistoryList;*/}

import React, { useEffect, useState } from 'react';
import axios from 'axios';
import HistoryItem from './HistoryItem';
import './HistoryList.css';

const groupByDate = (items) => {
  const today = new Date();
  const yesterday = new Date();
  yesterday.setDate(today.getDate() - 1);

  const formatDateOnly = (dateStr) => new Date(dateStr).toDateString();

  const groups = {
    Сегодня: [],
    Вчера: [],
    Ранее: [],
  };

  items.forEach((item) => {
    const itemDate = new Date(item.created_at);
    const itemDay = formatDateOnly(itemDate);

    if (itemDay === formatDateOnly(today)) {
      groups.Сегодня.push(item);
    } else if (itemDay === formatDateOnly(yesterday)) {
      groups.Вчера.push(item);
    } else {
      groups.Ранее.push(item);
    }
  });

  return groups;
};

const HistoryList = ({ token, onSelect, refreshTrigger }) => {
  const [historyGroups, setHistoryGroups] = useState(null);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchHistory = async () => {
      try {
        const res = await axios.get('http://localhost:8080/history', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
          validateStatus: () => true,
        });

        if (res.status === 500 || res.status >= 400) {
          setError('Ошибка при загрузке истории');
          return;
        }

        const sorted = [...res.data.results].sort(
          (a, b) => new Date(b.created_at) - new Date(a.created_at)
        );

        const grouped = groupByDate(sorted);
        setHistoryGroups(grouped);
      } catch (err) {
        console.error('Ошибка загрузки истории:', err);
        setError('Ошибка при загрузке истории');
      }
    };

    if (token) {
      fetchHistory();
    }
  }, [token, refreshTrigger]); // 👈 теперь следит за refreshTrigger

  return (
    <div className="historycontent">
      <h2 className="text-xl font-bold mb-4">История упрощений</h2>

      {error ? (
        <p className="text-red-500">{error}</p>
      ) : !historyGroups ? (
        <p className="text-gray-500">Загрузка...</p>
      ) : (
        Object.entries(historyGroups).map(([label, items]) =>
          items.length > 0 ? (
            <div key={label} className="mb-6">
              <h3 className="text-lg font-semibold mb-2">{label}</h3>
              {items.map((item, index) => (
                <HistoryItem
                  key={index}
                  original={item.original}
                  simplified={item.simplified}
                  createdAt={item.created_at}
                  onClick={() =>
                    onSelect({
                      original: item.original,
                      simplified: item.simplified,
                      createdAt: item.created_at,
                    })
                  }
                />
              ))}
            </div>
          ) : null
        )
      )}
    </div>
  );
};

export default HistoryList;

