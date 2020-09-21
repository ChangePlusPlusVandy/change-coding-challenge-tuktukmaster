import React, { Component } from 'react';
import './App.css';
import './GuessBtn.js';
import GuessBtn from './GuessBtn.js';

class App extends Component {

  construct(props){
    
    
  }

  async fetchTweets(){
    console.log("hello")
    const url = "https://api.twitter.com/1.1/statuses/user_timeline.json?screen_name=elonmusk&count=1";
    fetch(url, {
      method: "GET",
      headers: {
        "Bearer": "AAAAAAAAAAAAAAAAAAAAADlOHwEAAAAAr9Ncm6oJ8hhQ5U18GafR9ORycH4%3DDfgEkAO74v0YRMLdEbm5LkmhADayzqN26PkUTcf70A5A1v5D56"
      }
    }).then(response => response.json())
    .then(data => console.log(data));  
  }

  render(){
    //this.fetchTweets();
    return (
      <div className="App">
        <header className="App-header">
          <h1>Do You Know Your Meme-Lords?</h1>
          <p>
            Lorem ipsum dolar sit amet
          </p>
          <div className="Button-container">
            <GuessBtn
              name = "Elon Musk"
              className = "button success"
            />
            <GuessBtn
              name = "Kanye West"
              className = "button error"
            />
          </div>
          

        </header>
      </div>
    );
  }
  
}

export default App;
