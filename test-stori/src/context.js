import React, { useState } from "react";
  
export const Context = React.createContext();
export const ContextProvider = ({ children }) => {
    const [showNotification, setshowNotification] = useState(false);
    const [showAddRecipients, setshowAddRecipients] = useState(false);
    const [showAddRecipient, setshowAddRecipient] = useState(false);
    const [showAddEmailSuccess, setShowAddEmailSuccess] = useState(false);
      
    return (
        <Congittext.Provider value={[showNotification, setshowNotification, showAddRecipients, setshowAddRecipients,showAddRecipient, setshowAddRecipient, showAddEmailSuccess, setShowAddEmailSuccess]}>
            {children}
        </Context.Provider>
    );
};