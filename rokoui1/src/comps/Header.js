import React, { Component } from 'react';


class Header extends Component {
  render() {
   
    return (
    <div class="jumbotron text-center" style={{marginBottom:0,marginTop:0,backgroundImage:`url('header.png')`}}>
        <h1 style={{color:`#CA9932`}}>RoKoBook</h1>
        <p style={{marginBottom:0,marginTop:0,color:`#CA9932`,backgroundColor:`#DCECED`}}>Booking para artistas, bookers, managers y concert venues</p>
    </div>
     
    );
  }
}

export default Header;
