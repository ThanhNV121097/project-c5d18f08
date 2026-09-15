export type GreetingResponse = {
  greeting: string;
};

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "/api";

export async function saveGreeting(greeting: string): Promise<GreetingResponse> {
  const response = await fetch(`${apiBase}/v1/greeting`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ greeting }),
  });

  if (!response.ok) {
    throw new Error("Failed to save greeting");
  }

  return response.json();
}
