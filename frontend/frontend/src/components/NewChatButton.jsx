import React from 'react';
import { Plus } from 'lucide-react'; 
import './NewChatButton.css';
const NewChatButton = ({ onNewChat }) => {
  return (
    <div className='NewChatButton'>
      <div className="field cursor-pointer" onClick={onNewChat}>
        <div className='ArrowUpRight'>
          <Plus size={24}/>
        </div>
        <div>
          <p className="mb-2 text-gray-900">Новый чат</p>
        </div>
      </div>
    </div>
    
    
  );
};

export default NewChatButton;
