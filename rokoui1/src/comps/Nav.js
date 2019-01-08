import React, { Component } from 'react';
import Menuitem from './Menuitem';

class Nav extends Component {
    render() {
      return (
        <nav class="navbar navbar-expand-sm bg-dark navbar-dark" align="center">
            {/* <a class="navbar-brand" href="#">Navbar</a> */}
            <h4 style={{color:'#EBEBEB'}}>Menú:</h4>&nbsp;&nbsp;&nbsp;
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#collapsibleNavbar">
            <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="collapsibleNavbar" align="center">
            <ul class="navbar-nav">
                <Menuitem 
                    texto="Iniciar sesión"
                    liga="#"
                />
                <Menuitem 
                    texto="Artistas"
                    liga="#"
                />
                <Menuitem 
                    texto="Bookers"
                    liga="#"
                />
                <Menuitem 
                    texto="Concert Venues"
                    liga="#"
                />
                <Menuitem 
                    texto="Salir"
                    liga="#"
                />
            </ul>
            </div>  
        </nav>
        );
    }
  }
  
  export default Nav;
