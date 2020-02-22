import React from 'react'

class AddUser extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      id: null,
      name: '',
      age: null
    }
  }

  handleChange = event => {
    this.setState({
      [event.target.id]: event.target.value
    });
  };

  handleSubmit = event => {
    event.preventDefault();
    this.props.addUser(this.state);
  };

  render() {
    return (
      <div className="add-user">
        <h3>Add User:</h3>
        <form onSubmit={ this.handleSubmit }>
          <div>
            <label htmlFor="name">Name: { this.state.name }</label><br />
            <input type="text" id="name" onChange={ this.handleChange }/>
          </div>
          <div>
            <label htmlFor="age">Age: { this.state.age }</label><br />
            <input type="text" id="age" onChange={ this.handleChange }/>
          </div>
          <button>Submit</button>
        </form>
      </div>
    );
  }
}

export default AddUser