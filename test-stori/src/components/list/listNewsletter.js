import React,{useState, useEffect, useContext} from 'react';
import { useNavigate  } from "react-router-dom";
import { Context } from '../../context';

export default function ListNewsletter() {
    const history = useNavigate ();
    const [newsletters, setNewsletters] = useState([]);

   
    useEffect(() => {
        const url = "http://app:8080/newsletters";
    
        const fetchData = async () => {
          try {
            const response = await fetch(url);
            const json = await response.json();
            setNewsletters(json.data)
          } catch (error) {
            console.log("error", error);
          }
        };
    
        fetchData();
    }, []);
if (newsletters != null){
    return(
<>
        <div class="flex justify-center">
                    <div class="mb-3 w-9/12">
            <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                    <tr>
                        <th scope="col" className="py-3 px-6">
                            Nombre 
                        </th>
                        <th scope="col" className="py-3 px-6">
                           Correos enviados
                        </th>
                        <th scope="col" className="py-3 px-6">
                            Emails registrados
                        </th>
                    </tr>
                </thead>
                <tbody>

                {newsletters.map((newsletter) => (
                <tr className="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                    <th scope="row" className="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white dark:hover:bg-blue-700">
                    <a href='#'  onClick={e=>{e.preventDefault();history('/newsletter/'+newsletter.ID)}} >
                        {newsletter.Name}
                    </a>
                    </th>
                    <td className="py-4 px-6">
                    {newsletter.SentEmails}
                    </td>
                    <td className="py-4 px-6">
                    {newsletter.Recipients.length}
                    </td>    
                </tr>
                ))}

                </tbody>
            </table>
        </div>
        </div>
        </>
    )}else{
        return(
            <div class="flex justify-center">
            <div class="mb-3 w-9/12">
            <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                    <tr>
                        <th scope="col" className="py-3 px-6">
                            Nombre del Newsletter
                        </th>
                        <th scope="col" className="py-3 px-6">
                            Numeros de correos enviados
                        </th>
                        <th scope="col" className="py-3 px-6">
                            Numero de emails registrados
                        </th>
                    </tr>
                </thead>
                <tbody>
                </tbody>
            </table>
        </div>
        </div>
        )
    }
}