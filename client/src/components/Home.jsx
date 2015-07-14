import React from 'react';
import { Router, Link } from 'react-router';

import MinesStore from '../stores/minestore.js';

class Home extends React.Component {
  handleChange(e) {
    MinesStore.addUsername(e.target.value);
  }

  verifyName() {
    if (MinesStore.getUsername() == "") {
	alert('Please enter the nickname');
	return false;
    }
    return true;
  }

  render() {
    return (
      <div className="home">
        <h1 className="home__title">Entrapped v0.0</h1>
        <blockquote>Tiny little minefield of our dreams.</blockquote>
        <div className="nickname">
          <input className="nickname__input" placeholder="nickname" onChange={this.handleChange}/>
        </div>
        <Link to="/minefield" onClick={this.verifyName}>Play</Link>
      </div>
    )
  }
};

export default Home;
