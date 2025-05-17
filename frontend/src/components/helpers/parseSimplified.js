export const parseSimplifiedText = (simplified) => {
  try {
    console.log('RAW simplified:', simplified); // 👈 добавим лог

    if (typeof simplified === 'object' && simplified !== null) {
      return simplified.text || '';
    }

    if (typeof simplified === 'string') {
      const trimmed = simplified.trim();
      const looksLikeJson = trimmed.startsWith('{') && trimmed.endsWith('}');

      if (looksLikeJson) {
        const parsed = JSON.parse(trimmed);
        console.log('Parsed simplified:', parsed); // 👈 лог после парсинга
        return parsed?.text || '';
      }

      return simplified; // это уже финальный текст
    }

    return '';
  } catch (e) {
    console.error('Ошибка парсинга simplified:', e);
    return '';
  }
};
