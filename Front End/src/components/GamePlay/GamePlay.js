import React, { Component } from 'react';
import GuessBtn from './../GuessBtn/GuessBtn'
import './GamePlay.css'

class GamePlay extends Component {

  constructor(props){
    super(props)

    this.state = {
      nameOne: this.props.nameOne,
      nameTwo: this.props.nameTwo,
      correct: this.props.correct,
      tweet: this.props.tweet
    }

    this.handleClick = this.handleClick.bind(this);
    
  }

  handleClick() {
    this.state.startCallback();
  }

  render(){
    return (
      <div>
        <h1>Which Twitter User Said This?</h1>
            <p className="tweetDisplay">
              {this.state.tweet.Text}
            </p>
            <div className="Button-container">
              <GuessBtn
                name = {this.state.nameOne}
                className = {this.state.correct === "0" ? "button success" : "button error"}
              />
              <GuessBtn
                name = {this.state.nameTwo}
                className = {this.state.correct === "1" ? "button success" : "button error"}
              />
            </div>
      </div>
    );
  }
  
}

export default GamePlay;


