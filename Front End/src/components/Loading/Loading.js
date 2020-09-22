import React, { Component } from 'react';
import LoadAnimation from './../LoadAnimation/LoadAnimation'

class Loading extends Component {

  constructor(props){
    super(props)
    this.state = {
      twitterHandleOne: this.props.twitterHandleOne,
      twitterHandleTwo: this.props.twitterHandleTwo,
      setTwitterName: this.props.setTwitterNames,
      setTweets: this.props.setTweets
    }
  }

  

  render(){
    return (
      <div>
        <LoadAnimation/>
      </div>
    );
  }
  
}

export default Loading;
