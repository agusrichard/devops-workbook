import React from 'react';
import axios from 'axios';

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
        <div className="post card">
          <div className="card-content">
            <span className="card-title">{ post.title }</span>
            <p className="card-body">{ post.body }</p>
          </div>
        </div>
      );
    } else if (!this.state.post) {
      return (
        <div className="post card">
          <div className="card-content">
            <span className="card-title">Loading...</span>
          </div>
        </div>
      );
    } else {
      return (
        <div className="post card">
          <div className="card-content">
            <span className="card-title">{ this.state.msg }</span>
          </div>
        </div>
      );
    }
  }
}

export default Post;