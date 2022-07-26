import React,{useState, useEffect, useContext} from 'react';
import {StoreFile} from '../api/apiStoreFile'
import {SendNewsletter} from '../api/apiSendNewsletter'
import { useParams } from 'react-router-dom';
import Modal from '../modal/modal';
import  ModalAddEmails  from '../modal/addEmails'
import ModalAddEmail from '../modal/addEmail';
import { Context } from "../../context";


export default function Newsletter() {
    const { id } = useParams();
    const [selectedFile, setSelectedFile] = useState();
    const [showNotification, setshowNotification, showAddRecipients, setshowAddRecipients,showAddRecipient, setshowAddRecipient] = useContext(Context)
    let titleModal = "Newsletter Enviado con exito"
    let descriptionModal = "Se envio correctamente el newsletter, a la lista de contactos"

    let titleModalEmail = "Se agrego el email"
    let descriptionModalEmail = "Se agrego correctamente el Email"

    const changeHandler = (event) => {
        event.preventDefault(); 
		setSelectedFile(event.target.files[0]);
	};

    const changeModalAddEmails = (event) => {
        event.preventDefault(); 
		setshowAddRecipients(true)
	};
    const changeModalAddEmail = (event) => {
        event.preventDefault(); 
		setshowAddRecipient(true)
	};

    const handleSubmission =async (e) => {
        e.preventDefault(); 

        let { contentForm, subject } = document.forms[0];
        let data = {"content":contentForm.value,"subject":subject.value}

		const formData = new FormData();
		formData.append('file', selectedFile, selectedFile.name);

        StoreFile(formData)
        .then(response => {
            let idFile = response.file.id  
            SendNewsletter(id, idFile, data)
            .then(response => {
                setshowNotification(true)
              })
              .catch ( e=> {
                console.log("fracaso rotundo")
                console.log(e)
              })
          })
          .catch ( e=> {
            console.log(e)
          })
	};
    
	
	return(
        <div> 
            <ModalAddEmails idNewsletter={id}/>
            <ModalAddEmail title={titleModalEmail} description={descriptionModalEmail} idNewsletter={id} /> 
            <Modal title={titleModal} description={descriptionModal} />          
            <form onSubmit={handleSubmission} className="space-y-6">
                <div class="flex justify-center">
                        <div class="mb-3 w-9/12">
                    <label htmlFor="first_name" className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Asunto del newsletter</label>
                    <input name="subject" type="subject" id="subject" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Ingrese el Asunto con el que se enviara el newsletter" required/>
                </div>
                </div>
                <div class="flex justify-center">
                    <div class="mb-3 w-9/12">
                        <label htmlFor="message" className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-400">Escribe el contenido escrito del Newsletter</label>
                        <textarea id="contentForm" name="contentForm" type="contentForm" rows="4" class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="El newsletter de este dia..."></textarea>
                    </div>
                </div>
                <div class="flex justify-center">
                    <div class="mb-3 w-9/12">
                        <label for="formFile" class="form-label inline-block mb-2 text-gray-700">Archivo PDF o PNG del newsletter</label>
                        <input class="form-control
                        block
                        w-full
                        px-3
                        py-1.5
                        text-base
                        font-normal
                        text-gray-700
                        bg-white bg-clip-padding
                        border border-solid border-gray-300
                        rounded
                        transition
                        ease-in-out
                        m-0
                        focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none" type="file" name="file" id="formFile" onChange={changeHandler}/>
                    </div>
                </div>
                <div class="flex justify-center">
                    <div class="mb-1 ">
                        <button type="submit"  class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                            Enviar newsletter
                        </button>
                    </div>
                </div>
            </form>
            <div class="flex justify-center">
                <div class="mb-2 w-9/12">
                    <button onClick={changeModalAddEmails} class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded">
                        Agregar lista de Emails
                    </button>
                    <button onClick={changeModalAddEmail} class="bg-stone-500 hover:bg-stone-700 text-white font-bold py-2 px-4 rounded">
                        Agregar un solo email
                    </button>
                </div>
            </div>
        </div>
         )
}