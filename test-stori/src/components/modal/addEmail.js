/* This example requires Tailwind CSS v2.0+ */
import { Fragment, useState, useContext } from 'react'
import { Dialog, Transition } from '@headlessui/react'
import { CheckIcon } from '@heroicons/react/outline'
import { Context } from '../../context'

export default function ModalAddEmail(props) {
    const [showNotification, setshowNotification] = useContext(Context)
    return (
        <Transition.Root show={true} as={Fragment}>
        <Dialog as="div" className="fixed z-10 inset-0 overflow-y-auto" onClose={setshowNotification}>
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
                    <div class="mb-3 w-9/12">      
                        <form>
                            <div>
                            <div class="flex justify-center">
                                            <div class="mb-1 ">
                                <label htmlFor="first_name" className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Agregar Email</label>
                                </div>
                                </div>
                                <input name="name" type="name" id="name" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Ingrese el Email" required/>
                            </div>
                            <div class="flex justify-center">
                                <div class="mb-1 ">
                                    <button type="submit" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">Agregar</button>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
                </div>
            </Transition.Child>
            </div>
        </Dialog>
        </Transition.Root>
    )
}
