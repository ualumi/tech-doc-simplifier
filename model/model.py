


import pandas as pd
import torch
from sklearn.model_selection import train_test_split
from transformers import (
    AutoTokenizer,
    AutoModelForSeq2SeqLM,
    Seq2SeqTrainer,
    Seq2SeqTrainingArguments,
    DataCollatorForSeq2Seq,
)
import evaluate
import numpy as np
from datasets import Dataset
from peft import LoraConfig, get_peft_model

class CFG:
    MODEL_NAME = "IlyaGusev/rut5_base_sum_gazeta"
    MAX_SOURCE_LENGTH = 800
    MAX_TARGET_LENGTH = 200
    BATCH_SIZE = 4
    GRAD_ACCUM_STEPS = 4
    LEARNING_RATE = 1e-4
    EPOCHS = 12
    DEVICE = "cuda" if torch.cuda.is_available() else "cpu"

tokenizer = AutoTokenizer.from_pretrained(CFG.MODEL_NAME)
model = AutoModelForSeq2SeqLM.from_pretrained(CFG.MODEL_NAME)

from peft import prepare_model_for_kbit_training

model = AutoModelForSeq2SeqLM.from_pretrained(CFG.MODEL_NAME)
model.gradient_checkpointing_enable()

model = prepare_model_for_kbit_training(model)

lora_config = LoraConfig(
    r=8,
    lora_alpha=32,
    target_modules=["q", "v"],
    lora_dropout=0.05,
    bias="none",
    task_type="SEQ_2_SEQ_LM",
    inference_mode=False
)
model = get_peft_model(model, lora_config)

model.to(CFG.DEVICE)
model.print_trainable_parameters()

df = pd.read_excel("/kaggle/input/data-proekt/data_proekt.xlsx")

def clean_text(text):
    if pd.isnull(text): return ""
    text = str(text)
    text = ' '.join(text.split())
    return text.replace("\\n", " ").replace("\xa0", " ").strip()

df = df.dropna(subset=["text", "summary"])
df["text"] = df["text"].apply(clean_text)
df["summary"] = df["summary"].apply(clean_text)
df = df[(df["text"].str.len() > 50) & (df["summary"].str.len() > 20)]

train_df, val_df = train_test_split(df, test_size=0.2, random_state=42)

def preprocess_data(examples):
    examples["text"] = [
        f"Сделай структурированную выжимку: цель, мероприятия, бюджеты, нормативные акты: {text[:2000]}"
        for text in examples["text"]
    ]

    inputs = tokenizer(
        examples["text"],
        max_length=CFG.MAX_SOURCE_LENGTH,
        truncation=True,
        padding="max_length"
    )

    with tokenizer.as_target_tokenizer():
        labels = tokenizer(
            examples["summary"],
            max_length=CFG.MAX_TARGET_LENGTH,
            truncation=True,
            padding="max_length"
        )["input_ids"]

    labels = [
        [lab if lab != tokenizer.pad_token_id else -100 for lab in seq]
        for seq in labels
    ]
    inputs["labels"] = labels
    return inputs

train_dataset = Dataset.from_pandas(train_df[["text", "summary"]])
val_dataset = Dataset.from_pandas(val_df[["text", "summary"]])
train_dataset = train_dataset.map(preprocess_data, batched=True)
val_dataset = val_dataset.map(preprocess_data, batched=True)

rouge = evaluate.load("rouge")

def compute_metrics(eval_preds):
    preds, labels = eval_preds
    if isinstance(preds, tuple):
        preds = preds[0]

    preds = np.nan_to_num(preds, nan=tokenizer.pad_token_id)
    labels = np.where(labels != -100, labels, tokenizer.pad_token_id)

    decoded_preds = tokenizer.batch_decode(preds, skip_special_tokens=True)
    decoded_labels = tokenizer.batch_decode(labels, skip_special_tokens=True)

    decoded_preds = [p.strip() for p in decoded_preds]
    decoded_labels = [l.strip() for l in decoded_labels]

    result = rouge.compute(
        predictions=decoded_preds,
        references=decoded_labels,
        use_stemmer=True
    )
    return {k: round(v * 100, 2) for k, v in result.items()}
training_args = Seq2SeqTrainingArguments(
    output_dir="./results",
    per_device_train_batch_size=CFG.BATCH_SIZE,
    per_device_eval_batch_size=1,
    gradient_accumulation_steps=CFG.GRAD_ACCUM_STEPS,
    learning_rate=CFG.LEARNING_RATE,
    num_train_epochs=CFG.EPOCHS,
    evaluation_strategy="epoch",
    save_strategy="epoch",
    save_total_limit=2,
    logging_strategy="steps",
    logging_steps=10,
    predict_with_generate=True,
    generation_max_length=CFG.MAX_TARGET_LENGTH,
    fp16=False,                      
    gradient_checkpointing=True,   
    max_grad_norm=1.0,
    report_to="none",
    optim="adamw_torch"
)

data_collator = DataCollatorForSeq2Seq(tokenizer, model=model)

trainer = Seq2SeqTrainer(
    model=model,
    args=training_args,
    train_dataset=train_dataset,
    eval_dataset=val_dataset,
    tokenizer=tokenizer,
    data_collator=data_collator,
    compute_metrics=compute_metrics
)

trainer.train()