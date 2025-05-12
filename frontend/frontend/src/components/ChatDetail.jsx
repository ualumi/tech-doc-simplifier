import React from 'react';

const ChatDetail = ({ original, simplified }) => {
  return (
    <div className="p-4 border rounded-md shadow bg-white">
      <div className="mb-4">
        <h3 className="text-lg font-bold text-gray-700 mb-1">Вы:</h3>
        <div className="bg-gray-100 p-3 rounded text-gray-900">{original}</div>
      </div>
      <div>
        <h3 className="text-lg font-bold text-green-700 mb-1">Упрощённый текст:</h3>
        <div className="bg-green-100 p-3 rounded text-green-900">{simplified}</div>
      </div>
    </div>
  );
};

export default ChatDetail;
