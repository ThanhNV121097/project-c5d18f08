"use client";

import { FormEvent, useEffect, useRef, useState } from "react";
import { readGreeting, saveGreeting } from "../lib/mock/persisted-editable-greeting";
import styles from "./PersistedEditableGreeting.module.css";

type PersistedEditableGreetingProps = {
  initialGreeting: string;
};

export function PersistedEditableGreeting({ initialGreeting }: PersistedEditableGreetingProps) {
  const [greeting, setGreeting] = useState(initialGreeting);
  const [inputValue, setInputValue] = useState(initialGreeting);
  const inputRef = useRef<HTMLInputElement>(null);

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const nextGreeting = inputValue.trim();
    if (!nextGreeting) {
      inputRef.current?.focus();
      return;
    }

    const savedGreeting = saveGreeting(nextGreeting).greeting;
    setGreeting(savedGreeting);
    setInputValue(savedGreeting);
  }

  return (
    <main className={styles.shell}>
      <section aria-labelledby="greeting-heading" className={styles.section}>
        <h1 className={styles.heading} id="greeting-heading">
          {greeting}
        </h1>
        <form className={styles.form} onSubmit={handleSubmit}>
          <label className={styles.label} htmlFor="greeting-input">
            Greeting
          </label>
          <input
            ref={inputRef}
            className={styles.input}
            id="greeting-input"
            name="greeting"
            type="text"
            value={inputValue}
            autoComplete="off"
            required
            onChange={(event) => setInputValue(event.target.value)}
          />
          <button className={styles.button} type="submit">
            Save
          </button>
        </form>
      </section>
    </main>
  );
}
