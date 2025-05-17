import React from 'react';
import './HistoryItem.css';
import { ArrowUpRight} from 'lucide-react'; 

const truncateText = (text, wordLimit = 3) => {
  if (!text) return '';
  const words = text.split(' ');
  return words.length <= wordLimit
    ? text
    : words.slice(0, wordLimit).join(' ') + '...';
};

const HistoryItem = ({ original, simplified, createdAt, onClick }) => {
  return (
    <div className="field cursor-pointer" onClick={onClick}>
      <div className='ArrowUpRight'>
        <ArrowUpRight size={24}/>
      </div>
      <div>
        <p className="mb-2 text-gray-900">{truncateText(original)}</p>
        <p className="mb-2 text-green-800">{truncateText(simplified)}</p>
      </div>
    </div>
  );
};

export default HistoryItem;
