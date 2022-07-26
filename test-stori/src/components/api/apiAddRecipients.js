let BASE_URL = 'http://app:8080/newsletter/recipients';
  export async function AddRecipients(newsletter) {
    const response = await fetch(BASE_URL, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(newsletter)
      })
    return await response.json();
}
  