import React, { useState } from "react";
  
export const Context = React.createContext();
export const ContextProvider = ({ children }) => {
    const [showNotification, setshowNotification] = useState(false);
      
    return (
        <Context.Provider value={[showNotification, setshowNotification]}>
            {children}
        </Context.Provider>
    );
};