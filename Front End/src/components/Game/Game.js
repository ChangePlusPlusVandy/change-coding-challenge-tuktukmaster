import React, { Component } from 'react';
import StartBtn from './../StartBtn/StartBtn';
import TwitterHandleInput from './../TwitterHandleInput/TwitterHandleInput';
import Menu from './../Menu/Menu'

class Game extends Component {

  constructor(props){
    super(props)
    this.state = {
      twitterHandleOne: "",
      nameOne: "",
      tweetsOne: {},
      twitterHandleTwo: "",
      nameTwo: "",
      tweetsTwo: {},
      correctCount: 0,
      totalCount: 0
    }
    this.startGame = this.startGame.bind(this);
  }

  setHandleOne = (childData) => {
    this.setState({twitterHandleOne: childData})
    console.log(childData)
  }

  setHandleTwo = (childData) => {
    this.setState({twitterHandleTwo: childData})
    console.log(childData)
  }

  startGame = (handleOne, handleTwo) => {
    console.log(handleOne,handleTwo);
  }

  render(){
    return (
      <div>
        <Menu
          startGameCallback = {this.startGame}
        />
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
