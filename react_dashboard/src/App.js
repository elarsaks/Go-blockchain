import logo from "./logo.svg";
import "./App.css";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>

        <div class="ui animated button" tabindex="0">
          <div class="visible content">Next</div>
          <div class="hidden content">
            <i class="right arrow icon"></i>
          </div>
        </div>
      </header>
    </div>
  );
}

export default App;
