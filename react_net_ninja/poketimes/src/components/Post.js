import React from 'react';
import axios from 'axios';
import Pokeball from '../pokeball.png';

class Post extends React.Component {
  state = {
    post: null,
    msg: ''
  }

  componentDidMount() {
    const postId = this.props.match.params.post_id;
    axios.get('https://jsonplaceholder.typicode.com/posts' + '/' + postId)
      .then(res => {
        this.setState({ post: res.data });
      }).catch(() => {
        this.setState({ msg: 'There is no post with id:' + postId });
      });
  }

  render() {
    if (!this.state.msg && this.state.post) {
      let { post } = this.state;

      return (
        <div className="container">
          <div className="post card">
            <div className="card-content">
              <img src={ Pokeball } alt="A Pokeball Image"/>
              <span className="card-title red-text">{ post.title }</span>
              <p className="card-body black-text">{ post.body }</p>
            </div>
          </div>
        </div>
      );
    } else if (!this.state.post) {
      return (
        <div className="container">
          <div className="post card">
            <div className="card-content">
              <span className="card-title">Loading...</span>
            </div>
          </div>
        </div>
      );
    } else {
      return (
        <div className="container">
          <div className="post card">
            <div className="card-content">
              <span className="card-title">{ this.state.msg }</span>
            </div>
          </div>
        </div>
      );
    }
  }
}

export default Post;