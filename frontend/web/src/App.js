import React from 'react';
import './App.css';
import {Container} from "react-bootstrap";
import Leaderboard from "../src/components/leaderboard.js";
import Match from "./components/match.js";
import 'bootstrap/dist/css/bootstrap.min.css';


function App() {
  return (
    <div className="App">
        
        <Match />
        <div className="space"/>
        <p>Leaderboard</p>
        <Leaderboard />
    </div>
  );
}

export default App;
