import React from 'react';
import { connect } from 'react-redux';
import Pokeball from '../pokeball.png';
import { deletePost } from '../actions/postActions';

class Post extends React.Component {

  handleClick = () => {
    const postId = this.props.post.id;
    this.props.deletePost(postId)
    this.props.history.push('/');
  }

  render() {
    const post = this.props.post;

    return (
      <div className="container">
        <div className="post card">
          <div className="card-content">
            <img src={ Pokeball } alt="A Pokeball"/>
            <span className="card-title red-text">{ post.title }</span>
            <p className="card-body black-text">{ post.body }</p>
            <button className="btn gray" onClick={ this.handleClick }>Delete Post</button>
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

const mapDispatchToProps = (dispatch) => {
  return {
    deletePost: (postId) => dispatch(deletePost(postId))
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Post);