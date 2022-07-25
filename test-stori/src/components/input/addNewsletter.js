import React,{useContext } from 'react';
import { CreateNewsletter } from '../api/apiCreateNewsletter'
import Modal from '../modal/modal';
import { Context } from "../../context";
import { useNavigate  } from "react-router-dom";


export default function AddNewsletter() {
    const [showNotification, setshowNotification] = useContext(Context)
    let titleModal = "Se creo el nuevo newsletter"
    let descriptionModal = "Se creo el nuevo newsletter, ahora puedes agregar correos y enviarles información"
    const history = useNavigate ();

    const createNewsletter = (e) => {
        e.preventDefault(); 
        let { name } = document.forms[0];
        let data = {"name":name.value}
        CreateNewsletter(data)
        .then(response => {
            setshowNotification(true)
            setTimeout(function(){ history("/newsletter/"+response.newsletter.id);setshowNotification(false) }, 5000);         
          })
          .catch ( e=> {
            console.log(e)
          })
    } 

    return(
        <div>
        <Modal title={titleModal} description={descriptionModal}/>          
        <div class="flex justify-center">
        <div class="mb-3 w-9/12">      
            <form onSubmit={createNewsletter}>
                <div>
                <div class="flex justify-center">
                                <div class="mb-1 ">
                    <label htmlFor="first_name" className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Crear newsletter</label>
                    </div>
                    </div>
                    <input name="name" type="name" id="name" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Ingrese el nombre del newsletter" required/>
                </div>
                <div class="flex justify-center">
                    <div class="mb-1 ">
                        <button type="submit" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">Crear</button>
                    </div>
                </div>
            </form>
        </div>
        </div>
        </div>
    )
}