import React, { Component } from 'react';
import './App.scss';
import Home from '../Home';
import Search from '../Search';
import { getGIFs, upVote, downVote } from '../../util/api';

class App extends Component {
  
  constructor(props) {
    super(props);
    this.state = {
      searchOpen: false,
      gifs: {},
    };
  }
  
  componentDidMount() {
    this.toggleSearch();
  }
  
  toggleSearch = () => {
    const { searchOpen } = this.state;
    this.setState({ searchOpen: !searchOpen });
  }
  
  handleSearch = async (value, offset = 0) => {
    const { gifs } = this.state;
    const newGifs = await getGIFs(value, offset);
    if (gifs.query === value) {
      const result = { ...gifs };
      result.data = [ ...result.data, ...newGifs.data ];
      result.pagination = newGifs.pagination;
      this.setState({ gifs: result, searchOpen: false });
    } else {
      this.setState({ gifs: newGifs, searchOpen: false });
    } 
  }

  handleUpVote = (id) => {
    upVote(id);
    const { gifs } = this.state;
    const updatedGifs = { ...gifs };
    // should have included a limit in the API so I could get the real ranks
    for (let i = 0; i < updatedGifs.data.length; i++) {
      if (updatedGifs.data[i].id === id) {
        updatedGifs.data[i].rank.up_votes += 1;
        break;
      }
    }
    this.setState({ gifs: updatedGifs });
  };

  handleDownVote = (id) => {
    downVote(id);
    const { gifs } = this.state;
    const updatedGifs = { ...gifs };
    for (let i = 0; i < updatedGifs.data.length; i++) {
      if (updatedGifs.data[i].id === id) {
        updatedGifs.data[i].rank.down_votes += 1;
        break;
      }
    }
    this.setState({ gifs: updatedGifs });
  };
  
  render() {
    const { searchOpen, gifs } = this.state;
    const frameStyle = searchOpen ? "frame move" : "frame";
    return (
      <div className="App">
      {/* GIF Frames for search screen */}
      <div className={frameStyle}></div>
      <div className={frameStyle}></div>
      <div className={frameStyle}></div>
      <div className={frameStyle}></div>
      <div className={frameStyle}>
      <Home 
        page={this.handleSearch} 
        data={gifs} 
        upVote={this.handleUpVote} 
        downVote={this.handleDownVote} 
        toggleSearch={this.toggleSearch}
      />
      </div>
      <Search 
        search={this.handleSearch} 
        isOpen={searchOpen} 
        toggle={this.toggleSearch} 
      />
      </div>
      );
    }
  }
  
  export default App;
  