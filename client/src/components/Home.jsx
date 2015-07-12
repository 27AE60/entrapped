import React from 'react';
import { Router, Link } from 'react-router';

class Home extends React.Component {
  render() {
    return (
      <div>
        <h3>Entrapped v0.0</h3>
        <blockquote>Tiny little minefield of our dreams.</blockquote>
        <Link to="/minefield">Play</Link>
      </div>
    )
  }
};

export default Home;
