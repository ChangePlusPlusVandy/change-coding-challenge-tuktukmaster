import React, { Component } from 'react';
import './App.css';
import Game from './components/Game/Game';
//import './components/GuessBtn.js';
//import GuessBtn from './components/GuessBtn.js';

class App extends Component {

  construct(props){
    
    
  }

  async fetchTweets(){
    
  }

  render(){
    //this.fetchTweets();
    return (
      <div className="App">
        <header className="App-header">
          
          <Game/>
        </header>
      </div>
    );
  }
  
}

export default App;
