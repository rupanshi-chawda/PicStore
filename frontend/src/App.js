import React from 'react';
import Navbar from './components/Navbar';
import {BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import './App.css';
import Hero from './components/Hero';

function App() {
  return (
    <>
      <Router>
        <Navbar />
        <Hero />
        <Switch>
          <Route path='/' exact/>
        </Switch>
      </Router>
    </>
  );
}

export default App;
