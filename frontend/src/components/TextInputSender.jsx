





import React, { useState } from 'react';
import axios from 'axios';
import { parseModelResponse } from './helpers/parseModelResponse.js';
import mammoth from 'mammoth';
import * as pdfjsLib from 'pdfjs-dist/legacy/build/pdf';
import pdfWorker from 'pdfjs-dist/build/pdf.worker?url';
import './TextInputSender.css';
import { Paperclip, ArrowUp, Square} from 'lucide-react'; 

pdfjsLib.GlobalWorkerOptions.workerSrc = pdfWorker;




const TextInputSender = ({ token, onTriggerLogin, onResponse, onMessageAdd }) => {
  const [text, setText] = useState('');
  const [loading, setLoading] = useState(false);

  const handleInputClick = () => {
    if (!token) {
      console.log('Клик по input без токена');
      onTriggerLogin();
    }
  };

  const handleSend = async () => {
    if (!token || !text.trim()) return;

    const userInput = text;
    setText('');
    onResponse({ original: userInput }); // Предварительный отклик

    try {
      setLoading(true);

      const res = await axios.post(
        'http://87.228.89.190/main/simplify',
        { text: userInput },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
          validateStatus: () => true,
        }
      );

      if (res.status === 401) {
        onTriggerLogin();
        return;
      }

      if (res.status >= 200 && res.status < 300) {
        const { original, simplified } = parseModelResponse(res.data.model_response);

        onResponse({
          original: original || userInput,
          simplified: simplified || 'Не удалось извлечь упрощённый текст',
        });

        onMessageAdd?.({
          original: original || userInput,
          simplified: simplified || 'Не удалось извлечь упрощённый текст',
        });

      } else {
        onResponse({
          original: userInput,
          simplified: 'Ошибка при отправке текста',
        });
      }
    } catch (error) {
      console.error('❌ Ошибка отправки:', error);
      onResponse({
        original: userInput,
        simplified: 'Ошибка при отправке текста',
      });
    } finally {
      setLoading(false);
    }
  };

  const handleFileUpload = async (e) => {
    const file = e.target.files[0];
    if (!file) return;

    if (!token) {
      onTriggerLogin?.();
      return;
    }

    const ext = file.name.split('.').pop().toLowerCase();

    if (ext === 'txt') {
      const reader = new FileReader();
      reader.onload = (event) => {
        const fileText = event.target.result;
        setText(fileText);
        setTimeout(() => handleSend(), 100);
      };
      reader.readAsText(file);
    } else if (ext === 'pdf') {
      const reader = new FileReader();
      reader.onload = async (event) => {
        try {
          const typedArray = new Uint8Array(event.target.result);
          const pdf = await pdfjsLib.getDocument({ data: typedArray }).promise;
          let textContent = '';

          for (let i = 1; i <= pdf.numPages; i++) {
            const page = await pdf.getPage(i);
            const content = await page.getTextContent();
            const pageText = content.items.map(item => item.str).join(' ');
            textContent += pageText + '\n';
          }

          setText(textContent);
          setTimeout(() => handleSend(), 100);
        } catch (err) {
          console.error('Ошибка чтения PDF:', err);
          alert('Не удалось прочитать PDF-файл');
        }
      };
      reader.readAsArrayBuffer(file);
    } else if (ext === 'docx') {
      const reader = new FileReader();
      reader.onload = async (event) => {
        const arrayBuffer = event.target.result;
        const result = await mammoth.extractRawText({ arrayBuffer });
        setText(result.value);
        setTimeout(() => handleSend(), 100);
      };
      reader.readAsArrayBuffer(file);
    } else {
      alert('Поддерживаются только файлы .txt, .pdf и .docx');
    }
  };

  return (
    <div className="mt">
      <label className='paperclip'>
        <Paperclip size={24} />
        <input
          
          type="file"
          accept=".txt, .pdf, .docx, application/msword, application/vnd.openxmlformats-officedocument.wordprocessingml.document, application/pdf"
          onChange={handleFileUpload}
          className="file-input"
        />
      </label>
      
      <div className='InputTextTypeText'>
        <input
          type="text"
          className="InputText"
          placeholder="Введите текст..."
          value={text}
          onClick={handleInputClick}
          readOnly={!token}
          onChange={(e) => setText(e.target.value)}
        />
      </div>
      
      <div className='buttoning'>
        <button
          onClick={handleSend}
          disabled={!token || loading}
          className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 disabled:opacity-50"
        >
          {loading ? <Square size={24}/> : <ArrowUp size={24} />}
        </button>
      </div>
      
    </div>
  );
};

export default TextInputSender;

