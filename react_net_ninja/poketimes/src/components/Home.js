import React from 'react';
import axios from 'axios';
import Rainbow from '../hoc/Rainbow';

class Home extends React.Component {
  state = {
    posts: []
  }

  componentDidMount() {
    axios.get("https://jsonplaceholder.typicode.com/posts")
      .then(res => {
        this.setState({
          posts: res.data.slice(0, 10)
        });
      });
  }

  render() {
    const { posts } = this.state;
    let postsList = [];
    if (posts.length) {
      postsList = posts.map(post => {
        return (
          <div className="post card" key={ post.id }>
            <div className="card-content">
              <span className="card-title">{ post.title }</span>
              <p className="card-body">{ post.body }</p>
            </div>
          </div>
        );
      });
    } else {
      postsList = <div className="center">There is no posts</div>
    }

    return (
      <div className="container">
        <h3 className="center">Home</h3>
        <h4 className="center">Posts: </h4>
        { postsList }
      </div>
    );
  }
}

export default Home;