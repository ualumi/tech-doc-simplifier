/*export const parseModelResponse = (response) => {
  try {
    const parsed = JSON.parse(response?.model_response || '{}');
    return {
      original: parsed?.original?.text || '',
      simplified: parsed?.simplified?.text || '',
    };
  } catch (e) {
    console.error('Ошибка парсинга model_response:', e);
    return {
      original: '',
      simplified: '',
    };
  }
};*/

export function parseModelResponse(modelResponse) {
  try {
    const parsed = JSON.parse(modelResponse);
    return {
      original: parsed.original?.text || '',
      simplified: parsed.simplified?.text || '',
    };
  } catch (e) {
    console.error('❌ Ошибка парсинга model_response:', e, modelResponse);
    return { original: '', simplified: '' };
  }
}

