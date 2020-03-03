import * as React from 'react';

export interface HelloProps { name: string; }

class App extends React.Component<HelloProps, {}> {
  render() {
    return (
      <div>
        <h1>Helloo {this.props.name}</h1>
      </div>
    );
  }  
}

export default App;
