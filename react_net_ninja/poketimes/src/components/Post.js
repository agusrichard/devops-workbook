import React from 'react';
import { connect } from 'react-redux';
import Pokeball from '../pokeball.png';

class Post extends React.Component {
  render() {
    const post = this.props.post;

    return (
      <div className="container">
        <div className="post card">
          <div className="card-content">
            <img src={ Pokeball } alt="A Pokeball"/>
            <span className="card-title red-text">{ post.title }</span>
            <p className="card-body black-text">{ post.body }</p>
          </div>
        </div>
      </div>
    );
  }
}

const mapStateToProps = (state, ownProps) => {
  const postId = ownProps.match.params.post_id;
  return {
    post: state.posts.find(post => post.id === parseInt(postId))
  }
}

export default connect(mapStateToProps)(Post);