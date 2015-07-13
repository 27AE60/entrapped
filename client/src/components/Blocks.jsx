import React from 'react/addons';

import Api from '../api/api.js';
/**
 * Returns a random integer between min (inclusive) and max (inclusive)
 * Using Math.round() will give you a non-uniform distribution!
 */
function getRandomInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

class Block extends React.Component {
  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this);
  }

  render() {
    var cx = React.addons.classSet;
    var classes = cx({
      'block': true,
      'block__visited': this.props.status === -1,
      'block__dead': this.props.status === -2
    });

    return(
      <td className={classes} onClick={this.handleClick}></td>
    )
  }

  handleClick() {
    var msg = "data:open:[idx=" + this.props.index + "]:[help=9]";
    Api.send(msg);
  }
};
Block.propTypes = { status: React.PropTypes.number, index: React.PropTypes.number };
Block.defaultProps = { status: 0, index: 0 };

class Blocks extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {

    var nodes = [];
    var size = this.props.size;

    /* construct blocks */
    for (var i = 0; i < size; i++) {
      var _blocks = [];

      for (var j = 0; j < size; j++) {
        var index = i * size + j;
        _blocks.push(<Block key={'block-' + index} status={this.props.mines[index]} index={index} />);
      }

      nodes.push((<tr key={'blocks-' + i}>{_blocks}</tr>));
    }

    return (
      <table className="blocks">
        <tbody>
          {nodes}
        </tbody>
      </table>
    );
  }
};
Blocks.propTypes = { size: React.PropTypes.number, mines: React.PropTypes.array };
Blocks.defaultProps = { size: 7, mines: [] }

export default Blocks;
