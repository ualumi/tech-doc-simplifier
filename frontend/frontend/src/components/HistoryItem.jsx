import React from 'react';
import './HistoryItem.css';

const truncateText = (text, wordLimit = 3) => {
    if (!text) return '';
    const words = text.split(' ');
    return words.length <= wordLimit
      ? text
      : words.slice(0, wordLimit).join(' ') + '...';
  };

const HistoryItem = ({ original, simplified, createdAt }) => {
  return (
    <div className="field">
      {/*<p className="text-gray-700 font-semibold">Оригинал:</p>*/}
      <p className="mb-2 text-gray-900">{truncateText(original)}</p>
      <p className="mb-2 text-green-800">{truncateText(simplified)}</p>
    </div>
  );
};

export default HistoryItem;
