import { Route, Routes } from 'react-router-dom';
import './App.css';
import OpeningHandSim from './components/OpeningHandSim';

function App() {
  return (
    <div>
      <Routes>
        <Route path="/" element={<OpeningHandSim />} />
      </Routes>
    </div>
  );
}

export default App;
