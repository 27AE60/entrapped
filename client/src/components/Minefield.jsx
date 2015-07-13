import React from 'react';

import Blocks from './Blocks.jsx';
import MinesStore from '../stores/minestore.js';
import Api from '../api/api.js';

class Minefield extends React.Component {
  constructor() {
    super();

    this.state = {
      mines : MinesStore.getMines()
    };

    // connect
    Api.connect();
  }

  render() {
    return (
      <div>
        <h1 className="turn">Your Turn</h1>

        <div className="minefield__wrapper">
          <div className="minefield minefield__enemy">
            <blockquote>Enemy Minefield. <i>Click on a block to place a mine.</i></blockquote>
            <Blocks size={7} mines={this.state.mines} />
            <div className="life-status">
              <span>
                <i className="fa fa-2x fa-heartbeat heart-icons life-status__active"></i>
                <i className="fa fa-2x fa-heartbeat heart-icons life-status__active"></i>
                <i className="fa fa-2x fa-heartbeat heart-icons life-status__active"></i>
                <i className="fa fa-2x fa-heartbeat heart-icons life-status__active"></i>
                <i className="fa fa-2x fa-heartbeat heart-icons"></i>
              </span>
            </div>
          </div>

          <div className="minefield minefield__user">
            <blockquote>Your Minefield.</blockquote>
            <Blocks size={7} mines={this.state.mines} />
            <div className="life-status">
              <span>
                <i className="fa fa-2x fa-heartbeat heart-icons life-status__active"></i>
                <i className="fa fa-2x fa-heartbeat heart-icons life-status__active"></i>
                <i className="fa fa-2x fa-heartbeat heart-icons life-status__active"></i>
                <i className="fa fa-2x fa-heartbeat heart-icons life-status__active"></i>
                <i className="fa fa-2x fa-heartbeat heart-icons"></i>
              </span>
            </div>
          </div>
        </div>
      </div>
    )
  }
};

export default Minefield;
