let BASE_URL = 'http://app:8080/shipping/newsletter/';
  export async function SendNewsletter(id, idFile, data) {
    console.log("va a enviar el correo")
    const response = await fetch(`${BASE_URL}${id}/${idFile}`, {
        method: 'POST', 
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(data)
      })
    return await response.json();
}
  