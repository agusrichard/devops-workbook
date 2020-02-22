import React from 'react';
import Bio from './Bio';
import AddUser from './AddUser';

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      users: [
        { 
          id: 1,
          name:'Agus Richard',
          age: 22
        },
        { 
          id: 2,
          name:'Sherlock Holmes',
          age: 35
        },
        {
          id: 3,
          name:'John Watson',
          age: 34
        },
        { 
          id: 4,
          name:'Irene Adler',
          age: 27
        },
        { 
          id: 5,
          name:'James Moriarty',
          age: 37
        }
      ]
    }
  }

  addUser = newUser => {
    newUser.id = this.state.users.length + 1;
    let newUsersList = [...this.state.users, newUser];
    this.setState({
      users: newUsersList
    })
  };

  deleteUser = userId => {
    let newUsersList = this.state.users.filter(user => user.id !== userId);
    this.setState({
      users: newUsersList
    })
  }

  componentDidMount() {
    console.log('Mounted');
  }

  componentDidUpdate(prevProps, prevState) {
    console.log('Updated');
    console.log(prevProps, prevState);
  }

  render() {
    console.log('Rendered');

    const usersList = this.state.users.map(user => {
      return <Bio key={ user.id } user={ user } deleteUser={ this.deleteUser }/>
    });

    return (
      <div>
        <h1>Welcome!!!</h1>
        <AddUser addUser={ this.addUser } />
        <div className="list-of-users">
          <h3>List of Users:</h3>
          { usersList }
        </div>
      </div>  
    )
  }
}

export default App;
