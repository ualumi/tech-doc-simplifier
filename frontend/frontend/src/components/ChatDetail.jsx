import React from 'react';
import './ChatDetail.css';
const ChatDetail = ({ original, simplified }) => {
    console.log('👁 simplified в ChatDetail:', simplified);
  return (
    <div className="p-4 border rounded-md shadow bg-white">
      <div className="You">
        <h3 className="text-lg font-bold text-gray-700 mb-1">Вы:</h3>
        <div className="You_response">{original}</div>
      </div>
      <div className="model">
        <h3 className="text-lg font-bold text-green-700 mb-1">Упрощённый текст:</h3>
        <div className="You_response">{simplified}</div>
      </div>
    </div>
  );
};

export default ChatDetail;
