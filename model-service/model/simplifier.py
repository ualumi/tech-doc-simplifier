



from transformers import AutoModelForSeq2SeqLM, AutoTokenizer
import torch

class CFG:
    MODEL_PATH = "disemenova/proekt_sum"  
    MAX_SOURCE_LENGTH = 800
    MAX_TARGET_LENGTH = 200
    DEVICE = "cuda" if torch.cuda.is_available() else "cpu"

tokenizer = AutoTokenizer.from_pretrained(CFG.MODEL_PATH)
model = AutoModelForSeq2SeqLM.from_pretrained(CFG.MODEL_PATH)
model.to(CFG.DEVICE)
model.eval()

def generate_summary(text: str) -> str:
    prompt = (
        "Сделай структурированную выжимку: цель, мероприятия, бюджеты, нормативные акты. "
        f"Текст: {text[:2000]}"
    )

    if len(text)<= 100:
        return "Слишком маленький текст("
    
    if "программ" not in text:
        return "Данный текст не содержит документации по теме Стратегическое планирование и отчётность"
    if len(text)>2000:
        return "Слишком мольшой текст"

    inputs = tokenizer(
        prompt,
        max_length=CFG.MAX_SOURCE_LENGTH,
        truncation=True,
        padding="max_length",
        return_tensors="pt"
    ).to(CFG.DEVICE)

    with torch.no_grad():
        summary_ids = model.generate(
            input_ids=inputs["input_ids"],
            attention_mask=inputs["attention_mask"],
            max_length=CFG.MAX_TARGET_LENGTH,
            num_beams=5,
            repetition_penalty=3.0,
            length_penalty=0.7,
            no_repeat_ngram_size=3,
            early_stopping=False
        )

    summary = tokenizer.decode(summary_ids[0], skip_special_tokens=True)

    summary = summary.replace("программа", "Программа").replace("приложение", "Приложение")
    return summary

'''example_text = """Государственная программа Свердловской области «Развитие культуры и искусства» на 2024–2030 годы утверждена постановлением Правительства Свердловской области от 01.02.2024 № 22-п. Программа направлена на повышение доступности и качества культурных услуг, сохранение культурного наследия и развитие творческих индустрий.

В программу включены мероприятия по реставрации памятников архитектуры и музеев, развитию сетей культурных центров, поддержке творческих коллективов и фестивалей. Запланировано внедрение цифровых технологий для популяризации культурных проектов, развитие программ дополнительного образования в области искусств, а также создание условий для развития креативных индустрий и творческого бизнеса.

Финансирование составляет 13 млрд рублей: 7 млрд федеральных, 4 млрд региональных и 2 млрд внебюджетных. Ключевые показатели — рост посещаемости культурных мероприятий на 30%, увеличение числа творческих проектов на 25%, а также повышение удовлетворённости населения качеством культурных услуг."""  # можешь вставить тот же текст

print(generate_summary(example_text))'''