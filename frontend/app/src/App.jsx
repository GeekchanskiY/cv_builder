import './App.css';
import logo from './images/logo.png'
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';

function App() {
  return (
    <BrowserRouter>
      <header>
        <img src={logo} alt="logo" />
        <Link to={"/1"}>Emoloyees</Link>
        <Link to={"/1"}>Domains</Link>
        <Link to={"/1"}>Companies</Link>
        <Link to={"/1"}>Vacancies</Link>
        <Link to={"/1"}>CVs</Link>
        <Link to={"/1"}>Skills</Link>
        <Link to={"/1"}>Responsibilities</Link>
      </header>
      <Routes>
        <Route path="/" element={<h1>asd</h1>} />
        <Route path="/1" element={<h1>asd1</h1>} />
      </Routes>
      <footer>
        CVbuilder
      </footer>
    </BrowserRouter>
  );
}

export default App;
