import React from 'react';

const NewChatButton = ({ onNewChat }) => {
  return (
    <button
      className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 transition"
      onClick={onNewChat}
    >
      Новый чат
    </button>
  );
};

export default NewChatButton;
