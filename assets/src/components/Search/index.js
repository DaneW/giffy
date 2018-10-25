import React, { Component } from 'react';
import PropTypes from 'prop-types';

class Search extends Component {
  static propTypes = {
    isOpen: PropTypes.bool,
    toggle: PropTypes.func,
    search: PropTypes.func,
  }

  constructor(props) {
    super(props);
    this.state = {
      value: '',
    };
  }

  componentDidUpdate() {
    const { isOpen } = this.props;
    if (isOpen) this.input.focus();
  }

  handleChange = (e) => this.setState({value: e.target.value});

  handleSubmit = e => {
    const { search } = this.props;
    const { value } = this.state;
    e.preventDefault();
    search(value);
  }

  render() {
    const { isOpen, toggle } = this.props;
    const searchStyle = isOpen ? "search" : "search close";
    return (
      <div className={searchStyle}>
        <form onSubmit={this.handleSubmit}>
          <input 
            onChange={this.handleChange}
            ref={input => (this.input = input)}
            name="search" 
            type="search" 
            placeholder="Giffy..." 
            autoComplete="off" 
            autoCorrect="off"
            autoCapitalize="off" 
            spellCheck="false" 
          />
          <span>Hit enter to search</span>
        </form>
      </div>
    );
  }
}



export default Search;
