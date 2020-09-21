import React, { Component } from 'react';
//import './Game.css';
//import './StartBtn'
import StartBtn from './StartBtn';
import TwitterHandleInput from './TwitterHandleInput';

class Game extends Component {

  construct(props){
    console.log(props);
  }

  render(){
    return (
      <div>
        <h1>Change++ Coding Challenge</h1>
        <TwitterHandleInput
          InputName="Twitter Handle One"
        />
        <TwitterHandleInput
          InputName="Twitter Handle Two"
        />
        <StartBtn/>
      </div>
    );
  }
  
}
/*
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
*/

export default Game;
