import React from 'react';
import Rainbow from '../hoc/Rainbow';

function Home() {
  return (
    <div className="container">
      <h3 className="center">Home</h3>
      <p>
        Lorem ipsum dolor sit amet, consectetur adipisicing elit. 
        Vel officia, facilis inventore iure explicabo saepe numquam id 
        laudantium iusto adipisci perspiciatis quibusdam reiciendis assumenda, 
        quo tempora nobis culpa ex ipsam?
      </p>
    </div>
  );
}

export default Rainbow(Home);