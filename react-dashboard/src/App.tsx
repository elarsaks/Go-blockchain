import "./App.css";
import { Button } from "./components/Button";
import AppHeader from "./components/AppHeader";

function App() {
  return (
    <div className="App">
      <AppHeader title="Go Blockchain" />
      <Button>Normal</Button>
      <Button $primary>Primary</Button>
    </div>
  );
}

export default App;
