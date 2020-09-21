import React, { Component } from 'react';
import './TwitterHandleInput.scss';

class TwitterHandleInput extends Component {

  construct(props){
    console.log(props);
  }

  render(){
    return (
      <div>
        <p>
          <span className="twitterHandleInput">
            <input className="twitterHandleInput" type="text" placeholder={this.props.InputName}/>
            <span></span>	
          </span>
        </p>
        
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

export default TwitterHandleInput;
