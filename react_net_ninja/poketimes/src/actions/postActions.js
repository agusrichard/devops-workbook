export const deletePost = (postId) => {
  return {
    type: 'DELETE_POST',
    postId
  }
};