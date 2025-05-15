import React from 'react';

const ChatHistory = ({ messages }) => {
  if (!messages.length) {
    return <p className="text-gray-400"></p>;
  }

  return (
    <div className="space-y-6">
      {messages.map((msg, index) => (
        <div key={index} className="bg-gray-50 p-4 rounded shadow">
          <p className="text-sm text-gray-500 mb-1">Вы:</p>
          <div className="bg-white border p-3 rounded mb-3">{msg.original}</div>

          <p className="text-sm text-gray-500 mb-1">Ответ:</p>
          <div className="bg-green-50 border border-green-200 p-3 rounded">
            {msg.simplified}
          </div>
        </div>
      ))}
    </div>
  );
};

export default ChatHistory;
