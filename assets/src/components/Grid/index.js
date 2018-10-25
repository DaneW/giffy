import React, { Component } from 'react';
import PropTypes from 'prop-types';

import GIF from '../GIF';
import './styles.scss';

class Home extends Component {
  static propTypes = {
    images: PropTypes.array,
    upVote: PropTypes.func,
    downVote: PropTypes.func,
  }

  render() {
    const { images, downVote, upVote } = this.props;
    return (
      <main className="grid">
        {images.sort((a, b) => {
          const aScore = (a.rank.up_votes - a.rank.down_votes);
          const bScore = (b.rank.up_votes - b.rank.down_votes);
          if (aScore > bScore) return -1;
          if (aScore < bScore) return 1;
          if (a.import_datetime < b.import_datetime) return -1;
          if (a.import_datetime > b.import_datetime) return 1;
          return 0;
        }).map(i => <GIF key={i.id} upVote={upVote} downVote={downVote} image={i} />)}
      </main>
    );
  }
}

export default Home;