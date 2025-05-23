import React, { useEffect, useState } from 'react';
import './ResponseViewer.css'; 

const ResponseViewer = ({ response }) => {
  const [visible, setVisible] = useState(false);

  useEffect(() => {
    if (response) {
      setVisible(true);
    }
  }, [response]);

  if (!response) return null;

  return (
    <div className={`response-container ${visible ? 'fade-in' : ''}`}>
      {response.error ? (
        <span className="text-red-500">{response.error}</span>
      ) : (
        <pre className="Text">{JSON.stringify(response, null, 2)}</pre>
      )}
    </div>
  );
};

export default ResponseViewer;
