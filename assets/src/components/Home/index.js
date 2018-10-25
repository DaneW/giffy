import React, { Component } from 'react';
import PropTypes from 'prop-types';

import Grid from '../Grid';
import styles from './styles.module.scss';

class Home extends Component {
  static propTypes = {
    page: PropTypes.func,
    data: PropTypes.object,
    toggleSearch: PropTypes.func,
    upVote: PropTypes.func,
    downVote: PropTypes.func, // should have used the context API
  }

  page = () => {
    const { page, data: { query, pagination: { offset, count }} } = this.props;
    page(query, offset + count);
  };

  renderGrid = () => {
    if (!this.props.data.data) return null;
    const { 
      data: { data: images }, 
      data: { pagination: { count, total_count } }, 
      upVote, 
      downVote 
    } = this.props;
    const canPage = count < total_count;
    return (
      <>
        <Grid images={images} upVote={upVote} downVote={downVote} />
        <button disabled={!canPage} onClick={this.page}>Load More</button>
      </>
    );
  };

  render() {
    const { toggleSearch } = this.props;
    return (
      <main className={styles.main}>
        <span className={styles.search} onClick={toggleSearch} ><svg><use xlinkHref="#icon-search"/></svg></span>
        {this.renderGrid()}
      </main>
    );
  }
}

export default Home;
