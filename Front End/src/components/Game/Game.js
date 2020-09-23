import React, { Component } from 'react';
import Loading from './../Loading/Loading'
import Menu from './../Menu/Menu'
import GamePlay from './../GamePlay/GamePlay'

class Game extends Component {

  constructor(props){
    super(props)
    this.state = {
      gameState: "start",
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
    this.setTwitterData = this.setTwitterData.bind(this);
  }

  setHandleOne = (childData) => {
    this.setState({twitterHandleOne: childData})
    console.log(childData)
  }

  setHandleTwo = (childData) => {
    this.setState({twitterHandleTwo: childData})
    console.log(childData)
  }

  setTwitterData = (nameOne, tweetsOne, nameTwo, tweetsTwo) => {
    this.setState({
      gameState: "gameplay",
      nameOne: nameOne,
      tweetsOne: tweetsOne,
      nameTwo: nameTwo,
      tweetsTwo: tweetsTwo
    })
  }

  startGame = (handleOne, handleTwo) => {
    this.setState({
      gameState: "loading",
      twitterHandleOne: handleOne,
      twitterHandleTwo: handleTwo
    })
  }

  render(){
      if(this.state.gameState === "start"){
        return (
          <div>
            <Menu
              startGameCallback = {this.startGame}
            />
          </div>
        )
      } else if(this.state.gameState === "loading"){
        return (
          <div>
        <Loading
          twitterHandleOne = {this.state.twitterHandleOne}
          twitterHandleTwo = {this.state.twitterHandleTwo}
          setTwitterData = {this.setTwitterData}
        />
        
      </div>
        )
      } else if(this.state.gameState === "gameplay"){
        const chooseHandle = (Math.floor(Math.random() * 2) === 0)
        let correct;
        let index;
        let tweet;
        if(chooseHandle){
          correct = "0"
          index = Math.floor(Math.random() * this.state.tweetsOne.length)
          tweet = this.state.tweetsOne[index]
        }else{
          correct = "1"
          index = Math.floor(Math.random() * this.state.tweetsTwo.length)
          tweet = this.state.tweetsTwo[index]
        }
        return (
          <div>
            <GamePlay
              nameOne = {this.state.nameOne}
              nameTwo = {this.state.nameTwo}
              correct = {correct}
              tweet = {tweet}
            />
          </div>
        )
      }
    
  }
  
}

export default Game;
