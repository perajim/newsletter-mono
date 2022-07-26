import React, { useState } from "react";
  
export const Context = React.createContext();
export const ContextProvider = ({ children }) => {
    const [showNotification, setshowNotification] = useState(false);
    const [showAddRecipients, setshowAddRecipients] = useState(false);
    const [showAddRecipient, setshowAddRecipient] = useState(false);
      
    return (
        <Context.Provider value={[showNotification, setshowNotification, showAddRecipients, setshowAddRecipients,showAddRecipient, setshowAddRecipient]}>
            {children}
        </Context.Provider>
    );
};