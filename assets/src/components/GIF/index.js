import React, { Component } from 'react';
import PropTypes from 'prop-types';

import styles from './styles.module.scss';

class GIF extends Component {
  static propTypes = {
    image: PropTypes.object,
    upVote: PropTypes.func,
    downVote: PropTypes.func,
  }

  constructor(props) {
    super(props);
    this.state = {
      play: false
    }
  }

  render() {
    const { 
      image: { id, caption, images: { original, original_still }}, 
      upVote, 
      downVote,
    } = this.props;
    const { play } = this.state;
    return (
      <div 
        className={styles.image}
        onMouseEnter={() => this.setState({ play: true })}
        onMouseLeave={() => this.setState({ play: false })} 
      >
        <img 
          alt={caption}
          src={play ? original.url : original_still.url}
        />
        <span>
          <button title="Up Vote" onClick={() => upVote(id)} ><svg><use xlinkHref="#icon-plus" /></svg></button>
          <button title="Down Vote" onClick={() => downVote(id)} ><svg><use xlinkHref="#icon-minus" /></svg></button>
        </span>
      </div>
    );
  }
}

export default GIF;
