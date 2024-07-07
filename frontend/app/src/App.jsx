import './App.css';
import logo from './images/logo.png'
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import Home from './components/Home';
import Employee from './components/Employee';

function App() {
  return (
    <BrowserRouter>
      <header>
        <img src={logo} alt="logo" />
      </header>
      <nav>
      <ul>
          <li><a href="#">Home</a></li>
          <li><a href="#">About</a></li>
          <li><a href="#">Pricing</a></li>
          <li><a href="#">Terms of use</a></li>
          <li><a href="#">Contact</a></li>
      </ul>
      </nav>
      <main>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/employee" element={<Employee />} />
        </Routes>
      </main>
      <footer>
        CVbuilder
      </footer>
    </BrowserRouter>
  );
}

export default App;
