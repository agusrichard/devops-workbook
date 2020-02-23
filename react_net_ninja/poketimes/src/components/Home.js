import React from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import Pokeball from '../pokeball.png';

class Home extends React.Component {
  
  render() {
    const { posts } = this.props;
    let postsList = [];
    if (posts.length) {
      postsList = posts.map(post => {
        return (
          <div className="post card" key={ post.id }>
            <Link to={ '/' + post.id }>
              <div className="card-content">
                <img src={ Pokeball } alt="A Pokeball" />
                <span className="card-title red-text">{ post.title }</span>
                <p className="card-body black-text">{ post.body }</p>
              </div>
            </Link>
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

const mapStateToProps = (state) => {
  return {
    posts: state.posts
  }
}

export default connect(mapStateToProps)(Home);