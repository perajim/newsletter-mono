let BASE_URL = 'http://localhost:8080/newsletter/recipients';
  export async function AddRecipients(newsletter) {
    const response = await fetch(BASE_URL, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(newsletter)
      })
    return await response.json();
}
  