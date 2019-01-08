import React from 'react';

const Liitem = (props) => {
    return(
        <li class="nav-item" style={{marginBottom:10}}>
            <a class={props.status} href={props.liga}>{props.titulo}</a>&nbsp; {props.desc}
        </li>    
    )  
}

export default Liitem;