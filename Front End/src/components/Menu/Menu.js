import React, { Component } from 'react';
import StartBtn from './../StartBtn/StartBtn';
import TwitterHandleInput from './../TwitterHandleInput/TwitterHandleInput';

class Menu extends Component {

  constructor(props){
    super(props)
    this.state = {
      twitterHandleOne: "",
      twitterHandleTwo: "",
      startGame: this.props.startGameCallback
    }
  }

  setHandleOne = (childData) => {
    this.setState({twitterHandleOne: childData})
  }

  setHandleTwo = (childData) => {
    this.setState({twitterHandleTwo: childData})
  }

  startGame = () => {
    this.state.startGame(this.state.twitterHandleOne,this.state.twitterHandleTwo)
  }

  render(){
    return (
      <div>
        <h1>Change++ Coding Challenge</h1>
        <TwitterHandleInput
          InputName="Twitter Handle One"
          parentCallback = {this.setHandleOne}
        />
        <TwitterHandleInput
          InputName="Twitter Handle Two"
          parentCallback = {this.setHandleTwo}
        />
        <StartBtn
          parentCallback = {this.startGame}
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

export default Menu;
