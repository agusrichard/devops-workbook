import React from 'react';

function Rainbow(WrappedComponent) {

  const colours = ['red', 'blue', 'yellow', 'green', 'orange', 'pink'];
  const colour = colours[Math.floor(Math.random() * (colours.length + 1))];
  const className = colour + '-text';

  return (props) => {
    return (
      <div className={ className }>
        <WrappedComponent {...props} />
      </div>
    );
  }
}

export default Rainbow;