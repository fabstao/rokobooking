import React, { Component } from 'react';
import Liitem from './Liitem';

class Container extends Component {
    render() {
        let activa="nav-link active";
        let inactiva="nav-item";
        let desactivada="nav-link disabled";
        return (
          <div class="container" style={{marginTop:20}}>
            <div class="row">
                <div class="col-sm-8">
                    <h3>Lista de opciones</h3>
                        <p>Descripción de la lista</p>
                        <ul class="nav nav-pills flex-column">
                            <Liitem 
                                status={inactiva}
                                titulo="Concierto"
                                desc="Beneficiencia"
                                liga="#"
                                />
                            <Liitem 
                                status={inactiva}
                                titulo="Entrevista"
                                desc="Descripción del evento Lorem impsum Sata Pata"
                                liga="#"
                                />
                            <Liitem 
                                status={inactiva}
                                titulo="Tocada bar"
                                desc="Descripción del evento"
                                liga="#"
                                />
                        </ul>
                </div>
                <div class="col-sm-4" style={{backgroundColor:`#CDCDCD`,marginBottom:20}}>
                    <h4>Notificaciones</h4>
                        <ul class="nav nav-pills flex-column">
                            <Liitem 
                                status={inactiva}
                                titulo="Concierto"
                                desc="Beneficiencia"
                                liga="#"
                                />
                            <Liitem 
                                status={inactiva}
                                titulo="Entrevista"
                                desc="Descripción del evento"
                                liga="#"
                                />
                            <Liitem 
                                status={inactiva}
                                titulo="Tocada bar"
                                desc="Descripción del evento"
                                liga="#"
                                />
                        </ul>
                </div>
            </div>
          </div>
      );
    }
}


export default Container;