import React from 'react';
import { Router, Link } from 'react-router';

class Home extends React.Component {
  render() {
    return (
      <div className="home">
        <h1 className="home__title">Entrapped v0.0</h1>
        <blockquote>Tiny little minefield of our dreams.</blockquote>
        <div className="nickname">
          <input className="nickname__input" placeholder="nickname"/>
        </div>
        <Link to="/minefield">Play</Link>
      </div>
    )
  }
};

export default Home;
