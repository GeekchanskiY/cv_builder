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
        <li><a href="#">Main page</a></li>
        <li><a href="#">Build CV</a></li>
        <li><a href="#">Create models</a></li>
        <li><a href="#">List all cv's</a></li>
        
        <li><a href="#">Healthcheck</a></li>
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
