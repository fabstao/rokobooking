import React from 'react';

const Menuitem = (props) => {
    return(
        <li class="nav-item">
            <a class="nav-link" href={props.liga}>{props.texto}</a>
        </li>    
    )  
}

export default Menuitem;