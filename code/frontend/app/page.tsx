import { PersistedEditableGreeting } from "../components/PersistedEditableGreeting";
import { readGreeting } from "../lib/mock/persisted-editable-greeting";

export default function Page() {
  const { greeting } = readGreeting();

  return <PersistedEditableGreeting initialGreeting={greeting} />;
}
