import React, { Component } from 'react';
import Header from './Header';
import Nav from './Nav';
import Container from './Container';


class App extends Component {
  render() {
    const menues=[
      {titulo: "Iniciar sesión", liga: "#"},
      {titulo: "Artistas", liga: "#"},
      {titulo: "Concert Venues", liga: "#"},
      {titulo: "Managers", liga: "#"},
      {titulo: "Salir", liga: "#"}
    ];
    const menuen=[
      {titulo: "", liga: "#"},
      {titulo: "", liga: "#"},
      {titulo: "", liga: "#"}
    ];
    return (
      <div>
        <Header 
        />
        <Nav 
          menu={menues}
        />
        <Container />
      </div>
    );
  }
}

export default App;
