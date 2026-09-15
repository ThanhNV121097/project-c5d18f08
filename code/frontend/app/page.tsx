import { PersistedEditableGreeting } from "../components/PersistedEditableGreeting";

type GreetingResponse = {
  greeting: string;
};

export default async function Page() {
  const apiOrigin = process.env.API_ORIGIN ?? "http://backend:8080";
  const response = await fetch(`${apiOrigin}/v1/greeting`, { cache: "no-store" });
  const { greeting } = (await response.json()) as GreetingResponse;

  return <PersistedEditableGreeting initialGreeting={greeting} />;
}
