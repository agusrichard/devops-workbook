import React from 'react';

function Bio(props) {
  return (
    <div>
      <p>
        Bio: <button onClick={ () => props.deleteUser(props.user.id) }>Delete</button>
      </p>
      <p>Name: { props.user.name }</p>
      <p>Age: { props.user.age }</p>
      <br />
    </div>
  );
}

export default Bio;