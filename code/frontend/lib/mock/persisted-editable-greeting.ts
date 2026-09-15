export type GreetingResponse = {
  greeting: string;
};

const STORAGE_KEY = "persisted-editable-greeting";
export const initialGreeting = "Hello, World!";

export function readGreeting(): GreetingResponse {
  if (typeof window === "undefined") {
    return { greeting: initialGreeting };
  }

  return { greeting: window.localStorage.getItem(STORAGE_KEY) || initialGreeting };
}

export function saveGreeting(greeting: string): GreetingResponse {
  const nextGreeting = greeting.trim();

  if (!nextGreeting) {
    return readGreeting();
  }

  window.localStorage.setItem(STORAGE_KEY, nextGreeting);
  return { greeting: nextGreeting };
}
