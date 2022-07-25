import './App.css';
import MainView from './components/mainView';
import Newsletter from './components/newsletter/newsletter';
import {
    BrowserRouter,
    Routes,
    Route,
  } from "react-router-dom";

function App() {
  return (
    <BrowserRouter>
        <Routes>
            <Route path="/" element={<MainView/>} />            
            <Route path="/newsletter/:id" element={<Newsletter/>} />            
        </Routes>
    </BrowserRouter>    
  );
}

export default App;
