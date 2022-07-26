/* This example requires Tailwind CSS v2.0+ */
import { Fragment, useState, useContext } from 'react'
import { Dialog, Transition } from '@headlessui/react'
import { CheckIcon } from '@heroicons/react/outline'
import { Context } from '../../context'
import { AddRecipients } from '../api/apiAddRecipients'

export default function ModalAddEmail(props) {
    const [showNotification, setshowNotification, showAddRecipients, setshowAddRecipients,showAddRecipient, setshowAddRecipient, showAddEmailSuccess, setShowAddEmailSuccess] = useContext(Context)
    const [formFields, setFormFields] = useState([
        { name: ''},
      ])
    
      const handleFormChange = (event, index) => {
        let data = [...formFields];
        data[index][event.target.name] = event.target.value;
        setFormFields(data);
      }
    
      const submit = (e) => {
        e.preventDefault();
        let data = {"_id":"62df829358eb73ba6ec7a42b","EmailRecipients":[]} 
        formFields.map(field =>{
            data.EmailRecipients.push(field.name)
        })
        AddRecipients(data)
        .then(response => { 
            console.log(response)
           
          })
          .catch ( e=> {
            console.log(e)
          })
        console.log(data)
      }
    
      const addFields = () => {
        let object = {
          name: ''
        }
    
        setFormFields([...formFields, object])
      }
    
      const removeFields = (index) => {
        let data = [...formFields];
        data.splice(index, 1)
        setFormFields(data)
      }
    return (
        <>
        <Transition.Root show={showAddRecipients} as={Fragment}>
        <Dialog as="div" className="fixed z-10 inset-0 overflow-y-auto" onClose={setshowAddRecipients}>
            <div className="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
            <Transition.Child
                as={Fragment}
                enter="ease-out duration-300"
                enterFrom="opacity-0"
                enterTo="opacity-100"
                leave="ease-in duration-200"
                leaveFrom="opacity-100"
                leaveTo="opacity-0"
            >
                <Dialog.Overlay className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
            </Transition.Child>

            {/* This element is to trick the browser into centering the modal contents. */}
            <span className="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">
                &#8203;
            </span>
            <Transition.Child
                as={Fragment}
                enter="ease-out duration-300"
                enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                enterTo="opacity-100 translate-y-0 sm:scale-100"
                leave="ease-in duration-200"
                leaveFrom="opacity-100 translate-y-0 sm:scale-100"
                leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
            >
                <div className="relative inline-block align-bottom bg-white rounded-lg px-4 pt-5 pb-4 text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-sm sm:w-full sm:p-6">
                <div class="flex justify-center">
                <div className="App">
                <form onSubmit={submit}>
                    {formFields.map((form, index) => {
                    return (
                        <div key={index}>
                        <input
                            name='name'
                            placeholder='Name'
                            onChange={event => handleFormChange(event, index)}
                            value={form.name}
                        />
                        <button onClick={() => removeFields(index)} className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">Remove</button>
                        </div>
                    )
                    })}
                </form>
                <button onClick={addFields} className="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded">+</button>
                <br />
                <button onClick={submit} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">Agregar Correos</button>
            </div>
                </div>
                </div>
            </Transition.Child>
            </div>
        </Dialog>
        </Transition.Root>
        </>
    )
}
